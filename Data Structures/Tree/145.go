// 145. Binary Tree Postorder Traversal

// Given the root of a binary tree, return the postorder traversal of its nodes' values.

// Input: root = [1,null,2,3]
// Output: [3,2,1]


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    treePostorderTraversal := []int{}
    if root == nil {
        return treePostorderTraversal
    }
    treePostorderTraversal = append(treePostorderTraversal, postorderTraversal(root.Left)...)
    treePostorderTraversal = append(treePostorderTraversal, postorderTraversal(root.Right)...)
    treePostorderTraversal = append(treePostorderTraversal, root.Val)
    return treePostorderTraversal
}