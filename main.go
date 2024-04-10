package main

import (
    "fmt"
    "log"
    "os"
)

func main() {

    newFileName := getFileName("Old file: ")
    oldFileName := getFileName("New file: ")
    
    newFileData := readFile(newFileName)
    oldFileData := readFile(oldFileName)

    fmt.Printf(string(newFileData))
    fmt.Printf(string(oldFileData))
}

func getFileName(msg string) string {
    var fileName string
    fmt.Printf(msg)
    fmt.Scan(&fileName)

    return fileName
}

func readFile(fileName string) []byte {
    fileData, err := os.ReadFile(fileName) 

    if err != nil {
        log.Fatal("file error: ", fileName)
    }
    
    return fileData
}
