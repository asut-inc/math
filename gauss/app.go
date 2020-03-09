package main

type GaussImpl struct {
	Matrix Matrix
	Vector []float64
}

type Gauss interface {
	Solve() int
}

type Coordinates struct {
	x int
	y int
}

type Matrix map[Coordinates]int

var MatrixData [][]int = [][]int{
	[]int{1, 2, 3},
	[]int{4, 5, 6},
	[]int{7, 8, 9},
}

func NewGauss(matrixData [][]int, vector []float64) *GaussImpl {
	var gaussData *GaussImpl

	gaussData = &GaussImpl{
		Matrix: newMatrix(matrixData),
		Vector: vector,
	}

	return gaussData
}

func (m Matrix) PouringData(rowsCount int, columnsCount int, matrixData [][]int) {
	for columnIdx := 0; columnIdx < columnsCount-1; columnIdx++ {
		for rowIdx := 0; rowIdx < rowIdx-1; rowIdx++ {
			m[Coordinates{x: columnIdx, y: rowIdx}] = matrixData[columnIdx][rowIdx]
		}
	}
}

func (m Matrix) getCoordVal(x int, y int) int {
	return m[Coordinates{x, y}]
}

func NewCoord(x int, y int) *Coordinates {
	return &Coordinates{x, y}
}

func (m Matrix) SetValue(coord *Coordinates, matrixValue int) {
	m[*coord] = matrixValue
}

func MatrixMapper(matrix Matrix, matrixData [][]int) Matrix {
	columnsCount := len(matrixData[0])
	rowsCount := len(matrixData)

	for columnIdx := 0; columnIdx < columnsCount-1; columnIdx++ {
		for rowIdx := 0; rowIdx < rowsCount-1; rowsCount++ {
			coord := NewCoord(columnIdx, rowIdx)
			matrix.SetValue(coord, matrixData[columnIdx][rowIdx])
		}
	}

	return matrix
}

func newMatrix(matrixData [][]int) Matrix {
	matrix := make(Matrix)
	matrix = MatrixMapper(matrix, matrixData)

	return matrix
}

func main() {

	//jsonFile, err := os.Open("input.json")
	//if err != nil {
	//    fmt.Println(err)
	//}
	//
	//defer jsonFile.Close()
	//
	//byteValue, _ := ioutil.ReadAll(jsonFile)
	//
	//json.Unmarshal(byteValue, &slau)
	//
	//if solve() {
	//	var j int = 0
	//	for i := len(slau.B)-1; i >= 0; i-- {
	//		fmt.Printf("x[%v]: %v \n", j+1, x[i])
	//		j++
	//	}
	//}
}

// 9 18 10 -16

func solve() bool {
	//for i := 0; i < len(slau.A); i++ {
	//	if slau.A[i][i] == 0 {
	//		// Делаем перестановку столбцов, подставляем наибольший коэффициент в диоганальный элемент
	//		var notZeroElementIndex int = 0
	//		for j:= i; j < len(slau.A[i]); j++ {
	//			if math.Abs(slau.A[i][j]) > 0 {
	//				notZeroElementIndex = j
	//				break // находим первый не нулевой элемент, не будем проходить все элементы
	//			}
	//		}
	//		fmt.Println("not zero element index: ", notZeroElementIndex)
	//		if notZeroElementIndex != 0 {
	//			swapColumns(i, notZeroElementIndex)
	//		}
	//	}
	//
	//	if slau.A[i][i] == 0 {
	//		if slau.B[i] == 0 {
	//			fmt.Println("Infite solutions");
	//			return false;
	//		}else{
	//			fmt.Println("No solutions");
	//			return false;
	//		}
	//	}
	//
	//	for l := 0; l < len(slau.A); l++ {
	//		if l != i { // проходимся по остальным строкам
	//			var k float64 = slau.A[l][i] / slau.A[i][i]
	//			for j := i; j < len(slau.A[i]); j++ {
	//				slau.A[l][j] = slau.A[l][j] - slau.A[i][j]*k
	//			}
	//			slau.B[l] = slau.B[l] - slau.B[i]*k
	//		}
	//	}
	//}
	//
	//for i := len(slau.A)-1; i>=0; i-- {
	//	var summ float64 = 0
	//	for j := len(slau.A)-1; j>i; j-- {
	//		summ = summ + slau.A[i][j]
	//	}
	//	x = append(x, (slau.B[i] - summ)/slau.A[i][i])
	//	for l:= i; l>=0; l-- {
	//		slau.A[l][i] = slau.A[l][i] * (slau.B[i] - summ)/slau.A[i][i];
	//	}
	//}
	return true
}

func swapColumns(j int, notZeroElementIndex int) {
	//for i := 0; i < len(slau.A); i++ {
	//	slau.A[i][j], slau.A[i][notZeroElementIndex] = slau.A[i][notZeroElementIndex], slau.A[i][j]
	//}
}
