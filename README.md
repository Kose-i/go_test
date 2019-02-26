# golang

## compile[実行ファイルができる]

go build ${SRC}

## compile and run[実行ファイルなし]

go run ${SRC}

# format

gofmt -w filename.go

- before
```
var numbers = map[string] int{
  "one":1,
  "two":2,
  "three":3,
}
```

- after

```
var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
}
```

# 参考

[はじめてのGo言語](https://www.kohgakusha.co.jp/books/detail/978-4-7775-1559-2)
[Go言語による並行処理](https://www.oreilly.co.jp/books/9784873118468/)

- Web
[github-golang](https://github.com/golang)
[go-git](https://go.googlesource.com/go)
[go公式](http://golang.org)
[go公式jp](http://golang.jp)
[go talks](https://talks.golang.org)
[go testcode](https://play.golang.org)
[golang-group](https://groups.google.com/forum/#!forum/golang-nuts)
