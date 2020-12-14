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

    public List<List<String>> groupAnagrams2(String[] strs) {
        Map<String, List<String>> map = new HashMap<>();
        for (String str : strs) {
            char[] array = str.toCharArray();
            Arrays.sort(array);
            List<String> list = map.computeIfAbsent(new String(array), key -> new ArrayList<String>());
            list.add(str);
        }
        return new ArrayList<List<String>>(map.values());
    }

    public List<List<String>> groupAnagrams(String[] strs) {
        List<List<String>> ans = new ArrayList<>();
        Map<Long, Integer> hashIndexMap = new HashMap<>();
        int size = 0;
        for(String str : strs){
            long hash = getHash(str);
            List<String> list;
            if(hashIndexMap.containsKey(hash)){
                int index = hashIndexMap.get(hash);
                list = ans.get(index);
            }else{
                list = new ArrayList<>();
                ans.add(list);
                hashIndexMap.put(hash, size++);
            }
            list.add(str);
        }
        return ans;
    }
    long getHash(String str){
        long hash = 0;
        long sum = 0;
        long prod = 1;
        for(char c : str.toCharArray()){
            sum += (long) (c);
            prod *= (long) (c);
        }
        hash = sum + prod;
        return hash;
    }
}
