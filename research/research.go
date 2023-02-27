package research

import (
	"fmt"
	"time"
)

type Levels struct {
	Words   [][][]int
	Fingers [][][]int
}

func performance(columnNum int) {
	var fingersLevel level
	var averageVal float64
	for i := 0; i < 100; i++ {
		fingersLevel.init(columnNum)
		fingersLevel.kind = fingersParamGet(columnNum)
		start := time.Now()
		fingersLevel.new()
		duration := time.Since(start)
		averageVal += duration.Seconds()
	}
	fmt.Printf("average time:%v\n", averageVal/100)
}

func GetFingerLevel(columnNum int) ([][]int, float64) {
	var fingersLevel level
	fingersLevel.init(columnNum)
	fingersLevel.kind = fingersParamGet(columnNum)
	start := time.Now()
	mas := fingersLevel.new()
	duration := time.Since(start)
	return mas, duration.Seconds()
}

func Get() Levels {
	var levels Levels
	levels.Words = make([][][]int, 8)
	levels.Fingers = make([][][]int, 8)

	levels.Words[0] = getWords3()
	levels.Words[1] = getWords4()

	var wordsLevel level
	var fingersLevel level

	for i := 2; i < 7; i++ {
		wordsLevel.init(i + 3)
		wordsLevel.kind = wordsParamGet(i + 3)
		levels.Words[i] = wordsLevel.new()
	}

	for i := 0; i < 7; i++ {
		fingersLevel.init(i + 3)
		fingersLevel.kind = fingersParamGet(i + 3)
		levels.Fingers[i] = fingersLevel.new()
	}
	return levels
}

func GetLevel(len int) [][]int {
	if len < 3 {
		fmt.Println("incorrect Level len")
	} else if len == 3 {
		return getWords3()
	} else if len == 4 {
		return getWords4()
	} else {
		var wordsLevel level
		wordsLevel.init(len)
		wordsLevel.kind = wordsParamGet(len)
		return wordsLevel.new()
	}
	return nil
}

func GetWords() [][][]int {
	mas := make([][][]int, 8)

	mas[0] = getWords3()
	mas[1] = getWords4()

	var wordsLevel level

	for i := 2; i < 8; i++ {
		wordsLevel.init(i + 3)
		wordsLevel.kind = wordsParamGet(i + 3)
		mas[i] = wordsLevel.new()
	}
	return mas
}

func GetFingers() [][][]int {
	mas := make([][][]int, 8)

	var fingersLevel level

	for i := 0; i < 7; i++ {
		fingersLevel.init(i + 3)
		fingersLevel.kind = fingersParamGet(i + 3)
		mas[i] = fingersLevel.new()
	}
	return mas
}

type level struct {
	lines   int
	columns int
	matrix  [][]int
	kind    kind
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

func (l *level) init(columns int) {
	l.columns = columns
	l.lines = 5
	l.matrix = make([][]int, l.lines)
	for i := 0; i < l.lines; i++ {
		l.matrix[i] = make([]int, l.columns)
	}
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

func printInfo(m [][]int, pos coordinate) {
	fmt.Println("pos: ", pos.y, pos.x)
	for _, r := range m {
		fmt.Println(r)
	}
}

func (l *level) addLine(pos coordinate, line []int) bool {
	for i := 0; i < len(line); i++ {
		l.matrix[pos.y][pos.x] = line[i]

		if l.kind.check(l.matrix, pos) {
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
	fmt.Println("level print function:")
	for i := 0; i < len(l.matrix); i++ {
		fmt.Println(l.matrix[i])
	}
	fmt.Println()
}

func arrayPrint(a [][]int) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println()
}

func PrintFingersLevWithTime(columnNum int) {
	mas, t := GetFingerLevel(columnNum)
	arrayPrint(mas)
	fmt.Println("t:", t)
}

func PrintAllLevels() {
	levels := Get()

	for i := 0; i < len(levels.Words) && i < len(levels.Fingers); i++ {
		arrayPrint(levels.Words[i])
		arrayPrint(levels.Fingers[i])
	}
}
