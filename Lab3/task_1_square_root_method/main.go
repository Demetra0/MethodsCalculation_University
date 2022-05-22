package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
)

/**
example 1
Unix
go build ./main.go && ./main -a="[[-1,0,0],[0,-1,0],[0,0,-1]]" -b="[-1,-2,-3]"

Windows
go build ./main.go ; ./main -a="[[-1,0,0],[0,-1,0],[0,0,-1]]" -b="[-1,-2,-3]"

example 2
go build ./main.go ; ./main -a="[[3,1,1],[1,3,1],[1,1,3]]" -b="[8,10,12]"

example 3 (variant 11)
go build ./main.go ; ./main -a="[[2.56,0.67,-1.78],[0.67,-2.67,1.35],[-1.78,1.35,-0.55]]" -b="[1.14,0.66,1.72]"
*/

const N int = 11

type matrix [][]float64
type vector []float64

func main() {
	// Матрица из терминала
	//A, b := getMatrixFromTerminal()
	//println("Решение x")
	//A.squareRootMethod(b).print()

	// Матрица Гильберта
	var n = 1000
	println(n)
	H, b := getHilbertMatrix(n)

	println("Решение x")
	H.squareRootMethod(b).printEveryNSteps(5)
}

func getMatrixFromTerminal() (matrix, vector) {
	tempA, tempB, err := parseInput()
	n := tempA.getN()

	// проверка ошибок парсинга
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var A = make(matrix, n+1)
	var b = make(vector, n+1)

	for i := range A {
		A[i] = make(vector, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			A[i][j] = tempA[i-1][j-1]
		}
		b[i] = tempB[i-1]
	}

	return A, b
}

func getHilbertMatrix(n int) (matrix, vector) {
	var H = make(matrix, n+1)
	for i := range H {
		H[i] = make(vector, n+1)
	}

	var x = make(vector, n+1)
	var b = make(vector, n+1)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			H[i][j] = 1.0 / float64(i+j-1)
		}
		x[i] = float64(i * N)
		b[i] = 0.0
	}

	for j := 0; j <= n; j++ {
		for i := 0; i <= n; i++ {
			b[j] += H[i][j] * x[i]
		}
	}

	return H, b
}

func (A matrix) squareRootMethod(b vector) vector {
	var n = len(b) - 1

	var x = make(vector, n+1)
	var y = make(vector, n+1)

	var D = make(vector, n+1)
	var S = make(matrix, n+1)
	for i := range A {
		S[i] = make(vector, n+1)
	}

	D[1] = getSing(A[1][1])
	S[1][1] = math.Sqrt(math.Abs(A[1][1]))

	for j := 2; j <= n; j++ {
		S[1][j] = A[1][j] / (S[1][1] * D[1])
	}

	// Рассчитываем оставшиеся значения S и D
	for i := 2; i <= n; i++ {
		var sum float64 = 0
		for l := 1; l <= i-1; l++ {
			sum += S[l][i] * S[l][i] * D[l]
		}

		D[i] = getSing(A[i][i] - sum)
		S[i][i] = math.Sqrt(math.Abs(A[i][i] - sum))

		var k = 1 / (S[i][i] * D[i])
		for j := i + 1; j <= n; j++ {
			sum = 0
			for l := 1; l <= i-1; l++ {
				sum += S[l][i] * D[l] * S[l][j]
			}
			S[i][j] = (A[i][j] - sum) * k
		}
	}

	// Ищем решение системы (s^t * d)y = b
	y[1] = b[1] / S[1][1] * D[1]

	for i := 2; i <= n; i++ {
		var sum float64 = 0
		for j := 1; j <= i-1; j++ {
			sum += S[j][i] * y[j] * D[j]
		}
		y[i] = (b[i] - sum) / (S[i][i] * D[i])
	}

	//	Используя значения y, найденные в предыдущем уравнении, решаем второе уравнение Sx = y
	x[n] = y[n] / S[n][n]
	for i := n - 1; i >= 1; i-- {
		var sum float64 = 0
		for l := i + 1; l <= n; l++ {
			sum += S[i][l] * x[l]
		}
		x[i] = (y[i] - sum) / S[i][i]
	}

	return x
}

func getSing(x float64) float64 {
	var sing float64 = -1

	if x == 0.0 {
		sing = 0
	} else if x > 0.0 {
		sing = 1
	}

	return sing
}

func (A matrix) getN() int {
	return len(A)
}

func (A matrix) print() {
	for i := 1; i < len(A); i++ {
		for j := 1; j < len(A[i]); j++ {
			if A[i][j] == 0 { // избавление от -0
				fmt.Printf("[0] ")
			} else {
				fmt.Printf("[%v] ", A[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func (b vector) print() {
	for i := 1; i < len(b); i++ {
		fmt.Printf("[%v] ", b[i])
	}

	fmt.Println()
	fmt.Println()
}

func parseInput() (matrix, vector, error) {
	var a matrix
	var b vector

	// описание флагов командной строки
	aJson := flag.String("a", "[3,-9,3],[2,-4,4],[1,8,-18]", "квадратная матрица размером N на N")
	bJson := flag.String("b", "[-18,-10,35]", "числовой вектор-столбец размером N")

	// парсинг флагов командной строки
	flag.Parse()

	// парсинг матрицы a из Json
	if err1 := json.Unmarshal([]byte(*aJson), &a); err1 != nil {
		return nil, nil, err1
	}

	// парсинг вектора b из Json
	if err2 := json.Unmarshal([]byte(*bJson), &b); err2 != nil {
		return nil, nil, err2
	}

	// вылидация данных
	if len(a) < 2 || len(a) != len(b) {
		return nil, nil, errors.New("не верный формат данных")
	}

	// вылидация данных
	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b) {
			return nil, nil, errors.New("не верный формат данных")
		}
	}

	return a, b, nil
}

func (b vector) printEveryNSteps(n int) {
	for i := 1; i <= len(b)-1; i += n {
		fmt.Printf("x%v = %v\n", i, b[i])
	}
}
