# a hello world in Golang

```shell
cd src
go tests
go run ./main.go "from a simple app"
GOOS="darwin" GOARCH="arm64" go build ./main.go -o ../builds/hello-darwin-arm64
../builds/hello-darwin-arm64 "from the compiled binary"
```
