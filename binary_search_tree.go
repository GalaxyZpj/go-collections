package go_collections

import constraints "golang.org/x/exp/Constraints"

type (
	BinarySearchTree[T constraints.Ordered] struct {
		root *bstNode[T]
		Size int
	}

	bstNode[T constraints.Ordered] struct {
		data      T
		left      *bstNode[T]
		right     *bstNode[T]
		occurance int
	}
)

func newNode[T constraints.Ordered](element T) *bstNode[T] {
	return &bstNode[T]{data: element, occurance: 1}
}

func findSmallestNode[T constraints.Ordered](root *bstNode[T]) (smallestNode *bstNode[T]) {
	smallestNode = root
	for smallestNode.left.left != nil {
		smallestNode = smallestNode.left
	}
	temp := smallestNode.left
	smallestNode.left = nil
	smallestNode = temp
	return
}

func insert[T constraints.Ordered](root *bstNode[T], element T) *bstNode[T] {
	if root == nil {
		return newNode(element)
	}

	if element > root.data {
		root.right = insert(root.right, element)
	} else if element < root.data {
		root.left = insert(root.left, element)
	} else {
		root.occurance++
	}

	return root
}

func remove[T constraints.Ordered](root *bstNode[T], element T) *bstNode[T] {
	if root == nil {
		return nil
	}

	if element > root.data {
		root.right = remove(root.right, element)
	} else if element < root.data {
		root.left = remove(root.left, element)
	} else if root.occurance > 1 {
		root.occurance--
	} else if root.left == nil {
		return root.right
	} else if root.right == nil {
		return root.left
	} else {
		smallestNode := findSmallestNode(root.right)
		smallestNode.left = root.left
		smallestNode.right = root.right
		root = smallestNode
	}

	return root
}

func search[T constraints.Ordered](root *bstNode[T], element T) *bstNode[T] {
	if root == nil {
		return nil
	}

	if element < root.data {
		return search(root.left, element)
	} else if element > root.data {
		return search(root.right, element)
	}

	return root
}

func traverseInorder[T constraints.Ordered](root *bstNode[T], c chan *bstNode[T]) {
	if root == nil {
		return
	}

	traverseInorder(root.left, c)
	c <- root
	traverseInorder(root.right, c)
}

func traverseInorderReverse[T constraints.Ordered](root *bstNode[T], c chan *bstNode[T]) {
	if root == nil {
		return
	}

	traverseInorder(root.right, c)
	c <- root
	traverseInorder(root.left, c)
}

func (t *BinarySearchTree[T]) Add(elements ...T) {
	if len(elements) == 0 {
		return
	}

	if t.Size == 0 {
		t.root = insert(t.root, elements[0])
		t.Size++
	}

	for _, e := range elements[1:] {
		t.root = insert(t.root, e)
		t.Size++
	}
}

func (t *BinarySearchTree[T]) Delete(elements ...T) {
	if len(elements) == 0 || t.Size == 0 {
		return
	}

	for _, e := range elements {
		t.root = remove(t.root, e)
		t.Size--
	}
}

func (t *BinarySearchTree[T]) Find(element T) *bstNode[T] {
	if t.Size == 0 {
		return nil
	}
	return search(t.root, element)
}

func (t *BinarySearchTree[T]) Iterator() chan *bstNode[T] {
	c := make(chan *bstNode[T])
	go func() {
		traverseInorder(t.root, c)
		close(c)
	}()
	return c
}

func (t *BinarySearchTree[T]) IteratorReverse() chan *bstNode[T] {
	c := make(chan *bstNode[T])
	go func() {
		traverseInorderReverse(t.root, c)
		close(c)
	}()
	return c
}
