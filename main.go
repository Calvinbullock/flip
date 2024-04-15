package main

import (
    "fmt"
    "log"
    "os"
    "strings"
)

/* ======================================== \\
||		Storage struc		    ||
\\ ======================================== */

type Commit struct {
    fileName string
    fileData string
    commitMsg string
}

// takes user input and parses it into a commit instance 
func makeCommit() Commit {
    var newFileName string
    var newCommitMsg string
    
    // get command line args 
    if len(os.Args) < 2 {
	newFileName = getStringInput("File name: ")
	newCommitMsg = getStringInput("commit msg: ")
    } else {
	newFileName = os.Args[1]
	newCommitMsg = os.Args[2]
    }
    
    newFileData := readFile(newFileName)
    commit := Commit {fileName: newFileName, fileData: string(newFileData), commitMsg: newCommitMsg }
    
    return commit
}


/* ======================================== \\
||		main code		    ||
\\ ======================================== */
func main() {
    commit := makeCommit()
    fmt.Println(commit)
}

// takes in a msg to prompt for input, returns a fileName as a string. 
func getStringInput(msg string) string {
    var input string
    fmt.Printf(msg)
    fmt.Scan(&input )

    return input
}

// takes a file name and returns the file data as byte slice.
func readFile(fileName string) []byte {
    fileData, err := os.ReadFile(fileName) 

    if err != nil {
        log.Fatal("file error: ", fileName)
    }
    
    return fileData
}
