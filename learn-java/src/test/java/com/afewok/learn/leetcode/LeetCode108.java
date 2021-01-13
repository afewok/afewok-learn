package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * 108. 将有序数组转换为二叉搜索树
 */
public class LeetCode108 {
    @Test
    public void leetCode() {
        Flow.executor(() -> sortedArrayToBST(new int[] { -10, -3, 0, 5, 9 }));
        Flow.executor(() -> sortedArrayToBST(new int[] { -1, 0, 1, 2 }));
    }

    public TreeNode sortedArrayToBST(int[] nums) {
        return null;
    }

    @Data
    @NoArgsConstructor
    @AllArgsConstructor
    public class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
    }
}
