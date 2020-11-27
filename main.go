package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go getRepeatedWords(&wg)

	file, err := os.Open("file1.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 500)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	str1 := string(data[:count])
	fmt.Println("String Content of File1::", str1)
	for index, value := range countWords(str1) {
		fmt.Println("File1's word: ", index, "---->", value)
	}
	wg.Wait()

}

func getRepeatedWords(wg *sync.WaitGroup) {

	file, err := os.Open("file2.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 500)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	str2 := string(data[:count])

	fmt.Println("String Content of File2::", str2)

	for index, value := range countWords(str2) {
		fmt.Println("File2's Word: ", index, "---->", value)
	}

	wg.Done()

}

func countWords(str string) map[string]int {
	textContent := strings.Fields(str)
	mapForCounts := make(map[string]int)
	for _, val := range textContent {
		_, ok := mapForCounts[val]
		if ok {
			mapForCounts[val]++
		} else {
			mapForCounts[val] = 1
		}
	}
	return mapForCounts
}
