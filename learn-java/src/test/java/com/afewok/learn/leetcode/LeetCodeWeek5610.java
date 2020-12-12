package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

import java.util.HashMap;
import java.util.Map;

/**
 * 5610. 有序数组中差绝对值之和
 */
public class LeetCodeWeek5610 {
    @Test
    public void leetCode() {
        Flow.executor(() -> getSumAbsoluteDifferences(new int[]{2, 3, 5}));
    }

    public int[] getSumAbsoluteDifferences(int[] nums) {
        int length = nums.length, sum = 0, temp = 0;
        int[] result = new int[length];
        for (int i = 0; i < length; i++) {
            sum = sum + nums[i];
        }
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < length; i++) {
            if (map.get(nums[i]) != null) {
                result[i] = map.get(nums[i]);
                continue;
            }
            temp = 0;
            for (int j = 0; j < i; j++) {
                temp = temp + nums[j];
            }
            result[i] = (nums[i] * i - temp)-nums[i] +(sum - temp - nums[i] * (length - i - 1));
            map.put(nums[i], result[i]);
        }
        return result;
    }
}
