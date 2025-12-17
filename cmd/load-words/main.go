package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"wordle-helper/internal"
)

func main() {
	// ctx := context.Background()
	words, err := parseWords()
	if err != nil {
		panic(err)
	}

	fmt.Printf("We have found %d words in the list.\n", len(words))

	var wordList []internal.Word
	for _, w := range words {
		wordList = append(wordList, internal.Word{
			Word:      w,
			Frequency: 1,
		})
	}

	solver := internal.NewSolver()
	solver.LoadWords(wordList)

	fmt.Println("Words loaded into solver.")

	constraints := internal.WordContraints{
		ExcludedLetters:   make(map[byte]bool),
		Positions:         make(map[int]*byte),
		ExcludedPositions: make(map[int]map[byte]bool),
		LetterCount:       make(map[byte]int),
	}

	constraints.ExcludedLetters['A'] = true
	constraints.ExcludedLetters['E'] = true
	constraints.ExcludedLetters['I'] = true
	constraints.ExcludedLetters['O'] = true

	suggestions := solver.SuggestWords(constraints)
	fmt.Printf("Found %d possible words:\n", len(suggestions))

	constraints.Positions[0] = ptrByte('P')
	suggestions = solver.SuggestWords(constraints)
	fmt.Printf("Found %d possible words:\n", len(suggestions))

	for _, suggestion := range suggestions {
		fmt.Println(suggestion)
	}
}

func ptrByte(b byte) *byte {
	return &b
}

func parseWords() ([]string, error) {
	fh, err := os.Open("words.txt")
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)
	var words []string
	for scanner.Scan() {
		words = append(words, strings.ToUpper(strings.TrimSpace(scanner.Text())))
	}

	return words, nil
}
