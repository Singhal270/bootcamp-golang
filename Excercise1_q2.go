package main

type node struct {
	data  string
	left  *node
	right *node
}

func preorder(root *node) {
	if root == nil {
		return
	}
	print(root.data)
	preorder(root.left)
	preorder(root.right)
}

func inorder(root *node) {
	if root == nil {
		return
	}
	inorder(root.left)
	print(root.data)
	inorder(root.right)
}

func postorder(root *node) {
	if root == nil {
		return
	}
	postorder(root.left)
	postorder(root.right)
	print(root.data)
}

func main() {
	var root, root1, root2, root3, root4 node
	root.data = "+"
	root1.data = "a"
	root2.data = "-"
	root3.data = "b"
	root4.data = "c"
	root.left = &root1
	root.right = &root2
	root2.left = &root3
	root2.right = &root4

	preorder(&root)
	println("")
	postorder(&root)
	println("")
	inorder(&root)
	println("")

}
