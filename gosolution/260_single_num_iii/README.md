
## 基本操作
- a=0^a=a^0
- 0=a^a

由上面两个推导出：
- a=a^b^b
- 交换两个数 
a=a^b
b=a^b
a=a^b
- 移除最后一个 1
n&(n-1)
- 获取最后一个 1 
diff=(n&(n-1))^n 或者 n&(-n)

## 思路

有两个数只出现了一次记为 x、y， 初始化为 0, 其余的数出现了两次,
我们可以对所有的数进行一次异或操作, 易知这个结果就等于 x^y (相同的数异或结果都为 0, 0和任意数异或结果都为那个数).

那么可以考虑异ans的某个非 0 位,如最后一个非 0 位, 因为我们知道只有当 x、y 在该位不一样的时候才会出现异或结果为 1. 
所以以该位是否为 1 对数组进行划分(这样就划分成两个数组，每个数组序列中都存在一个只出现一次的数字，既可以使用简单的异或直接找到结果)
只要该位为 1 就和 x 异或, 只要该位为 0就和 y 异或, 这样最终得到就是只出现过一次的两个数(其他在该位为 1 或0 的数必然出现 0/2 次对异或结果无影响)
