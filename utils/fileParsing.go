package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadAndBuildWordsMap() (string, map[string][]string, error) {
	var words []string
	var textContent strings.Builder
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		textContent.WriteString(line + " ")
		words = append(words, strings.Fields(line)...)
	}

	if err := scanner.Err(); err != nil {
		return "", nil, fmt.Errorf("error reading from stdin: %v", err)
	}
	if Prefixlen < 0 {
		fmt.Println("Error: Prefix can not be negative.")
		os.Exit(1)
	} else if Prefixlen == 0 {
		fmt.Println("Prefix equal to zero.")
		os.Exit(0)
	}
	return textContent.String(), BuildTextSaver(words), nil
}

func HasInput() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	return fi.Mode()&os.ModeNamedPipe != 0
}

func ValidatePrefix(prefix string) []string {
	prefixWords := strings.Fields(prefix)
	if len(prefixWords) > 5 {
		fmt.Println("Error: Prefix must contain a maximum of 5 words.")
		os.Exit(1)
	}
	return prefixWords
}
