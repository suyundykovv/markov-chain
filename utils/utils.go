package utils

import (
	"math/rand"
	"strings"
)

func getPrefixKey(prefixWords []string) string {
	if len(prefixWords) < Prefixlen {
		return ""
	}
	if WordCount < Prefixlen {
		Prefixlen = WordCount
	}
	return strings.Join(prefixWords[:Prefixlen], " ")
}

func getRandomSuffix(suffixes []string) string {
	randomIndex := rand.Intn(len(suffixes))
	return suffixes[randomIndex]
}

func updatePrefixWords(prefixWords []string, newSuffix string) []string {
	return append(prefixWords[1:], newSuffix)
}
