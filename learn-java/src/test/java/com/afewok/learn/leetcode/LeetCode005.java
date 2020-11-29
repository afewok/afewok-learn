package com.afewok.learn.leetcode;

import java.util.*;

import org.testng.annotations.Test;

/**
 * 5. 最长回文子串
 * 
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 * 
 * 示例 1：输入: "babad" 输出: "bab" 注意: "aba" 也是一个有效答案。
 * 
 * 示例 2：输入: "cbbd" 输出: "bb"
 * 
 * 思路：
 */
public class LeetCode005 {

    @Test
    public void leetCode005() {
        String str1="babad";
        String str2="cbbd";

        System.out.println(longestPalindrome1(str1));
        System.out.println(longestPalindrome1(str2));

    }

    public String longestPalindrome1(String s) {



        return s;
    }
}
