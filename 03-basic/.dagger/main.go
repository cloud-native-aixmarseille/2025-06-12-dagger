// A generated module for Meetup functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/meetup/internal/dagger"
)

type Meetup struct {
	Source *dagger.Directory
}

func New(
	// +defaultPath="src"
	source *dagger.Directory,
) *Meetup {
	return &Meetup{Source: source}
}

func (m *Meetup) Build(ctx context.Context) *dagger.File {
	ctr := m.BuildAppleSiliecon(ctx)
	return ctr.File("/src/hello")
}

// compile hello program for apple silicon
func (m *Meetup) BuildAppleSiliecon(ctx context.Context) *dagger.Container {
	return dag.Container().
		From("golang:1.21").
		WithDirectory("/src", m.Source).
		WithEnvVariable("GOOS", "darwin").
		WithEnvVariable("GOARCH", "arm64").
		WithWorkdir("/src").
		WithExec([]string{"go", "build", "-o", "hello"})
}
