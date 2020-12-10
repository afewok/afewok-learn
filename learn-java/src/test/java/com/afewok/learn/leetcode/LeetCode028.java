package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

/**
 * 28. 实现 strStr()
 *
 * 思路：首字母比对、KMP 算法（Knuth-Morris-Pratt 算法）
 */
public class LeetCode028 {
    @Test
    public void leetCode() {

    }

    public int strStr(String haystack, String needle) {
        int hlen = haystack.length(), nlen = needle.length();
        if (nlen == 0) {
            return 0;
        } else if (hlen == 0 || hlen < nlen) {
            return -1;
        }
        for (int i = 0; i <= hlen - nlen; i++) {
            if (haystack.substring(i, i + nlen).equals(needle)) {
                return i;
            }
        }
        return -1;
    }

}
