package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

/**
 * 477. 汉明距离总和
 */
public class LeetCode477 {
    @Test
    public void leetCode() {
        Flow.executor(() -> totalHammingDistance(new int[]{4, 14, 2}));
    }

    public int totalHammingDistance(int[] nums) {
        int max=0;
        for(int i=0;i<nums.length;i++){
            for(int j=i+1;j<nums.length;j++){
                if(i==j){
                    continue;
                }
                max +=Integer.bitCount(nums[i] ^ nums[j]);
            }
        }
        return max;
    }

    //同一位的0和1相乘
    public int totalHammingDistance1(int[] nums) {
        int ans = 0, n = nums.length;
        for (int i = 0; i < 30; ++i) {
            int c = 0;
            for (int val : nums) {
                c += (val >> i) & 1;
            }
            ans += c * (n - c);
        }
        return ans;
    }
}
