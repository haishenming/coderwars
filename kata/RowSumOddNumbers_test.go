package kata

import "testing"

func TestRowSumOddNumbers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{1}, 1},
		{"1", args{2}, 8},
		{"1", args{4}, 13 + 15 + 17 + 19},
		{"1", args{13}, 2197},
		{"1", args{19}, 6859},
		{"1", args{41}, 68921},
		{"1", args{42}, 74088},
		{"1", args{74}, 405224},
		{"1", args{86}, 636056},
		{"1", args{93}, 804357},
		{"1", args{101}, 1030301},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RowSumOddNumbers(tt.args.n); got != tt.want {
				t.Errorf("RowSumOddNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
