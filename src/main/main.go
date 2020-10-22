package main

import (
	"./command"
	"./repository"
	"flag"
	"fmt"
	"os"
)

/**
 *chrome履歴を抽出するコマンド
 */
func main() {
	dbPath := os.Getenv("HOME") + "/.chc/History"
	config := os.Getenv("HOME") + "/.chc/chc.config"

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		// デフォルト
		command.CopyHistoryDB(dbPath, config)
		repository.GetHistory(dbPath)
	} else if args[0] == "pull" {
		command.CopyHistoryDB(dbPath, config)
	} else {
		fmt.Printf("引数はまだ実装中です。見逃して")
	}
}
