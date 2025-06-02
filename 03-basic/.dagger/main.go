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

// custom type i want to abstract
type Meetup struct {
	Source *dagger.Directory
}

// constructor set a default source path
func New(
	// +defaultPath="src"
	source *dagger.Directory,
) *Meetup {
	return &Meetup{Source: source}
}

// base
func (m *Meetup) Base(ctx context.Context) *dagger.Container {
	return dag.Container().
		From("golang:1.21").
		WithDirectory("/src", m.Source).
		WithWorkdir("/src")
}

// will run all Go tests
func (m *Meetup) GoTests(ctx context.Context) *dagger.Container {
	return m.Base(ctx).
		WithExec([]string{"go", "test"})
}

// build will centralize all the different builds
func (m *Meetup) BuildAll(ctx context.Context) *dagger.File {
	ctr := m.BuildAppleSiliecon(ctx)
	return ctr.File("/src/hello")
}

// compile the program for apple silicon
func (m *Meetup) BuildAppleSiliecon(ctx context.Context) *dagger.Container {
	return m.Base(ctx).
		WithEnvVariable("GOOS", "darwin").
		WithEnvVariable("GOARCH", "arm64").
		WithExec([]string{"go", "build", "-o", "hello"})
}
