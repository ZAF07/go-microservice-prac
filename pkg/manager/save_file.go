package manager

import (
	"fmt"
	"os"
	"time"
)


func Save()(err error){
	// Check if file exists
	if _, err := os.Stat("hello.txt"); err == nil {
		fmt.Println("File exists")
		content, _ := os.ReadFile("hello.txt")
		fmt.Println(string(content))
	}

	time.AfterFunc(3*time.Second, func() {
		data := "hello world"
		err = os.WriteFile("hello.txt", []byte(data), 0600)
		fmt.Println("Saved file")
	})
	return
}

