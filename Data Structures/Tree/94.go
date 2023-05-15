// 94. Binary Tree Inorder Traversal

// Given the root of a binary tree, return the inorder traversal of its nodes' values.

// Input: root = [1,null,2,3]
// Output: [1,3,2]


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
    treeInorderTraversal := []int{}
    if root == nil {
        return treeInorderTraversal
    }
    treeInorderTraversal = append(treeInorderTraversal, inorderTraversal(root.Left)...)
    treeInorderTraversal = append(treeInorderTraversal, root.Val)
    treeInorderTraversal = append(treeInorderTraversal, inorderTraversal(root.Right)...)
    return treeInorderTraversal 
}