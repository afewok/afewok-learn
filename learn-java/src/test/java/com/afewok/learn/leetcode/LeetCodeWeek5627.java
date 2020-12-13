package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

/**
 * 5627. 石子游戏 VII
 */
public class LeetCodeWeek5627 {
    @Test
    public void leetCode() {
        Flow.executor(() -> stoneGameVII(new int[]{5, 3, 1, 4, 2}));
        Flow.executor(() -> stoneGameVII(new int[]{7, 90, 5, 1, 100, 10, 10, 2}));
    }

    public int stoneGameVII(int[] stones) {
        if (stones.length < 3) {
            return Math.abs(stones[0] - stones[1]);
        }
        int l = 0, r = stones.length - 1, aCount = 0, bCount = 0;
        while (l == r) {
            if (stones[l] < stones[r]) {
                l++;
            } else {
                r--;
            }
            int count = 0;
            for (int i = l; i <= r; i++) {
                count = count + stones[i];
            }
            aCount = aCount + count;

            if (stones[l] > stones[r]) {
                l++;
            } else {
                r--;
            }

            count = 0;
            for (int i = l; i <= r; i++) {
                count = count + stones[i];
            }
            bCount = bCount + count;

        }
        return Math.abs(aCount - bCount);
    }
}
