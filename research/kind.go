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
	check(matrix [][]int, pos coordinate) bool
	newRandLine(int) []int
}

type param struct {
	identPairsAmount_M  int
	identPairsMinDist_M int
	identNumsAmount_C   int
	identNumsMinDist_C  int

	identTrios_M          bool
	identPairs_R          bool
	threeAscendingOrder_R bool
	identNums_P           bool
}

type words struct {
	param param
}
type fingers struct {
	param param
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

func (w words) check(matrix [][]int, pos coordinate) bool {
	minDist, amount := identNums_C(matrix, pos)
	if minDist < w.param.identNumsMinDist_C || amount > w.param.identNumsAmount_C {
		return false
	}

	minDist, amount = identPairs_M(matrix, pos)
	if minDist < w.param.identPairsMinDist_M || amount > w.param.identPairsAmount_M {
		return false
	}
	return true
}

func (f fingers) check(pos coordinate, matrix [][]int) bool {
	if f.param.identNums_P == false && identNums_P(matrix, pos) {
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

	minDist, amount := identNums_C(matrix, pos)
	if minDist < f.param.identNumsMinDist_C || amount > f.param.identNumsAmount_C {
		return false
	}
	minDist, amount = identPairs_M(matrix, pos)
	if minDist < f.param.identPairsMinDist_M || amount > f.param.identPairsAmount_M {
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

func identPairs_M(matrix [][]int, pos coordinate) (minDist int, amount int) {
	minDist = math.MaxInt
	if pos.x == 0 {
		return minDist, amount
	}

	for y := pos.y - 1; y >= 0; y-- {
		for x := 1; x < len(matrix[y]); x++ {
			if matrix[pos.y][pos.x] == matrix[y][x] &&
				matrix[pos.y][pos.x-1] == matrix[y][x-1] {
				if minDist == math.MaxInt {
					minDist = pos.y - y
				}
				amount++
			}
		}
	}
	return minDist, amount
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
