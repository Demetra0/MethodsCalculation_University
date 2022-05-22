package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
)

/**
example 1
go build ./task_1.go && ./main -a="[[1,0,0],[0,1,0],[0,0,1]]" -b="[1,2,3]"

example 2
go build ./task_1.go && ./main -a="[[-1,0,0],[0,-1,0],[0,0,-1]]" -b="[-1,-2,-3]"

example 3
go build ./task_1.go && ./main -a="[[3,1,1],[1,3,1],[1,1,3]]" -b="[8,10,12]"

example 4
go build ./task_1.go && ./main -a="[[3,1,2],[3,1,0],[0,1,0]]" -b="[11,5,2]"

example 5
go build ./task_1.go && ./main -a="[[1,2,3],[2,1,0],[0,0,1]]" -b="[14,4,3]"
*/

type matrix [][]float64
type vector []float64

func main() {

	// парсинг входных данных
	a, err := parseInput()

	// проверка ошибок парсинга
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// отображаем исходные данные
	fmt.Println("Матрица A")
	a.print()

	fmt.Println("Обратная матрица")
	a.inverse().print()
}

func (a matrix) getN() int {
	return len(a)
}

// Прямой ход (приведение к верхней треугольной форме)
func (a matrix) upperTriangular(b vector) int {
	numberPermutations := 0 // кол-во перестановок

	for k := 0; k < a.getN(); k++ {
		// главный элемент
		max := k

		// двигаемся вправо от диаганаотного элемента, для поиска максимального по модулю элемента
		for i := k + 1; i < len(a); i++ {
			if math.Abs(a[i][k]) > math.Abs(a[max][k]) {
				max = i
			}
		}
		if max != k {
			numberPermutations++
		}

		a[k], a[max] = a[max], a[k]
		b[k], b[max] = b[max], b[k]

		if math.Abs(a[k][k]) == 0 {
			if b[k] == 0 {
				fmt.Println("Система имеет множество решений")
			} else {
				fmt.Println("Система не имеет решений")
			}
			os.Exit(1)
		}

		for m := k + 1; m < a.getN(); m++ {
			var c_m_k = a[m][k] / a[k][k]
			for i := k; i < a.getN(); i++ {
				a[m][i] = a[m][i] - c_m_k*a[k][i]
			}
			b[m] = b[m] - c_m_k*b[k]
		}
	}

	return numberPermutations
}

// Обраный ход
func (a matrix) getSolution(b vector) vector {
	var x vector = make(vector, a.getN())
	var sum float64 = 0

	for k := a.getN() - 1; k >= 0; k-- {
		sum = 0
		for j := k + 1; j < a.getN(); j++ {
			sum += a[k][j] * x[j]
		}
		x[k] = (b[k] - sum) / a[k][k]
	}

	return x
}

func (a *matrix) gaussMainElement(b vector) vector {
	a.upperTriangular(b)
	return a.getSolution(b)
}

func (a matrix) inverse() matrix {
	var itr vector = make(vector, a.getN())
	var E vector = make(vector, a.getN())

	var Am1 matrix = make(matrix, a.getN())
	for i := range a {
		Am1[i] = make(vector, a.getN())
	}

	for i := 0; i < a.getN(); i++ {
		for j := 0; j < a.getN(); j++ {
			if i == j {
				E[j] = 1
			} else {
				E[j] = 0
			}
		}
		itr = a.gaussMainElement(E)

		for k := 0; k < a.getN(); k++ {
			Am1[k][i] = itr[k]
		}
	}

	return Am1
}

// Вывод матрицы в stdout
func (a matrix) print() {
	for i := range a {
		for j := range a[i] {
			if a[i][j] == 0 { // избавление от -0
				fmt.Printf("[0] ")
			} else {
				fmt.Printf("[%v] ", a[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Вывод вектора в stdout
func (b vector) print() {
	for i := 0; i < len(b); i++ {
		fmt.Printf("[%v] ", b[i])
	}

	fmt.Println()
	fmt.Println()
}

// Парсинг входных данных
func parseInput() (matrix, error) {
	var a matrix

	// описание флагов командной строки
	aJson := flag.String("a", "[3,-9,3],[2,-4,4],[1,8,-18]", "квадратная матрица размером N на N")

	// парсинг флагов командной строки
	flag.Parse()

	// парсинг матрицы a из Json
	if err1 := json.Unmarshal([]byte(*aJson), &a); err1 != nil {
		return nil, err1
	}

	return a, nil
}
