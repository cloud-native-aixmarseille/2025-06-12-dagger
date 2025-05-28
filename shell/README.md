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

