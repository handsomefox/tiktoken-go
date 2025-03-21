package tiktoken

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

const TEST_FILE = "test/test.txt"

func ReadTestFile() ([]byte, error) {
	// open and read TEST_FILE
	return os.ReadFile(TEST_FILE)
}

func BenchmarkEncoding(b *testing.B) {
	fileContent, err := ReadTestFile()
	if err != nil {
		panic(err)
	}

	tkm, err := EncodingForModel("gpt-4o")
	if err != nil {
		panic(err)
	}

	text := string(fileContent)

	for range 4 {
		// do actual encoding
		fmt.Printf("Encoding %d bytes\n", len(text))
		tkm.Encode(text, nil, nil)

		stringBuilder := strings.Builder{}
		for range 10 {
			stringBuilder.WriteString(text)
		}

		text = stringBuilder.String()
	}
}
