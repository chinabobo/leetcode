class Solution {
    public String reverseLeftWords(String s, int n) {
        String resStr = s.substring(n, s.length()) + s.substring(0, n);
        return resStr;
    }
}

// https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/