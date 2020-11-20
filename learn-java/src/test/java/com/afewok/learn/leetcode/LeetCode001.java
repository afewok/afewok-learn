package com.afewok.learn.leetcode;

import java.util.Map;

import com.afewok.learn.util.Json;
import org.testng.annotations.Test;

import java.util.HashMap;

/**
 * 1. 两数之和
 *
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 * 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
 * 
 * 示例: 给定 nums = [2, 7, 11, 15], target = 9 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 *
 * 思路：暴力2次循环、头尾递进、hash字典、数组存储差值
 */
public class LeetCode001 {

    @Test
    public void leetCode001() {
        int[] nums = { 2, 7, 11, 15 };
        int target = 9;

        LeetCode001 leetCode001 = new LeetCode001();
        System.out.println(Json.toJSONString(leetCode001.twoSum1(nums, target)));
        System.out.println(Json.toJSONString(leetCode001.twoSum2(nums, target)));
    }

    public int[] twoSum1(int[] nums, int target) {
        int[] temps = new int[nums.length];
        for (int i = 0; i < nums.length; i++) {
            for (int j = 0; j < i; j++) {
                if (nums[i] == temps[j]) {
                    return new int[] { j, i };
                }
            }
            temps[i] = target - nums[i];
        }
        throw new IllegalArgumentException("Illegal nums or target");
    }

    public int[] twoSum2(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>(nums.length);
        for (int i = 0; i < nums.length; i++) {
            if (map.containsKey(nums[i])) {
                return new int[] { map.get(nums[i]), i };
            }
            map.put(target - nums[i], i);
        }
        throw new IllegalArgumentException("Illegal nums or target");
    }
}