import React from 'react';

interface WordCellProps {
    letter?: string;
    status: 'correct' | 'present' | 'absent' | 'empty';
    disabled?: boolean;
    handleLetterStatusChange?: (row: number, col: number, status: 'correct' | 'present' | 'absent' | 'empty') => void;
}

const statusClasses: Record<WordCellProps['status'], string> = {
    correct: 'bg-green-500 text-white border-green-700',
    present: 'bg-yellow-500 text-white border-yellow-700',
    absent: 'bg-gray-500 text-white border-gray-700',
    empty: 'bg-white text-black border-neutral-700',
};

// const handleLetterStatusChange = (row: number, col: number, status: 'correct' | 'present' | 'absent' | 'empty') => {
//     // Logic to update the status of the letter in the grid
// }

const WordCell: React.FC<WordCellProps> = ({ letter, status, disabled, handleLetterStatusChange }) => {
    const [value, setValue] = React.useState('');

    console.log("Rendering WordCell:", { letter, status, disabled });

    return (
        <div
            className={`w-14 h-14 sm:w-16 sm:h-16 border-2 rounded-lg flex items-center justify-center ${statusClasses[status]}`}
            role="gridcell"
            aria-label={`Cell with letter ${letter} and status ${status}`}
        >
            <span className="text-3xl font-bold">{
                (status === 'empty') ? <input
                value={value}
                onChange={(e) => {
                    const value = e.target.value;
                    if (value.length <= 1 && /^[a-zA-Z]?$/.test(value)) {
                        setValue(e.target.value)
                    }
                }}
                disabled={disabled}
                maxLength={1}
                inputMode="text"
                autoComplete="off"
                autoCorrect="off"
                autoCapitalize="characters"
                spellCheck={false}
                className="w-full h-full bg-transparent text-center text-2xl sm:text-3xl font-bold tracking-wide uppercase outline-none"
                /> : letter
            }</span>
        </div>
    );
}

export default WordCell;
