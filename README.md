# RUNNING

to develop, you can make a local go.mod.

say you have `libmonpos` cloned as a submodule at `../libmonpos`.

first copy `go.mod` to `local.mod`:

```
$ cp go.mod local.mod
```

then you can  put in a replace for the module

```
go mod edit -modfile local.mod -replace github.com/almondheil/libmonpos=../libmonpos
```

then when you run or whatever, just specify that other modfile.

```
go run -modfile local.mod .
```

blammo. those files will be `.gitignore`-d for you too
