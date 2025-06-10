# using daggerverse

```shell
dagger init --sdk=go --name meetup
dagger install github.com/sagikazarmark/daggerverse/go@v0.9.0
```

```golang
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
func (m *Meetup) GoBuildDarwinArm64(ctx context.Context) *dagger.File {
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
go-build-darwin-arm-64 | export builds/hello
```
