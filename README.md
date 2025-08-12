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
GET    /api/tasks/:id      - タスク取得
PUT    /api/tasks/:id      - タスク更新
DELETE /api/tasks/:id      - タスク削除
```

**主要機能**:
- GORMを利用したSQLiteでのデータ永続化
- RESTful API設計
- JSON形式でのデータ交換
- MVCを意識したレイヤードアーキテクチャ（handler, service, repository）
- go-ginを利用したルーティング

**習得技術**:
- Webサーバー構築（gin）
- JSON エンコード/デコード
- ORM（GORM）
- SQLiteの基本操作
- ルーティング設計
- テスト（ユニットテスト）
  

## Week 4: 並列Webクローラー開発
### 作成アプリ: log-collector（分散ログ収集システム）

**概要**

本プロジェクトは、複数サーバーから並列でログを収集し、構造化・正規化する **分散ログ収集システム** です。  
Go の並列処理（goroutine / channel）やファイル監視（fsnotify）、HTTP クライアントなどの技術を用いて効率的にログを取得します。

1. 複数サーバーからの並列収集
goroutine を利用して、複数サーバーのログ取得を同時に実行します。  
これにより、順次処理より大幅に処理時間を短縮できます。

2. channel によるデータ受け渡し
goroutine 間で結果やエラーを安全に受け渡すために channel を使用します。

3. sync.WaitGroup による同期制御
全ての goroutine の処理完了を待ってから次の処理へ進めます。

4. ワーカープールパターン
同時並列数を制限して、サーバーやクライアントの負荷を軽減します。

5. HTTPクライアント実装
net/http を使い、タイムアウト付きでログを取得します。

6. ファイル監視（fsnotify）
ローカルディレクトリを監視し、ログファイルが更新・追加されたら自動処理します。

7. レート制限
アクセス間隔を調整して、安定した稼働を実現します。

**主な機能**

- **複数サーバーからの並列ログ収集**
- **HTTP/HTTPS経由でのファイルダウンロード**
- **リアルタイムログ監視（fsnotify）**
- **収集データの構造化・正規化**
- **ワーカープールによる負荷制御**
- **レート制限による安定稼働**

**習得技術**:
* goroutine による並列処理
* channel を使った安全な通信
* sync.WaitGroup による同期制御
* ワーカープールパターン
* HTTP クライアント実装
* ファイル監視（fsnotify）
* レート制限・負荷制御