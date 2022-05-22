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
	a, b, err := parseInput()

	// проверка ошибок парсинга
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// отображаем исходные данные
	fmt.Println("Матрица A")
	a.print()
	fmt.Println("Cтолбец свободных членов b")
	b.print()

	numberPermutations := a.upperTriangular(b)

	determinant := a.getDeterminant(numberPermutations)
	fmt.Printf("Определитель равен: [%v] ", determinant)
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

		fmt.Println("--------------------------")
		fmt.Println("Матрица A")
		a.print()
		fmt.Println("Столбец свободных членов b")
		b.print()
	}

	return numberPermutations
}

func (a matrix) getDeterminant(numberPermutations int) float64 {
	result := 1.0

	for i := 0; i < a.getN(); i++ {
		result = result * a[i][i]
	}
	return result * math.Pow(1, float64(numberPermutations))
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
