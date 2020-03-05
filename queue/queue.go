package queue

import "errors"

type (
	queue []int
)

func (q *queue) que(value int) {
	*q = append(*q, value)
}

func (q *queue) deque() (int, error) {
	if len(*q) == 0 {
		return -1, errors.New("No values in que")
	}
	value, que := (*q)[0], (*q)[1:]
	*q = que
	return value, nil
}
