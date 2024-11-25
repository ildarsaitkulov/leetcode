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
	return (row/3)*3 + col/3
}
