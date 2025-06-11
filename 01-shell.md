# dagger shell

## hello

```daggershell
container |
  from alpine |
  with-exec -- ash -c 'echo Hello meetup 👋 > ./hello.dat' |
  file ./hello.dat |
  contents
```

```daggershell
base=$(container | from alpine | with-exec apk add curl)
$base | with-exec -- curl -4 ifconfig.me | stdout
$base | terminal
```

## realworld

```shell
dagger < ./instructions.dagger
./builds/hello-darwin-arm64 youou
```
