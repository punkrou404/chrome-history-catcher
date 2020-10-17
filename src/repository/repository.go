package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

var DbConnection *sql.DB

type History struct {
    url string
	title string
	date string
}

func GetHistory() {
	DbConnection, _ := sql.Open("sqlite3", "./History")
    cmd := `SELECT
	title,
	url,
    DATETIME(last_visit_time / 1000000 + (strftime('%s', '1601-01-01') ), 'unixepoch', '+9 hours') AS date
FROM
	urls
GROUP BY
    title
ORDER BY
    date DESC
LIMIT 
    10000`

	rows, _  := DbConnection.Query(cmd)
	if rows == nil {
		panic("なんかqueryの結果nil")
	}
	
	// データ取得
	var history []History
	for rows.Next() {
		var h History
		err := rows.Scan(&h.title, &h.url, &h.date)
		if err != nil {
		fmt.Println(err)
		}
		history = append(history, h)
	}
	defer rows.Close()

	for _, h := range history {
		fmt.Println(fmt.Sprintf("%s\t%s\t%s",h.title,h.url,h.date))
	}
}