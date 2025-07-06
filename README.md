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