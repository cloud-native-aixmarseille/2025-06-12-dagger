# dagger module

```shell
dagger init --sdk=go --name meetup-go
dagger develop --sdk=go
```

```golang
// main.go
package main

import (
	"context"
	"dagger/meetup-go/internal/dagger"
)

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

// base
func (m *MeetupGo) Base(ctx context.Context) *dagger.Container {
	return dag.Container().
		From("golang:1.21").
		WithDirectory("/src", m.Source).
		WithWorkdir("/src")
}

// will run all Go tests
func (m *MeetupGo) GoTests(ctx context.Context) *dagger.Container {
	return m.Base(ctx).
		WithExec([]string{"go", "test"})
}

// build will centralize all the different builds
func (m *MeetupGo) BuildAll(ctx context.Context) *dagger.File {
	ctr := m.BuildAppleSiliecon(ctx)
	return ctr.File("/src/hello")
}

// compile the program for apple silicon
func (m *MeetupGo) BuildAppleSiliecon(ctx context.Context) *dagger.Container {
	return m.Base(ctx).
		WithEnvVariable("GOOS", "darwin").
		WithEnvVariable("GOARCH", "arm64").
		WithExec([]string{"go", "build", "-o", "hello"})
}
```
