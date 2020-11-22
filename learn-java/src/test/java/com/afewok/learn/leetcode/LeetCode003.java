package com.afewok.learn.leetcode;

import java.util.HashMap;
import java.util.HashSet;
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
 * 思路：双指针+滑动窗口、动态规划+哈希表（虚拟-1下标）、利用数组（桶）代替hashmap、暴力双循环.....其核心是记左下标，还是记上一次长度
 */
public class LeetCode003 {

    @Test
    public void leetCode003() {
        String s1 = "abcabcbb";
        System.out.println(s1 + " " + lengthOfLongestSubstring0(s1));
        System.out.println(s1 + " " + lengthOfLongestSubstring1(s1));
        System.out.println(s1 + " " + lengthOfLongestSubstring2(s1));
        System.out.println(s1 + " " + lengthOfLongestSubstring3(s1));

        String s2 = "pwwkew";
        System.out.println(s2 + " " + lengthOfLongestSubstring0(s2));
        System.out.println(s2 + " " + lengthOfLongestSubstring1(s2));
        System.out.println(s2 + " " + lengthOfLongestSubstring2(s2));
        System.out.println(s2 + " " + lengthOfLongestSubstring3(s2));

        String s3 = "abba";
        System.out.println(s3 + " " + lengthOfLongestSubstring0(s3));
        System.out.println(s3 + " " + lengthOfLongestSubstring1(s3));
        System.out.println(s3 + " " + lengthOfLongestSubstring2(s3));
        System.out.println(s3 + " " + lengthOfLongestSubstring3(s3));
    }

    public int lengthOfLongestSubstring0(String s) {
        char c;
        int sub = -1, max = 0, length = s.length();
        Map<Character, Integer> map = new HashMap<>(s.length());
        for (int i = 0; i < length; i++) {
            c = s.charAt(i);
            sub = Math.max(sub,  map.getOrDefault(c, -1));
            max = Math.max(max, i - sub );
            map.put(c, i);
        }
        return max;
    }

    public int lengthOfLongestSubstring1(String s) {
        char c;
        int count = 0, max = 0, length = s.length();
        Map<Character, Integer> map = new HashMap<>(s.length());
        for (int i = 0; i < length; i++) {
            c = s.charAt(i);
            count = Math.min(count + 1, i - map.getOrDefault(c, -1));
            max = Math.max(max, count);
            map.put(c, i);
        }
        return max;
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

    public int lengthOfLongestSubstring3(String s) {
        char c;
        int sub, temp = 0, max = 0, length = s.length();
        Map<Character, Integer> map = new HashMap<>(s.length());
        for (int i = 0; i < length; i++) {
            c = s.charAt(i);
            sub = map.getOrDefault(c, -1);
            map.put(c, i);

            temp = temp < i - sub ? temp + 1 : i - sub; // 上一次的差距与本次差距比较，大于则上一次的差距+1，小于则为本次差距
            max = Math.max(max, temp);
        }
        return max;
    }
}
