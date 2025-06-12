package main

import (
	"context"
	"dagger/meetup-go/internal/dagger"
)

// custom type i want to abstract
type MeetupGo struct {
	Source *dagger.Directory
}

// constructor set a default source path
func New(
	// +defaultPath="src"
	source *dagger.Directory,
) *MeetupGo {
	return &MeetupGo{Source: source}
}

// run go test
func (m *MeetupGo) GoTests(ctx context.Context) (string, error) {
	return dag.Go().
		WithSource(m.Source).
		Exec([]string{"go", "test"}).
		Stdout(ctx)
}

// build linux - amd64 binary
func (m *MeetupGo) GoBuildLinuxAmd64(ctx context.Context) *dagger.File {
	return dag.Go().
		WithSource(m.Source).
		WithCgoDisabled().
		WithEnvVariable("GOOS", "linux").
		WithEnvVariable("GOARCH", "amd64").
		Build()
}

// use the binary
func (m *MeetupGo) End2End(ctx context.Context) (string, error) {
	binary := m.GoBuildLinuxAmd64(ctx)
	return dag.Container().
		From("alpine:latest").
		WithFile("/app/hello", binary).
		WithEntrypoint([]string{"/app/hello"}).
		WithDefaultArgs([]string{"world"}).
		WithExec([]string{}).Stdout(ctx)
}
