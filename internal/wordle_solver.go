package internal

type LetterLevel map[byte]LetterLevel

type Solver struct {
	words LetterLevel
}

func NewSolver() *Solver {
	return &Solver{
		words: make(LetterLevel),
	}
}

const WordleWordLength = 5

func (s *Solver) LoadWords(wordList []Word) {
	for _, word := range wordList {
		currentLevel := s.words
		for i := 0; i < len(word.Word); i++ {
			char := word.Word[i]
			if _, exists := currentLevel[char]; !exists {
				currentLevel[char] = make(LetterLevel)
			}
			currentLevel = currentLevel[char]
		}
	}
}

func (s *Solver) GenerateValidStarterWords() []string {
	var results []string

	level := s.words
	excludedLetters := make(map[byte]bool)

	s.searchForStarterWords(level, "", excludedLetters, &results)

	return results
}

func (s *Solver) searchForStarterWords(level LetterLevel, prefix string, excludedLetters map[byte]bool, results *[]string) {
	if len(prefix) == WordleWordLength {
		*results = append(*results, prefix)
		return
	}

	for char, nextLevel := range level {
		if excludedLetters[char] {
			continue
		}
		excludedLetters[char] = true
		s.searchForStarterWords(nextLevel, prefix+string(char), excludedLetters, results)
		delete(excludedLetters, char)
	}
}

func (s *Solver) SuggestWords(constraints WordContraints) []string {
	var results []string
	s.traverse(s.words, "", constraints, &results)
	return results
}

func (s *Solver) traverse(level LetterLevel, prefix string, constraints WordContraints, results *[]string) {
	if len(prefix) == WordleWordLength {
		*results = append(*results, prefix)
		return
	}

	position := len(prefix)
	if constraints.Positions[position] != nil {
		if nextLevel, exists := level[*constraints.Positions[position]]; exists {
			s.traverse(nextLevel, prefix+string(*constraints.Positions[position]), constraints, results)
		} else {
			// No valid words in this path
			// This should in theory never happen because any valid word should exist in our words trie
			return
		}
	} else {
		for char, nextLevel := range level {
			// Check for letters that we know are not in the word or in specific positions
			if constraints.ExcludedLetters[char] || (constraints.ExcludedPositions[position] != nil && constraints.ExcludedPositions[position][char]) {
				continue
			}

			s.traverse(nextLevel, prefix+string(char), constraints, results)
		}
	}
}
