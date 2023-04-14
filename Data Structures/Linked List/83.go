// 83. Remove Duplicates from Sorted List

// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. 
// Return the linked list sorted as well.

// Input: head = [1,1,2,3,3]
// Output: [1,2,3]


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
    curr := head
    for curr != nil && curr.Next != nil {
        for curr.Next != nil && curr.Val == curr.Next.Val {
            curr.Next = curr.Next.Next
        }
        curr = curr.Next
    }
    return head
}