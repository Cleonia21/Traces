package research

import (
	"fmt"
	"sort"
	"time"
)

type Levels struct {
	Words   [][][]int
	Fingers [][][]int
	Visual  [][][]int
}

func performance(columnNum int) {
	var fingersLevel level
	var averageVal float64
	for i := 0; i < 100; i++ {
		fingersLevel.init(columnNum, nil, fingersParamGet(columnNum))
		start := time.Now()
		fingersLevel.new()
		duration := time.Since(start)
		averageVal += duration.Seconds()
	}
	fmt.Printf("average time:%v\n", averageVal/100)
}

func GetFingerLevel(columnNum int) ([][]int, float64) {
	var fingersLevel level
	fingersLevel.init(columnNum, nil, fingersParamGet(columnNum))
	start := time.Now()
	mas := fingersLevel.new()
	duration := time.Since(start)
	return mas, duration.Seconds()
}

func Get() Levels {
	var levels Levels
	levels.Words = GetWords()
	levels.Fingers = GetFingers()
	levels.Visual = GetEmptyLevels()
	return levels
}

func GetEmptyLevels() [][][]int {
	mas := make([][][]int, 8)

	var wordsLevel level

	for i := 0; i < 8; i++ {
		wordsLevel.init(i+3, nil, nil)
		mas[i] = wordsLevel.matrix
	}

	return mas
}

func GetWords() [][][]int {
	mas := make([][][]int, 8)

	mas[0] = getWords3()
	mas[1] = getWords4()

	var wordsLevel level

	for i := 2; i < 8; i++ {
		wordsLevel.init(i+3, nil, wordsParamGet(i+3))
		mas[i] = wordsLevel.new()
	}
	return mas
}

func GetFingers() [][][]int {
	mas := make([][][]int, 8)

	var fingersLevel level

	for i := 0; i < 8; i++ {
		fingersLevel.init(i+3, nil, fingersParamGet(i+3))
		mas[i] = fingersLevel.new()
	}
	return mas
}

func GenerateAllPermutation(columns int) {
	var fingersLevel level
	fingersLevel.init(columns, nil, fingersParamGet(columns))
	fingersLevel.allPermutation()
}

type level struct {
	lines       int
	columns     int
	matrix      [][]int
	beforeTrios map[int]map[int]map[int]int
	kind        kind
}

type coordinate struct {
	y int // |
	x int //---->
	/*
		||||| y=0
		||||| y=1
		||||| y=2
		x=0 x=1 x=2 ------>
	*/
}

func (l *level) fillBeforeMatrixTrios(beforeMatrix [][]int) {
	for i := 0; i < len(beforeMatrix); i++ {
		for j := 0; j < len(beforeMatrix[i])-2; j++ {
			l.beforeTrios[beforeMatrix[i][j]][beforeMatrix[i][j]][beforeMatrix[i][j]]++
		}
	}
	//fmt.Println(l.beforeTrios)
}

func (l *level) init(columns int, beforeMatrix [][]int, k kind) {
	l.columns = columns
	l.lines = 5
	l.matrix = make([][]int, l.lines)
	for i := 0; i < l.lines; i++ {
		l.matrix[i] = make([]int, l.columns)
	}
	//l.fillBeforeMatrixTrios(beforeMatrix)
	l.kind = k
}

func (l *level) new() [][]int {
	var pos coordinate
	var replNum int
	for pos.y < l.lines {
		tmp := l.kind.newRandLine(l.columns)
		//fmt.Println("Want add line", tmp, "Matrix\n", l.matrix)
		if l.addLine(pos, tmp) {
			pos.y++
		} else {
			pos.y--
			replNum++
		}
		if replNum == 5 {
			//fmt.Println(pos)
			replNum = 0
			pos.y = 0
			pos.x = 0
		}
	}
	//l.print()
	return l.matrix
}

func (l *level) addLine(pos coordinate, line []int) bool {
	for i := 0; i < len(line); i++ {
		l.matrix[pos.y][pos.x] = line[i]

		if l.kind.check(l.matrix, pos, l.beforeTrios) {
			if pos.x == l.columns-1 {
				return true
			}

			tmp := l.delete(line, line[i])
			if l.addLine(coordinate{pos.y, pos.x + 1}, tmp) {
				return true
			}
		}
	}
	return false
}

// delete removes the element val from a by value
// and returns a new copy of the array.
func (l *level) delete(a []int, val int) []int {
	newA := make([]int, len(a))
	i := 0
	for _, v := range a {
		if v != val {
			newA[i] = v
			i++
		} else {
			val = -1
		}
	}
	newA = newA[:i]
	return newA
}

func (l *level) print() {
	//fmt.Println("level print function:")
	for i := 0; i < len(l.matrix); i++ {
		fmt.Println(l.matrix[i])
	}
	fmt.Println()
}

func ArrayPrint(a [][]int) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println()
}

func PrintFingersLevWithTime(columnNum int) {
	mas, t := GetFingerLevel(columnNum)
	ArrayPrint(mas)
	fmt.Println("t:", t)
}

func PrintAllLevels() {
	levels := Get()

	for i := 0; i < len(levels.Words) && i < len(levels.Fingers); i++ {
		ArrayPrint(levels.Words[i])
		ArrayPrint(levels.Fingers[i])
	}
}

func printInfo(m [][]int, pos coordinate) {
	fmt.Println("pos: ", pos.y, pos.x)
	for _, r := range m {
		fmt.Println(r)
	}
}

func (l *level) allPermutation() {
	var base = make([]int, 5)
	for i := 0; i < 5; i++ {
		base[i] = i + 1
	}
	var additionalBase = make([]int, l.columns-5)
	for i := 0; i < l.columns-5; i++ {
		additionalBase[i] = i + 1
	}
	base = append(base, additionalBase...)
	sort.Ints(base)
	copy(l.matrix[0], base)

	var y int
	var num int
	var allNum int

	for y >= 0 {
		allNum++
		if l.checkLine(y) {
			if y == l.lines-1 {
				l.print()
				num++
			} else {
				y++
				copy(l.matrix[y], base)
				continue
			}
		}
		for y >= 0 && nextPermutation(l.matrix[y]) == false {
			y--
		}
	}
	fmt.Printf("These are all permutations (%d)(%d)\n", allNum, num)
}

func (l *level) checkLine(y int) bool {
	var pos = coordinate{y, 0}
	for ; pos.x < l.columns; pos.x++ {
		if l.kind.check(l.matrix, pos, l.beforeTrios) == false {
			return false
		}
	}
	return true
}

func TestNextPermutation(columns int) {
	var base = make([]int, 5)
	for i := 0; i < 5; i++ {
		base[i] = i + 1
	}
	var additionalBase = make([]int, columns-5)
	for i := 0; i < columns-5; i++ {
		additionalBase[i] = i + 1
	}
	base = append(base, additionalBase...)
	sort.Ints(base)

	for {
		fmt.Println(base)
		if !nextPermutation(base) {
			break
		}
	}
}

func nextPermutation(b []int) bool {
	l := len(b)
	result := false

	for i := l - 1; i > 0; i-- {
		if b[i-1] < b[i] {
			pivot := i
			for j := pivot; j < l; j++ {
				if b[j] <= b[pivot] && b[i-1] < b[j] {
					pivot = j
				}
			}

			b[i-1], b[pivot] = b[pivot], b[i-1]

			for j := l - 1; i < j; i, j = i+1, j-1 {
				b[i], b[j] = b[j], b[i]
			}
			result = true
			break
		}
	}
	if result == false {
		return nextBase(b)
	}
	return result
}

func TestNextBase(columns int) {
	var base = make([]int, 5)
	for i := 0; i < 5; i++ {
		base[i] = i + 1
	}
	var additionalBase = make([]int, columns-5)
	for i := 0; i < columns-5; i++ {
		additionalBase[i] = i + 1
	}
	base = append(base, additionalBase...)
	sort.Ints(base)

	for {
		fmt.Println(base)
		if !nextBase(base) {
			break
		}
	}
}

func nextBase(b []int) bool {
	sort.Ints(b)
	biggestPair := 6

	for i := len(b) - 1; i > 0; {
		if b[i] == b[i-1] && b[i] == biggestPair-1 {
			biggestPair--
			i -= 2
		} else {
			break
		}
	}

	for i := len(b) - 1; i > 0; i-- {
		if b[i] == b[i-1] && b[i] < biggestPair {
			b[i]++
			return true
		}
	}
	return false
}
