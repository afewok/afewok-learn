package com.afewok.learn.sort;

import org.testng.annotations.Test;

import java.util.Arrays;

/**
 * 选择排序
 * 空间复杂度
 * 时间复杂度
 * 稳定性：
 */
public class SelectionSort extends Sort {

    @Test
    public void testSelectionSort1() {
        Arrays.stream(arrs).forEach(arr -> {
            int[] temp = new int[arr.length];
            System.arraycopy(arr, 0, temp, 0, temp.length);
            executor(() -> selectionSort1(temp));
        });
    }

    /**
     * 1、基本实现
     */
    public int[] selectionSort1(int[] nums) {

        return nums;
    }


    @Test
    public void testSelectionSort2() {
        Arrays.stream(arrs).forEach(arr -> {
            int[] temp = new int[arr.length];
            System.arraycopy(arr, 0, temp, 0, temp.length);
            executor(() -> selectionSort2(temp));
        });
    }

    /**
     * 2、轻微优化
     */
    public int[] selectionSort2(int[] nums) {
        return nums;
    }
}
