package queue

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	q := queue{}
	q.que(1)
	q.que(2)
	q.que(4)

	want := queue{1, 2, 4}
	if !reflect.DeepEqual(want, q) {
		t.Errorf("result: \n %v, want: \n %v", q, want)
	}

}

func TestQueueDeque(t *testing.T) {
	q := queue{}
	q.que(1)
	_, _ = q.deque()
	q.que(4)
	q.que(5)
	q.que(6)
	_, _ = q.deque()
	q.que(7)

	want := queue{5, 6, 7}
	if !reflect.DeepEqual(want, q) {
		t.Errorf("result: \n %v, want: \n %v", q, want)
	}

}

func TestEmptyQueue(t *testing.T) {
	q := queue{}

	_, err := q.deque()

	if err == nil {
		t.Errorf("Should return error when queue empty")
	}
}

func TestQueueOneValue(t *testing.T) {
	q := queue{}
	q.que(1)

	v, err := q.deque()

	if err == nil && v != 1 {
		t.Errorf("should return no error and 1")
	}
}
