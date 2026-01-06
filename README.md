# Gin Fleamarket

Go言語とGin Frameworkを使用したフリーマーケットアプリケーションのREST APIです。

## 技術スタック

- **言語**: Go 1.25.5
- **Webフレームワーク**: Gin
- **ORM**: GORM
- **データベース**: PostgreSQL 16
- **コンテナ**: Docker / Docker Compose
- **データベース管理**: pgAdmin4

## プロジェクト構成

```
gin-fleamarket/
├── controllers/       # HTTPリクエストハンドラー
├── dto/              # データ転送オブジェクト
├── infra/            # インフラストラクチャ層（DB接続など）
├── migrations/       # データベースマイグレーション
├── models/           # データモデル
├── repositories/     # データアクセス層
├── services/         # ビジネスロジック層
├── docker/           # Docker関連ファイル
├── main.go           # エントリーポイント
└── docker-compose.yaml
```

## 機能

- ✅ 商品一覧取得
- ✅ 商品詳細取得
- ✅ 商品作成
- ✅ 商品更新
- ✅ 商品削除
- ✅ 商品検索

## API エンドポイント

| メソッド | エンドポイント | 説明 |
|---------|--------------|------|
| GET     | /items       | 商品一覧を取得 |
| GET     | /items/:id   | 指定IDの商品を取得 |
| POST    | /items       | 新規商品を作成 |
| PUT     | /items/:id   | 指定IDの商品を更新 |
| DELETE  | /items/:id   | 指定IDの商品を削除 |

## セットアップ

### 前提条件

- Go 1.25.5以上
- Docker & Docker Compose

### 環境構築

1. リポジトリをクローン

```bash
git clone <repository-url>
cd gin-fleamarket
```

2. 環境変数ファイルを作成

```bash
cp .env.example .env
```

`.env`ファイルを編集して、必要な環境変数を設定してください。

3. Dockerコンテナを起動

```bash
docker-compose up -d
```

これにより以下のサービスが起動します：
- PostgreSQL: `localhost:5432`
- pgAdmin: `http://localhost:81`

4. 依存関係をインストール

```bash
go mod download
```

5. データベースマイグレーションを実行

```bash
go run migrations/migration.go
```

6. アプリケーションを起動

```bash
go run main.go
```

アプリケーションは `http://localhost:8080` で起動します。

## pgAdmin アクセス

- URL: `http://localhost:81`
- Email: `gin@example.com`
- Password: `ginpassword`

### PostgreSQL接続情報
- Host: `postgres`
- Port: `5432`
- Database: `fleamarket`
- Username: `ginuser`
- Password: `ginpassword`

## API使用例

### 商品一覧を取得

```bash
curl http://localhost:8080/items
```

### 商品を作成

```bash
curl -X POST http://localhost:8080/items \
  -H "Content-Type: application/json" \
  -d '{
    "name": "商品名",
    "price": 1000,
    "description": "商品の説明",
    "soldOut": false
  }'
```

### 商品を更新

```bash
curl -X PUT http://localhost:8080/items/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "更新後の商品名",
    "price": 1500,
    "description": "更新後の説明",
    "soldOut": true
  }'
```

### 商品を削除

```bash
curl -X DELETE http://localhost:8080/items/1
```

## 開発

### ホットリロード

開発時のホットリロードには [air](https://github.com/cosmtrek/air) の使用を推奨します。

```bash
go install github.com/cosmtrek/air@latest
air
```

### テストの実行

```bash
go test ./...
```

## データモデル

### Item (商品)

| フィールド | 型 | 説明 |
|-----------|-------|------|
| ID | uint | 商品ID（自動生成） |
| Name | string | 商品名（必須） |
| Price | uint | 価格（必須） |
| Description | string | 商品説明 |
| SoldOut | bool | 売り切れフラグ（デフォルト: false） |
| CreatedAt | time.Time | 作成日時 |
| UpdatedAt | time.Time | 更新日時 |
| DeletedAt | *time.Time | 削除日時（ソフトデリート） |

## ライセンス

MIT
