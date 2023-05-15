// 144. Binary Tree Preorder Traversal

// Given the root of a binary tree, return the preorder traversal of its nodes' values.

// Input: 
// root = [1,2,3,4,5,6,7,8,9,10]
// Output: [1,2,4,8,9,5,10,3,6,7]


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    treePreorderTraversal := []int{}
    if root == nil {
        return treePreorderTraversal 
    }
    treePreorderTraversal = append(treePreorderTraversal, root.Val)
    treePreorderTraversal = append(treePreorderTraversal, preorderTraversal(root.Left)...)
    treePreorderTraversal = append(treePreorderTraversal, preorderTraversal(root.Right)...)
    return treePreorderTraversal 
}