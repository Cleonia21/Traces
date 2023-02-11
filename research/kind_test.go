package research

import (
	"math"
	"testing"
)

func Test_identNums_C(t *testing.T) {
	type args struct {
		matrix [][]int
		c      coordinate
	}
	tests := []struct {
		name        string
		args        args
		wantMinDist int
		wantAmount  int
	}{
		{
			"empty",
			args{
				[][]int{
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
				},
				coordinate{0, 0},
			},
			math.MaxInt,
			0,
		},
		{
			"",
			args{
				[][]int{
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 1},
					{0, 0, 0, 0, 1},
				},
				coordinate{4, 4},
			},
			1,
			1,
		},
		{
			"",
			args{
				[][]int{
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2},
					{0, 0, 0, 0, 2},
					{0, 0, 0, 0, 2},
				},
				coordinate{4, 4},
			},
			1,
			2,
		},
		{
			"",
			args{
				[][]int{
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2},
					{0, 0, 0, 0, 2},
					{0, 0, 0, 0, 1},
				},
				coordinate{4, 4},
			},
			math.MaxInt,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMinDist, gotAmount := identNums_C(tt.args.matrix, tt.args.c)
			if gotMinDist != tt.wantMinDist {
				t.Errorf("identNums_C() gotMinDist = %v, want %v", gotMinDist, tt.wantMinDist)
			}
			if gotAmount != tt.wantAmount {
				t.Errorf("identNums_C() gotAmount = %v, want %v", gotAmount, tt.wantAmount)
			}
		})
	}
}

func Test_identPairs_M(t *testing.T) {
	type args struct {
		matrix [][]int
		pos    coordinate
	}
	tests := []struct {
		name        string
		args        args
		wantMinDist int
		wantAmount  int
	}{
		{
			"",
			args{
				[][]int{
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2},
					{0, 0, 0, 0, 2},
					{0, 0, 0, 0, 1},
				},
				coordinate{4, 4},
			},
			math.MaxInt,
			0,
		},
		{
			"",
			args{
				[][]int{
					{2, 1, 2, 1, 0},
					{0, 2, 1, 0, 0},
					{0, 0, 2, 0, 1},
					{0, 0, 0, 2, 1},
					{0, 0, 0, 2, 1},
				},
				coordinate{4, 4},
			},
			1,
			4,
		},
		{
			"",
			args{
				[][]int{
					{0, 1, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2},
					{0, 0, 0, 0, 2},
					{0, 0, 0, 0, 1},
				},
				coordinate{4, 4},
			},
			4,
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMinDist, gotAmount := identPairs_M(tt.args.matrix, tt.args.pos)
			if gotMinDist != tt.wantMinDist {
				t.Errorf("identPairs_M() gotMinDist = %v, want %v", gotMinDist, tt.wantMinDist)
			}
			if gotAmount != tt.wantAmount {
				t.Errorf("identPairs_M() gotAmount = %v, want %v", gotAmount, tt.wantAmount)
			}
		})
	}
}

func Test_words_check(t *testing.T) {
	type fields struct {
		param param
	}
	type args struct {
		matrix [][]int
		pos    coordinate
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"",
			fields{
				param{
					identPairsAmount_M: 1, identPairsMinDist_M: 3,
				},
			},
			args{
				[][]int{
					{3, 5, 4, 1, 2},
					{1, 3, 2, 4, 5},
					{4, 2, 5, 3, 1},
					{5, 4, 1, 2, 3},
					{2, 1, 3, 5, 4},
				},
				coordinate{0, 0},
			},
			true,
		},
	}
	for _, tt := range tests {
		for i := 0; i < len(tt.args.matrix); i++ {
			for j := 0; j < len(tt.args.matrix[i]); j++ {
				tt.args.pos = coordinate{i, j}
				t.Run(tt.name, func(t *testing.T) {
					w := words{
						param: tt.fields.param,
					}
					if got := w.check(tt.args.matrix, tt.args.pos); got != tt.want {
						t.Errorf("check() = %v, want %v, pos {%v/%v}", got, tt.want, tt.args.pos.y, tt.args.pos.x)
					}
				})
			}
		}
	}
}
