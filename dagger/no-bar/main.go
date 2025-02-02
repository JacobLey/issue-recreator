package main

import (
	"context"
	"dagger/no-bar/internal/dagger"
	"errors"
)

type NoBar struct{}

// Removes bar from source directory and returns it.
// Interally asserts that bar is removed for good measure
func (m *NoBar) RemoveBar(
	ctx context.Context,
	// +ignore=["bar"]
	source *dagger.Directory,
) (*dagger.Directory, error) {

	// Expected to fail because foo shouldn't exist
	_, err := source.Directory("bar").Digest(ctx)

	if err != nil {
		return source, nil
	}

	return nil, errors.New("bar still exists")
}