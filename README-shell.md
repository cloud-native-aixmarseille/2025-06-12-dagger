# dagger shell

## hello

```daggershell
container |
  from alpine |
  with-exec -- ash -c 'echo Hello meetup 👋 > ./hello.dat' |
  file ./hello.dat |
  contents
```

## realworld

```shell
dagger < ./instructions.dagger
./builds/hello-darwin-arm64 youyou
```
