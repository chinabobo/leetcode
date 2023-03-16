
https://leetcode.cn/problems/longest-substring-without-repeating-characters/

思路:
利用一个window。一直维持这个window里面没有重复的字符，记录这个window的最大长度。
维持的方法：
不断的加入字符，每次加入字符，检查此时window有没有重复的，如果有，就从window的左边踢出一个字符，直到window没有重复字符，继续从右边加入一个字符。最终记录window的最大长度。