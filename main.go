package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Printf("%v\n", "hola mundo genial!!!!")

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
}
