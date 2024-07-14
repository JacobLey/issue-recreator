# Issue - NX Run-Commands uses wrong cwd

Highly recommended to open this issue in VSCode/DevContainer. Guarantees any regressions remain consistent (e.g. not accidentally depending on globally installed packages).

See [Dockerfile](./.devcontainer/Dockerfile) for relevant installation/version information if not using DevContainer.

## Issue Description

When defining a task graph that has tasks defined in both `nx.json` and `project.json`, it is expected that the `dependsOn` is inherited by the `nx.json` implemention if not otherwise overridden by `project.json`.

In fact, this is consistent with the task graph displayed by `nx graph`.

However, when overriding targets in `project.json`, the dependency order is not always respected. This conflicts with the graph which claims it still is.

Create two "meta-targets" `build` and `test`. These are simply `nx:noop` executors that point to a series of real targets for user convenicence.

Building is bound to a single target `build-impl`.
Testing is bound to two targets `test-impl` and `report-impl`. These both depend on `build` being complete.

For simplicity, these default implementations are a 1 second wait, then logging their name.

`report-impl` is overwritten in `project.json` to just immediately log the name.

Inspect the task graph `nx graph` appears to confirm that a test command would:
1. Run build
2. Run test
3. Generate report

### Expected behavior

Running `nx run foo:test` should result in the following logs:

1. > BUILD
2. > TEST
3. > REPORT

### Actual behavior

Running `nx run foo:test` actually results in the following logs:

1. > REPORT
2. > BUILD
3. > TEST

The report step is executed immediately, despite the dependency graph claiming otherwise.

## Steps to recreate issue

* `pnpm i`
  * Install necessary packages
* `nx run foo:test`
  * Inspect the console output shows `REPORT` _before_ `BUILD`, meaning that it did not properly block on a dependency.