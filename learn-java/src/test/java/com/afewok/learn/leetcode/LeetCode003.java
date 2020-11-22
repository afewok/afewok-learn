package com.afewok.learn.leetcode;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;

import org.testng.annotations.Test;

/**
 * 3. 无重复字符的最长子串
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:输入: "abcabcbb" 输出: 3 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 * 示例 2:输入: "bbbbb" 输出: 1 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 * 示例 3:输入: "pwwkew" 输出: 3 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。 请注意，你的答案必须是子串
 * 的长度，"pwke" 是一个子序列，不是子串。
 *
 * 思路：双指针+滑动窗口、动态规划+哈希表、利用数组（桶）代替hashmap、暴力双循环
 */
public class LeetCode003 {

    @Test
    public void leetCode003() {
        String s1 = "abcabcbb";
        System.out.println(s1 + " " + lengthOfLongestSubstring0(s1));
        System.out.println(s1 + " " + lengthOfLongestSubstring1(s1));
        System.out.println(s1 + " " + lengthOfLongestSubstring2(s1));

        String s2 = "pwwkew";
        System.out.println(s2 + " " + lengthOfLongestSubstring0(s2));
        System.out.println(s2 + " " + lengthOfLongestSubstring1(s2));
        System.out.println(s2 + " " + lengthOfLongestSubstring2(s2));

        String s3 = "abbbba";
        System.out.println(s3 + " " + lengthOfLongestSubstring0(s3));
        System.out.println(s3 + " " + lengthOfLongestSubstring1(s3));
        System.out.println(s3 + " " + lengthOfLongestSubstring2(s3));
    }

    public int lengthOfLongestSubstring0(String s) {
        Map<Character, Integer> dic = new HashMap<>();
        int res = 0, tmp = 0;
        for (int j = 0; j < s.length(); j++) {
            int i = dic.getOrDefault(s.charAt(j), -1); // 获取索引 i
            dic.put(s.charAt(j), j); // 更新哈希表
            tmp = tmp < j - i ? tmp + 1 : j - i; // dp[j - 1] -> dp[j]
            res = Math.max(res, tmp); // max(dp[j - 1], dp[j])
        }
        return res;
    }

    public int lengthOfLongestSubstring1(String s) {
        int sub = 0, temp, max = 0, length = s.length();
        Map<Character, Integer> map = new HashMap<>(s.length());
        for (int i = 0; i < length; i++) {
            char c = s.charAt(i);
            if (map.containsKey(c)) {
                max = Math.max(max, i - sub);
                sub = Math.max(sub, map.get(c) + 1);
            }
            map.put(c, i);
        }
        temp = length - sub;
        return max > temp ? max : temp;
    }

    public int lengthOfLongestSubstring2(String s) {
        int rsub = 0, max = 0, length = s.length();
        Set<Character> set = new HashSet<>(length);
        for (int i = 0; i < length; i++) {
            if (i > 0) {
                set.remove(s.charAt(i - 1));
            }
            while (rsub < length && !set.contains(s.charAt(rsub))) {
                set.add(s.charAt(rsub));
                rsub++;
                max = Math.max(max, rsub - i);
            }
            if (max > length - i) {
                break;
            }
        }
        return max;
    }
}
