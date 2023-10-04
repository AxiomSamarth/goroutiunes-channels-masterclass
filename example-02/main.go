package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

func getWords() ([]string, error) {
	file, err := os.Open("./sample.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	textBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	textString := string(textBytes)
	return strings.Split(textString, " "), nil
}

func regularWordCount(words []string) map[string]int {
	result := map[string]int{}
	for _, word := range words {
		result[word]++
	}
	return result
}

func wordCount(words []string, wcChannel chan<- map[string]int) {
	wcChannel <- regularWordCount(words)
}

func main() {
	// get all the words from the text file
	words, err := getWords()
	if err != nil {
		panic(err)
	}

	var (
		endIndex       int
		iterationCount int
	)

	// create the channel
	wcChannel := make(chan map[string]int)

	for i := 0; i < len(words); i += 50 {
		if i+50 < len(words) {
			endIndex = i + 50
		} else {
			endIndex = len(words)
		}
		go wordCount(words[i:endIndex], wcChannel)
		iterationCount++
	}

	concurrentResult := map[string]int{}
	for i := 0; i < iterationCount; i++ {
		for key, value := range <-wcChannel {
			concurrentResult[key] += value
		}
	}

	regularResult := regularWordCount(words)

	if reflect.DeepEqual(regularResult, concurrentResult) {
		fmt.Println("Perfecto!")
	} else {
		fmt.Println("Error")
	}
}
