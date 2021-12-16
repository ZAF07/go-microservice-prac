package manager

import (
	"fmt"
	"os"
)


func Save()(err error){
	// Check if file exists
	if _, err := os.Stat("hello.txt"); err == nil {
		fmt.Println("File exists")
		content, _ := os.ReadFile("hello.txt")
		fmt.Println(string(content))
	}
	data := "hello world"
	err = os.WriteFile("hello.txt", []byte(data), 0600)
	return
}