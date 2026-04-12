package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Go programming!"
	println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	input, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s", input)
}
