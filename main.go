package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	var directoryFlag string

	flag.StringVar(&directoryFlag, "d", "", "Directory to look for files")
	flag.Parse()

	if directoryFlag == "" {
		log.Fatal("directoryFlag is mandatory")
	}

	files, err := ioutil.ReadDir(directoryFlag)
	if err != nil {
		log.Fatal(err)
	}

	epubFiles := getFilesWithExtension(files, "epub")

	for _, file := range epubFiles {
		fmt.Printf("%v\n", file.Name())
	}
}

func getFilesWithExtension(files []fs.FileInfo, ext string) []fs.FileInfo {
	var filtered []fs.FileInfo

	for _, file := range files {
		if strings.HasSuffix(file.Name(), "."+ext) {
			filtered = append(filtered, file)
		}
	}

	return filtered
}
