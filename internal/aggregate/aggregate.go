package aggregate

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/okb97/go-log-platform/internal/model"
)

func AggregateLogs(dateTime string, level, format *string) error {
	date, err := time.Parse("2006-01-02", dateTime)
	if err != nil {
		return fmt.Errorf("dateTimeの形式が不正です: %w", err)
	}

	folderPath := filepath.Join("parsed", date.Format("2006-01-02"))
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return fmt.Errorf("対象フォルダ読み込み失敗: %w", err)
	}

	var allLogs []model.LogEntry

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
			continue
		}
		filePath := filepath.Join(folderPath, file.Name())
		f, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("ファイルオープン失敗: %w", err)
		}
		defer f.Close()

		var logs []model.LogEntry
		if err := json.NewDecoder(f).Decode(&logs); err != nil {
			return fmt.Errorf("JSONデコード失敗 (%s): %w", file.Name(), err)
		}
		allLogs = append(allLogs, logs...)
	}

	counts := map[string]int{}
	for _, log := range allLogs {
		logLevel := strings.ToLower(log.Level)
		if level != nil && *level != "" {
			if logLevel == strings.ToLower(*level) {
				counts[log.Level]++
			}
		} else {
			counts[log.Level]++
		}
	}

	outputFormat := "json"
	if format != nil && *format != "" {
		outputFormat = strings.ToLower(*format)
	}

	switch outputFormat {
	case "json":
		out, err := json.MarshalIndent(counts, "", "  ")
		if err != nil {
			return fmt.Errorf("JSONマーシャル失敗: %w", err)
		}
		fmt.Println(string(out))

	case "csv":
		writer := csv.NewWriter(os.Stdout)
		defer writer.Flush()

		if err := writer.Write([]string{"Level", "Count"}); err != nil {
			return fmt.Errorf("CSV書き込み失敗: %w", err)
		}
		for lvl, cnt := range counts {
			if err := writer.Write([]string{lvl, fmt.Sprintf("%d", cnt)}); err != nil {
				return fmt.Errorf("CSV書き込み失敗: %w", err)
			}
		}
	default:
		return errors.New("対応していないフォーマットです。json または csv を指定してください。")
	}

	return nil
}
