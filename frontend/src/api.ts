import type { Cell } from "./models";

export const API_BASE_URL = 'http://localhost:5000/api';

export const getPossibleWords = async (grid: Cell[][]) => {
    const response = await fetch(`${API_BASE_URL}/words`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ grid }),
    });
    const data = await response.json();
    return data.possibleWords;
}
