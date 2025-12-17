package internal

type LetterCell struct {
	Letter string `json:"letter" validate:"omitempty,uppercase,len=1,alpha"`
	Status string `json:"status" validate:"required,oneof=correct present absent empty"`
}

type Row struct {
	Cells []LetterCell `json:"cells" validate:"required,len=5,dive"`
}

type GetWordsRequest struct {
	Words []Row `json:"words"`
}

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
