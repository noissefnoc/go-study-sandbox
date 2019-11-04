# modules

Golang の `go mod` の話。

これを書いた時点で Golang のバージョンは `1.13` で

* `GO111MODULES` はディフォルト設定 `auto` 
* ただし、 `GOPATH` 配下であっても `go.mod` があれば module-aware mode になる

という状態。

以前はセマンティックバージョニングしてないやつに厳しくて、 `dep` を使い続ける話を聞いたけど、今はブランチ/コミットハッシュも指定できるので特に問題ないはず。


## Golang のバージョンを go.mod に反映する方法

```
$ go mod edit -go=<GOVERSION>
```

## モジュールの依存を追加する方法

モジュールモードになっている時なら `go get` するだけで `go.mod` が書き換わる

```
$ go get path/to/repository             # これは path/to/repository@latest のエイリアス
$ go get path/to/repository@v0.0.1      # バージョン指定
$ go get path/to/repository@branch      # ブランチ指定
$ go get path/to/repository@commit-hash # コミット指定
```

その後、 `go build` / `go test` などのタイミングでレポジトリを取りに行くが、IDEだとビルド/テストするまえに取ってきてほしいのだが、設定あるのだろうか？

## 依存関係の正しい整理

以下で実施。長いこと使ってアップデート何回もしているようなものは仕掛けておいたほうがいいのかも。

```
$ go mod tidy
```

## 依存関係一覧の表示

依存関係全部表示するコマンドは以下

```
$ go list -m all

```
