package research

import (
	"math/rand"
	"time"
)

var patterns = [][][]int{
	{
		{1, 2, 3, 4},
		{2, 1, 4, 3},
		{3, 4, 1, 2},
		{4, 3, 2, 1},
		{1, 2, 4, 3},
	},
	{
		{1, 2, 3, 4},
		{2, 1, 4, 3},
		{3, 4, 1, 2},
		{4, 2, 3, 1},
		{1, 3, 2, 4},
	},
	{
		{1, 2, 3, 4},
		{2, 1, 4, 3},
		{3, 4, 2, 1},
		{4, 3, 1, 2},
		{2, 1, 3, 4},
	},
	{
		{1, 2, 3, 4},
		{2, 4, 1, 3},
		{3, 1, 4, 2},
		{1, 3, 2, 4},
		{4, 2, 3, 1},
	},
	{
		{1, 2, 3, 4},
		{2, 4, 1, 3},
		{3, 1, 4, 2},
		{1, 3, 2, 4},
		{4, 2, 1, 3},
	},
	{
		{1, 2, 3, 4},
		{4, 3, 2, 1},
		{3, 4, 1, 2},
		{2, 1, 4, 3},
		{1, 3, 2, 4},
	},
	{
		{1, 2, 3, 4},
		{4, 3, 2, 1},
		{3, 4, 1, 2},
		{2, 1, 3, 4},
		{1, 2, 4, 3},
	},
	{
		{1, 2, 3, 4},
		{4, 3, 2, 1},
		{3, 1, 4, 2},
		{2, 4, 1, 3},
		{4, 2, 3, 1},
	},
	{
		{1, 2, 3, 4},
		{4, 3, 1, 2},
		{3, 4, 2, 1},
		{1, 2, 4, 3},
		{2, 1, 3, 4},
	},
	{
		{1, 2, 3, 4},
		{4, 3, 1, 2},
		{3, 4, 2, 1},
		{1, 2, 4, 3},
		{2, 3, 1, 4},
	},
}

func getWords3() [][]int {
	lev := getWords4()
	for i, _ := range lev {
		lev[i] = lev[i][:3]
	}
	return lev
}

func getWords4() [][]int {
	mask := words{}.newRandLine(4)

	rand.Seed(time.Now().UnixNano())
	return PutMask(rand.Intn(len(patterns)), mask)
}

func PutMask(index int, mask []int) [][]int {
	tmp := make([][]int, 5)
	for i, _ := range patterns[index] {
		tmp[i] = make([]int, 4)
		for j, _ := range patterns[index][i] {
			tmp[i][j] = mask[patterns[index][i][j]-1]
		}
	}
	return tmp
}

/*
var l int

func mapCheck3(m map[int][]int) bool {
	if m[0][0] == m[3][0] && m[0][1] == m[3][1] && m[0][2] == m[3][2] {
		return false
	}
	if m[0][0] == m[4][0] && m[0][1] == m[4][1] && m[0][2] == m[4][2] {
		return false
	}
	if m[1][0] == m[4][0] && m[1][1] == m[4][1] && m[1][2] == m[4][2] {
		return false
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if m[i][j] == m[i+2][j] {
				return false
			}
		}
		for j := 0; j < 2; j++ {
			if m[i][j] == m[i+2][0] && m[i][j+1] == m[i+2][1] ||
				m[i][j] == m[i+2][1] && m[i][j+1] == m[i+2][2] {
				return false
			}
		}
	}

	return true
}

type patterns struct {
}

func my(m map[int][]int, depth int) {
	if depth == 4 {
		if !mapCheck3(m) {
			return
		}
		fmt.Println("{")
		fmt.Printf("{%d, %d, %d, %d},\n", m[0][0], m[0][1], m[0][2], m[0][3])
		fmt.Printf("{%d, %d, %d, %d},\n", m[1][0], m[1][1], m[1][2], m[1][3])
		fmt.Printf("{%d, %d, %d, %d},\n", m[2][0], m[2][1], m[2][2], m[2][3])
		fmt.Printf("{%d, %d, %d, %d},\n", m[3][0], m[3][1], m[3][2], m[3][3])
		fmt.Printf("{%d, %d, %d, %d},\n", m[4][0], m[4][1], m[4][2], m[4][3])
		fmt.Println("},")
		return
	}
	lev := m[depth]
	depth++
	m[depth] = []int{lev[1], lev[0], lev[3], lev[2]}
	my(m, depth)
	m[depth] = []int{lev[1], lev[3], lev[0], lev[2]}
	my(m, depth)
	m[depth] = []int{lev[3], lev[2], lev[1], lev[0]}
	my(m, depth)
	m[depth] = []int{lev[3], lev[2], lev[0], lev[1]}
	my(m, depth)

	//fmt.Println(i, ")")
	//i++

}

func main() {
	m := make(map[int][]int, 5)
	m[0] = []int{1, 2, 3, 4}
	my(m, 0)

}
*/
