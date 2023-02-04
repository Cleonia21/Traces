package research

import (
	"math/rand"
	"time"
)

type fingersParam struct {
	lines   int
	columns int
	matrix  []int
}

func (l *fingersParam) generateLine() []int {
	line := make([]int, l.columns)
	for i := 0; i < l.columns; i++ {
		line[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(line),
		func(i, j int) { line[i], line[j] = line[j], line[i] })
	return line
}
