package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

/**
 * 37. 解数独
 * 
 * 编写一个程序，通过填充空格来解决数独问题。
 * 一个数独的解法需遵循如下规则：数字 1-9 在每一行只能出现一次。数字 1-9 在每一列只能出现一次。数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
 * 
 * 空白格用 '.' 表示。
 * 
 * 提示：给定的数独序列只包含数字 1-9 和字符 '.' 。你可以假设给定的数独只有唯一解。给定数独永远是 9x9 形式的。
 */
public class LeetCode037 {
    @Test
    public void leetCode037() {
        char[][] board={
            {'5','3','.','.','7','.','.','.','.'},
            {'6','.','.','1','9','5','.','.','.'},
            {'.','9','8','.','.','.','.','6','.'},
            {'8','.','.','.','6','.','.','.','3'},
            {'4','.','.','8','.','3','.','.','1'},
            {'7','.','.','.','1','.','.','.','6'},
            {'.','6','.','.','.','.','2','8','.'},
            {'.','.','.','4','1','9','.','.','5'},
            {'.','.','.','.','8','.','.','7','9'}
        };
        solveSudoku(board);
        print(board);
    }

    private void print(char[][] board){
        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[i].length; j++) {
                System.out.print(board[i][j]+" ");
            }
            System.out.println();
        }
    }

    public void solveSudoku(char[][] board) {

    }
}
