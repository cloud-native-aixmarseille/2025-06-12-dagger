# using daggerverse

```shell
dagger init --sdk=go --name meetup-go
dagger install github.com/sagikazarmark/daggerverse/go@v0.9.0
code .dagger/main.go
```

```golang
// .dagger/main.go
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

// base
func (m *MeetupGo) GoBuildDarwinArm64(ctx context.Context) *dagger.File {
	return dag.Go().
		WithSource(m.Source).
		WithCgoDisabled().
		WithEnvVariable("GOOS", "darwin").
		WithEnvVariable("GOARCH", "arm64").
		Build()
}
```

dagger shell code :

```daggershell
go-build-darwin-arm-64 | export ./builds/hello-darwin-arm64
```

```shell
./builds/hello-darwin-arm64 'daggernauts 🥋 ceinture noire !'
```
