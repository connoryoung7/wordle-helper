package internal

type Word struct {
	Word      string `json:"word"`
	Frequency int    `json:"frequency"`
}

type WordContraints struct {
	ExcludedLetters   map[byte]bool
	Positions         map[int]*byte
	ExcludedPositions map[int]map[byte]bool
	LetterCount       map[byte]int
}
