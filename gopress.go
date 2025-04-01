package main

import "fmt"

func compress() {
	// This function will compress the data
}

func decompress() {
	// This function will decompress the data
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
		fmt.Println("Compressing...")
		compress()
	} else if choice == 1 {
		fmt.Println("Decompressing...")
		decompress()
	} else {
		fmt.Println("Invalid choice")
		return
	}

}