# XML

* `encoding/xml` の使い方を確認する
* [mattn/docx2md](https://github.com/mattn/docx2md) の流れを把握する


### mattn/docx2md の把握

コミットハッシュは [82ba1ae](https://github.com/mattn/docx2md/commit/82ba1ae25a7eb42f8e302546f705b28f3031215f) で確認


#### main関数からのコールグラフ

1. `main` 関数
    * コマンドライン引数 `embed (boolean)` をパース
    * 引数ない場合、利用方法表示して異常ステータスで終了
    * `docx2md` 関数を呼ぶ。引数はファイルパスとフラグ。エラーが出たらログ出して終了
2. `docx2md` 関数
    * 引数で受け取ったファイルパスで `zip.OpenReader` でファイルを開く (エラー処理)
    * zipファイル内のファイルをループ
        * ファイル名によって分岐
            * `word/_rels/document.xml.rels` のパス
                * ファイルを開いて、XMLパースして構造体に割り当て
            * `word/numbering.xml` のパス
                * ファイルを開いて、XMLパースして構造体に割り当て
    * `word/document.xml` があるかを調べる
    * ファイルを読んで `file` 構造体を作り、 `walk` を読んで、木構造ブロックごとにNodeを作成して随時読み込む
    * バッファに貯めたmarkdown変換したテキストを出力
3. `walk` 関数
    * `node.XMLName.Local` (多分タグ名)で分岐をしているが、 L119の `Node` struct の定義だと `Local` が何を表現しているのか分からない。ここ、 `encoding/xml` を読まないとダメか
        * `hyperlink` タグの場合 (リンク記法)
            * markdown の `[title](link_url)` になるように、角括弧始まりを出力
            * タイトル部分のテキストを `walk` を再起的に取得
            * タイトルに使われている角括弧をエスケープして閉じカッコをつける
            * リンクの始まり丸括弧を出力
            * `hyperlink` タグの `id` 属性の値を、リンクがまとまっている別XML ( `Relationship`) から参照して出力
            * 丸括弧をエスケープ
            * 閉じカッコをつける
        * `t` タグの場合 (中のテキスト)
            * テキスト値を取得
        * `pPr` タグの場合
            * いったんパス
        * `tbl` (テーブルタグ)
            * 文字列の二次元配列 `rows` に `tr` と `tc` を詰めて行って、最後 `rows` を出力
        * `r` (文字飾り)
            * イタリック、ボールド、打ち消し線の要素をそれぞれ探して、前後につける
        * `p` (段落)
            * 読み進めてるだけに読めるが、これ `default` と混ぜてない理由はなんなのだろうか
            * (追記)内容書き出ししてた
       

