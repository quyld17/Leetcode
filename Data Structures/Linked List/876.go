// 876. Middle of the Linked List

// Given the head of a singly linked list, return the middle node of the linked list.
// If there are two middle nodes, return the second middle node.

// Input: head = [1,2,3,4,5]
// Output: [3,4,5]
// Explanation: The middle node of the list is node 3.

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func middleNode(head *ListNode) *ListNode {
    count := 0
    originalHead := head
    for head != nil {
        head = head.Next
        count++
    } 
    run := 0
    for run != (count/2) {
        originalHead = originalHead.Next
        run++
    }
    return originalHead
}