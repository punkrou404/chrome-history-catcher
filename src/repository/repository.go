package repository

import (
	"os"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

var LATEST_DB_PATH string = os.Getenv("HOME") + "/.chc/History"
const DELIMITER = "\t"

type TrashScanner struct{}

func (TrashScanner) Scan(interface{}) error {
    return nil
}

var DbConnection *sql.DB

type History struct {
    url string
	title string
	date string
}

func GetHistory() {
	DbConnection, _ := sql.Open("sqlite3", LATEST_DB_PATH)
	cmd := `
	select 
	urls.id,
	datetime(visits.visit_time/1000000-11644473600,'unixepoch','localtime') as visit_time, 
	urls.title, 
	urls.url 
	from visits 
	left join urls 
	on visits.url = urls.id 
	order by visits.id desc;
	`

	rows, _  := DbConnection.Query(cmd)
	if rows == nil {
		panic("Query failed.: RE-001")
	}
	
	// データ取得
	var history []History
	for rows.Next() {
		var h History
		err := rows.Scan(
			TrashScanner{},
			&h.date,
			&h.title, 
			&h.url,
		)
		if err != nil {
		fmt.Println(err)
		}
		history = append(history, h)
	}
	defer rows.Close()

	for _, h := range history {
		fmt.Println(h.title + DELIMITER + h.url + DELIMITER + h.date)
	}
}