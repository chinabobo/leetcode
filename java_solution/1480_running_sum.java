class Solution {
    public int[] runningSum(int[] nums) {
        int[] outputNums = new int[nums.length];
        int tmp = 0;
        for (int i = 0; i < nums.length; i++) {
            outputNums[i] = nums[i] + tmp;
            tmp = outputNums[i];
        }
        return outputNums;
    }
}

// https://leetcode.cn/problems/running-sum-of-1d-array/