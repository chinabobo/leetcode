按位与运算的性质, 对于一系列的位，只要有一个0，那么一系列的按位与的运算结果都是0.

9  | 0 0 0 0 1 0 0 1
10 | 0 0 0 0 1 0 1 0
11 | 0 0 0 0 1 0 1 1
12 | 0 0 0 0 1 1 0 0 

注意到前5位一样，由于数字是连续的，所以后面一定是有0的，一系列的按位与的运算结果是0.

所以最后的答案即为二进制字符串的公共前缀再用零补上后面的剩余位。
最终的问题转化为，给定两个整数，要找到它们对应的二进制字符串的公共前缀
