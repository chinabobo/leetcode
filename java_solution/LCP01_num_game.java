class Solution {
    public int game(int[] guess, int[] answer) {
        int cot = 0;
        int i = 0;
        for(int j = 0; j < guess.length;j++){
            if (guess[j] == answer[i]){
                cot++;
            }
            i++;
        }
        return cot;
    }

}

// https://leetcode.cn/problems/guess-numbers/