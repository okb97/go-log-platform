package parse

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/okb97/go-log-platform/internal/model"
)

func ParseLog(inputPath string, outputDir string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("ファイルオープン失敗: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("CSV読み込み失敗: %w", err)
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
	return nil
}
