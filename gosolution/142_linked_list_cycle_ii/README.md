README

![](https://assets.leetcode-cn.com/solution-static/142/142_fig1.png)

a + n ( b + c ) + b - 1 = 2(a + b)

a = (n - 1) (b + c) + c - 1

快慢指针，快慢相遇之后，慢指针回到头，快指针向后一个节点（由于一开始快指针是从`fast := head.Next`开始的），快慢指针步调一致一起移动，相遇点即为入环点。
第一次相交后，快指针需要从下一个节点开始和头指针一起匀速移动

另外一种方式是 fast=head,slow=head
```go
func detectCycle(head *ListNode) *ListNode {
    // 思路：快慢指针，快慢相遇之后，其中一个指针回到头，快慢指针步调一致一起移动，相遇点即为入环点
    // nb+a=2nb+a
    if head == nil {
        return head
    }
    fast := head
    slow := head

    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            // 指针重新从头开始移动
            fast = head
            for fast != slow {
                fast = fast.Next
                slow = slow.Next
            }
            return slow
        }
    }
    return nil
}
```
这两种方式不同点在于，一般用 fast=head.Next 较多，因为这样可以知道中点的上一个节点，可以用来删除等操作。
- fast 如果初始化为 head.Next 则中点在 slow.Next
- fast 初始化为 head,则中点在 slow