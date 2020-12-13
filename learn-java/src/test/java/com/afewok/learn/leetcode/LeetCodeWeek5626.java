package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

/**
 * 5626. 十-二进制数的最少数目
 */
public class LeetCodeWeek5626 {
    @Test
    public void leetCode() {
        Flow.executor(() -> minPartitions("32"));
        Flow.executor(() -> minPartitions("82732"));
        Flow.executor(() -> minPartitions("27346209830709182346"));
    }

    public int minPartitions(String n) {
        int max = 0;
        for (int i = 0; i < n.length(); i++) {
            char ch = n.charAt(i);
            int p = Integer.parseInt(ch + "");
            if (max < p) {
                max = p;
            }
        }
        return max;
    }
}
