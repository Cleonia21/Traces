package levels

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Level struct {
	lines         int
	columns       int
	matrix        []int
	pseudoColumns bool
}

func (l *Level) GetWords() [][]string {
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

func (l *Level) GetLetters() [][]string {
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

func (l *Level) GetNums() [][]string {
	m := make(map[int]string, l.columns)
	for i := 0; i < l.columns; i++ {
		m[i] = strconv.Itoa(i + 1)
	}
	//fmt.Println("GetNums map:", m)
	return l.getByAlphabet(m)
}

func (l *Level) GetFingers() [][]string {
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

func (l *Level) getByAlphabet(m map[int]string) [][]string {
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

func (l *Level) New(columns int) {
	//fmt.Println("New call with columns ", columns)
	l.init(columns)

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

func (l *Level) addLine(index int, line []int, lineNext []int) bool {
	if len(line) == 0 {
		if index%l.columns == 0 && len(lineNext) == 0 {
			return true
		}
		return false
	}
	l.matrix[index] = line[0]
	line = line[1:]

	if l.columnCheck(index) && l.pairCheck(index) {
		tmp := l.delete(lineNext, l.matrix[index])
		if l.addLine(index+1, tmp, tmp) {
			return true
		}
	}
	return l.addLine(index, line, lineNext)
}

func (l *Level) init(columns int) {
	if columns == 3 {
		l.pseudoColumns = true
		l.columns = 4
	} else {
		l.pseudoColumns = false
		l.columns = columns
	}
	l.lines = 5
	l.matrix = make([]int, l.lines*l.columns)
}

func (l *Level) columnCheck(i int) bool {
	elem := l.matrix[i]
	identicalNum := 0
	for i >= l.columns {
		i -= l.columns
		if elem == l.matrix[i] {
			identicalNum++
		}
		if identicalNum > 0 && l.columns > 4 {
			return false
		} else if identicalNum > 1 {
			return false
		}
	}
	return true
}

func (l *Level) pairCheck(i int) bool {
	if i < 0 || i > l.columns*l.lines {
		return false
	}
	identicalNum := 0
	for j := i - 2; j > 0; j-- {
		if j%l.columns != 0 && l.matrix[j] == l.matrix[i] &&
			l.matrix[j-1] == l.matrix[i-1] {
			identicalNum++
		}
		if identicalNum > 0 && l.columns > 5 {
			return false
		} else if identicalNum > 1 {
			return false
		}
	}
	return true
}

// generateLine returns an array of shuffled numbers
func (l *Level) generateLine() []int {
	line := make([]int, l.columns)
	for i := 0; i < l.columns; i++ {
		line[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(line),
		func(i, j int) { line[i], line[j] = line[j], line[i] })
	return line
}

// delete removes the element val from a by value
// and returns a new copy of the array.
func (l *Level) delete(a []int, val int) []int {
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

func (l *Level) print() {
	fmt.Println("Level print function:")
	for i := 0; i < len(l.matrix); i++ {
		if i%l.columns == 0 && i != 0 {
			fmt.Println()
		}
		fmt.Print(l.matrix[i], " ")
	}
	fmt.Println()
}
