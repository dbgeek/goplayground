package sortalgo

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestInsertSort(t *testing.T) {
	for _, v := range []int{10, 100, 1000, 10000, 50000} {
		t.Run(fmt.Sprintf("test_%v", v), func(t *testing.T) {
			in := rand.Perm(v)
			var want []int
			want = append(want, in...)
			result := doInsertSort(in)
			sort.Ints(want)
			if !reflect.DeepEqual(result, want) {
				t.Errorf("want: %v \n gott: %v\n", want, result)
			}
		})
	}
}
