# basic

A basic example of a dagger CI can be found here :

- [CI](.dagger/main.go)

```shell
# dagger shell
go-tests | stdout
build-all | export builds/hello-darwin-arm64
./builds/hello-darwin-arm64 'daggernauts 🥋 ceinture blanche !'
```
