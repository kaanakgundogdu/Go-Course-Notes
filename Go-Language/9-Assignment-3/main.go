package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
        fmt.Println("Missing parameter, provide file name!")
        return
    }

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)
}