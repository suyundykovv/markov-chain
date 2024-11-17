package main

import (
	"fmt"
	"os"
	"strings"

	. "markov-chain/utils"
)

func main() {
	if !HasInput() && Prefix == "" {
		fmt.Println("Error: no input text")
		os.Exit(1)
	}

	var inputText string
	var wordsMap map[string][]string
	var err error

	if HasInput() {
		inputText, wordsMap, err = ReadAndBuildWordsMap()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		if Prefix == "" {
			prefixWords := strings.Fields(inputText)
			if len(prefixWords) >= Prefixlen {
				Prefix = strings.Join(prefixWords[:Prefixlen], " ")
			} else {
				fmt.Println("Error: Not enough words to form a starting prefix.")
				os.Exit(1)
			}
		}
	} else {
		wordsMap = BuildTextSaver(strings.Fields(Prefix))
		if err != nil {
			fmt.Printf("Error building words map from prefix: %v\n", err)
			os.Exit(1)
		}
	}

	prefixWords := ValidatePrefix(Prefix)
	if len(prefixWords) > Prefixlen {
		fmt.Println("Error: prefix len not equal to 2, change default prefix len")
		os.Exit(1)
	}

	err = Marcovform(prefixWords, WordCount, wordsMap)
	if err != nil {
		fmt.Printf("Error generating Markov text: %v\n", err)
		os.Exit(1)
	}
}
