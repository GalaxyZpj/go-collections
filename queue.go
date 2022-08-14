package go_collections

type (
	Queue[T comparable] struct {
		first *queueNode[T]
		last *queueNode[T]
		Size int
	}
	
	queueNode[T comparable] struct {
		data T
		next *queueNode[T]
	}
)

func (q *Queue[T]) Add(elements ...T) {
	if len(elements) == 0 {
		return
	}

	if q.first == nil {
		q.first = &queueNode[T]{data: elements[0]}
		q.last = q.first
		q.Size++
	}

	for _, e := range elements[1:] {
		q.last.next = &queueNode[T]{data: e}
		q.last = q.last.next
		q.Size++
	}
}

func (q *Queue[T]) Remove() (res *T) {
	if q.Size == 0 {
		return nil
	}

	res = &q.first.data
	q.first = q.first.next
	q.Size--
	if q.Size == 0 {
		q.last = nil
	}
	return
}

func (q *Queue[T]) Peek() (res *T) {
	if q.Size == 0 {
		return nil
	}
	return &q.first.data
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size == 0
}
