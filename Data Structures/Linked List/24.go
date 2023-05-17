// 24. Swap Nodes in Pairs

// Given a linked list, swap every two adjacent nodes and return its head. 
// You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

// Input: head = [1,2,3,4]
// Output: [2,1,4,3]


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    curr := head
    for curr != nil && curr.Next != nil {
        temp := curr.Val
        curr.Val = curr.Next.Val
        curr.Next.Val = temp
        curr = curr.Next.Next
    }
    return head
}