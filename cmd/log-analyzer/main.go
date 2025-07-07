package main

import (
	"flag"
	"fmt"
	"os"

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
		parse.ParseLog(*inputPath, *outputDir)
	default:
		fmt.Println("対応していないサブコマンドです")
		os.Exit(1)
	}

}
