package levels

import (
	"reflect"
	"testing"
)

func TestLevel_New(t *testing.T) {
	type fields struct {
		lines    int
		columns  int
		alphabet []int
		matrix   []int
	}
	type args struct {
		columns int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Level{
				lines:    tt.fields.lines,
				columns:  tt.fields.columns,
				alphabet: tt.fields.alphabet,
				matrix:   tt.fields.matrix,
			}
			l.New(tt.args.columns)
		})
	}
}

func TestLevel_addLine(t *testing.T) {
	type fields struct {
		lines    int
		columns  int
		alphabet []int
		matrix   []int
	}
	type args struct {
		index    int
		line     []int
		lineNext []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Level{
				lines:    tt.fields.lines,
				columns:  tt.fields.columns,
				alphabet: tt.fields.alphabet,
				matrix:   tt.fields.matrix,
			}
			if got := l.addLine(tt.args.index, tt.args.line, tt.args.lineNext); got != tt.want {
				t.Errorf("addLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevel_columnCheck(t *testing.T) {
	type fields struct {
		lines    int
		columns  int
		alphabet []int
		matrix   []int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Level{
				lines:    tt.fields.lines,
				columns:  tt.fields.columns,
				alphabet: tt.fields.alphabet,
				matrix:   tt.fields.matrix,
			}
			if got := l.columnCheck(tt.args.i); got != tt.want {
				t.Errorf("columnCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevel_generateLine(t *testing.T) {
	type fields struct {
		lines    int
		columns  int
		alphabet []int
		matrix   []int
	}
	tests := []struct {
		name     string
		fields   fields
		wantLine []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Level{
				lines:    tt.fields.lines,
				columns:  tt.fields.columns,
				alphabet: tt.fields.alphabet,
				matrix:   tt.fields.matrix,
			}
			if gotLine := l.generateLine(); !reflect.DeepEqual(gotLine, tt.wantLine) {
				t.Errorf("generateLine() = %v, want %v", gotLine, tt.wantLine)
			}
		})
	}
}

func TestLevel_init(t *testing.T) {
	type fields struct {
		lines    int
		columns  int
		alphabet []int
		matrix   []int
	}
	type args struct {
		columns int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Level{
				lines:    tt.fields.lines,
				columns:  tt.fields.columns,
				alphabet: tt.fields.alphabet,
				matrix:   tt.fields.matrix,
			}
			l.init(tt.args.columns)
		})
	}
}

func TestLevel_pairCheck(t *testing.T) {
	type Fields struct {
		lines    int
		columns  int
		alphabet []int
		matrix   []int
	}
	tests := []struct {
		name   string
		fields Fields
		i      int
		want   bool
	}{
		{
			"default",
			Fields{
				2,
				3,
				[]int{}, //not use
				[]int{
					1, 2, 4,
					4, 1, 3,
				},
			},
			0,
			true,
		},
		{
			"matrix error",
			Fields{
				2,
				3,
				[]int{}, //not use
				[]int{
					1, 2, 4,
					4, 1, 2,
				},
			},
			5,
			false,
		},
		{
			"iterator max error",
			Fields{
				2,
				3,
				[]int{}, //not use
				[]int{
					1, 2, 4,
					4, 1, 3,
				},
			},
			10,
			false,
		},
		{
			"iterator min error",
			Fields{
				2,
				3,
				[]int{}, //not use
				[]int{
					1, 2, 4,
					4, 1, 3,
				},
			},
			-5,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Level{
				lines:   tt.fields.lines,
				columns: tt.fields.columns,
				matrix:  tt.fields.matrix,
			}
			if got := l.pairCheck(tt.i); got != tt.want {
				t.Errorf("pairCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		n    []int
		item int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"default",
			args{
				[]int{1, 2, 3, 4, -5, -17, 0, 100, 21, 934, -1},
				1,
			},
			[]int{2, 3, 4, -5, -17, 0, 100, 21, 934, -1},
		},
		{
			"empty array",
			args{
				[]int{},
				1,
			},
			[]int{},
		},
		{
			"item not found",
			args{
				[]int{1, 2, 3, 4, -5, -17, 0, 100, 21, 934, -1},
				13,
			},
			[]int{1, 2, 3, 4, -5, -17, 0, 100, 21, 934, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.args.n, tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
