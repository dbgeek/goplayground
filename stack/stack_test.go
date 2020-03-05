package stack

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	s := stack{}
	s.push(1)
	s.push(2)
	s.push(4)

	want := stack{1, 2, 4}
	if !reflect.DeepEqual(want, s) {
		t.Fatalf("result: \n %v, want: \n %v", s, want)
	}

}

func TestStackPushPop(t *testing.T) {
	s := stack{}
	s.push(1)
	_, _ = s.pop()
	s.push(4)
	s.push(5)
	s.push(6)
	_, _ = s.pop()
	s.push(7)

	want := stack{4, 5, 7}
	if !reflect.DeepEqual(want, s) {
		t.Fatalf("result: \n %v, want: \n %v", s, want)
	}

}

func TestEmptyStackPop(t *testing.T) {
	s := stack{}

	_, err := s.pop()

	if err == nil {
		t.Fatalf("Should return error when stack empty")
	}
}
