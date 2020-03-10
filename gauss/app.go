package main

import(
	"fmt"
	"math"
	"encoding/json"
    "io/ioutil"
    "os"
)

type SLAU struct {
	A [][]float64 `json:"a"`
	B []float64   `json:"b"`
}

func main(){
	var slau SLAU
	var x[]float64

    jsonFile, err := os.Open("input.json")
    if err != nil {
        fmt.Println(err)
    }

    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &slau)

	x = solve(&slau)

	var j int = 0
	for i := len(x)-1; i >= 0; i-- {
		fmt.Printf("x[%v]: %v \n", j+1, x[i])
		j++
	}
}

func solve(slau *SLAU) []float64 {
	var _x[]float64
	for i := 0; i < len(slau.A); i++ {
		if slau.A[i][i] == 0 {
			// Делаем перестановку столбцов, подставляем наибольший коэффициент в диоганальный элемент
			var notZeroElementIndex int = 0
			for j:= i; j < len(slau.A[i]); j++ {
				if math.Abs(slau.A[i][j]) > 0 {
					notZeroElementIndex = j
					break // находим первый не нулевой элемент, не будем проходить все элементы
				}
			}
			fmt.Println("not zero element index: ", notZeroElementIndex)
			if notZeroElementIndex != 0 {
				swapColumns(i, notZeroElementIndex, &*slau)
			}
		}

		if slau.A[i][i] == 0 {
			if slau.B[i] == 0 {
				fmt.Println("Infite solutions");
			}else{
				fmt.Println("No solutions");
			}
		}

		for l := 0; l < len(slau.A); l++ {
			if l != i { // проходимся по остальным строкам
				var k float64 = slau.A[l][i] / slau.A[i][i]
				for j := i; j < len(slau.A[i]); j++ {
					slau.A[l][j] = slau.A[l][j] - slau.A[i][j]*k
				}
				slau.B[l] = slau.B[l] - slau.B[i]*k
			}
		}
	}
	
	for i := len(slau.A)-1; i>=0; i-- {
		var summ float64 = 0
		for j := len(slau.A)-1; j>i; j-- {
			summ = summ + slau.A[i][j]
		}
		_x = append(_x, (slau.B[i] - summ)/slau.A[i][i])
		for l:= i; l>=0; l-- {
			slau.A[l][i] = slau.A[l][i] * (slau.B[i] - summ)/slau.A[i][i];
		}
	}
	return _x
}

func swapColumns(j int, notZeroElementIndex int, slau *SLAU){
	for i := 0; i < len(slau.A); i++ {
		slau.A[i][j], slau.A[i][notZeroElementIndex] = slau.A[i][notZeroElementIndex], slau.A[i][j]
	}
}