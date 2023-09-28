package binarysearchtree

type Node struct {
	Left  *Node
	Right *Node
	Value *int
}

type BinarySearchTree struct {
	Root *Node
}

func newNode(val int) Node {
	return Node{Value: &val}
}

func NewBinarySearchTree() BinarySearchTree {
	return BinarySearchTree{}
}
func (bst *BinarySearchTree) Insert(val int) bool {
	newNode := newNode(val)
	if bst.Root == nil {
		bst.Root = &newNode
		return true
	}
	curr := bst.Root
	var assigned bool
	for curr != nil && !assigned {
		if *curr.Value == val {
			return false
		}
		if val < *curr.Value {
			if curr.Left == nil {
				curr.Left = &newNode
				assigned = true
			} else {
				curr = curr.Left
			}
		} else {
			if curr.Right == nil {
				curr.Right = &newNode
				assigned = true
			} else {
				curr = curr.Right
			}
		}
	}
	return true
}

func (bst *BinarySearchTree) Find(val int) (bool, Node) {
	if bst.Root == nil {
		return false, Node{}
	}
	if val == *bst.Root.Value {
		return true, *bst.Root
	}
	var curr = bst.Root
	var found bool
	for curr != nil && !found {
		if val < *curr.Value {
			curr = curr.Left
		} else if val > *curr.Value {
			curr = curr.Right
		} else {
			found = true
		}
	}
	return found, *curr
}

// Example Tree
//
//	         10
//	    6         15
//	3      8           20
//
// Breadth First Search
// [10 6 15 3 8 20 1]
func (bst *BinarySearchTree) BFS() []int {
	var queue []*Node
	var visited []int
	if bst.Root == nil {
		return []int{}
	}
	queue = append(queue, bst.Root)

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		if item.Left != nil {
			queue = append(queue, item.Left)
		}
		if item.Right != nil {
			queue = append(queue, item.Right)
		}
		visited = append(visited, *item.Value)
	}
	return visited
}

// Depth First Search - PreOrder - traverse from the root to the left subtree then to the right subtree.
// [10 6 3 1 8 15 20]
func (bst *BinarySearchTree) DFSPreOrder() []int {
	var visited []int
	var stack []*Node
	if bst.Root == nil {
		return []int{}
	}
	stack = append(stack, bst.Root)

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1] //Pop from stack
		if item.Right != nil {
			stack = append(stack, item.Right)
		}
		if item.Left != nil {
			stack = append(stack, item.Left)
		}
		visited = append(visited, *item.Value)
	}

	return visited
}

// Depth First Search - PostOrder- you traverse from the left subtree to the right subtree then to the root.
// [3,8,6,20,15,10]
func (bst *BinarySearchTree) DFSPostOrder() []int {
	var visited []int
	var stack []*Node
	if bst.Root == nil {
		return []int{}
	}
	stack = append(stack, bst.Root)

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1] //Pop from stack
		if item.Left != nil {
			stack = append(stack, item.Left)
		}
		if item.Right != nil {
			stack = append(stack, item.Right)
		}
		//Prepending here
		visited = append([]int{*item.Value}, visited...)
	}

	return visited

}

// Depth First Search - InOrder - traverse from the left subtree to the root then to the right subtree.
// [3,6,8,10,15,20]
func (bst *BinarySearchTree) DFSInOrder() []int {
	var visited []int
	var stack []*Node
	var stack2 []int

	if bst.Root == nil {
		return []int{}
	}
	stack = append(stack, bst.Root)
	stack2 = append(stack2, *bst.Root.Value)
	curr := bst.Root
	for len(stack) > 0 || curr != nil {
		if curr != nil {
			if curr.Left != nil {
				stack = append(stack, curr.Left)
				stack2 = append(stack2, *curr.Left.Value)
			}
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			visited = append(visited, *curr.Value)
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
			curr = curr.Right
		}

	}
	return visited
}

//[]

//                10
//           6           15
//       3      8    13       20
//
//     1   5  4   9

//[10,6,3,1,5,8,4,9,15,20]

//[3,8,6,20,15,10]
