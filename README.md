# go-dev-template

GO + PostgreSQLでの開発を、DockerとAirにより快適なものとするためのテンプレートリポジトリです。

Goで開発するAPIサーバー用のコンテナと、PostgreSQLによるDB用のコンテナをDocker Composeにより立ち上げます。また、Airによるホットリロードのおかげで、Goのソースコードに対する変更は、立ち上げているAPIサーバー用のコンテナに反映されます。

詳しい説明については、[こちらの記事]()を参照してください。

## 環境構築

1. [Docker](https://www.docker.com/ja-jp/get-started/)をインストール
2. 本リポジトリをクローンするか、本リポジトリをテンプレートとして選択してリポジトリを作成する
3. `docker network create api-db-network`
4. `docker network create db-pgadmin-network`
5. `docker compose build`

## 開発の流れ

下記コマンドで docker コンテナを起動する。これにより、開発中のAPIサーバー、PostgresのDBコンテナ、データベースをGUI操作できるPGAdmin、の3つのコンテナが立ち上がる

```
docker compose up
```

docker コンテナを閉じる

```
docker compose down
```

テンプレートでは、APIサーバは `localhost:8080` 、 PgAdminは `localhost:81` で立ち上がる

## 使用技術

テンプレートでは、以下のライブラリやフレームワークを使用していますが、開発者各自の判断で使いたいライブラリ等は追加してください。

- Go
- Gin
- Air
- PostgreSQL
- PgAdmin
- Docker
