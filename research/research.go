package research

type Levels struct {
	Words [][][]int
}

func Get() Levels {
	var levels Levels
	levels.Words = make([][][]int, 8)

	levels.Words[0] = getWords3()
	levels.Words[1] = getWords4()

	var l words
	for i := 2; i < 8; i++ {
		levels.Words[i] = l.new(i + 3)
	}
	return levels
}
