# Issue - NX Packages without default export are excluded from dependencies

## Steps to recreate

1. Open this repo in a devcontainer
2. Run `pnpm i`
3. Run `nx graph --file ./graph.json`
4. Open `nx graph` to open web UI as well

Note that the `foobar` has a dependency on `foo`.

`bar` is not declared as a dependent, despite the fact that it is equally declared as a dependency in [package.json](./foobar/package.json).

This means that targets that require previous projects to be fully built or tested may fail, as the dependency order will not be enforced.
It can also impact plugins like [nx-update-ts-references](https://www.npmjs.com/package/nx-update-ts-references) which will fail to detect the project as a local dependency.

## Issue explained

This repo has 3 packages: `foo`, `bar`, and `foobar`.

`foo` and `bar` are both simple packages that export a single javascript file (`foo.js` and `bar.js` respectively).
`foobar` depends on both of these files.

Each file has a bare-minimum `project.json` that includes the `echo` command (which just echos "hello world"). 
This target is specified to [depend on upstream projects](./nx.json) to help visualize these dependencies.

Both export via the `exports` field in package.json.

However `foo` exports this file as a index file: [".": "/.foo.js"](./foo/package.json).
`bar` exports this file under its name: ["./bar": "./bar.js"](./bar/package.json).

As a result, nx does pick up this export as a primary dependency.

This is because the logic to declare a dependency as "local" as opposed to "external" is implemented in [this line](https://www.npmjs.com/package/nx-update-ts-references).

```ts
const localProject = targetProjectLocator.findDependencyInWorkspaceProjects(d);
```

The implementation logic for this method is [fairly straightforward](https://github.com/nrwl/nx/blob/master/packages/nx/src/plugins/js/project-graph/build-dependencies/target-project-locator.ts#L257).
If the name of the project exists in the map, return in it

```ts
return this.packageEntryPointsToProjectMap[dep]?.name ?? null;
```

However it will also calculate this map (memoized), which [exposes the bug](https://github.com/nrwl/nx/blob/master/packages/nx/src/plugins/js/utils/packages.ts#L18-L24):
```ts
if (!packageExports || typeof packageExports === 'string') {
    // no `exports` or it points to a file, which would be the equivalent of
    // an '.' export, in which case the package name is the entry point
    result[packageName] = project;
} else {
    for (const entryPoint of Object.keys(packageExports)) {
        result[join(packageName, entryPoint)] = project;
    }
}
```

`result['foo']` will get written because `join('foo', '.')` returns `foo`.
`result['foobar']` would also work because it lacks an exports field.

Howecer `result['bar/bar']` is the result of bar because it only has an export for that path.
But the dependency is checked against `'bar'`, and therefore appears to be an external dependency.