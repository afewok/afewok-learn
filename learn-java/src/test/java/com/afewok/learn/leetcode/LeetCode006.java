package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

/**
 * 6. Z 字形变换
 * 
 * 将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
 * 
 * 比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
 * 
 * L _ C _ I _ R
 * 
 * E T O E S I I G
 * 
 * E _ D _ H _ N
 * 
 * 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
 * 
 * 请你实现这个将字符串进行指定行数变换的函数：string convert(string s, int numRows);
 * 
 * 示例 1: 输入: s = "LEETCODEISHIRING", numRows = 3 输出: "LCIRETOESIIGEDHN"
 * 
 * 示例 2: 输入: s = "LEETCODEISHIRING", numRows = 4 输出: "LDREOEIIECIHNTSG"
 * 
 * 解释:
 * 
 * L _ _ D _ _ R
 * 
 * E _ O E _ I I
 * 
 * E C _ I H _ N
 * 
 * T S G
 */
public class LeetCode006 {

    @Test
    public void leetCode006() {
        String s = "LEETCODEISHIRING";
        System.out.println(convert1(s, 3).equals("LCIRETOESIIGEDHN"));
        System.out.println(convert1(s, 4).equals("LDREOEIIECIHNTSG"));
    }

    public String convert1(String s, int numRows) {
        if (numRows < 2) {
            return s;
        }
        char[] chats = s.toCharArray();
        int length = chats.length, interval = numRows*2-2,group=(length+interval-1)/interval,sub=0;
        char[] result = new char[length];
        // for (int i = 0; i < group; i++) {
        //     int m=i*interval;
        //     int n=m+interval;

        //     result[sub++]=chats[];
        //     int temp=i+interval;

        // }

        // for (int i = 0; i < interval; i++) {
        //     result[sub++]=chats[i*interval];
        // }
        return String.valueOf(result);
    }
}
