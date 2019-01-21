# Golang の Validator

## [go-playground/validator](https://github.com/go-playground/validator)

対象バージョンは `v9` 。

必須で必要なのは構造体の定義側にJSONやYAMLを取り扱うときのように `validate` タグでアノテートする。

アノテーションは

* 標準タグ ([一覧](https://godoc.org/gopkg.in/go-playground/validator.v9) ※ `Baked In Validators and Tags` 以降)
* カスタムタグ (ユーザが作成。作成方法は後述)

の二つに別れる。

手順としては標準タグでバリデーションルールが満たせるか確認したのち

* 標準タグでバリデーションルールが ***満たせる***
    * バリデーション対象の構造体にタグを記載
    * `validate.Validator.Struct` などでバリデーション (してエラーチェック)
* 標準タグでバリデーションルールが ***満たせない***
    * `bool` を返却するバリデーション用の関数を作成
    * `validate.Validator.RegisterValdidation` にタグ名とバリデーションの関数を定義
    * バリデーション対象の構造体にタグを記載
    * `validate.Validator.Struct` などでバリデーション (してエラーチェック)

という流れ。

## Author

noissefnoc <noissefnoc@gmail.com>