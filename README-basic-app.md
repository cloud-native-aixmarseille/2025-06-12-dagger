# a hello world in Golang

```shell
cd src
go test
go run ./main.go "from a simple app"
GOOS="darwin" GOARCH="arm64" go build -o ../builds/hello-darwin-arm64 ./main.go
../builds/hello-darwin-arm64 "from the compiled binary"
```
