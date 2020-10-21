package command

import (
	"os"
	"io"
	"io/ioutil"
	"regexp"
)

func CopyHistoryDB(copiedDBPath string, configPath string) {
	_, err := os.Stat(copiedDBPath)
	if err == nil {
		os.Remove(copiedDBPath)
	}

	// open db file
	path, err := ioutil.ReadFile(configPath)
	if err != nil {
        panic(err)
	}

	// open db file
	srcDBPath := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`).ReplaceAllString(string(path), "")
	src, err := os.Open(srcDBPath)
	if err != nil {
        panic(err)
    }
	defer src.Close()
	
	// create empty file
	dest, err := os.Create(copiedDBPath)
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