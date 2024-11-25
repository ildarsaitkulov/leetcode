package arrays_and_hashing

import "unicode"

// сложность по времени и памяти O(n), потому что размеры судоку фиксированы
func isValidSudoku(board [][]byte) bool {
	colValues := make(map[int]map[byte]struct{})
	duplicateBox := make(map[int]map[byte]struct{})
	for row := 0; row <= 8; row++ {
		rowValues := make(map[byte]struct{})
		for col := 0; col <= 8; col++ {
			elem := board[row][col]
			if !unicode.IsDigit(rune(elem)) {
				continue
			}
			_, duplicateRow := rowValues[elem]
			if duplicateRow {
				return false
			}
			rowValues[elem] = struct{}{}
			_, duplicateCol := colValues[col][elem]
			if duplicateCol {
				return false
			}
			if colValues[col] == nil {
				colValues[col] = make(map[byte]struct{})
			}
			colValues[col][elem] = struct{}{}
			boxNumber := getBoxNumber(row, col)
			_, duplicateInBox := duplicateBox[boxNumber][elem]
			if duplicateInBox {
				return false
			}
			if duplicateBox[boxNumber] == nil {
				duplicateBox[boxNumber] = make(map[byte]struct{})
			}
			duplicateBox[boxNumber][elem] = struct{}{}

		}
	}

	return true
}

func getBoxNumber(row, col int) int {
	if row <= 2 {
		if col <= 2 {
			return 1
		}
		if col <= 5 {
			return 2
		}
		return 3
	}

	if row <= 5 {
		if col <= 2 {
			return 4
		}
		if col <= 5 {
			return 5
		}
		return 6
	}

	if col <= 2 {
		return 7
	}
	if col <= 5 {
		return 8
	}
	return 9
}
