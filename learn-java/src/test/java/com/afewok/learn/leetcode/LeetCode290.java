package com.afewok.learn.leetcode;

import java.util.HashMap;
import java.util.Map;

import org.testng.annotations.Test;

/**
 * 290. 单词规律
 * 
 */
public class LeetCode290 {

    @Test
    public void leetCode() {
        Flow.executor(() -> wordPattern("abba", "dog cat cat dog"));
        Flow.executor(() -> wordPattern("abba", "dog cat cat fish"));
        Flow.executor(() -> wordPattern("aaaa", "dog cat cat dog"));
        Flow.executor(() -> wordPattern("abba", "dog dog dog dog"));
    }

    public boolean wordPattern(String pattern, String s) {
        char[] chars = pattern.toCharArray();
        String[] strings = s.split(" ");
        int length1 = chars.length, length2 = strings.length;
        if (length1 != length2) {
            return false;
        }
        Map<Character, String> map1 = new HashMap<>();
        Map<String, Character> map2 = new HashMap<>();
        for (int i = 0; i < length1; i++) {
            if ( map1.get(chars[i]) == null && map2.get(strings[i]) == null) {
                map1.put(chars[i], strings[i]);
                map2.put(strings[i], chars[i]);
            } else if (!strings[i].equals( map1.get(chars[i])) || chars[i] != map2.get(strings[i])) {
                return false;
            }
        }
        return true;
    }
}
