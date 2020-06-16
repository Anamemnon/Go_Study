package main

import "testing"

type testData struct {
	list []int
	want int
}

func TestSum(t *testing.T) {
	tests := [...]testData{
		{
			list: []int{1, 2},
			want: 3,
		},
		{
			list: []int{1},
			want: 1,
		},
	}

	for _, test := range tests {
		got := sum(test.list...)
		if got != test.want {
			t.Errorf("sum(%#v...) = \"%d\", want \"%d\"", test.list, got, test.want)
		}
	}


}
