package command

import (
	"os"
	"io"
	"io/ioutil"
	"regexp"
	"../repository"
)

const CONFIG_PATH string = "./chc.config"

func CopyHistoryDB() {
	_, err := os.Stat(repository.LATEST_DB_PATH)
	if err == nil {
		os.Remove(repository.LATEST_DB_PATH)
	}

	// open db file
	path, err := ioutil.ReadFile(CONFIG_PATH)
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
	dest, err := os.Create(repository.LATEST_DB_PATH)
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