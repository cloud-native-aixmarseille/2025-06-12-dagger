# using a module

```shell
dagger init --sdk=go --name test
dagger install ../04-module
```

```golang
// .dagger/main.go
package main

import (
	"dagger/test/internal/dagger"
)

type Test struct {
	Source *dagger.Directory
}

// constructor set a default source path
func New(
	// +defaultPath="src"
	source *dagger.Directory,
) *Test {
	return &Test{Source: source}
}

// Run Go tests
func (m *Test) RunGoTests() *dagger.Container {
	return dag.MeetupGo(dagger.MeetupGoOpts{Source: m.Source}).GoTests()
}

// Build binary for Apple Silicon
func (m *Test) RunGoBuild() *dagger.File {
	return dag.MeetupGo(dagger.MeetupGoOpts{Source: m.Source}).BuildAll()
}
```

```daggershell
run-go-tests
run-go-build | export ./builds/hello-darwin-arm64
```

```shell
./builds/hello-darwin-arm64 'daggernauts 🥋 ceinture verte !'
```
