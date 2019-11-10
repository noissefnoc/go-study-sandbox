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
            * `ind` タグの場合 (インデント？)
                * `left` attribute を 360 で割った数だけ半角スペースを入れる
            * `pStyle` タグの場合
                * `val` attribute の値による判定
                    * `Heading` から始まる場合、`Heading` を取り除いた次の数だけ `#` をつける
                    * `Code` の場合、 `code` の値を `true` にする
                    * そのほかの場合、数値に変換できたら、その数だけ `#` をつける
            * `numPr` タグの場合
                * `numId` タグの場合、 `numId` 変数に `val` の値を入れる
                * `ilvl` タグの場合、 `ilvl` 変数に `val` の値を入れる
                * インデントの深さを計算する
                * インデントをつける
                * `numFmt` の値によって、数値見出しか、箇条書き見出しかを決める
            * `code` の値が `true` ならバッククォートで結果をくくる
            * 再起的にノード読み込み
        * `tbl` (テーブルタグ)
            * 文字列の二次元配列 `rows` に `tr` と `tc` を詰めて行って、最後 `rows` を出力
        * `r` (文字飾り)
            * イタリック、ボールド、打ち消し線の要素をそれぞれ探して、前後につける
        * `p` (段落)
            * 読み進めてるだけに読めるが、これ `default` と混ぜてない理由はなんなのだろうか
            * (追記)内容書き出ししてた
        * `blip` (画像)
            * `embed` オプションの有無で分岐
                * `relationship.xml` にあれば `file.extract` を呼ぶ
        * `Fallback` と `txbxContent` (コードブロック)
            * 中のテキストを再起的に読み取って周りを三点バッククォートで囲む
        * `default` それ以外
            * 再起的にXMLを読み進める

#### ユーティリティー関数

* `Node.UnmarchalXML`
    * 明示的な呼び出しがないので、これ何かの `interface` なんだろうな
    * 引数の `xml.StartElement` を `Node.Attrs` にセットして
    * `xml.Decoder` でデコードした結果を `Node` にセットしてる
* `escape`
    * 第一引数( `s` )に第二引数( `set` )が入っていたら、バックスラッシュを前につける
* `file.extract`
    * リファレンスが書いてある XML からパスを取り出してファイルオープン
    * `embed` オプションの有無で分岐
        * `embed` あり：画像ヘッダをつけて base64 で埋め込む
        * `embed` なし：画像を書き出し
 * `attr`
    * XML の attribute の中から、該当する文字列があるか検索し、あれば値と `true` なければ空文字と `false` を返す

