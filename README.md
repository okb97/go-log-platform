## Week 1: CLIバッチ処理ツール開発
### 作成アプリ: log-analyzer（ログ解析CLIツール）

**主要機能**:
- CSVログファイルの読み込み・解析
- 日別/時間別/レベル別でのデータ集計
- JSON/CSV形式でのレポート出力
- エラーログの抽出・カウント機能

**コマンド例**:
```
./log-analyzer parse --input logs/access.log --output reports/
./log-analyzer aggregate --date 2025-01-01 --format json
./log-analyzer filter --level error --count
```

**習得技術**:
- Go基礎文法（構造体、スライス、マップ）
- ファイル操作（os、bufio パッケージ）
- CSV処理（encoding/csv）
- コマンドライン引数処理（flag パッケージ）
- エラーハンドリング


## Week 2-3: REST API サーバー開発
### 作成アプリ: task-api（タスク管理Web API）

**API エンドポイント**:
```
POST   /api/tasks          - タスク作成
GET    /api/tasks          - タスク一覧取得
PUT    /api/tasks/:id      - タスク更新
DELETE /api/tasks/:id      - タスク削除
GET    /api/reports/daily  - 日次レポート取得
```

**主要機能**:
- SQLiteでのデータ永続化
- RESTful API設計
- JSON形式でのデータ交換
- 基本的な認証機能（APIキー）
- 入力値バリデーション

**習得技術**:
- HTTP サーバー構築（net/http）
- JSON エンコード/デコード
- データベース操作（database/sql）
- SQLiteの基本操作
- ルーティング設計
- 認証・認可の基礎