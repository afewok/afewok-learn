package com.afewok.learn.leetcode;

import java.util.*;

import org.testng.annotations.Test;

/**
 * 4. 寻找两个正序数组的中位数
 * 
 * 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。 进阶：你能设计一个时间复杂度为
 * O(log (m+n)) 的算法解决此问题吗？
 *
 * 示例 1：输入：nums1 = [1,3], nums2 = [2] 输出：2.00000 解释：合并数组 = [1,2,3] ，中位数 2
 *
 * 示例 2：输入：nums1 = [1,2], nums2 = [3,4] 输出：2.50000 解释：合并数组 = [1,2,3,4] ，中位数 (2 +
 * 3) / 2 = 2.5
 *
 * 示例 3：输入：nums1 = [0,0], nums2 = [0,0] 输出：0.00000
 *
 * 示例 4：输入：nums1 = [], nums2 = [1] 输出：1.00000
 *
 * 示例 5：输入：nums1 = [2], nums2 = [] 输出：2.00000
 *
 * 提示：nums1.length == m nums2.length == n 0 <= m <= 1000 0 <= n <= 1000 1 <= m +
 * n <= 2000 -106 <= nums1[i], nums2[i] <= 106
 *
 * 思路：合并排序取值、下标记录、二分查找、划分数组
 */
public class LeetCode004 {

    @Test
    public void leetCode004() {
        int[] nums1 = { 1, 2 };
        int[] nums2 = { 3, 4 };
        System.out.println(findMedianSortedArrays1(nums1, nums2));
        System.out.println(findMedianSortedArrays2(nums1, nums2));
        System.out.println(findMedianSortedArrays3(nums1, nums2));
        System.out.println(findMedianSortedArrays4(nums1, nums2));
    }

    public double findMedianSortedArrays1(int[] nums1, int[] nums2) {
        int length = nums1.length + nums2.length;
        int pair = length / 2;
        List<Integer> list = new ArrayList<>(length);
        Arrays.stream(nums1).forEach(list::add);
        Arrays.stream(nums2).forEach(list::add);
        list.sort(Comparator.comparingDouble(t -> t));
        return pair * 2 == length ? (list.get(pair - 1) + list.get(pair)) / 2.0 : list.get(pair);
    }

    public double findMedianSortedArrays2(int[] nums1, int[] nums2) {
        int length = nums1.length + nums2.length;
        int pair = length / 2;
        int last = 0, current = 0;
        for (int i = 0, n1 = 0, n2 = 0; i < length; i++) {
            last = current;
            if (nums1.length > n1 && nums2.length > n2) {
                if (nums1[n1] <= nums2[n2]) {
                    current = nums1[n1];
                    n1++;
                } else {
                    current = nums2[n2];
                    n2++;
                }
            } else if (nums1.length > n1) {
                current = nums1[n1];
                n1++;
            } else {
                current = nums2[n2];
                n2++;
            }

            if (pair == i) {
                return pair * 2 == length ? (last + current) / 2.0 : current;
            }
        }
        throw new IllegalArgumentException("Illegal nums1 or nums2");
    }

    public double findMedianSortedArrays3(int[] nums1, int[] nums2) {

        return 0;
    }

    public double findMedianSortedArrays4(int[] nums1, int[] nums2) {

        return 0;
    }
}
