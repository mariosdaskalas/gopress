package main

import (
    "compress/gzip"
    "fmt"
    "io"
    "os"
)

func compress(sourceFile, destinationFile string) {
    // Open the source file
    source, err := os.Open(sourceFile)
    if err != nil {
        fmt.Println("Error opening source file:", err)
        return
    }
    defer source.Close()

    // Create the destination file
    destination, err := os.Create(destinationFile)
    if err != nil {
        fmt.Println("Error creating destination file:", err)
        return
    }
    defer destination.Close()

    // Create a gzip writer
    gzipWriter := gzip.NewWriter(destination)
    defer gzipWriter.Close()

    // Copy the contents of the source file to the gzip writer
    _, err = io.Copy(gzipWriter, source)
    if err != nil {
        fmt.Println("Error compressing file:", err)
        return
    }
    fmt.Println("File compressed successfully")
}

func decompress(sourceFile, destinationFile string) {

	source, err := os.Open(sourceFile)
    if err != nil {
        fmt.Println("Error opening source file:", err)
        return
    }
    defer source.Close()

	 // Create a gzip reader
	gzipReader, err := gzip.NewReader(source)
    if err != nil {
        fmt.Println("Error creating gzip reader:", err)
        return
    }
    defer gzipReader.Close()

	 // Create the destination file
	 destination, err := os.Create(destinationFile)
	 if err != nil {
		 fmt.Println("Error creating destination file:", err)
		 return
	 }
	 defer destination.Close()

	 // Copy the decompressed contents to the destination file
	 _, err = io.Copy(destination, gzipReader)
	 if err != nil {
		 fmt.Println("Error decompressing file:", err)
		 return
	 }
	 fmt.Println("File decompressed successfully")
}

func main() {
    fmt.Println(` 

    ██████   ██████  ██████  ██████  ███████ ███████ ███████ 
    ██       ██    ██ ██   ██ ██   ██ ██      ██      ██      
    ██   ███ ██    ██ ██████  ██████  █████   ███████ ███████ 
    ██    ██ ██    ██ ██      ██   ██ ██           ██      ██ 
     ██████   ██████  ██      ██   ██ ███████ ███████ ███████ 
                                                              
                                                              
                                    
    `)
    var choice int
    fmt.Println("[0] : Compress")
    fmt.Println("[1] : Decompress")
    fmt.Println()

    fmt.Printf("Select your choice: ")
    fmt.Scan(&choice)

    if choice == 0 {
        var sourceFile, destinationFile string
        fmt.Print("Enter the source file path: ")
        fmt.Scan(&sourceFile)
        fmt.Print("Enter the destination file path: ")
        fmt.Scan(&destinationFile)

        fmt.Println("Compressing...")
        compress(sourceFile, destinationFile)
    } else if choice == 1 {
		var sourceFile, destinationFile string
        fmt.Print("Enter the compressed file path: ")
        fmt.Scan(&sourceFile)
        fmt.Print("Enter the destination decompressed file path: ")
        fmt.Scan(&destinationFile)

        fmt.Println("Decompressing...")
        decompress(sourceFile, destinationFile)
    } else {
        fmt.Println("Invalid choice")
        return
    }
}