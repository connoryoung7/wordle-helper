package internal_test

import (
	"fmt"
	"testing"
	"wordle-helper/internal"
)

func TestSolver_LoadWordsAndSuggestWords(t *testing.T) {
	words := []internal.Word{
		{Word: "APPLE", Frequency: 10},
		{Word: "ANGLE", Frequency: 5},
		{Word: "BANJO", Frequency: 3},
		{Word: "BINGO", Frequency: 8},
		{Word: "CRANE", Frequency: 7},
	}

	solver := internal.NewSolver()
	solver.LoadWords(words)

	tests := []struct {
		name        string
		constraints internal.WordContraints
		expected    []string
	}{
		{
			name: "No constraints",
			constraints: internal.WordContraints{
				ExcludedLetters:   make(map[byte]bool),
				Positions:         make(map[int]*byte),
				ExcludedPositions: make(map[int]map[byte]bool),
				LetterCount:       make(map[byte]int),
			},
			expected: []string{"APPLE", "ANGLE", "BANJO", "BINGO", "CRANE"},
		},
		{
			name: "Exclude letters A and E",
			constraints: internal.WordContraints{
				ExcludedLetters: map[byte]bool{
					'A': true,
					'E': true,
				},
				Positions:         make(map[int]*byte),
				ExcludedPositions: make(map[int]map[byte]bool),
				LetterCount:       make(map[byte]int),
			},
			expected: []string{"BANJO"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suggestions := solver.SuggestWords(tt.constraints)
			if len(suggestions) != len(tt.expected) {
				t.Errorf("expected %d suggestions, got %d", len(tt.expected), len(suggestions))
			}
			suggestionMap := make(map[string]bool)
			for _, s := range suggestions {
				suggestionMap[s] = true
			}
		})
	}
}

func TestSolver_GenerateValidStarterWords(t *testing.T) {
	words := []internal.Word{
		{Word: "APPLE", Frequency: 10},
		{Word: "ANGLE", Frequency: 5},
		{Word: "BANJO", Frequency: 3},
		{Word: "BINGO", Frequency: 8},
		{Word: "CRANE", Frequency: 7},
	}

	solver := internal.NewSolver()
	solver.LoadWords(words)

	starterWords := solver.GenerateValidStarterWords()

	for _, word := range starterWords {
		fmt.Println(word)
	}

	// println("Generated Starter Words:", starterWords)

	expectedStarters := []string{"ANGLE", "BANJO", "BINGO", "CRANE"}

	if len(starterWords) != len(expectedStarters) {
		t.Errorf("expected %d starter words, got %d", len(expectedStarters), len(starterWords))
	}
}
