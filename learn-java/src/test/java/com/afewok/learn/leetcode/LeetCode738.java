package com.afewok.learn.leetcode;

import org.checkerframework.checker.units.qual.min;
import org.testng.annotations.Test;

/**
 * 738. 单调递增的数字
 * 
 */
public class LeetCode738 {

    @Test
    public void leetCode() {
        Flow.executor(() -> monotoneIncreasingDigits(989998));
        Flow.executor(() -> monotoneIncreasingDigits(120));
        Flow.executor(() -> monotoneIncreasingDigits(121));
        Flow.executor(() -> monotoneIncreasingDigits(20));
        Flow.executor(() -> monotoneIncreasingDigits(10));
        Flow.executor(() -> monotoneIncreasingDigits(1234));
        Flow.executor(() -> monotoneIncreasingDigits(332));
    }

    public int monotoneIncreasingDigits(int N) {
        char[] chars = String.valueOf(N).toCharArray();
        int length = chars.length;

        for (int i = length - 2; i >= 0; i--) {
            if (chars[i] < chars[i + 1]) {
                continue;
            } else if (chars[i] > chars[i + 1]) {
                chars[i] = (char) ((int) chars[i] - 1);
                for (int j = i + 1; j < length; j++) {
                    chars[j] = '9';
                }
            }
        }
        return Integer.parseInt(new String(chars));
    }
}