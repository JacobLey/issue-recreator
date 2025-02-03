# Issue - Dagger UI shows repeated steps

Dagger traces are shown in [Dagger UI](https://v3.dagger.cloud) and help see every call against dagger native methods
as well as external modules.

Ideally this is roughly a 1-1, with one "span" per call to dagger.

However when referencing resources that may have already "resolved" and then are used later, dagger will repeat the span.

When dealing with directories that may get repeated many times in future, this can result in major multiplier to number of spans shown compared to the actual execution size.

Not only does this result in noise, this actually bogs down the dagger site, and can eventually crash the page, depite the fact that the resources in use are minimal.

## Steps to reproduce

1. Clone this repo.
2. Open in devcontainer (optional if dagger is installed locally)
3. `dagger login`
4. `dagger call --mod ./dagger/exponential-delay with-delays --delay=2 --num-iterations=10 --source=.`
5. Open trace in [dagger.cloud](https://v3.dagger.cloud)
6. Count number of times Delay.wait() was called (Shows ~45 as opposed to 10)

Given the synchronous delay per call (mapped to final duration log), it is fairly safe to assume that we are not _actually_ calling the module as many times as it claims.
This can be future confirmed by inspecting the span and noting the inputs/outputs/durations are exactly the same.