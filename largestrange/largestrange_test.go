package largestrange

import (
	"reflect"
	"testing"
)

func TestLargestrange(t *testing.T) {
	tt := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "test1",
			in:   []int{2, 1, 3},
			want: []int{1, 3},
		},
		{
			name: "test2",
			in:   []int{1, 11, 4, 5, 9, 14, 10, 12, 13, 2, 16, 17},
			want: []int{9, 14},
		},
		{
			name: "test3",
			in:   []int{1, 19, 11, 100, 101, 4, 500, 50, 5, 34, 9, 33, 20, 14, 66, 10, 12, 77, 56, 52, 13, 2, 16, 17, 55, 110, 240},
			want: []int{9, 14},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := largestRange(tc.in)
			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("result: \n %v, want: \n %v", result, tc.want)
			}

		})
	}
}
