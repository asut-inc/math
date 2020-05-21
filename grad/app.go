package main

import(
	"fmt"
	"math"
	"encoding/json"
    "io/ioutil"
	"os"
	"time"
)

type SLAU struct {
	A [][]float64 `json:"a"`
	B []float64   `json:"b"`
}

func main(){

	var slau SLAU
	var eps float64 = 0.0001

    jsonFile, err := os.Open("input.json")
    if err != nil {
        fmt.Println(err)
    }

    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &slau)

	prepare(&slau) // привести к сходимому виду
	// Если в функции prepare условия сходимости не выполняются, нужно останавливать алгоритм

	x := solve(&slau, &eps)

	for i := 0; i < len(x); i++ {
		fmt.Printf("x[%v]: %v \n", i+1, x[i])
	}
}

func solve(slau *SLAU, eps *float64) []float64 {
	var p[]float64
	var x[]float64

	for i := 0; i < len(slau.A); i++ {
		x = append(x, 0.1)
		p = append(p, 0.1)
	}
	t0 := time.Now()
	for {	
		for i := 0; i < len(x); i++ {
			p[i] = x[i]
		}
		for i := 0; i < len(slau.A); i++ {
			var sum float64 = 0
			for j := 0; j < i; j++ {
				sum = sum + slau.A[i][j] * x[j]
			}
			for j := i + 1; j < len(slau.A[i]); j++ {
				sum = sum + slau.A[i][j] * p[j]
			}
			x[i] = (slau.B[i] - sum) / slau.A[i][i]
		}
		
		if converge(x, p, eps) {
			break
		}
		
	}
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	return x
}

func converge(xk[] float64, xkp[] float64, eps *float64) bool {
    var norm float64 = 0
    for i := 0; i < len(xk); i++ {
        norm = norm + (xk[i] - xkp[i])*(xk[i] - xkp[i])
	}
	fmt.Println(math.Sqrt(norm))
	return math.Sqrt(norm) < *eps
}

/*
 * Для сходимости итерационного процесса достаточно, чтобы модули диагональных коэффициентов для каждого уравнения системы
 * были не меньше сумм модулей всех остальных коэффициентов (преобладание диагональных элементов).
 * При этом хотя бы для одного уравнения неравенство (25) должно выполняться строго.
 */
func prepare(slau *SLAU) {
	for i:=0; i<len(slau.A); i++ {
		var max float64 = 0
		var maxkey int = 0
		var minsSum float64 = 0
		for j:=0; j<len(slau.A[i]); j++ {
			minsSum += math.Abs(slau.A[i][j])
			if math.Abs(slau.A[i][j]) > max {
				minsSum += max
				max = math.Abs(slau.A[i][j])
				maxkey = j
				minsSum -= max
			}
		}
		if max > minsSum {
			// Поменять местами A[i,maxkey] с A[i,i]
			for l:=0; l<len(slau.A); l++ {
				slau.A[l][i], slau.A[l][maxkey] = slau.A[l][maxkey], slau.A[l][i]
			}
			//break
			// Можно прервать подготовку массива
		}else{
			// Это уравнение не приводит к сходимости
			fmt.Println("Уравнение не сходится")
		}
	}
	// fmt.Println(slau.A)
}