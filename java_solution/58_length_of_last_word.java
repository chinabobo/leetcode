class Solution {
    private static String nullStr = " ";
    public int lengthOfLastWord(String s) {
        if (nullStr.equals(s)) {
            return 0;
        }
        if(s.length()==1){
            return 1;
        }
        for (int i = s.length()-1; i >= 0; i--) {
            if(!s.substring(i, i+1).equals(nullStr)){
                for(int j = i; j >= 0; j--){
                    if(j == 0){
                        return i-j+1;
                    }
                    else if(s.substring(j-1, j).equals(nullStr)){
                        return i-j+1;
                    }
                }
            }
        }
        return 0;
    }
}

// https://leetcode.cn/problems/length-of-last-word/