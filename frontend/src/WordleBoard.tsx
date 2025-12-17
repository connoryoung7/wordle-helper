import { useState } from "react";

import WordCell from "./WordCell";

const DEFAULT_ROWS = 6;
const DEFAULT_COLS = 5;

function WordleGrid({ rows = DEFAULT_ROWS, cols = DEFAULT_COLS }) {
    const [currGuess, setCurrGuess] = useState(0);

    const handleLetterStatusChange = (row: number, col: number, status: 'correct' | 'present' | 'absent' | 'empty') => {
        // Logic to update the status of the letter in the grid
    }

    const isReadyForNextGuess = () => {
    }

    const toggleLetterStatus = (row: number, col: number) => {
        
    }

  return (
    <div className="min-h-screen flex items-center justify-center p-6">
      <div
        className="grid gap-2"
        style={{
          gridTemplateRows: `repeat(${rows}, 1fr)`,
          gridTemplateColumns: `repeat(${cols}, 1fr)`,
        }}
        aria-label="Wordle grid"
        role="grid"
      >
        {
            Array.from({ length: Math.min(rows, currGuess + 1) }).map((_, i) => {
                return (
                    <>
                        {Array.from({ length: cols }).map((_, j) => (
                            <WordCell
                                letter={undefined}
                                status="empty"
                                disabled={i > currGuess}
                                handleLetterStatusChange={handleLetterStatusChange}
                            />
                        ))}
                    </>
                )
            })
        }
      </div>
    </div>
  );
}

export default function App() {
  return <WordleGrid />;
}
