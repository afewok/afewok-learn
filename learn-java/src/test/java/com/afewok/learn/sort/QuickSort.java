package com.afewok.learn.sort;

import org.testng.annotations.Test;

import java.util.Arrays;

/**
 * 快速排序
 * 空间复杂度
 * 时间复杂度
 * 稳定性：
 */
public class QuickSort extends Sort {

    @Test
    public void testQuickSort1() {
        Arrays.stream(arrs).forEach(arr -> {
            int[] temp = new int[arr.length];
            System.arraycopy(arr, 0, temp, 0, temp.length);
            executor(() -> quickSort1(temp));
        });
    }

    /**
     * 1、填坑法
     */
    public int[] quickSort1(int[] nums) {

        return nums;
    }


    @Test
    public void testQuickSort2() {
        Arrays.stream(arrs).forEach(arr -> {
            int[] temp = new int[arr.length];
            System.arraycopy(arr, 0, temp, 0, temp.length);
            executor(() -> quickSort2(temp));
        });
    }

    /**
     * 2、交换法
     */
    public int[] quickSort2(int[] nums) {
        return nums;
    }

    @Test
    public void testQuickSort3() {
        Arrays.stream(arrs).forEach(arr -> {
            int[] temp = new int[arr.length];
            System.arraycopy(arr, 0, temp, 0, temp.length);
            executor(() -> quickSort3(temp));
        });
    }

    /**
     * 3、顺序遍历法
     */
    public int[] quickSort3(int[] nums) {
        return nums;
    }

    @Test
    public void testQuickSort4() {
        Arrays.stream(arrs).forEach(arr -> {
            int[] temp = new int[arr.length];
            System.arraycopy(arr, 0, temp, 0, temp.length);
            executor(() -> quickSort4(temp));
        });
    }

    /**
     * 4、另类交换法
     */
    public int[] quickSort4(int[] nums) {
        return nums;
    }
}
