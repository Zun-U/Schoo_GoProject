Go言語学習
==========

>### 進捗状況

| 授業    | 学んだこと|
|--------|-----------|
|Lesson1 |変数、定数、制御構造、配列|
|Lesson2 |関数とメソッド|
|Lesson3 |並行処理|
|Lesson4 |GoのインストールとHTML|
|Lesson5 |SQL|
|Lesson6 |テスト駆動開発|
|Lesson7 |テーブル駆動テスト|
|Lesson8 |テスト用ヘルパー関数、DB接続テスト|
|Lesson9 |select、create、delete|
|Lesson10|`template、embed、デプロイ...学習中`|

--------------------


>### 環境

ローカル環境汚染を最低限にするために、Dockerにて開発環境を構築しました。<br>
Go言語の実行環境や、VSCodeによる補完機能やフォーマッターなどが準備されています。<br>
git cloneを行うと、そのまま使えるようになっています。
- Go 1.21.4
- MySQL 8.0
- VScodeの拡張機能
  - Go
    - goimports
    - hey
    - staticcheck
    - golangci-lint
    - gopls
  - trailing-spaces

> [!NOTE]
> 実行環境はWSL(Ubuntu)で、VSCodeの利用を想定しています。

-------------

>### 環境構築手順
- Docker Desktopを起動します。
- `git clone` したソースコード上でVSCodeを起動します。
- VSCodeのコマンド `Dev Containers:Reopen in Container` でDockerを起動します。

<br>

> [!NOTE]
> VSCodeの拡張機能の `Remote Development`を事前にインストールする必要があります。

<br>


-------------

>### build実行
権限がない場合、WSL内(ubuntu)内でビルド作成するとエラーが発生します。
- `ローカルホストのオーナーを変更`してからビルドを行います。

<br>

buildしたファイルの実行は以下の通りです。
- `./ビルドしたファイル名` または `./ディレクトリ名/ビルドしたファイル名`