package main

import (
	"context"
	"dagger/no-foo-or-bar/internal/dagger"
)

type NoFooOrBar struct{}

// Removes foo+Bar from source directory and returns it.
// Interally asserts that foo+bar are removed for good measure
func (m *NoFooOrBar) RemoveFooAndBar(
	ctx context.Context,
	source *dagger.Directory,
) *dagger.Directory {

	return dag.NoBar().RemoveBar(
		dag.NoFoo().RemoveFoo(
			source,
		),
	)
}
