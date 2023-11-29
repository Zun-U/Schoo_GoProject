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
|Lesson7 |`学習途中`|
|Lesson8 |`未学習`|
|Lesson9 |`未学習`|
|Lesson10|`未学習`|

--------------------


>### 環境構築

ローカル環境汚染を最低限にするために、Dockerにて開発環境を構築しました。<br>
Go言語の実行環境や、VSCodeによる補完機能やフォーマッターなどが準備されています。<br>
git cloneを行うと、そのまま使えるようになっています。
- Go 1.21.4
- VScodeの拡張機能
  - Go
    - goimports
    - hey
    - staticcheck
    - golangci-lint
    - gopls
  - trailing-spaces

> [!NOTE]
> 実行環境はLinux(Ubuntu)で、VSCodeの利用を想定しています。

-------------

>### 構築手順
- Docker Desktopを起動します。
- `git clone` したソースコード上でVSCodeを起動します。
- VSCodeのコマンド `Dev Containers:Reopen in Container` でDockerを起動します。

<br>

> [!NOTE]
> VSCodeの拡張機能の `Remote Development`を事前にインストールする必要があります。

<br>
