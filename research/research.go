package research

import (
	"fmt"
	"math/rand"
	"time"
)

type Levels struct {
	l3  [][]int
	l4  [][]int
	l5  [][]int
	l6  [][]int
	l7  [][]int
	l8  [][]int
	l9  [][]int
	l10 [][]int
}

func GetResearch() Levels {
	var l Levels
	l.l3 = getWords3()
	l.l4 = getWords4()
	return l
}

type LevelParam struct {
	lines       int
	columns     int
	alphabetLen int
	matrix      []int
}

func (l *LevelParam) init(columns int) {
	l.alphabetLen = columns
	l.columns = columns
	l.lines = 5
	l.matrix = make([]int, l.lines*l.columns)
}

func (l *LevelParam) new() {
	var i int
	var replNum int
	for i < l.lines*l.columns {
		line := l.generateLine()
		if l.addLine(i, line, line) {
			i += l.columns
		} else {
			i -= l.columns
			replNum++
		}
		if replNum == 5 {
			i = 0
			replNum = 0
		}
	}
	//l.print()
}

// generateLine returns an array of shuffled numbers
func (l *LevelParam) generateLine() []int {
	line := make([]int, l.alphabetLen)
	for i := 0; i < l.alphabetLen; i++ {
		line[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(line),
		func(i, j int) { line[i], line[j] = line[j], line[i] })
	return line
}

func (l *LevelParam) addLine(index int, line []int, lineNext []int) bool {
	if len(line) == 0 {
		if index%l.columns == 0 && len(lineNext) == 0 {
			return true
		}
		return false
	}
	l.matrix[index] = line[0]
	line = line[1:]

	if l.check(index) {
		tmp := l.delete(lineNext, l.matrix[index])
		if l.addLine(index+1, tmp, tmp) {
			return true
		}
	}
	return l.addLine(index, line, lineNext)
}

func (l *LevelParam) check(i int) bool {
	if l.columnCheck(i) > 0 || l.pairCheck(i) > 0 {
		return false
	}
	return true
}

func (l *LevelParam) columnCheck(i int) int {
	elem := l.matrix[i]
	identicalNum := 0
	for i >= l.columns {
		i -= l.columns
		if elem == l.matrix[i] {
			identicalNum++
		}
	}
	return identicalNum
}

func (l *LevelParam) pairCheck(i int) int {
	identicalNum := 0
	for j := i - 2; j > 0; j-- {
		if j%l.columns != 0 && l.matrix[j] == l.matrix[i] &&
			l.matrix[j-1] == l.matrix[i-1] {
			identicalNum++
		}
	}
	return identicalNum
}

type levelWords3 struct {
	LevelParam
}

func (l *levelWords3) init() {
	l.alphabetLen = 4
	l.columns = 3
	l.lines = 5
	l.matrix = make([]int, l.lines*l.columns)
}

func (l *levelWords3) check(i int) bool {
	if l.columnCheck(i) > 1 || l.pairCheck(i) > 1 {
		return false
	}
	return true
}

func (l *levelWords3) preColumnCheck(i int) bool {
	elem := l.matrix[i]
	if i >= l.columns {
		i -= l.columns
		if elem == l.matrix[i] {
			return false
		}
	}
	return true
}

func (l *levelWords3) trioCheck(i int) bool {
	if (i+1)%3 != 0 {
		return true
	}
	for j := i/3 - 1; j >= 0; j-- {
		if j%l.columns != 0 && l.matrix[j] == l.matrix[i] &&
			l.matrix[j-1] == l.matrix[i-1] {
			//identicalNum++
		}
	}
	return true
}

type levelWords4 struct {
	LevelParam
}

type levelWords5 struct {
	LevelParam
}

type levelWordsM struct {
	LevelParam
}

type levelFingers struct {
	LevelParam
}

//func (l *LevelParam) columnCheck(i int) bool {
//	elem := l.matrix[i]
//	identicalNum := 0
//	for i >= l.columns {
//		i -= l.columns
//		if elem == l.matrix[i] {
//			identicalNum++
//		}
//		if identicalNum > 0 && l.alphabetLen > 4 {
//			return false
//		} else if identicalNum > 1 {
//			return false
//		}
//	}
//	return true
//}

// delete removes the element val from a by value
// and returns a new copy of the array.
func (l *LevelParam) delete(a []int, val int) []int {
	newA := make([]int, len(a))
	i := 0
	for _, v := range a {
		if v != val {
			newA[i] = v
			i++
		}
	}
	newA = newA[:i]
	return newA
}

func (l *LevelParam) print() {
	fmt.Println("LevelParam print function:")
	for i := 0; i < len(l.matrix); i++ {
		if i%l.columns == 0 && i != 0 {
			fmt.Println()
		}
		fmt.Print(l.matrix[i], " ")
	}
	fmt.Println()
}

/*
func (l *LevelParam) GetWords() [][]string {
	m := map[int]string{
		0: "дом",
		1: "кот",
		2: "лес",
		3: "стол",
		4: "звон",
		5: "брат",
		6: "мост",
		7: "час",
		8: "вол",
		9: "рев",
	}
	return l.getByAlphabet(m)
}

func (l *LevelParam) GetLetters() [][]string {
	m := map[int]string{
		0: "д",
		1: "к",
		2: "л",
		3: "с",
		4: "з",
		5: "б",
		6: "м",
		7: "ч",
		8: "в",
		9: "р",
	}
	return l.getByAlphabet(m)
}

func (l *LevelParam) GetNums() [][]string {
	m := make(map[int]string, l.columns)
	for i := 0; i < l.columns; i++ {
		m[i] = strconv.Itoa(i + 1)
	}
	//fmt.Println("GetNums map:", m)
	return l.getByAlphabet(m)
}

func (l *LevelParam) GetFingers() [][]string {
	m := make(map[int]string, l.columns)
	for i := 0; i < l.columns; i++ {
		if i > 4 {
			m[i] = strconv.Itoa(i - 5 + 1)
		} else {
			m[i] = strconv.Itoa(i + 1)
		}
	}
	//fmt.Println("GetNums map:", m)
	return l.getByAlphabet(m)
}

func (l *LevelParam) getByAlphabet(m map[int]string) [][]string {
	strArray := make([][]string, l.lines)
	for i := 0; i < l.lines; i++ {
		strArray[i] = make([]string, l.columns)
		for j := 0; j < l.columns; j++ {
			if l.pseudoColumns && j+1 == l.columns {
				strArray[i][j] = ""
			} else {
				strArray[i][j] += m[l.matrix[i*l.columns+j]]
			}
		}
	}
	return strArray
}
*/

func main() {
	var l levelWords3
	l.init()
}

/*
func (l *LevelParam) generateLine() []int {
	line := make([]int, l.alphabetLen)
	for i := 0; i < l.alphabetLen; i++ {
		line[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	for true {
		rand.Shuffle(len(line), func(i, j int) { line[i], line[j] = line[j], line[i] })
		for i := 0; i < l.alphabetLen; i++ {
			if i < l.alphabetLen-2 {
				if line[i] < line[i+1] && line[i+1] < line[i+2] {
					continue
				} else if line[i] > line[i+1] && line[i+1] > line[i+2] {
					continue
				}
				break
			}
		}
	}

	if l.columns != l.alphabetLen {
		return line[:l.columns]
	} else {
		return line
	}
}

*/
