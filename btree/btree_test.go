package btree

import (
	"reflect"
	"testing"
)

//var test1 = newBST(10, 5, 15, 5, 2, 14, 22)
//var test2 = newBST(10, 15, 11, 22).delete(10)
//var test3 = newBST(10, 5, 7, 2).delete(10)
//var test4 = newBST(10, 5, 15, 22, 17, 34, 7, 2, 5, 1, 35, 27, 16, 30).delete(22).delete(17)

func TestNewBST(t *testing.T) {
	tt := []struct {
		name       string
		inBtree    []int
		outInOrder []int
	}{
		{
			name:       "test1",
			inBtree:    []int{2, 1, 3},
			outInOrder: []int{1, 2, 3},
		},
		{
			name:       "test2",
			inBtree:    []int{5, 6, 4, 2, 3, 1, 7, 8, 9},
			outInOrder: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:       "test3",
			inBtree:    []int{100, 50, 60, 30, 10, 55, 80, 90, 75, 1},
			outInOrder: []int{1, 10, 30, 50, 55, 60, 75, 80, 90, 100},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tree := newBST(tc.inBtree[:1][0], tc.inBtree[1:])
			result := []int{}
			result = tree.inOrderTraverseRecursive(result)
			if !reflect.DeepEqual(result, tc.outInOrder) {
				t.Fatalf("result: \n %v, want: \n %v", result, tc.outInOrder)
			}
		})
	}
}

func TestBstInsert(t *testing.T) {
	tree := btree{value: 10}
	tree.insertRecursive(9).insertRecursive(11)

	if tree.left.value != 9 {
		t.Errorf("tree.left.value wrong. Got %v, want: 9", tree.left.value)
	}
	if tree.right.value != 11 {
		t.Errorf("tree.right.value wrong. Got %v, want: 11", tree.right.value)
	}
}

func TestDeleteEmptyRightLeafeNodeRecursive(t *testing.T) {
	tree := newBST(50, []int{25, 35, 75, 15})
	want := []int{15, 25, 50, 75}
	tree.deleteRecursive(35)
	result := tree.inOrderTraverseRecursive([]int{})
	if !reflect.DeepEqual(result, want) {
		t.Errorf("result: \n %v, want: \n %v", result, want)
	}
}
func TestDeleteEmptyLeftLeafeNodeRecursive(t *testing.T) {
	tree := newBST(50, []int{25, 35, 75, 15})
	want := []int{25, 35, 50, 75}
	tree.deleteRecursive(15)
	result := tree.inOrderTraverseRecursive([]int{})
	if !reflect.DeepEqual(result, want) {
		t.Errorf("result: \n %v, want: \n %v", result, want)
	}
}

func TestDeleteLeftBothLeftRightLeafeRecursive(t *testing.T) {
	tree := newBST(25, []int{35, 75, 30, 15, 55})
	want := []int{15, 25, 30, 55, 75}
	tree.deleteRecursive(35)
	result := tree.inOrderTraverseRecursive([]int{})
	if !reflect.DeepEqual(result, want) {
		t.Errorf("result: \n %v, want: \n %v", result, want)
	}
	if tree.right.value != 55 {
		t.Errorf("right now should be 55 got: %v", tree.right.value)
	}
}

func TestDeleteParentRecursive(t *testing.T) {
	tree := newBST(25, []int{35, 75})
	want := []int{35, 75}
	tree.deleteRecursive(25)
	result := tree.inOrderTraverseRecursive([]int{})
	if !reflect.DeepEqual(result, want) {
		t.Errorf("result: \n %v, want: \n %v", result, want)
	}
	if tree.right.value != 75 {
		t.Errorf("right now should be 75 got: %v", tree.right.value)
	}
}

func TestMinValueRecursive(t *testing.T) {
	tt := []struct {
		name    string
		inBtree []int
		want    int
	}{
		{
			name:    "test1",
			inBtree: []int{2, 1, 3},
			want:    1,
		},
		{
			name:    "test2",
			inBtree: []int{60, 100, 30, 90, 27, 3, 5, 150, 5, 2, 0},
			want:    0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tree := newBST(tc.inBtree[:1][0], tc.inBtree[1:])
			result := tree.minValueRecursive()
			if result != tc.want {
				t.Errorf("got: %v, want: %v\n", result, tc.want)
			}
		})
	}
}

func TestMaxValueRecursive(t *testing.T) {
	tt := []struct {
		name    string
		inBtree []int
		want    int
	}{
		{
			name:    "test1",
			inBtree: []int{2, 1, 3},
			want:    3,
		},
		{
			name:    "test2",
			inBtree: []int{60, 100, 30, 90, 27, 3, 5, 150, 5, 2, 0},
			want:    150,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tree := newBST(tc.inBtree[:1][0], tc.inBtree[1:])
			result := tree.maxValueRecursive()
			if result != tc.want {
				t.Errorf("got: %v, want: %v\n", result, tc.want)
			}
		})
	}
}
