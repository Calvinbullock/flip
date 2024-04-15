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

// NOTE ?????
func compareCommitData(newCommit Commit, oldCommit Commit) {
    newLines := strings.Split(newCommit.fileData, "/n")
    oldLines := strings.Split(oldCommit.fileData, "/n")

    for _, line := range newLines {
	fmt.Print("L:", line)
    }

    for _, line := range oldLines {
	fmt.Print("L:", line)
    }
}

func testCompareCommitData() bool {
    commit1 := Commit {fileName: "test", fileData: string(readFile("test")), commitMsg: "commit 1"}
    commit2 := Commit {fileName: "test2", fileData: string(readFile("test2")), commitMsg: "commit 2"}

    compareCommitData(commit1, commit2)

    return true
}

/* ======================================== \\
||		main code		    ||
\\ ======================================== */

// runs all test functions
func testRunner() bool {
    if testCompareCommitData() == false {
	return false
    }

    fmt.Println()
    return true
}

func main() {

    if testRunner() == false {
	return
    }

    commitList := []Commit {}
    
    for true {
	commit := makeCommit()
	commitList = append(commitList, commit)   

	fmt.Println(commit.fileName, commit.commitMsg) 
	fmt.Println()
    }
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
