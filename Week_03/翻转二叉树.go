func invertTree(root *TreeNode) *TreeNode {

	//return
	if root == nil {
		//process
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	//next
	invertTree(root.Left)
	invertTree(root.Right)

	//clear
	return root
}