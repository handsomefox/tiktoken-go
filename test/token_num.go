package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/handsomefox/tiktoken-go"
)

// main
func main() {
	textList, modelList, encodingList := ReadTestFile()
	testTokenByModel(textList, modelList)
	fmt.Println("=========================================")
	testTokenByEncoding(textList, encodingList)
}

// read all columns from a file
func ReadTestFile() (textList []string, modelList []string, encodingList []string) {
	file, err := os.Open("test/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	textList = strings.Split(lines[0], ",")
	modelList = strings.Split(lines[1], ",")
	encodingList = strings.Split(lines[2], ",")

	return
}

// getTokenByModel
func getTokenByModel(text string, model string) (num_tokens int) {
	tkm, err := tiktoken.EncodingForModel(model)
	if err != nil {
		err = fmt.Errorf(": %w", err)
		return
	}

	token := tkm.Encode(text, nil, nil)

	return len(token)
}

// getTokenByEncoding
func getTokenByEncoding(text string, encoding string) (num_tokens int) {
	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		err = fmt.Errorf(": %w", err)
		return
	}

	token := tke.Encode(text, nil, nil)

	return len(token)
}

// testTokenByModel
func testTokenByModel(textList []string, modelList []string) {
	for i := range textList {
		for j := range modelList {
			fmt.Printf("text: %s, model: %s, token: %d\n", textList[i], modelList[j], getTokenByModel(textList[i], modelList[j]))
		}
	}
}

// testTokenByEncoding
func testTokenByEncoding(textList []string, encodingList []string) {
	for i := range textList {
		for j := range encodingList {
			fmt.Printf("text: %s, encoding: %s, token: %d\n", textList[i], encodingList[j], getTokenByEncoding(textList[i], encodingList[j]))
		}
	}
}
