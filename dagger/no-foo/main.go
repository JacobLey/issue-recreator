package main

import (
	"context"
	"dagger/no-foo/internal/dagger"
	"errors"
)

type NoFoo struct{}

// Removes foo from source directory and returns it.
// Interally asserts that foo is removed for good measure
func (m *NoFoo) RemoveFoo(
	ctx context.Context,
	// +ignore=["foo"]
	source *dagger.Directory,
) (*dagger.Directory, error) {

	// Expected to fail because foo shouldn't exist
	_, err := source.Directory("foo").Digest(ctx)

	if err != nil {
		return source, nil
	}

	return nil, errors.New("foo still exists")
}
