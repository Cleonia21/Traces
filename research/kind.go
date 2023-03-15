package research

import (
	"math"
	"math/rand"
	"time"
)

/*
function postfixes:
  M - matrix
  R - row
  C - column
  P - pair
*/

type kind interface {
	check(matrix [][]int, pos coordinate, previous map[int]map[int]map[int]int) bool
	newRandLine(int) []int
}

type param struct {
	identPairsAmount_M  int
	identPairsMinDist_M int
	identNumsAmount_C   int
	identNumsMinDist_C  int

	identPairsByColumn_M  bool
	identTrios_M          bool
	identPairs_R          bool
	threeAscendingOrder_R bool
	threeSuccessively_R   bool
	identNums_P           bool
	findTriosInPrevious_M bool
}

type words struct {
	param param
}
type fingers struct {
	param param
}

func fingersParamGet(columnNum int) fingers {
	m := map[int]param{
		3:  param{},
		4:  param{identPairsAmount_M: 1, identPairsMinDist_M: 2, identNumsAmount_C: 1, identNumsMinDist_C: 4},
		5:  param{identPairsAmount_M: 1, identPairsMinDist_M: 2, identNumsAmount_C: 1, identNumsMinDist_C: 3},
		6:  param{identPairsAmount_M: 1, identPairsMinDist_M: 2, identNumsAmount_C: 1, identNumsMinDist_C: 3},
		7:  param{identPairsAmount_M: 1, identPairsMinDist_M: 2, identNumsAmount_C: 1, identNumsMinDist_C: 2},
		8:  param{identPairsAmount_M: 1, identPairsMinDist_M: 1, identNumsAmount_C: 1, identNumsMinDist_C: 2},
		9:  param{identPairsAmount_M: 2, identPairsMinDist_M: 1, identNumsAmount_C: 1, identNumsMinDist_C: 2},
		10: param{identPairsAmount_M: 2, identPairsMinDist_M: 1, identNumsAmount_C: 2, identNumsMinDist_C: 2, identTrios_M: true},
	}
	return fingers{m[columnNum]}
}

func wordsParamGet(columnNum int) words {
	m := map[int]param{
		5:  param{identPairsAmount_M: 1, identPairsMinDist_M: 2},
		6:  param{},
		7:  param{},
		8:  param{},
		9:  param{},
		10: param{},
	}
	return words{m[columnNum]}
}

func (w words) newRandLine(length int) []int {
	line := make([]int, length)
	for i := 0; i < length; i++ {
		line[i] = i + 1
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(line),
		func(i, j int) { line[i], line[j] = line[j], line[i] })
	return line
}

func (f fingers) newRandLine(length int) []int {
	line := append(words(f).newRandLine(5), words(f).newRandLine(5)...)
	return line[0:length]
}

func (w words) check(matrix [][]int, pos coordinate, previous map[int]map[int]map[int]int) bool {
	minDist, amount := identNums_C(matrix, pos)
	if minDist < w.param.identNumsMinDist_C || amount > w.param.identNumsAmount_C {
		return false
	}

	minDist, amount, _ = identPairs_M(matrix, pos)
	if minDist < w.param.identPairsMinDist_M || amount > w.param.identPairsAmount_M {
		return false
	}

	//if w.param.findTriosInPrevious_M == true && findTriosInPrevious_M(matrix, pos, previous) == true {
	//	return false
	//}
	return true
}

func (f fingers) check(matrix [][]int, pos coordinate, previous map[int]map[int]map[int]int) bool {
	if f.param.identNums_P == false && identNums_P(matrix, pos) {
		return false
	}
	if f.param.threeSuccessively_R == false && threeSuccessively_R(matrix, pos) {
		return false
	}
	if f.param.threeAscendingOrder_R == false && threeAscendingOrder_R(matrix, pos) {
		return false
	}
	if f.param.identPairs_R == false && identPairs_R(matrix, pos) {
		return false
	}
	if f.param.identTrios_M == false && identTrios_M(matrix, pos) {
		return false
	}

	minDist, amount, identPairsByColumn_M := identPairs_M(matrix, pos)
	if minDist < f.param.identPairsMinDist_M || amount > f.param.identPairsAmount_M {
		return false
	}
	if f.param.identPairsByColumn_M == false && identPairsByColumn_M == true {
		return false
	}

	minDist, amount = identNums_C(matrix, pos)
	if minDist < f.param.identNumsMinDist_C || amount > f.param.identNumsAmount_C {
		return false
	}

	return true
}

func identNums_C(matrix [][]int, c coordinate) (minDist int, amount int) {
	minDist = math.MaxInt
	y := c.y
	for y > 0 {
		y--
		if matrix[y][c.x] == matrix[c.y][c.x] {
			if minDist == math.MaxInt {
				minDist = c.y - y
			}
			amount++
		}
	}
	return minDist, amount
}

func identPairs_M(matrix [][]int, pos coordinate) (minDist int, amount int, identPairsByColumn bool) {
	minDist = math.MaxInt
	if pos.x == 0 {
		return minDist, amount, identPairsByColumn
	}

	for y := pos.y - 1; y >= 0; y-- {
		for x := 1; x < len(matrix[y]); x++ {
			if matrix[pos.y][pos.x] == matrix[y][x] &&
				matrix[pos.y][pos.x-1] == matrix[y][x-1] {
				if pos.x == x {
					identPairsByColumn = true
				}

				if minDist == math.MaxInt {
					minDist = pos.y - y
				}
				amount++
			}
		}
	}
	return minDist, amount, identPairsByColumn
}

func findTriosInPrevious_M(matrix [][]int, pos coordinate, previous map[int]map[int]map[int]int) bool {
	if pos.x < 2 || pos.x+2 >= len(matrix[0]) {
		return false
	}

	n1 := matrix[pos.y][pos.x]
	n2 := matrix[pos.y][pos.x+1]
	n3 := matrix[pos.y][pos.x+2]
	if previous[n1][n2][n3] > 0 {
		return true
	}
	return false
}

func identTrios_M(matrix [][]int, pos coordinate) bool {
	if pos.x < 2 {
		return false
	}

	for y := pos.y - 1; y >= 0; y-- {
		for x := 2; x < len(matrix[y]); x++ {
			if matrix[pos.y][pos.x] == matrix[y][x] &&
				matrix[pos.y][pos.x-1] == matrix[y][x-1] &&
				matrix[pos.y][pos.x-2] == matrix[y][x-2] {
				return true
			}
		}
	}
	return false
}

func threeSuccessively_R(matrix [][]int, pos coordinate) bool {
	if pos.x < 2 {
		return false
	}

	if matrix[pos.y][pos.x] < matrix[pos.y][pos.x-1] &&
		matrix[pos.y][pos.x-1] < matrix[pos.y][pos.x-2] {
		return true
	}
	if matrix[pos.y][pos.x] > matrix[pos.y][pos.x-1] &&
		matrix[pos.y][pos.x-1] > matrix[pos.y][pos.x-2] {
		return true
	}
	return false
}

func threeAscendingOrder_R(matrix [][]int, pos coordinate) bool {
	if pos.x < 2 {
		return false
	}

	if matrix[pos.y][pos.x] == matrix[pos.y][pos.x-1]-1 &&
		matrix[pos.y][pos.x] == matrix[pos.y][pos.x-2]-2 {
		return true
	}
	if matrix[pos.y][pos.x] == matrix[pos.y][pos.x-1]+1 &&
		matrix[pos.y][pos.x] == matrix[pos.y][pos.x-2]+2 {
		return true
	}
	return false
}

func identNums_P(matrix [][]int, pos coordinate) bool {
	if pos.x == 0 {
		return false
	}

	if matrix[pos.y][pos.x] == matrix[pos.y][pos.x-1] {
		return true
	}

	return false
}

func identPairs_R(matrix [][]int, pos coordinate) bool {
	if pos.x == 0 {
		return false
	}

	tempPos := pos
	tempPos.x -= 2
	for ; tempPos.x > 0; tempPos.x-- {
		if matrix[pos.y][pos.x] == matrix[tempPos.y][tempPos.x] &&
			matrix[pos.y][pos.x-1] == matrix[tempPos.y][tempPos.x-1] {
			return true
		}
	}

	return false
}
