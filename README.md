# Package - Mocha

## Issue - Modules in errors cause unhandledRejection, tests quit early

https://github.com/mochajs/mocha/issues/4887

Open this repo in a devcontainer (not strictly necessary, but guarantees you have all the working parts setup without polluting the rest of your workspaces).

Then run:

* `pnpm i`
  * Installs the latest mocha + chai packages
* `pnpm run test`

You should see an incomplete output, that looks like:

```
issue-recreator@<id>:/workspace$ pnpm run test

> @ test /workspace
> mocha ./test.js


Compare imports
```
