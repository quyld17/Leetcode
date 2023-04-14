// 19. Remove Nth Node From End of List

// Given the head of a linked list, 
// remove the nth node from the end of the list and return its head.

// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    curr := head
    length := 0
    for curr != nil {
        length += 1
        curr = curr.Next
    }
    if length == n {
        return head.Next
    }
    curr = head
    for curr != nil && curr.Next != nil {
        if length == n+1 {
            curr.Next = curr.Next.Next
            return head
        }
        curr = curr.Next
        length -= 1
    }
    return head
}