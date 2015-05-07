package main

import (
    "fmt"
    "flag"
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }   
}

func main() {

    // A welcome message
    fmt.Printf("Welcome to File Handler Service\n")

    // Define Command Line flags
    opcodePtr := flag.String("o", "read", "opcode for file operations - read, write")
    msgPtr := flag.String("m", "Dont be lazy type something", "message to be written toa file")
    filePtr := flag.String("p", "null", "file path to read/write data")
    
    // Parse the command line argurments
    flag.Parse()

    fmt.Printf("opcode: %s\n", *opcodePtr)
    if *opcodePtr != "read" {
        fmt.Printf("message: %s\n", *msgPtr)
    }

    if *filePtr == "null" {
        fmt.Printf("You must enter a file path. See fileio -h for usage\n")
        return
    } else {
        fmt.Printf("file path: %s\n", *filePtr)
    }     

    switch *opcodePtr {
        case "read":
            rdContents, err := ioutil.ReadFile("/Users/sambati/tmp")
            check(err)
            fmt.Printf("Reading contents of file - /Users/sambati/tmp\n")
            fmt.Printf("======================================\n")
            fmt.Print(string(rdContents))
            fmt.Printf("\nDone\n")
            fmt.Printf("======================================\n") 
        case "write":
            fmt.Printf("Writing message %s to file - /Users/sambati/tmp\n", *msgPtr)
            fmt.Printf("=================================================\n")
            d1 := []byte(*msgPtr)        
            err := ioutil.WriteFile("/Users/sambati/tmp", d1, 0644)
            check(err)
            fmt.Printf("Done\n")
            fmt.Printf("=================================================\n")
        default:
            panic("Invalid opcode")            
    }
} 
