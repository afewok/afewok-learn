package com.afewok.learn.leetcode;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

import org.testng.annotations.Test;

public class LeetCode049 {
    @Test
    public void leetCode() {
        Flow.executor(() -> groupAnagrams(new String[] { "eat", "tea", "tan", "ate", "nat", "bat" }));
        Flow.executor(() -> groupAnagrams(new String[] { "cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc" }));
    }

    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> map = new HashMap<>();
        for (String str : strs) {
            StringBuilder stringBuilder = new StringBuilder("000000000000000000000000000");
            for (int i = 0; i < str.length(); i++) {
                stringBuilder.setCharAt(((int) str.charAt(i) - 97), '1');
            }

            List<String> list = map.computeIfAbsent(stringBuilder.toString(), key -> new ArrayList<String>());
            list.add(str);
        }
        return map.values().stream().collect(Collectors.toList());
    }
}
