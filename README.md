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

- Web
[github-golang](https://github.com/golang)
[go公式](http://golang.org)
[go公式jp](http://golang.jp)
