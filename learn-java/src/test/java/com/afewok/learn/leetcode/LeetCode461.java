package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

/**
 * 461. 汉明距离
 */
public class LeetCode461 {
    @Test
    public void leetCode() {
        Flow.executor(() -> hammingDistance(1, 4));
    }

    public int hammingDistance(int x, int y) {
        return Integer.bitCount(x ^ y);
    }

    public int hammingDistance1(int x, int y) {
        int xor = x ^ y;
        int max = 0;
        while (xor > 0) {
            max += xor & 1;
            xor = xor >> 1;
        }
        return max;
    }
}
