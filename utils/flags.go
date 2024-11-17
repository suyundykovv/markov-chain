package utils

import (
	"flag"
	"fmt"
	"os"
)

const (
	DEFAULT_PREFIX_LEN = 2
	DEFAULT_WORD_COUNT = 100
)

var (
	WordCount int
	Prefix    string
	Prefixlen int
	Help      bool
)

func init() {
	flag.IntVar(&WordCount, "w", DEFAULT_WORD_COUNT, "Number of words to generate")
	flag.StringVar(&Prefix, "p", "", "Prefix to start the text generation")
	flag.IntVar(&Prefixlen, "l", DEFAULT_PREFIX_LEN, "len of prefics to generate")

	// error handling
	if Prefixlen <= 0 {
		fmt.Printf("Error : len of prefix can not be more 5 and negative or zero ")
		os.Exit(1)
	}

	flag.Usage = func() {
		fmt.Println("Markov Chain text generator.")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  markovchain [-w <N>] [-p <S>] [-l <N>]")
		fmt.Println("  markovchain --help")
		fmt.Println()
		fmt.Println("Options:")
		fmt.Println("    --help  Show this screen.")
		fmt.Println("    -w N    Number of maximum words")
		fmt.Println("    -p S    Starting prefix")
		fmt.Println("    -l N    Prefix length")
	}

	flag.Parse()
}
