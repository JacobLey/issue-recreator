# Issue - NX run parallelism does not work in DevContainer

## Steps to repoduce issue

1. Open this package in a DevContainer.
  * This code was personally tested in VSCode.
2. Run `npm i` to install Nx
  * A "global" version was already installed during DevContainer spin up
3. Run `nx run example:c --verbose`

The terminal looks like the following:
```
$ nx run example:c --verbose

 NX   Running target c for project example and 2 tasks it depends on:

—————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————

> nx run example:a


> nx run example:b


—————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————

 NX   Running target c for project example and 2 tasks it depends on failed

Failed tasks:

- example:b

Hint: run the command with --verbose for more details.
```

Inspecting the `nx.json` file, every task is a simple `nx:noop` executor.
The selected task `C` depends on both `A` + `B`. `A` is successful, but `B` fails.

No extra information is outputted with the failure.

## Avoiding the issue

Editing either `A` or `B` to depend on the other fixes the issue. This is not a true solution though, it simply prevents the parallelism that breaks this code.

It also does not appear to happen in non-devcontainers.

## Diagnosing the issue

The error is swallowed here: https://github.com/nrwl/nx/blob/master/packages/nx/src/tasks-runner/task-orchestrator.ts#L495

Editing the line of code locally to print the failure message results in

```
Error: Operation not permitted (os error 1)

at PseudoTerminal.fork (/workspace/node_modules/nx/src/tasks-runner/pseudo-terminal.js:42:73)

at ForkedProcessTaskRunner.forkProcessWithPseudoTerminal (/workspace/node_modules/nx/src/tasks-runner/forked-process-task-runner.js:141:45)


at ForkedProcessTaskRunner.forkProcess (/workspace/node_modules/nx/src/tasks-runner/forked-process-task-runner.js:127:25)

at TaskOrchestrator.runTaskInForkedProcess (/workspace/node_modules/nx/src/tasks-runner/task-orchestrator.js:260:54)
                                             at TaskOrchestrator.applyFromCacheOrRunTask (/workspace/node_modules/nx/src/tasks-runner/task-orchestrator.js:243:61)

at process.processTicksAndRejections (node:internal/process/task_queues:95:5)
                                                                                          at async TaskOrchestrator.executeNextBatchOfTasksUsingTaskSchedule (/workspace/node_modules/nx/src/tasks-runner/task-orchestrator.js:76:13)

at async Promise.all (index 1)
                                            at async TaskOrchestrator.run (/workspace/node_modules/nx/src/tasks-runner/task-orchestrator.js:53:9)

at async defaultTasksRunner (/workspace/node_modules/nx/src/tasks-runner/default-tasks-runner.js:18:16) {
 code: 'GenericFailure'
}
```

The forking happens on a [native binding](https://github.com/nrwl/nx/blob/master/packages/nx/src/native/index.js#L261), which is why the stack trace does not go farther.

It appears that more than one fork at a time is not supported.