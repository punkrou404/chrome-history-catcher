package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type TrashScanner struct{}

func (TrashScanner) Scan(interface{}) error {
	return nil
}

var DbConnection *sql.DB

type History struct {
	url   string
	title string
	date  string
}

func GetHistory(DBPath string) error {
	DbConnection, _ := sql.Open("sqlite3", DBPath)
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

	rows, _ := DbConnection.Query(cmd)
	if rows == nil {
		return fmt.Errorf("err %s", "Query failed.: RE-001")
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
			return err
		}
		history = append(history, h)
	}
	defer rows.Close()

	const DELIMITER string = "\t"
	for _, h := range history {
		fmt.Println(h.title + DELIMITER + h.url + DELIMITER + h.date)
	}

	return nil
}
