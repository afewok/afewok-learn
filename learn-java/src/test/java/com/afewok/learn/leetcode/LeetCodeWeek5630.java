package com.afewok.learn.leetcode;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;

import org.testng.annotations.Test;

/**
 * 5630. 删除子数组的最大得分
 */
public class LeetCodeWeek5630 {

    @Test
    public void leetCode() {
        Flow.executor(() -> maximumUniqueSubarray(new int[] { 4, 2, 4, 5, 6 }));
        Flow.executor(() -> maximumUniqueSubarray(new int[] { 5, 2, 1, 2, 5, 2, 1, 2, 5 }));
        Flow.executor(() -> maximumUniqueSubarray(new int[] { 20, 19, 20, 1, 2, 3, 4, 5 }));
        Flow.executor(() -> maximumUniqueSubarray(new int[] { 187, 470, 25, 436, 538, 809, 441, 167, 477, 110, 275, 133,
                666, 345, 411, 459, 490, 266, 987, 965, 429, 166, 809, 340, 467, 318, 125, 165, 809, 610, 31, 585, 970,
                306, 42, 189, 169, 743, 78, 810, 70, 382, 367, 490, 787, 670, 476, 278, 775, 673, 299, 19, 893, 817,
                971, 458, 409, 886, 434 }));
    }

    public int maximumUniqueSubarray(int[] nums) {
        int maxSum = 0, sum = 0, sub = 0;
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            if (map.containsKey(nums[i])) {
                if (maxSum < sum) {
                    maxSum = sum;
                }
                int temp = map.get(nums[i]);
                for (int j = sub; j <= temp; j++) {
                    map.remove(nums[j]);
                    sum = sum - nums[j];
                }
                sub=temp+1;
            }
            map.put(nums[i], i);
            sum = sum + nums[i];
        }
        if (maxSum < sum) {
            maxSum = sum;
        }
        return maxSum;
    }

}