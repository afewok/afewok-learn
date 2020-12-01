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
 * 思路：暴力解法、中心扩散、动态规划、Manacher 算法
 */
public class LeetCode005 {

    @Test
    public void leetCode005() {
        String str1 = "sbv";
        String str2 = "babad";
        String str3 = "ccc";
        String str4 = "cbbd";
        String str5 = "abcabccbacba";

        System.out.println(longestPalindrome1(str1));
        System.out.println(longestPalindrome1(str2));
        System.out.println(longestPalindrome1(str3));
        System.out.println(longestPalindrome1(str4));
        System.out.println(longestPalindrome1(str5));

    }

    public String longestPalindrome1(String s) {
        char[] chats = s.toCharArray(), maxStr = { chats[0] };
        int m, n, temp, length = chats.length, maxHalf = maxStr.length / 2;
        for (int i = 0; i < length && i + maxHalf < length; i++) {
            for (m = i - 1, n = i + 1; m >= 0 && n < length; m--, n++) {
                if (chats[m] != chats[n]) {
                    break;
                }
                temp = n - m + 1;
                if (temp > maxStr.length) {
                    maxStr = new char[temp];
                    maxHalf = maxStr.length / 2;
                    System.arraycopy(chats, m, maxStr, 0, temp);
                }
            }
            for (m = i - 1, n = i; m >= 0 && n < length; m--, n++) {
                if (chats[m] != chats[n]) {
                    break;
                }
                temp = n - m + 1;
                if (temp > maxStr.length) {
                    maxStr = new char[temp];
                    maxHalf = maxStr.length / 2;
                    System.arraycopy(chats, m, maxStr, 0, temp);
                }
            }
        }
        return String.valueOf(maxStr);
    }
}
