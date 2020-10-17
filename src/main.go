 
package main

import (
	"flag"
	"fmt"
	"./repository"
	"./command"
)

/*
chrome履歴を抽出するコマンド
*/
func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		// デフォルト
		command.CopyHistoryDB()
		repository.GetHistory()
	} else if args[0] == "pull" {
		command.CopyHistoryDB()
	} else {
		fmt.Printf("引数はまだ実装中です。見逃して")
	}
}