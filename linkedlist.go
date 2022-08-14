package go_collections

type (
	LinkedList[T comparable] struct {
		head *node[T]
		Size int
	}

	node[T comparable] struct {
		data T
		next *node[T]
	}
)

func (l *LinkedList[T]) traverseToEnd() *node[T] {
	if l.head == nil {
		return nil
	}
	tail := l.head
	for tail.next != nil {
		tail = tail.next
	}
	return tail
}

func (l *LinkedList[T]) Iterator() chan T {
	c := make(chan T)
	go func() {
		temp := l.head
		for temp != nil {
			c <- temp.data
			temp = temp.next
		}
		close(c)
	}()
	return c
}

func (l *LinkedList[T]) Append(elements ...T) {
	if len(elements) == 0 {
		return
	}

	if l.head == nil {
		l.head = &node[T]{data: elements[0]}
		l.Size++
	}

	tail := l.traverseToEnd()
	for _, e := range elements[1:] {
		tail.next = &node[T]{data: e}
		tail = tail.next
		l.Size++
	}
}

func (l *LinkedList[T]) Remove(element T) {
	var prev *node[T]
	for temp := l.head; temp != nil; prev, temp = temp, temp.next {
		if temp.data == element {
			if prev == nil {
				l.head = l.head.next
			} else {
				prev.next = temp.next
			}
			l.Size--
			break
		}
	}
}

func (l *LinkedList[T]) RemoveAtPosition(pos int) (res *T) {
	if pos < 1 || pos > l.Size {
		return nil
	}

	var prev *node[T]
	temp := l.head
	for i := 1; i != pos; i, prev, temp = i+1, temp, temp.next {
	}
	res = &temp.data
	if prev == nil {
		l.head = l.head.next
	} else {
		prev.next = temp.next
	}
	l.Size--
	return
}

func (l *LinkedList[T]) Reverse() {
	if l.Size == 0 {
		return
	}

	var prev, cur *node[T]
	temp := l.head
	for temp != nil {
		cur = temp
		temp = temp.next
		cur.next = prev
		prev = cur
	}
	l.head = cur
}
