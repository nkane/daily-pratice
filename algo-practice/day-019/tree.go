package practice

import (
	"container/list"
	"fmt"
)

type TreeNode[T any] struct {
	Left   *TreeNode[T]
	Right  *TreeNode[T]
	Parent *TreeNode[T]
	Data   T
}

func CreateTestTree() *TreeNode[int] {
	root := &TreeNode[int]{
		Data: 1,
	}
	// left subtree
	root.Left = &TreeNode[int]{
		Data:   2,
		Parent: root,
	}
	root.Left.Left = &TreeNode[int]{
		Data:   4,
		Parent: root.Left,
	}
	root.Left.Left.Left = &TreeNode[int]{
		Data:   8,
		Parent: root.Left.Left,
	}
	root.Left.Right = &TreeNode[int]{
		Data:   5,
		Parent: root.Left,
	}
	// right subtree
	root.Right = &TreeNode[int]{
		Data:   3,
		Parent: root,
	}
	root.Right.Left = &TreeNode[int]{
		Data:   6,
		Parent: root.Right,
	}
	root.Right.Left.Left = &TreeNode[int]{
		Data:   9,
		Parent: root.Right.Left,
	}
	root.Right.Left.Right = &TreeNode[int]{
		Data:   10,
		Parent: root.Right.Left,
	}
	root.Right.Right = &TreeNode[int]{
		Data:   7,
		Parent: root.Right,
	}
	return root
}

/*
Inorder traversal is one of the most used variants of DFS traversal of a tree.
As DFS suggest, we first focus on the depth of the chosen node and then go breadth
at that level. Therefore, we will start from the root node of the tree and go deeper
and deeper into the left subtree in a recursive manner.

When we reach the left-most node, we then will visit that current ndoe and go to
the left-most node of its right subtree, if it exists.

Same steps should be followed in a recursive manner to complete the inorder traversal.
The order of those steps will be similar, in recursive function:
1. Go to the left subtree.
2. Visit node.
3. Go to the right subtree.
*/
func InorderTraversal[T any](root *TreeNode[T]) {
	if root != nil {
		InorderTraversal(root.Left)
		fmt.Print(root.Data, " ")
		InorderTraversal(root.Right)
	}
}

/*
Preorder traversal is another variant of DFS. The atomic operations in a recursive
function are the same inorder traversal but in a different order.

Here, we visit the current node first and then go to the left subtree. After covering
every node of the left subtree, we will move toward the right subtree and visit in a
similar fashion.

Order of the steps will be:
1. Visit node.
2. Go to the left subtree.
3. Go to the right subtree.
*/
func PreorderTraversal[T any](root *TreeNode[T]) {
	if root != nil {
		fmt.Print(root.Data, " ")
		PreorderTraversal(root.Left)
		PreorderTraversal(root.Right)
	}
}

/*
A similar process goes for the postorder traversal, where we visit the left subtree
and right subtree before visiting the current node in recursion.

So, the sequence of the steps will be:
1. Go to the left subtree.
2. Go to the right subtree.
3. Visit the node.
*/
func PostorderTraversal[T any](root *TreeNode[T]) {
	if root != nil {
		PostorderTraversal(root.Left)
		PostorderTraversal(root.Right)
		fmt.Print(root.Data, " ")
	}
}

/*
This is a different traversal than what we have covered above. Level order traversal follows breadth-first
search to visit/modify every node of the tree.

As BFS suggest, the breadth of the tree takes priority first and then moves to depth. In simple words, we will
visit all the nodes present at the same level one-by-one from left to right, and then it moves to the next
level to visit all the nodes of that level.

The implementation  of level order traversal is slightly more challenging than the previous three traversals. We
will use a Queue(FIFO) data structure to implement level order traversal, where after visiting a node, we put its
left and right children to a queue sequentially.

The order of adding children in the queue is important, as we have to traverse left-to-right at the same level.
*/
func LevelorderTraversal[T any](root *TreeNode[T]) {
	if root == nil {
		return
	}
	q := list.New()
	q.PushFront(root)
	for q.Len() > 0 {
		node, ok := q.Remove(q.Front()).(*TreeNode[T])
		if ok && node != nil {
			fmt.Print(node.Data, " ")
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
	}
}
