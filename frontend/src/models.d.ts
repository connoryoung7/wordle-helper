export type LetterStatus = 'correct' | 'present' | 'absent' | 'empty';

export interface Cell {
    letter?: string
    status: LetterStatus
}

export type Grid = Cell[][];
