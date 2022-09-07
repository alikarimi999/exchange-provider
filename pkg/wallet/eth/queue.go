package eth

import (
	"sort"
	"sync"
)

type queue struct {
	mux   *sync.Mutex
	slice []int
}

func newQ() *queue {
	return &queue{
		mux:   &sync.Mutex{},
		slice: []int{},
	}
}

func (q *queue) push(i int) {
	q.mux.Lock()
	defer q.mux.Unlock()
	for _, n := range q.slice {
		if n == i {
			return
		}
	}
	q.slice = append(q.slice, i)
	sort.Ints(q.slice)
}

func (q *queue) pop() int {
	q.mux.Lock()
	defer q.mux.Unlock()
	if len(q.slice) > 0 {
		i := q.slice[0]
		q.slice = q.slice[1:]
		return i
	}
	return 0
}

func (q *queue) remove(i int) {
	q.mux.Lock()
	defer q.mux.Unlock()

	for a, n := range q.slice {
		if n == i {
			q.slice = append(q.slice[:a], q.slice[a+1:]...)
			break
		}
	}
}

func (q *queue) size() int {
	q.mux.Lock()
	defer q.mux.Unlock()
	return len(q.slice)
}
