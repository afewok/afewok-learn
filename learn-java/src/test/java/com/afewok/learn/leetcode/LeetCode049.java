package com.afewok.learn.leetcode;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collector;
import java.util.stream.Collectors;
import java.util.stream.Stream;

import org.testng.annotations.Test;

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
            int[] temps = { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 };
            for (int i = 0; i < str.length(); i++) {
                int temp = ((int) str.charAt(i) - 97);
                temps[temp] = temps[temp] + 1;
            }
            StringBuilder stringBuilder = new StringBuilder();
            for (int temp : temps) {
                stringBuilder.append(temp + ".");
            }

            List<String> list = map.computeIfAbsent(stringBuilder.toString(), key -> new ArrayList<String>());
            list.add(str);
        }
        return map.values().stream().collect(Collectors.toList());
    }

    public List<List<String>> groupAnagrams(String[] strs) {
        char[] chs;char c;int length=0;
        Map<String, List<String>> map = new HashMap<>();
        for (String str : strs) {
            chs = str.toCharArray();
            length=chs.length;
            for (int i = 0; i < length; i++) {
                for (int j = i + 1; j < length; j++) {
                    if (chs[i] > chs[j]) {
                        c = chs[i];
                        chs[i] = chs[j];
                        chs[j] = c;
                    }
                }
            }
            List<String> list = map.computeIfAbsent(new String(chs), key -> new ArrayList<String>());
            list.add(str);
        }
        return map.values().stream().collect(Collectors.toList());
    }

    public List<List<String>> groupAnagrams3(String[] strs) {
        int length=strs.length;
        Map<String, List<String>> map = new HashMap<>(length);
        for (String str : strs) {
            System.out.println( Stream.of(str.toCharArray()).sorted().collect(Collectors.toList()));
            List<String> list = map.computeIfAbsent( Stream.of(str.toCharArray()).sorted().toString(), key -> new ArrayList<String>());
            list.add(str);
        }
        return map.values().stream().collect(Collectors.toList());
    }
}
