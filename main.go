package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("file1.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 500)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File1 content:", string(data[:count]))

	file2, err := os.Open("file2.txt")
	if err != nil {
		log.Fatal(err)
	}
	data1 := make([]byte, 500)
	count1, err := file2.Read(data1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File2 content:", string(data1[:count1]))
}
