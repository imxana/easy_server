package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func main() {
	path := getCurrentDirectory()
	fsh := http.FileServer(http.Dir(getParentDirectory(path) + "/static"))

	// use static dirctory
	http.Handle("/", http.StripPrefix("/", fsh))
	log.Printf("Listening on *:9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
