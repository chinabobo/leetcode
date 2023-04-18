## searchMatrix by Myself

思路：先找每行第一个元素，二分搜索 ，找到target在哪一行，找到之后，再调用一次二分搜索，找到再这个行的哪一个。

## searchMatrix

思路：将2纬数组转为1维数组， 直接平铺， 进行二分搜索
```go
start := 0
end := row*col - 1
for start+1 < end {
    mid := start + (end-start)/2
    // 获取2纬数组对应值
    val := matrix[mid/col][mid%col]
    if val > target {
        end = mid
    } else if val < target {
        start = mid
    } else {
        return true
    }
}
if matrix[start/col][start%col] == target || matrix[end/col][end%col] == target{
    return true
}
```
