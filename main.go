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
	ch := make(chan string)
	wg.Add(1)

	go getRepeatedWordsForFile2(&wg, ch)

	file1, err := os.Open("file1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file1.Close()

	data := make([]byte, 1000)
	count, err := file1.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	file1content := string(data[:count])
	file2content := <-ch

	combinedContent := file2content + " " + file1content

	fmt.Println("Combined Content Of File1 and File2:", combinedContent)
	wg.Add(1)

	go getCombinedRepeatedWords(&wg, combinedContent)

	fmt.Println("String Content of File1::", file1content)
	file1count := 0
	sumOfDuplicatesFile1 := 0

	for index, value := range countWords(file1content) {
		if value > 1 {
			file1count++
			sumOfDuplicatesFile1 = sumOfDuplicatesFile1 + value

			fmt.Println("File1's word and repeated times: ", index, "---->", value)
		}
	}
	fmt.Println("No.of Repeated words in File1 is:", file1count)
	fmt.Println("Duplication count in File1 is:", sumOfDuplicatesFile1)

	wg.Wait()

}

func getRepeatedWordsForFile2(wg *sync.WaitGroup, ch chan string) {

	file2, err := os.Open("file2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()
	data := make([]byte, 1000)
	count, err := file2.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	str2 := string(data[:count])
	ch <- str2

	fmt.Println("String Content of File2::", str2)
	file2count := 0
	sumOfDuplicatesFile2 := 0
	for index, value := range countWords(str2) {
		if value > 1 {
			file2count++
			sumOfDuplicatesFile2 = sumOfDuplicatesFile2 + value

			fmt.Println("File2's Word and repeated times: ", index, "---->", value)
		}
	}

	fmt.Println("No.of Repeated words in File2 is:", file2count)
	fmt.Println("Duplication count in File2 is:", sumOfDuplicatesFile2)
	wg.Done()

}

func getCombinedRepeatedWords(wg *sync.WaitGroup, combined string) {
	combinedCount := 0
	sumOfDuplicatesCombined := 0
	for index, value := range countWords(combined) {
		if value > 1 {
			combinedCount++
			sumOfDuplicatesCombined = sumOfDuplicatesCombined + value

			fmt.Println("Combined file's Word and repeated times: ", index, "---->", value)
		}
	}

	fmt.Println("No.of Repeated words in Combined File  is:", combinedCount)
	fmt.Println("Duplication count in Combined File is:", sumOfDuplicatesCombined)
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
