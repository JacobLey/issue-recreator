# Issue - NX Run-Commands uses wrong cwd

Highly recommended to open this issue in VSCode/DevContainer. Guarantees any regressions remain consistent (e.g. not accidentally depending on globally installed packages).

See [Dockerfile](./.devcontainer/Dockerfile) for relevant installation/version information if not using DevContainer.

## Issue Description

When running commands with the [run-commands](https://nx.dev/nx-api/nx/executors/run-commands) executor by default the working directory is the root.

Often times it is desired to automatically move to the directorying containing the target project automatically. The `cwd` option set to `{projectRoot}` (`options: {"cwd": "{projectRoot}" }`).

From documentation: https://nx.dev/reference/nx-json#outputs

> When defining any options or configurations inside of a target default, you may use the {workspaceRoot} and {projectRoot} tokens. This is useful for defining options whose values are paths.

We have three simple projects, `foo`, `bar`, and `baz`.

All have a `project.json` which declares a single `print-pwd` target, which, you guessed it, prints the working directory.

It does this via the `run-commands` executor, and a command of `pwd`.

There are no cached results, other targets defined, or dependencies between these projects.

Both the `foo` and the `bar` project inherit their target definition from a shared `nx.json` definition.

The `baz` project redefines the target in `project.json` (with the exact same options).

### Expected behavior

Running the following commands should result in following output

* `nx run bar:print-pwd`
  * > /workspace/packages/bar
* `nx run baz:print-pwd`
  * > /workspace/packages/baz
* `nx run foo:print-pwd`
  * > /workspace/packages/foo

### Actual behavior

Running the following commands results in following output

* `nx run bar:print-pwd`
  * > /workspace/packages/bar
* `nx run baz:print-pwd`
  * > /workspace/packages/baz
* `nx run foo:print-pwd`
  * > /workspace/packages/bar
  * WRONG!

The foo project displays a working directory of bar!

I suspect this is because bar is prior to foo alphabetically (bar is just lucky to end up getting itself).

We do not see the issue re-created with `baz` who redefines the defintion.

From what I can see, having `bar` redefine the target has no impact on this issue.

## Steps to recreate issue

* `pnpm i`
  * Install necessary packages
* `nx run foo:print-pwd`
  * Inspect the console output for path to bar package

## Notes

This appears to be a regression due to `19.5.0-beta.2`.

I am not able to reproduce in `19.4.3`.