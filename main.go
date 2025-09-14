package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	mat := inputMatrix()

	if mat == nil {
		fmt.Println("Ошибка ввода матрицы")
		return
	}

	matDet := det(mat)
	rank := rang(mat)

	fmt.Println("Матрица А: ")
	for row := range mat {
		fmt.Println(mat[row])
	}

	fmt.Print("Детерминант матрицы А: ")
	fmt.Println(matDet)

	fmt.Print("Ранг матрицы А: ")
	fmt.Println(rank)

}

func inputMatrix() [][]float64 {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите количество строк матрицы: ")
	rowsStr, _ := reader.ReadString('\n')
	rowsStr = strings.TrimSpace(rowsStr)
	rows, err := strconv.Atoi(rowsStr)
	if err != nil || rows <= 0 {
		fmt.Println("Некорректное количество строк")
		return nil
	}

	fmt.Print("Введите количество столбцов матрицы: ")
	colsStr, _ := reader.ReadString('\n')
	colsStr = strings.TrimSpace(colsStr)
	cols, err := strconv.Atoi(colsStr)
	if err != nil || cols <= 0 {
		fmt.Println("Некорректное количество столбцов")
		return nil
	}

	matrix := make([][]float64, rows)

	fmt.Printf("Введите элементы матрицы построчно (%d строк по %d чисел):\n", rows, cols)
	fmt.Println("Числа в строке разделяйте пробелами:")

	for i := 0; i < rows; i++ {
		fmt.Printf("Строка %d: ", i+1)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		numbers := strings.Fields(line)

		if len(numbers) != cols {
			fmt.Printf("Ошибка: ожидалось %d чисел, получено %d\n", cols, len(numbers))
			return nil
		}

		matrix[i] = make([]float64, cols)
		for j, numStr := range numbers {
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				fmt.Printf("Ошибка преобразования числа: %s\n", numStr)
				return nil
			}
			matrix[i][j] = num
		}
	}

	return matrix
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

func rang(mat [][]float64) float64 {
	if !is_square_mat(mat) {
		return 0
	}
	rows := len(mat)
	cols := len(mat[0])

	matrix := make([][]float64, rows)
	for i := range mat {
		matrix[i] = make([]float64, cols)
		copy(matrix[i], mat[i])
	}

	rank := 0

	for col := 0; col < cols && rank < rows; col++ {
		pivotRow := -1
		for row := rank; row < rows; row++ {
			if math.Abs(matrix[row][col]) > 1e-10 {
				pivotRow = row
				break
			}
		}

		if pivotRow == -1 {
			continue
		}

		if pivotRow != rank {
			matrix[rank], matrix[pivotRow] = matrix[pivotRow], matrix[rank]
		}

		pivot := matrix[rank][col]
		for j := col; j < cols; j++ {
			matrix[rank][j] /= pivot
		}

		for i := rank + 1; i < rows; i++ {
			factor := matrix[i][col]
			for j := col; j < cols; j++ {
				matrix[i][j] -= factor * matrix[rank][j]
			}
		}

		rank++
	}

	return float64(rank)
}
