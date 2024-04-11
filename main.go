package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    var newFileName string
    var oldFileName string
    
    // Command line args 
    if len(os.Args) < 1 {
	newFileName = getFileName("Old file: ")
	oldFileName = getFileName("New file: ")
    } else {
	newFileName = os.Args[1]
	oldFileName = getFileName("New file: ")
    }
    
    newFileData := readFile(newFileName)
    oldFileData := readFile(oldFileName)

    fmt.Printf(string(newFileData))
    fmt.Printf(string(oldFileData))
}

// takes in a msg to prompt for input, returns a fileName as a string. 
func getFileName(msg string) string {
    var fileName string
    fmt.Printf(msg)
    fmt.Scan(&fileName)

    return fileName
}

// takes a file name and returns the file data as byte slice.
func readFile(fileName string) []byte {
    fileData, err := os.ReadFile(fileName) 

    if err != nil {
        log.Fatal("file error: ", fileName)
    }
    
    return fileData
}
