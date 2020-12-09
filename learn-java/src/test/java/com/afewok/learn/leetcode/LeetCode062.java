package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

/**
 * 62. 不同路径
 * 
 * 一个机器人位于一个 m x n 网格的左上角 。 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角。问总共有多少条不同的路径？
 * 
 * 示例 1:输入: m = 3, n = 2 输出: 3 解释:从左上角开始，总共有 3 条路径可以到达右下角。1. 向右 -> 向右 -> 向下2. 向右
 * -> 向下 -> 向右3. 向下 -> 向右 -> 向右
 * 
 * 示例 2:输入: m = 7, n = 3 输出: 28
 * 
 * 示例 3：输入：m = 3, n = 3 输出：6
 * 
 * 提示：1 <= m, n <= 100 题目数据保证答案小于等于 2 * 10 ^ 9
 * 
 * 思路：二叉树遍历，控制变量法
 */
public class LeetCode062 {

    @Test
    public void leetCode062() {
        Flow.executor(() -> uniquePaths1(2, 1));
        Flow.executor(() -> uniquePaths1(2, 2));
        Flow.executor(() -> uniquePaths1(2, 3));
        Flow.executor(() -> uniquePaths1(2, 4));
        Flow.executor(() -> uniquePaths1(2, 5));
        Flow.executor(() -> uniquePaths1(2, 6));
        System.out.println();
        Flow.executor(() -> uniquePaths1(3, 1));
        Flow.executor(() -> uniquePaths1(3, 2));
        Flow.executor(() -> uniquePaths1(3, 3));
        Flow.executor(() -> uniquePaths1(3, 4));
        Flow.executor(() -> uniquePaths1(3, 5));
        Flow.executor(() -> uniquePaths1(3, 6));
        System.out.println();
        Flow.executor(() -> uniquePaths1(4, 1));
        Flow.executor(() -> uniquePaths1(4, 2));
        Flow.executor(() -> uniquePaths1(4, 3));
        Flow.executor(() -> uniquePaths1(4, 4));
        Flow.executor(() -> uniquePaths1(4, 5));
        Flow.executor(() -> uniquePaths1(4, 6));
        Flow.executor(() -> uniquePaths1(10, 10));
    }

    public int uniquePaths(int m, int n) {
        return uniquePaths(1, 1, m, n);
    }

    public int uniquePaths(int y, int x, int m, int n) {
        if (x > n || y > m) {
            return 0;
        } else if (x == n && y == m) {
            return 1;
        }
        return uniquePaths(y + 1, x, m, n) + uniquePaths(y, x + 1, m, n);
    }

    public int uniquePaths1(int m, int n) {
        if(m==1){
            return 1;
        }
        int[] arrays=new int[n];
        for (int i = 0; i < n; i++) {
            arrays[i]=1;
        }
        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                arrays[j]=arrays[j-1]+arrays[j];
            }
        }
        return arrays[n-1];
    }

}
