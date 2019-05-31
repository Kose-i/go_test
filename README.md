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

[Go言語とは？特徴・できること・人気の理由について徹底解説](https://tech-camp.in/note/technology/49655/)

[Go 言語における「オブジェクト」](https://text.baldanders.info/golang/object-oriented-programming/)


- Qiita

[Go言語関連書籍のまとめ](https://qiita.com/yoskeoka/items/d07b60f755e8a9b30ccf)

[クラスとオブジェクトの関係性（Go 言語編）](https://qiita.com/spiegel-im-spiegel/items/2da5e5902aa2cb6d9e30)
