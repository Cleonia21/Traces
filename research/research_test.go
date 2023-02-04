package research

import "testing"

func Test_levelParam_columnCheck(t *testing.T) {
	type fields struct {
		lines                      int
		columns                    int
		matrix                     [][]int
		pair                       pairParam
		column                     columnParam
		repInLine                  bool
		tripleCombo                bool
		consectvDigits             bool
		ascendingDigitsNum         int
		distBetweenIdenticalDigits int
	}
	type args struct {
		c coordinate
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"no repeats",
			fields{
				lines:   5,
				columns: 4,
				matrix: [][]int{
					{1, 2, 3, 4},
					{2, 1, 4, 3},
					{3, 4, 1, 2},
					{4, 3, 2, 1},
					{1, 2, 4, 3},
				},
				pair:                       pairParam{-1, -1},
				column:                     columnParam{0, 0},
				ascendingDigitsNum:         -1,
				distBetweenIdenticalDigits: -1,
			},
			args{coordinate{2, 3}},
			true,
		},
		{
			"first position",
			fields{
				lines:   5,
				columns: 4,
				matrix: [][]int{
					{1, 2, 3, 4},
					{2, 1, 4, 3},
					{3, 4, 1, 2},
					{4, 3, 2, 1},
					{1, 2, 4, 3},
				},
				pair:                       pairParam{-1, -1},
				column:                     columnParam{0, 0},
				ascendingDigitsNum:         -1,
				distBetweenIdenticalDigits: -1,
			},
			args{coordinate{0, 0}},
			true,
		},
		{
			"last position",
			fields{
				lines:   5,
				columns: 4,
				matrix: [][]int{
					{1, 2, 3, 4},
					{2, 1, 4, 3},
					{3, 4, 1, 2},
					{4, 3, 2, 1},
					{1, 2, 4, 3},
				},
				pair:                       pairParam{-1, -1},
				column:                     columnParam{0, 0},
				ascendingDigitsNum:         -1,
				distBetweenIdenticalDigits: -1,
			},
			args{coordinate{0, 3}},
			true,
		},
		{
			"one repeat with maxNum=1",
			fields{
				lines:   5,
				columns: 4,
				matrix: [][]int{
					{1, 2, 3, 2},
					{2, 1, 4, 3},
					{3, 4, 1, 2},
					{4, 3, 2, 1},
					{1, 2, 4, 3},
				},
				pair:                       pairParam{-1, -1},
				column:                     columnParam{1, 0},
				ascendingDigitsNum:         -1,
				distBetweenIdenticalDigits: -1,
			},
			args{coordinate{2, 3}},
			true,
		},
		{
			"one repeat with maxNum=1 but repeat in ban radius",
			fields{
				lines:   5,
				columns: 4,
				matrix: [][]int{
					{1, 2, 3, 2},
					{2, 1, 4, 3},
					{3, 4, 1, 2},
					{4, 3, 2, 1},
					{1, 2, 4, 3},
				},
				pair:                       pairParam{-1, -1},
				column:                     columnParam{1, 2},
				ascendingDigitsNum:         -1,
				distBetweenIdenticalDigits: -1,
			},
			args{coordinate{2, 3}},
			false,
		},
		{
			"one repeat with maxNum=1 with ban radius",
			fields{
				lines:   5,
				columns: 4,
				matrix: [][]int{
					{1, 2, 3, 2},
					{2, 1, 4, 3},
					{3, 4, 1, 2},
					{4, 3, 2, 1},
					{1, 2, 4, 3},
				},
				pair:                       pairParam{-1, -1},
				column:                     columnParam{1, 1},
				ascendingDigitsNum:         -1,
				distBetweenIdenticalDigits: -1,
			},
			args{coordinate{2, 3}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &levelParam{
				lines:                      tt.fields.lines,
				columns:                    tt.fields.columns,
				matrix:                     tt.fields.matrix,
				pair:                       tt.fields.pair,
				column:                     tt.fields.column,
				repInLine:                  tt.fields.repInLine,
				tripleCombo:                tt.fields.tripleCombo,
				consectvDigits:             tt.fields.consectvDigits,
				ascendingDigitsNum:         tt.fields.ascendingDigitsNum,
				distBetweenIdenticalDigits: tt.fields.distBetweenIdenticalDigits,
			}
			if got := l.columnCheck(tt.args.c); got != tt.want {
				t.Errorf("columnCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_levelParam_pairCheck(t *testing.T) {
	type fields struct {
		lines                      int
		columns                    int
		matrix                     [][]int
		pair                       pairParam
		column                     columnParam
		repInLine                  bool
		tripleCombo                bool
		consectvDigits             bool
		ascendingDigitsNum         int
		distBetweenIdenticalDigits int
	}
	type args struct {
		c coordinate
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"no repeats",
			fields{
				lines:   5,
				columns: 4,
				matrix: [][]int{
					{1, 2, 3, 4},
					{2, 1, 4, 3},
					{3, 4, 1, 3},
					{4, 3, 2, 1},
					{1, 2, 4, 3},
				},
				pair:                       pairParam{0, 0},
				column:                     columnParam{-1, -1},
				ascendingDigitsNum:         -1,
				distBetweenIdenticalDigits: -1,
			},
			args{coordinate{2, 3}},
			true,
		},
		{
			"first position",
			fields{
				lines:   5,
				columns: 4,
				matrix: [][]int{
					{1, 2, 3, 4},
					{2, 1, 4, 3},
					{3, 4, 1, 2},
					{4, 3, 2, 1},
					{1, 2, 4, 3},
				},
				pair:                       pairParam{0, 0},
				column:                     columnParam{0, 0},
				ascendingDigitsNum:         -1,
				distBetweenIdenticalDigits: -1,
			},
			args{coordinate{0, 0}},
			true,
		},
		{
			"last position",
			fields{
				lines:   5,
				columns: 4,
				matrix: [][]int{
					{1, 2, 3, 4},
					{2, 1, 4, 3},
					{3, 4, 1, 2},
					{4, 3, 2, 1},
					{1, 2, 4, 3},
				},
				pair:                       pairParam{0, 0},
				column:                     columnParam{-1, -1},
				ascendingDigitsNum:         -1,
				distBetweenIdenticalDigits: -1,
			},
			args{coordinate{0, 3}},
			true,
		},
		{
			"one repeat with maxNum=1",
			fields{
				lines:   5,
				columns: 4,
				matrix: [][]int{
					{1, 2, 1, 4},
					{2, 1, 4, 3},
					{3, 4, 1, 3},
					{4, 3, 2, 1},
					{1, 2, 4, 3},
				},
				pair:                       pairParam{1, 0},
				column:                     columnParam{-1, -1},
				ascendingDigitsNum:         -1,
				distBetweenIdenticalDigits: -1,
			},
			args{coordinate{2, 3}},
			true,
		},
		{
			"one repeat with maxNum=1 but repeat in ban radius",
			fields{
				lines:   5,
				columns: 4,
				matrix: [][]int{
					{1, 2, 1, 2},
					{2, 1, 4, 3},
					{3, 4, 1, 2},
					{4, 3, 2, 1},
					{1, 2, 4, 3},
				},
				pair:                       pairParam{1, 2},
				column:                     columnParam{-1, -1},
				ascendingDigitsNum:         -1,
				distBetweenIdenticalDigits: -1,
			},
			args{coordinate{2, 3}},
			false,
		},
		{
			"one repeat with maxNum=1 and repeat do not in ban radius",
			fields{
				lines:   5,
				columns: 4,
				matrix: [][]int{
					{1, 2, 3, 2},
					{2, 1, 4, 3},
					{3, 4, 1, 2},
					{4, 3, 2, 1},
					{1, 2, 4, 3},
				},
				pair:                       pairParam{1, 1},
				column:                     columnParam{-1, -1},
				ascendingDigitsNum:         -1,
				distBetweenIdenticalDigits: -1,
			},
			args{coordinate{2, 3}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &levelParam{
				lines:                      tt.fields.lines,
				columns:                    tt.fields.columns,
				matrix:                     tt.fields.matrix,
				pair:                       tt.fields.pair,
				column:                     tt.fields.column,
				repInLine:                  tt.fields.repInLine,
				tripleCombo:                tt.fields.tripleCombo,
				consectvDigits:             tt.fields.consectvDigits,
				ascendingDigitsNum:         tt.fields.ascendingDigitsNum,
				distBetweenIdenticalDigits: tt.fields.distBetweenIdenticalDigits,
			}
			if got := l.pairCheck(tt.args.c); got != tt.want {
				t.Errorf("pairCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
