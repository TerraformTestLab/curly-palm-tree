package main

import "fmt"

const (
	welcomeMessage = "Hello, World!"
)

func main() {
	fmt.Println(HelloWorld())
}

// HelloWorld returns the classic greeting message
func HelloWorld() string {
	return welcomeMessage
}
