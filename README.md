# hiromi-go
Goの習作です...勘弁して下さい...

# ハマり

- `go get` でやると死ぬ

```
$  go get github.com/r-fujiwara/hiromi-go
can't load package: /usr/local/go/src/github.com/r-fujiwara/hiromi-go/hiromi.go:4:3: local import "./words" in non-local package
```

- ローカルで`go run`でやると動く

```
$ go run hiromi.go go
Achichi Achi Achi!

$ go run hiromi.go
Machaaki!
```
