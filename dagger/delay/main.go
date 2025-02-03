package main

import (
	"context"
	"dagger/delay/internal/dagger"
	"strconv"
)

type Delay struct{}

// Waits N seconds and then returns container
// with wait command in history
func (m *Delay) Wait(
	ctx context.Context,
	container *dagger.Container,
	delay int,
) (*dagger.Container, error) {

	waiting := container.WithExec([]string{"sleep", strconv.Itoa(delay)})

	return waiting.Sync(ctx)
}
