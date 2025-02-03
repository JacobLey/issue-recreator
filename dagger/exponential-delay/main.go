package main

import (
	"context"
	"dagger/exponential-delay/internal/dagger"
	"errors"
	"strconv"
	"time"
)

type ExponentialDelay struct{}

// Removes foo from source directory and returns it.
// Interally asserts that foo is removed for good measure
func (m *ExponentialDelay) WithDelays(
	ctx context.Context,
	source *dagger.Directory,
	delay int,
	numIterations int,
) (*dagger.Directory, error) {

	start := time.Now()

	baseContainer := dag.Container().
		From("debian:12.9").
		WithEnvVariable("USER", "leyman").
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "upgrade", "-y"}).
		WithWorkdir("/workspace").
		WithDirectory("foo", source)

	fooDirectories := make([]*dagger.Directory, 0, numIterations+1)
	fooDirectories = append(fooDirectories, baseContainer.Directory("foo"))

	var waitContainer *dagger.Container

	for i := 0; i < numIterations; i++ {
		waitContainer = baseContainer
		for fooNum, fooDir := range fooDirectories {
			waitContainer = waitContainer.WithDirectory(
				"foo"+strconv.Itoa(fooNum),
				fooDir,
			)
		}
		waitContainer = dag.Delay().
			Wait(
				waitContainer,
				delay,
			).
			WithDirectory("foo"+strconv.Itoa(i+1), baseContainer.Directory("foo")).
			WithExec([]string{"bash", "-c", "echo Hello " + strconv.Itoa(i+1) + "th world! > ./foo" + strconv.Itoa(i+1) + "/dynamic-hello-world.txt"})

		fooDirectories = append(fooDirectories, waitContainer.Directory("foo"+strconv.Itoa(i+1)))
	}

	waitedContainer, err := waitContainer.Sync(ctx)

	if err != nil {
		return nil, err
	}

	end := time.Now()
	// Expected time to complete, plus a generous x2 multiplier (accounting for dagger overhead)
	maxDelayInSeconds := delay * numIterations * int(time.Second) * 2

	if int(end.Sub(start)) > maxDelayInSeconds {
		return nil, errors.New("TOOK TOO LONG: " + end.Sub(start).String())
	}

	return waitedContainer.
		WithExec([]string{"bash", "-c", "echo '" + end.Sub(start).String() + "' > ./duration.txt"}).
		Terminal().
		Directory("."), nil
}
