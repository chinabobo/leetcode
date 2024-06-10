/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package removenthfromend

import "github.com/chinabobo/leetcode/model"

 func removeNthFromEnd(head *model.ListNode, n int) *model.ListNode {
    dummy := &model.ListNode{Val: 0, Next: head}
    left, right := dummy, dummy
    for i := 0; i < n; i++ {
        if right.Next == nil {
            return nil
        }
        right = right.Next
    }
    for right.Next != nil {
        left = left.Next
        right = right.Next
    }

    left.Next = left.Next.Next

    return dummy.Next
}
