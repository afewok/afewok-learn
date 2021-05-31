package com.afewok.learn.sort;

import org.testng.annotations.Test;

import java.util.Arrays;

/**
 * 冒泡排序
 * 空间复杂度:(1)
 * 时间复杂度:T(n)~T(n*n)
 * 稳定性：稳定
 */
public class BubbleSort extends Sort {

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
        for (int i = 0; i < nums.length; i++) {
            for (int j = 1; j < nums.length; j++) {
                if (nums[j - 1] > nums[j]) {
                    nums[j - 1] = nums[j - 1] + nums[j];
                    nums[j] = nums[j - 1] - nums[j];
                    nums[j - 1] = nums[j - 1] - nums[j];
                }
            }
        }
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
        boolean OK = true;
        for (int i = 1; i < nums.length && OK; i++) {
            OK = false;
            for (int j = 0; j < nums.length - i; j++) {
                if (nums[j] > nums[j + 1]) {
                    OK = true;
                    nums[j] = nums[j] + nums[j + 1];
                    nums[j + 1] = nums[j] - nums[j + 1];
                    nums[j] = nums[j] - nums[j + 1];
                }
            }
        }
        return nums;
    }
}
