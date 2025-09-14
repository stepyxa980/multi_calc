package main

import (
	"fmt"
	"math"
)

func main() {
	mat := [][]float64{{1, 2, 3},
		{4, 8, 6},
		{1, 1, 3}}

	matMinor := minor(mat, 2, 3)
	matDet := det(mat)

	fmt.Println("Матрица А: ")
	for row := range mat {
		fmt.Println(mat[row])
	}

	fmt.Println("Минор 2,3 матрицы А: ")

	for row := range matMinor {
		fmt.Println(matMinor[row])
	}

	fmt.Println("Детерминант матрицы А: ")
	fmt.Println(matDet)

}

func minor(mat [][]float64, i int, j int) [][]float64 {
	mat_len := len(mat)
	if mat_len == 0 {
		return nil
	}

	cMat := make([][]float64, mat_len-1)

	rowIdx := 0
	for row := range mat {
		if row == i-1 {
			continue
		}

		cMat[rowIdx] = make([]float64, len(mat[row])-1)

		colIdx := 0

		for col := range mat[row] {
			if col == j-1 {
				continue
			}

			cMat[rowIdx][colIdx] = mat[row][col]
			colIdx++
		}
		rowIdx++
	}

	return cMat
}

func is_square_mat(mat [][]float64) bool {
	for row := range mat {
		if len(mat[row]) != len(mat) {
			return false
		}
	}
	return true
}

func det(mat [][]float64) float64 {
	mat_len := len(mat)
	if !is_square_mat(mat) {
		return 0
	}

	if mat_len == 1 {
		return mat[0][0]
	}

	if mat_len == 2 {
		return mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]
	}

	var determinant float64 = 0

	for j := range mat_len {
		minor_mat := minor(mat, 1, j+1)

		determinant += math.Pow(-1, float64(j)) * mat[0][j] * det(minor_mat)
	}

	return determinant
}
