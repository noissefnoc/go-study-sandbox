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
    * ファイルを読んで `file` 

