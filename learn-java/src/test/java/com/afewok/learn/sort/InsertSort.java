package com.afewok.learn.sort;

import org.testng.annotations.Test;

import java.util.Arrays;

/**
 * 插入排序
 * 空间复杂度
 * 时间复杂度
 * 稳定性：
 */
public class InsertSort extends Sort {

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
        int[] result = new int[nums.length];
        int count = 0;

        result[count++] = nums[0];
        for (int i = 1; i < nums.length; i++) {
            int k = 0;
            for (; k < count; k++) {
                if (nums[i] < result[k]) {
                    for (int m = count; m > k; m--) {
                        result[m] = result[m] + result[m - 1];
                        result[m - 1] = result[m] - result[m - 1];
                        result[m] = result[m] - result[m - 1];
                    }
                    break;
                }
            }
            count++;
            result[k] = nums[i];
        }
        return result;
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
        for (int i = 1; i < nums.length; i++) {
            for (int k = i - 1; k >= 0; k--) {
                if (nums[k] > nums[k + 1]) {
                    nums[k] = nums[k] + nums[k + 1];
                    nums[k + 1] = nums[k] - nums[k + 1];
                    nums[k] = nums[k] - nums[k + 1];
                } else {
                    break;
                }
            }
        }
        return nums;
    }
}
