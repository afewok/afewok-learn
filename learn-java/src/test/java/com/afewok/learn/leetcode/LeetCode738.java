package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

/**
 * 738. 单调递增的数字
 * 
 */
public class LeetCode738 {

    @Test
    public void leetCode() {
        // Flow.executor(() -> monotoneIncreasingDigits(1000));
        // Flow.executor(() -> monotoneIncreasingDigits(10));
        // Flow.executor(() -> monotoneIncreasingDigits(1234));
        Flow.executor(() -> monotoneIncreasingDigits(332));
    }

    public int monotoneIncreasingDigits(int N) {
        if (N == 0) {
            return 0;
        }
        char[] chars = String.valueOf(N).toCharArray();
        int length = chars.length;
        int minValue = Integer.MAX_VALUE;
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < length; i++) {
            if (minValue <= chars[i]) {
                sb.append("9");
                continue;
            }
            int sbLength = sb.length();
            char[] charTemp = new char[sbLength];
            for (int j = 0; j < sbLength; j++) {
                charTemp[j] = chars[i];
            }
            sb.delete(0, sbLength).append(new String(charTemp)).append(chars[i]);
        }
        return Integer.parseInt(sb.toString());
    }
}
