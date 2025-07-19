package parse

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/okb97/go-log-platform/internal/model"
)

func ParseLog(inputPath string) ([]model.LogEntry, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("ファイルオープン失敗: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("CSV読み込み失敗: %w", err)
	}

	layout := "2006-01-02 15:04:05"
	var logs []model.LogEntry
	for i, row := range records {
		if i == 0 {
			continue
		}
		if len(row) < 3 {
			continue
		}
		timestamp, err := time.Parse(layout, row[0])
		if err != nil {
			fmt.Printf("row %d: time parse error: %v\n", i, err)
			continue
		}
		logs = append(logs, model.LogEntry{
			TimeStamp: timestamp,
			Level:     row[1],
			Message:   row[2],
		})
	}
	//fmt.Printf("%+v\n", logs)
	return logs, nil
}

func SaveParseLog(logs []model.LogEntry, outputDir string) error {
	dateDir := time.Now().Format("2006-01-02")
	fullDir := filepath.Join(outputDir, dateDir)

	if err := os.MkdirAll(fullDir, 0755); err != nil {
		return fmt.Errorf("出力ディレクトリ作成失敗; %w", err)
	}
	today := time.Now().Format("2006-01-02_15:04:05")
	outputPath := filepath.Join(fullDir, fmt.Sprintf("parsed_logs_%s.json", today))
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("JSONファイル作成:%w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(logs); err != nil {
		return fmt.Errorf("JSONエンコード失敗: %w", err)
	}
	fmt.Printf("ログを%sに保存しました\n", outputPath)
	return nil
}
