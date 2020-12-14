package com.afewok.learn.leetcode;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collector;
import java.util.stream.Collectors;
import java.util.stream.Stream;

import org.testng.annotations.Test;

/**
 * 49. 字母异位词分组
 * 
 * 排序,计数,，质数乘积
 */
public class LeetCode049 {
    @Test
    public void leetCode() {
        Flow.executor(() -> groupAnagrams(new String[] { "eat", "tea", "tan", "ate", "nat", "bat" }));
        Flow.executor(() -> groupAnagrams(
                new String[] { "cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc" }));
    }

    public List<List<String>> groupAnagrams1(String[] strs) {
        Map<String, List<String>> map = new HashMap<>();
        for (String str : strs) {
            int length = str.length();
            int[] temps = new int[26];
            for (int i = 0; i < length; i++) {
                temps[str.charAt(i) - 'a']++;
            }
            StringBuilder stringBuilder = new StringBuilder();
            for (int i = 0; i < 26; i++) {
                if (temps[i] != 0) {
                    stringBuilder.append((char) ('a' + i)).append(temps[i]);
                }
            }
            List<String> list = map.computeIfAbsent(stringBuilder.toString(), key -> new ArrayList<String>());
            list.add(str);
        }
        return map.values().stream().collect(Collectors.toList());
    }

    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> map = new HashMap<>();
        for (String str : strs) {
            char[] array = str.toCharArray();
            Arrays.sort(array);
            List<String> list = map.computeIfAbsent(new String(array), key -> new ArrayList<String>());
            list.add(str);
        }
        return new ArrayList<List<String>>(map.values());
    }

    public List<List<String>> groupAnagrams3(String[] strs) {
        int length = strs.length;
        Map<String, List<String>> map = new HashMap<>(length);
        for (String str : strs) {
            System.out.println(Stream.of(str.toCharArray()).sorted().collect(Collectors.toList()));
            List<String> list = map.computeIfAbsent(Stream.of(str.toCharArray()).sorted().toString(),
                    key -> new ArrayList<String>());
            list.add(str);
        }
        return map.values().stream().collect(Collectors.toList());
    }
}
