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
 * 示例:
 * 给定 nums = [2, 7, 11, 15], target = 9
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 */
public class LeetCode001 {

    @Test
    public void leetCode001() {
        int[] nums = { 2, 7, 11, 15 };
        int target = 9;

        LeetCode001 leetCode001 = new LeetCode001();
        System.out.println(Json.toJSONString(leetCode001.twoSum(nums, target)));
    }

    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>(nums.length);
        for (int i = 0; i < nums.length; i++) {
            int temp = target - nums[i];
            if (map.get(temp) != null) {
                return new int[] { map.get(temp),i };
            }
            map.put(nums[i], i);
        }

        throw new IllegalArgumentException("Illegal nums or target");
    }
}