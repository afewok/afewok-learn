package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

/**
 * 26. 删除排序数组中的重复项
 * 
 * 思路：双指针法
 */
public class LeetCode026 {
    @Test
    public void leetCode() {
        Flow.executor(() -> removeDuplicates(new int[] { 1, 2 }));
        Flow.executor(() -> removeDuplicates(new int[] { 1, 1, 2 }));
        Flow.executor(() -> removeDuplicates(new int[] { 0, 0, 1, 1, 1, 2, 2, 3, 3, 4 }));
    }

    public int removeDuplicates(int[] nums) {
        int sub = 0;
        for (int num : nums) {
            if (nums[sub] < num) {
                nums[++sub] = num;
            }
        }
        return sub + 1;
    }
}
