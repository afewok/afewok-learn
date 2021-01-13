package com.afewok.learn.leetcode;

import java.util.ArrayList;
import java.util.List;

import org.testng.annotations.Test;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * 103. 二叉树的锯齿形层序遍历
 */
public class LeetCode103 {
    @Test
    public void leetCode() {
        Flow.executor(() -> zigzagLevelOrder(
                new TreeNode(3, new TreeNode(9, new TreeNode(15, null, null), new TreeNode(7, null, null)),
                        new TreeNode(20, null, null))));
    }

    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> result = new ArrayList<>();
        if (root == null) {
            return result;
        }
        boolean b = true;
        List<TreeNode> list = new ArrayList<>();
        List<TreeNode> temp = new ArrayList<>();

        list.add(root);
        while (list.size() > 0) {
            ArrayList<Integer> res = new ArrayList<>();
            for (int i = 0; i < list.size(); i++) {
                if (b) {
                    res.add(list.get(i).val);
                } else {
                    res.add(list.get(list.size() - i - 1).val);
                }
                if (list.get(i).left != null) {
                    temp.add(list.get(i).left);
                }
                if (list.get(i).right != null) {
                    temp.add(list.get(i).right);
                }
            }

            result.add(res);
            b = !b;
            list = temp;
            temp = new ArrayList<>();
        }

        return result;
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
