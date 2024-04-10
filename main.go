package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
    var fileName string
    fmt.Printf("File Name: ")
    fmt.Scan(&fileName)

    fileData := readFile(fileName)
}

func readFile(fileName string) []byte {
    fileData, err := os.ReadFile(fileName) 

    if err != nil {
        log.Fatal("file error: ", fileName)
    }
    
    return fileData
}
