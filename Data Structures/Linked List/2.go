// 2. Add Two Numbers

// You are given two non-empty linked lists representing two non-negative integers. 
// The digits are stored in reverse order, and each of their nodes contains a single digit. 
// Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.



/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{}
    temp := result
    for l1 != nil || l2 != nil {
        if l1 != nil {
            temp.Val += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            temp.Val += l2.Val
            l2 = l2.Next
        }
        if temp.Val > 9 {
            temp.Val -= 10
            temp.Next = &ListNode{Val: 1}
        } else if l1 != nil || l2 != nil {
            temp.Next = &ListNode{}
        }
        temp = temp.Next
    }
    return result
}