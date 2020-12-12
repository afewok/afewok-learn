package com.afewok.learn.leetcode;

import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.stream.Collector;
import java.util.stream.Collectors;
import java.util.stream.Stream;

import org.testng.annotations.Test;

/**
 * 5609. 统计一致字符串的数目
 */
public class LeetCodeWeek5609 {
    @Test
    public void leetCode() {
        Flow.executor(() -> countConsistentStrings("ab", new String[]{"ad", "bd", "aaab", "baa", "badab"}));
    }

    public int countConsistentStrings(String allowed, String[] words) {
        int count = 0;
        boolean ex = false;
        for (String string : words) {
            ex = true;
            for (char c : string.toCharArray()) {
                if (!allowed.contains(c + "")) {
                    ex = false;
                    break;

                }
            }
            if (ex) {
                count++;
            }
        }
        return count;
    }
}
