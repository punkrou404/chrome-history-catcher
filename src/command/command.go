package command

import (
	"os"
	"io"
	"io/ioutil"
	"regexp"
)

func CopyHistoryDB() {
	_, err := os.Stat("./History")
	if err == nil {
		os.Remove("./History")
	}

	// open db file
	path, err := ioutil.ReadFile("./chc.config")
	if err != nil {
        panic(err)
	}

	// open db file
	dbPath := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`).ReplaceAllString(string(path), "")
	src, err := os.Open(dbPath)
	if err != nil {
        panic(err)
    }
	defer src.Close()
	
	// create empty file
	dest, err := os.Create("./History")
	if err != nil {
        panic(dest)
    }
	defer dest.Close()
	
	// copy db
	_, err = io.Copy(dest, src)
    if  err != nil {
        panic(err)
	}
}