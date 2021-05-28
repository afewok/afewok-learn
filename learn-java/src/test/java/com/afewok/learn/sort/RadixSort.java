package com.afewok.learn.sort;

import org.testng.annotations.Test;

import java.util.Arrays;

/**
 * 基数排序
 * 空间复杂度
 * 时间复杂度
 * 稳定性：
 */
public class RadixSort extends Sort {

    @Test
    public void testBubbleSort1() {
        Arrays.stream(arrs).forEach(arr -> {
            int[] temp = new int[arr.length];
            System.arraycopy(arr, 0, temp, 0, temp.length);
            executor(() -> bubbleSort1(temp));
        });
    }

    /**
     * 1、基本实现
     */
    public int[] bubbleSort1(int[] nums) {

        return nums;
    }


    @Test
    public void testBubbleSort2() {
        Arrays.stream(arrs).forEach(arr -> {
            int[] temp = new int[arr.length];
            System.arraycopy(arr, 0, temp, 0, temp.length);
            executor(() -> bubbleSort2(temp));
        });
    }

    /**
     * 2、轻微优化
     */
    public int[] bubbleSort2(int[] nums) {
        return nums;
    }
}
