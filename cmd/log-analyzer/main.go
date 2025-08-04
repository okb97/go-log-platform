package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/okb97/go-log-platform/internal/aggregate"
	"github.com/okb97/go-log-platform/internal/parse"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("サブコマンドを指定してください。")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "parse":
		parseCmd := flag.NewFlagSet("parse", flag.ExitOnError)
		inputPath := parseCmd.String("input", "", "入力ログファイルパス")
		outputDir := parseCmd.String("output", "", "出力ディレクトリ")
		parseCmd.Parse(os.Args[2:])
		if *inputPath == "" || *outputDir == "" {
			fmt.Println("inputとoutputを指定してください")
			os.Exit(1)
		}
		logs, err := parse.ParseLog(*inputPath)
		if err != nil {
			fmt.Println("ParseLogでエラー発生:", err)
			os.Exit(1)
		}
		if err := parse.SaveParseLog(logs, *outputDir); err != nil {
			fmt.Println("保存失敗：", err)
			os.Exit(1)
		}
	case "aggregate":
		aggregateCmd := flag.NewFlagSet("aggregate", flag.ExitOnError)
		date := aggregateCmd.String("date", "", "日時")
		level := aggregateCmd.String("level", "", "エラーレベル")
		format := aggregateCmd.String("format", "", "出力ファイルフォーマット")
		aggregateCmd.Parse(os.Args[2:])
		if *date == "" {
			fmt.Println("日時を指定してください")
			os.Exit(1)
		}
		if err := aggregate.AggregateLogs(*date, level, format); err != nil {
			fmt.Println("集計失敗：", err)
			os.Exit(1)
		}
	default:
		fmt.Println("対応していないサブコマンドです")
		os.Exit(1)
	}
}
