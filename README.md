## build and run
docker-compose up -d


## jwt key generate
ssh-keygen -t rsa -f revel
ssh-keygen -f revel.pub -e -m pkcs8 > revel.pub
https://qiita.com/AkiTakeU/items/e2133eeb94f57629b5e7

## 使うツールなど

- goose
  - https://github.com/pressly/goose
  - マイグレーションを実行する
- xo
  - https://github.com/xo/xo
  - データーベースからモデル定義ファイルを出力する
- gorm
  - https://github.com/jinzhu/gorm
  - シーディングを実行するために利用するORM
- sqlacodegen
  - https://pypi.org/project/sqlacodegen/
  - sqlalchemyで使うモデル定義を出力（未検証）

インストール方法や詳しい使い方は各ツールのドキュメント参照

## マイグレーション

マイグレーション（テーブルの作成やカラムの追加・変更）を行う場合は以下の手順に従うこと

※マイグレーションの作業は `migrations` ディレクトリで行うこと

### 0. `migrations` ディレクトリに移動

```sh
$ cd migrations
```

### 1. マイグレーションファイルを生成

```sh
$ goose create describe_your_change sql
```

`20181107165534_describe_your_change.sql` というファイルが生成される

※ファイル名には自動で実行時の日時が追加されるのでその時々で変わる

### 2. 変更内容をマイグレーションファイルに記述

```sql
-- +goose Up
CREATE TABLE post (
    id int NOT NULL,
    title text,
    body text,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE post;
```

`goos up` 以下に変更内容、 `goose down` 以下に変更内容をもとに戻すSQLを記述する

※1つのマイグレーションファイルにup, downの両方を記述する

### 3. マイグレーションを適用

```sh
$ goose mysql 'YOUR DB INFO' up
2018/11/07 16:56:49 OK    20181107165534_describe_your_change.sql
2018/11/07 16:56:49 goose: no migrations to run. current version: 20181107165534
```

自動で最新の状態までマイグレーションが実行されるので特定のバージョンにする場合は `up-to` コマンドを使う

### 4. マイグレーションを取り消す（ロールバック）

```sh
$ goose mysql 'YOUR_DB_INFO' down
2018/11/07 16:56:49 OK    20181107165534_describe_your_change.sql
```

デフォルトで一つ前の状態に戻るので特定のバージョンにする場合は `down-to` コマンドを使う

※変更する際はロールバックまで正常にできるか確認すること

## モデル定義ファイル出力

### Go

xoを使いMySQLの情報からモデル定義ファイルを出力する

以下のコマンドで `models' 以下に各テーブルに対応するモデル定義ファイルが出力される

```sh
$ xo --escape-all 'YOUR DB INFO' -o models
```

`users` テーブルに対応するモデルは `user.xo.go` のように出力される

※ `goosedbversion.xo.go` というファイルは削除すること（取り込み時にエラーになるため）

### Python

sqlacodegenを使いMySQLの情報からモデル定義ファイルを出力する

以下のコマンドで `models` 以下に `models.py` として出力

```sh
$ sqlacodegen 'YOUR_DB_INFO' > models/models.py
```

## シーディング

`seeder.go` を実行することによりローカル環境での確認やテスト時に使うデータの生成を行う

```sh
$ go run seeder.go
```

※この辺はやり方検討しつつ更新

# 以下古い情報なので後で消す

## migration tool

Install [migrate](https://github.com/golang-migrate/migrate) (see [installation](https://github.com/golang-migrate/migrate/tree/master/cli#installation)).

```
# check your version
$ migrate --version
```

## create migration file

Create up.sql, and also down.sql to rollback.

### create new table

Create `version_date_create_new_table.up.sql`. (version: increament latest version number)

```sql
CREATE TABLE `new_table` (
  ...
) ...;
```

Also create `version_date_create_new_table.down.sql`. (version: same as up file)

```sql
DROP TABLE `new_table`;
```

### change table

Create `version_date_update_existing_table.up.sql`. (version: increment latest version number)

```sql
ALTER TABLE `existing_table` ADD `new_column` ...;
```

Create `version_date_update_existing_table.down.sql`. (version: same as up file)

```sql
ALTER TABLE `existing_table` DROP `new_column`;
```

## migration (up)

```
$ migrate -database $DATABASE_URL -path ./path_directory up
```
ex)
```
$ migrate -database 'mysql://user:pass@tcp(127.0.0.1:3306)/database_name' -path ./ up
1/u 20181029_create_table_source_sites (173.220907ms)
```

## migration (down)

```
$ migrate -database $DATABASE_URL -path ./path_directory down
```
ex)
```
$ migrate -database 'mysql://user:pass@tcp(127.0.0.1:3306)/database_name' -path ./ down
1/d 20181029_create_table_source_sites (100.842044ms)
```

## migration (version check)

```
$ migrate -datebase $DATABASE_URL -path ./path_directory version
```
ex)
```
$ migrate -database 'mysql://user:pass@tcp(127.0.0.1:3306)/database_name' -path ./ version
1
```
ex) case of failure in the previous migration
```
$ migrate -database 'mysql://user:pass@tcp(127.0.0.1:3306)/database_name' -path ./ up
1/u 20181029_create_table_source_sites (173.220907ms)
error: migration failed in line 0:  (details: Error 1065: Query was empty)
$ migrate -database 'mysql://user:pass@tcp(127.0.0.1:3306)/database_name' -path ./ version
2 (dirty)
```

## migration (force)

```
$ migrate -datebase $DATABASE_URL -path ./path_directory force N
```
N: Set migration vertion

ex)
```
$ migrate -database 'mysql://user:pass@tcp(127.0.0.1:3306)/database_name' -path ./ force 1
$ migrate -database 'mysql://user:pass@tcp(127.0.0.1:3306)/database_name' -path ./ version
1
```

ex) Use case
```
$ migrate -database 'mysql://user:pass@tcp(127.0.0.1:3306)/database_name' -path ./ up
5/u 20181029_create_table_crawled_company_site_urls (424.952155ms)
error: migration failed in line 0: (~~ error contents ~~)
$ migrate -database 'mysql://user:pass@tcp(127.0.0.1:3306)/database_name' -path ./ up
error: Dirty database version 6. Fix and force version.
$ migrate -database 'mysql://user:pass@tcp(127.0.0.1:3306)/database_name' -path ./ force 5
$ migrate -database 'mysql://user:pass@tcp(127.0.0.1:3306)/database_name' -path ./ up
6/u 20181029_create_table_contents (79.379645ms)
```