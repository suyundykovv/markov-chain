package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Marcovform(prefixWords []string, wordCount int, wordsMap map[string][]string) error {
	if wordCount > 10000 || wordCount < 0 {
		return fmt.Errorf("number of words cannot be more than 10000 or negative")
	}

	rand.Seed(time.Now().UnixNano())
	prefixKey := getPrefixKey(prefixWords)
	var result strings.Builder

	result.WriteString(prefixKey + " ")

	for i := 0; i < wordCount-Prefixlen; i++ {
		suffixes, exists := wordsMap[prefixKey]
		if !exists || len(suffixes) == 0 {
			fmt.Println(strings.TrimSpace(result.String()))
			return fmt.Errorf("no more suffixes found for the current prefix")
		}
		randomSuffix := getRandomSuffix(suffixes)
		result.WriteString(randomSuffix + " ")
		prefixWords = updatePrefixWords(prefixWords, randomSuffix)
		prefixKey = getPrefixKey(prefixWords)
	}

	fmt.Println(strings.TrimSpace(result.String()))
	return nil
}

func BuildTextSaver(words []string) map[string][]string {
	wordMap := make(map[string][]string)
	for i := 0; i < len(words)-Prefixlen; i++ {
		key := strings.Join(words[i:i+Prefixlen], " ")
		if i+Prefixlen < len(words) {
			suffix := words[i+Prefixlen]
			wordMap[key] = append(wordMap[key], suffix)
		}
	}
	return wordMap
}
