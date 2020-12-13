package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

/**
 * 5625. 比赛中的配对次数
 */
public class LeetCodeWeek5625 {
    @Test
    public void leetCode() {
        Flow.executor(() -> numberOfMatches(5));
        Flow.executor(() -> numberOfMatches(7));
        Flow.executor(() -> numberOfMatches(14));
    }

    public int numberOfMatches(int n) {
        int count = 0, half = 0;
        boolean p = false;
        while (n > 0) {
            n=n + (p ? 1 : 0);
            half = n / 2;
            count = count + half;
            p = false;
            if (half * 2 < n) {
                p = true;
            }
            n = half;
        }
        return count;
    }
}
