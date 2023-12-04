package day4

import (
	"strconv"
	"strings"
)

func mapToIntArray(input []string) ([]int, error) {
	var output []int
	for _, value := range input {
		if value == "" {
			continue
		}
		number, err := strconv.Atoi(strings.TrimSpace(value))
		if err != nil {
			return nil, err
		}
		output = append(output, number)
	}
	return output, nil
}

type queue[T comparable] struct {
	pointer int
	items   []T
}

func (q *queue[T]) push(item T) {
	q.items = append(q.items, item)
}

func (q *queue[T]) next() T {
	if q.pointer > len(q.items)-1 {
		var defaultT T
		return defaultT
	}
	q.pointer++
	return q.items[q.pointer-1]
}

func (q queue[T]) len() int {
	return len(q.items)
}

func (q *queue[T]) hasMore() bool {
	return q.pointer < len(q.items)
}
