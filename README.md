# Issue - Dagger pre-call filtering does not apply cross-module

Dagger supports [pre-call filtering](https://docs.dagger.io/api/filters/#pre-call-filtering). This is implemented as annotations on the directory parameters to strip out certain files/directories before any processing happens.

However this _only_ applies to the top-level module that is called via CLI.

Every other module that is used internally as a dependency and may implement pre-call filtering is completely ignored.
Therefore it is unsafe to trust pre-call filtering as a way to reduce the size of directories, and filtering _must_ also be done via `WithDirectory -> Exclude` option.

## Steps to reproduce

1. Clone this repo.
2. Open in devcontainer (optional if dagger is installed locally)
3. `dagger call --mod ./dagger/no-bar/ remove-bar --source .`
    * Success. Pre-call filtering removes the `bar` directory.
4. `dagger call --mod ./dagger/no-foo/ remove-foo --source .`
    * Success. Pre-call filtering removes the `foo` directory.
5. Lets try combining the two! `dagger call --mod ./dagger/no-foo-or-bar/ remove-foo-and-bar --source .`
    * Error! No filter is ever applied!

This could be fixed by manually writing something like 
```
source.WithoutDirectory("foo")
```

but the fact that it _sometimes_ works without it appears to break the deterministic behavior that Dagger highlights.

## Feature Request

Ignore annotations on directories for custom functions should _always_ be respected, even for module calls that are not top-level.