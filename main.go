package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

var directoryFlag string

func main() {
	flag.StringVar(&directoryFlag, "d", "", "Directory to look for files")
	flag.Parse()

	if directoryFlag == "" {
		log.Fatal("directoryFlag is mandatory")
	}

	files, err := ioutil.ReadDir(directoryFlag)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf("%v\n", file.Name())
	}
}
