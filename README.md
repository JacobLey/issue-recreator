# Issue - pnpm false positive peer dependency warning when aliased

pnpm will flag a peer dependency as missing if:

* The package with the peer dependency is the one being installed
  * in this recreation, that is the [b](./b/package.json) packagem, declaring a peer dependency on `eslint`
  * Eslint is not "special" in this recreation. Virutally any package could be used for demonstration, but it is a common package to be referenced in peer dependencies.
* The package _installing_ the peer dependency, installs it under an alias
  * in this recreation, that is [c](./c/package.json), which installs both `b` and `eslint` as an aliased `custom-eslint`.
* This package is yet again installed by _another_ package.
  * in this recreation, that is [d](./d/package.json)

## Steps to reproduce

1. `pnpm --filter "./*" exec pnpm pack`
    * For sake of reproduction, simple pack of trivial packages is easiest. But in real world, these packages would probably already be published to npm.
2. `pnpm i`
    * Note the warning messages about "missing peer" (there is also an "unmet peer" warning. This is not the topic of the issue, and does not recreate when using npm instead of tgz)
```
d
└─┬ c 0.0.1
  └─┬ b 0.0.1
    └── ✕ missing peer a@file:../a/a-0.0.1.tgz
Peer dependencies that should be installed:
  a@file:../a/a-0.0.1.tgz
```
3. `pnpm --filter "./d" exec node ./d.js`
   * logs: `{ c: { eslint: [Object] } } }`
   * The command is successful! So the peer dependency must actually exist, and therefore be a false positive.

## Repo explained

There are 3 packages in this reproduction.

* [b](./b/package.json)
  * This package imports `eslint` and re-exports with a wrapped `b = { eslint }`.
  * `eslint` is a peer dependency. It is not installed directly, and as a result the `b` package has no `node_modules` file.
* [c](./c/package.json)
  * This package imports both `eslint` and `b`, although `eslint` is aliased as a different name `custom-eslint`. The resolution is the same though. 
  * It re-exports b with a wrapped `c = { b }`. This should only be possible if `b` is successfully loaded, which in turn means the peer dependency was satisfied.
* [d](./d/package.json)
  * This package imports `c` normally. Nothing is suspicious about this package directly, but it is the source of our error!

## Real world equivalent

The `b` package would be the actual plugin implementation. It needs to depend on eslint, but shouldn't bring it's own implemention, since it needs to be an exact match for the user's version.
When discovering this issue, that was my [haywire-launcher](https://github.com/JacobLey/leyman/blob/main/tools/haywire-launcher/package.json#L28-L29) package

The `c` package would be the user that is actually consuming the plugin. In most cases users installed dependencies based on the name (e.g. `"eslint": "9.17.0"` as opposed to `"eslint": "npm:eslint@9.17.0`), but that isn't strictly required.
A possible reason to use an alias is to indicate a package that _must_ come from npm, as opposed to a locally named package.
When discovering this issue, that was my [nx-update-ts-references](https://github.com/JacobLey/leyman/blob/main/apps/nx-update-ts-references/package.json#L39-L43) package, which loads those packages (all of them, so dependencies are all fulfilled) from npm to avoid circular dependencies (those packages in turn depend on the _local_ version, e.g. [haywire-launcher](https://github.com/JacobLey/leyman/blob/main/tools/haywire-launcher/package.json#L48))

The `d` package is simply a package that imports `c`. Maybe its yet another plugin wrapper.
When discovering this issue, there were multiple packages that fulfilled this, for example [common-proxy](https://github.com/JacobLey/leyman/blob/main/tools/common-proxy/package.json#L44)

## Notes

This issue does not occur when performing local links (either via the `workspace:^` keyword, or `link:../a`). This is a false negative though, since the script does not work.

The workaround I've used locally is to also install the peer dependency as a dev dependency. I have had no issues so far.

An even more simple reproduction is add a dependency on [nx-update-ts-references](https://www.npmjs.com/package/nx-update-ts-references) and see the error.
The root `package.json` has done that already.
```
.
└─┬ nx-update-ts-references 0.1.1
  └─┬ haywire-launcher 0.1.9
    ├── ✕ missing peer entry-script@^3.0.7
    └── ✕ missing peer haywire@^0.1.6
Peer dependencies that should be installed:
  entry-script@^3.0.7  haywire@^0.1.6
```