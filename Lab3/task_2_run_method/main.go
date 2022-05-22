package main

import (
	"fmt"
)

const N int = 11

type matrix [][]float64
type vector []float64

func main() {
	var n int = 101
	var d float64 = 2

	fmt.Printf("n = %v\n", n)
	fmt.Printf("d = %v\n", d)

	A, b := get3DiagonalMatrix(n, d)

	x := A.sweepMethod(b)
	x.printResult()
}

func get3DiagonalMatrix(n int, d float64) (matrix, vector) {
	var A = make(matrix, 3)
	for i := range A {
		A[i] = make(vector, n+1)
	}

	var b = make(vector, n+1)

	for i := 2; i < n; i++ {
		A[0][i] = 1
		A[1][i] = d
		A[2][i] = 1
		b[i] = (float64(N*(i-1)))/float64(n-1) + 1
	}

	A[0][1] = 0
	A[0][n] = 0

	A[1][1] = 1
	A[1][n] = 1

	A[2][1] = 0
	A[2][n] = 0

	b[1] = 1
	b[n] = float64(N + 1)

	//	Вывод диагонали
	for i := 1; i < n; i++ {
		fmt.Printf("[%v] ", A[0][i])
	}
	fmt.Println()

	//	Вывод диагонали
	for i := 1; i <= n; i++ {
		fmt.Printf("[%v] ", A[1][i])
	}
	fmt.Println()

	//	Вывод диагонали
	for i := 2; i <= n; i++ {
		fmt.Printf("[%v] ", A[2][i])
	}
	fmt.Println()

	//	Вывод вектора правой части
	b.print()

	return A, b
}

func (A matrix) sweepMethod(b vector) vector {
	var n = len(b) - 1

	// Прямой ход
	var x = make(vector, n+1)
	var alpha = make(vector, n+1)
	var beta = make(vector, n+1)

	alpha[1] = -A[0][1] / A[1][1]
	beta[1] = b[1] / A[1][1]

	for i := 2; i < n; i++ {
		alpha[i] = -A[0][i] / (A[1][i] + A[2][i]*alpha[i-1])
		beta[i] = (b[i] - A[2][i]*beta[i-1]) / (A[1][i] + A[2][i]*alpha[i-1])
	}

	//	Обратный ход
	x[n] = (b[n] + A[2][n]*beta[n-1]) / (A[1][n] - A[2][n]*alpha[n-1])
	for i := n - 1; i >= 1; i-- {
		x[i] = alpha[i]*x[i+1] + beta[i]
	}

	return x
}

func (A matrix) print() {
	for i := 1; i < len(A); i++ {
		for j := 1; j < len(A[i]); j++ {
			fmt.Printf("[%v]", A[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func (b vector) print() {
	for i := 1; i < len(b); i++ {
		fmt.Printf("[%.2f] ", b[i])
	}

	fmt.Println()
	//fmt.Println()
}

func (b vector) printResult() {
	for i := 1; i <= len(b)-1; i += 10 {
		fmt.Printf("x%v = %v\n", i, b[i])
	}
}
