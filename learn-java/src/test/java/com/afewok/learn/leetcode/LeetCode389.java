package com.afewok.learn.leetcode;

import java.util.HashMap;
import java.util.Map;

import org.testng.annotations.Test;

/**
 * 389. 找不同
 */
public class LeetCode389 {
    @Test
    public void leetCode() {
        Flow.executor(() -> findTheDifference("abcd", "abcde"));
    }

    public char findTheDifference1(String s, String t) {
        int length = s.length();
        int[] nums = new int[26];
        nums[t.charAt(length) - 97]--;
        for (int i = 0; i < length; i++) {
            nums[s.charAt(i) - 97]++;
            nums[t.charAt(i) - 97]--;
        }
        for (int i = 0; i < 26; i++) {
            if (nums[i] < 0) {
                return (char) (i + 97);
            }
        }
        throw new RuntimeException("panic panic panic");
    }

    public char findTheDifference2(String s, String t) {
        char[] sChar = s.toCharArray(), tChar = t.toCharArray();
        int length = s.length();
        Map<Character, Integer> map = new HashMap<>();
        map.put(tChar[length], map.getOrDefault(tChar[length], 0) - 1);
        for (int i = 0; i < length; i++) {
            map.put(sChar[i], map.getOrDefault(sChar[i], 0) + 1);
            map.put(tChar[i], map.getOrDefault(tChar[i], 0) - 1);
        }
        for (Map.Entry<Character, Integer> entry : map.entrySet()) {
            if (entry.getValue() < 0) {
                return entry.getKey();
            }
        }
        throw new RuntimeException("panic panic panic");
    }

    public char findTheDifference(String s, String t) {
        int length = s.length(), sum = t.charAt(length);
        for (int i = 0; i < length; i++) {
            sum = sum - s.charAt(i) + t.charAt(i);
        }
        return (char) sum;
    }
}
