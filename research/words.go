package research

import (
	"fmt"
	"math/rand"
	"time"
)

type words struct {
	lines   int
	columns int
	matrix  [][]int
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

func (l *words) init(columns int) {
	l.columns = columns
	l.lines = 5
	l.matrix = make([][]int, l.lines)
	for i := 0; i < l.lines; i++ {
		l.matrix[i] = make([]int, l.columns)
	}
}

func (l *words) new(columns int) [][]int {
	l.init(columns)
	var pos coordinate
	var replNum int
	for pos.y < l.lines {
		line := l.generateLine()
		if l.addLine(pos, line, line) {
			pos.y++
		} else {
			pos.y--
			replNum++
		}
		if replNum == 5 {
			pos.y = 0
			pos.x = 0
			replNum = 0
		}
	}
	//l.print()
	return l.matrix
}

// generateLine returns an array of shuffled numbers
func (l *words) generateLine() []int {
	line := make([]int, l.columns)
	for i := 0; i < l.columns; i++ {
		line[i] = i + 1
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(line),
		func(i, j int) { line[i], line[j] = line[j], line[i] })
	return line
}

func (l *words) addLine(pos coordinate, line []int, lineNext []int) bool {
	if len(line) == 0 {
		if (pos.y*pos.x)%l.columns == 0 && len(lineNext) == 0 {
			return true
		}
		return false
	}
	l.matrix[pos.y][pos.x] = line[0]
	line = line[1:]

	if l.check(pos) {
		tmp := l.delete(lineNext, l.matrix[pos.y][pos.x])
		if l.addLine(coordinate{pos.y, pos.x + 1}, tmp, tmp) {
			return true
		}
	}
	return l.addLine(pos, line, lineNext)
}

func (l *words) check(pos coordinate) bool {
	if !l.columnCheck(pos) {
		return false
	}

	res, dist := l.pairCheck(pos)
	if !res {
		if l.columns == 5 && dist > 1 {
			return true
		}
		return false
	}
	return true
}

func (l *words) columnCheck(c coordinate) bool {
	elem := l.matrix[c.y][c.x]
	for c.y > 0 {
		c.y--
		if elem == l.matrix[c.y][c.x] {
			return false
		}
	}
	return true
}

/*

func (l *Words) tripleComboCheck(a coordinate, b coordinate) bool {
	if a.x-2 < 0 || b.x-2 < 0 {
		return true
	}
	if l.matrix[a.y][a.x-2] == l.matrix[b.y][b.x-2] {
		return false
	}
	return true
}

func (l *Words) pairCheck(pos coordinate) bool {
	if pos.x == 0 {
		return true
	}

	for y := pos.y - 1; y >= 0; y-- {
		for x := 1; x < l.columns; x++ {
			if l.matrix[pos.y][pos.x] == l.matrix[y][x] &&
				l.matrix[pos.y][pos.x-1] == l.matrix[y][x-1] {

				if l.tripleCombo {
					if !l.tripleComboCheck(pos, coordinate{y, x}) {
						return false
					}
				}

				l.pair.maxNum--
				if l.pair.maxNum < 0 {
					return false
				}
				if pos.y-y < l.column.banRadius {
					return false
				}
			}
		}
	}
	return true
}
*/

func (l *words) tripleComboCheck(a coordinate, b coordinate) bool {
	if a.x-2 < 0 || b.x-2 < 0 {
		return true
	}
	if l.matrix[a.y][a.x-2] == l.matrix[b.y][b.x-2] {
		return false
	}
	return true
}

func (l *words) pairCheck(pos coordinate) (bool, int) {
	if pos.x == 0 {
		return true, 0
	}

	for y := pos.y - 1; y >= 0; y-- {
		for x := 1; x < l.columns; x++ {
			if l.matrix[pos.y][pos.x] == l.matrix[y][x] &&
				l.matrix[pos.y][pos.x-1] == l.matrix[y][x-1] {

				if !l.tripleComboCheck(pos, coordinate{y, x}) {
					return false, 0
				}

				return false, pos.y - y
			}
		}
	}
	return true, 0
}

// delete removes the element val from a by value
// and returns a new copy of the array.
func (l *words) delete(a []int, val int) []int {
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

func (l *words) print() {
	fmt.Println("words print function:")
	for i := 0; i < len(l.matrix); i++ {
		fmt.Println(l.matrix[i])
	}
	fmt.Println()
}

/*

type levelWords3 struct {
	words
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
	words
}

type levelWords5 struct {
	words
}

type levelWordsM struct {
	words
}

type levelFingers struct {
	words
}

//func (l *words) columnCheck(i int) bool {
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

func (l *words) GetWords() [][]string {
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

func (l *words) GetLetters() [][]string {
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

func (l *words) GetNums() [][]string {
	m := make(map[int]string, l.columns)
	for i := 0; i < l.columns; i++ {
		m[i] = strconv.Itoa(i + 1)
	}
	//fmt.Println("GetNums map:", m)
	return l.getByAlphabet(m)
}

func (l *words) GetFingers() [][]string {
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

func (l *words) getByAlphabet(m map[int]string) [][]string {
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


func main() {
	var l levelWords3
	l.init()
}


func (l *words) generateLine() []int {
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

/*

type pairParam struct {
	maxNum    int //Максимальное количество повторов пар во всей матрице
	banRadius int //Как далеко вообще запретить повторы
}

type columnParam struct {
	maxNum    int //Максимальное количество повторов числа в столбце
	banRadius int //Как далеко вообще запретить повторы
}

type Words struct {
	lines   int
	columns int
	matrix  [][]int

	pair        pairParam
	column      columnParam
	tripleCombo bool //Проверять ли повторы сочетаний из 3 элементов
}

func (l *Words) Init(columns int) {
	l.columns = columns
	l.lines = 5
	l.matrix = make([][]int, l.lines)
	for i := 0; i < l.lines; i++ {
		l.matrix[i] = make([]int, l.columns)
	}
	l.pair = pairParam{1, 0}
	l.column = columnParam{1, 0}
	l.tripleCombo = true
}

func (l *Words) New(columns int) {
	//l.init(columns)
	var pos coordinate
	var replNum int
	for pos.y < l.lines {
		line := l.generateLine()
		l.print()
		if l.addLine(pos, line, line) {
			pos.y++
		} else {
			pos.y--
			replNum++
		}
		if replNum == 5 {
			pos.y = 0
			pos.x = 0
			replNum = 0
		}
	}
	l.print()
}

// generateLine returns an array of shuffled numbers
func (l *Words) generateLine() []int {
	line := make([]int, l.columns)
	for i := 0; i < l.columns; i++ {
		line[i] = i + 1
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(line),
		func(i, j int) { line[i], line[j] = line[j], line[i] })
	return line
}

func (l *Words) addLine(pos coordinate, line []int, lineNext []int) bool {
	if len(line) == 0 {
		if (pos.y*pos.x)%l.columns == 0 && len(lineNext) == 0 {
			return true
		}
		return false
	}
	l.matrix[pos.y][pos.x] = line[0]
	line = line[1:]

	if l.columnCheck(pos) && l.pairCheck(pos) {
		tmp := l.delete(lineNext, l.matrix[pos.y][pos.x])
		if l.addLine(coordinate{pos.y, pos.x + 1}, tmp, tmp) {
			return true
		}
	}
	return l.addLine(pos, line, lineNext)
}

func (l *Words) columnCheck(c coordinate) bool {
	cBuf := c
	for cBuf.y > 0 {
		cBuf.y--
		if l.matrix[c.y][c.x] == l.matrix[cBuf.y][cBuf.x] {
			l.column.maxNum--
			if l.column.maxNum < 0 {
				return false
			}
			if c.y-cBuf.y <= l.column.banRadius {
				return false
			}
		}
	}
	return true
}

func (l *Words) tripleComboCheck(a coordinate, b coordinate) bool {
	if a.x-2 < 0 || b.x-2 < 0 {
		return true
	}
	if l.matrix[a.y][a.x-2] == l.matrix[b.y][b.x-2] {
		return false
	}
	return true
}

func (l *Words) pairCheck(pos coordinate) bool {
	if pos.x == 0 {
		return true
	}

	for y := pos.y - 1; y >= 0; y-- {
		for x := 1; x < l.columns; x++ {
			if l.matrix[pos.y][pos.x] == l.matrix[y][x] &&
				l.matrix[pos.y][pos.x-1] == l.matrix[y][x-1] {

				if l.tripleCombo {
					if !l.tripleComboCheck(pos, coordinate{y, x}) {
						return false
					}
				}

				l.pair.maxNum--
				if l.pair.maxNum < 0 {
					return false
				}
				if pos.y-y < l.column.banRadius {
					return false
				}
			}
		}
	}
	return true
}

func (l *Words) delete(a []int, val int) []int {
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

func (l *Words) print() {
	fmt.Println("Words print function:")
	for i := 0; i < len(l.matrix); i++ {
		fmt.Println(l.matrix[i])
	}
	fmt.Println()
}

/*
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

// delete removes the element val from a by value
// and returns a new copy of the array.


func (l *Words) GetWords() [][]string {
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

func (l *Words) GetLetters() [][]string {
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

func (l *Words) GetNums() [][]string {
	m := make(map[int]string, l.columns)
	for i := 0; i < l.columns; i++ {
		m[i] = strconv.Itoa(i + 1)
	}
	//fmt.Println("GetNums map:", m)
	return l.getByAlphabet(m)
}

func (l *Words) GetFingers() [][]string {
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

func (l *Words) getByAlphabet(m map[int]string) [][]string {
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

func (l *Words) columnCheck(i int) bool {
	elem := l.matrix[i]
	identicalNum := 0
	for i >= l.columns {
		i -= l.columns
		if elem == l.matrix[i] {
			identicalNum++
		}
		if identicalNum > 0 && l.alphabetLen > 4 {
			return false
		} else if identicalNum > 1 {
			return false
		}
	}
	return true
}


func (l *Words) generateLine() []int {
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

func (l *Words) check(i int) bool {
	if l.columnCheck(i) > 0 || l.pairCheck(i) > 0 {
		return false
	}
	return true
}

*/
