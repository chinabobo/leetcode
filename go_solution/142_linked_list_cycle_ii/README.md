README

![](https://assets.leetcode-cn.com/solution-static/142/142_fig1.png)

a + n ( b + c ) + b - 1 = 2(a + b)

a = (n - 1) (b + c) + c - 1

快慢指针，快慢相遇之后，慢指针回到头，快指针向后一个节点（由于一开始快指针是从`fast := head.Next`开始的），快慢指针步调一致一起移动，相遇点即为入环点。


