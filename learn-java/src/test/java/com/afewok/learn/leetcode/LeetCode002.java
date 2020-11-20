package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

import lombok.AllArgsConstructor;

/**
 * 2. 两数相加
 * 
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 * 
 * 示例： 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4) 输出：7 -> 0 -> 8 原因：342 + 465 = 807
 * 
 * 思路：
 */
public class LeetCode002 {

    @AllArgsConstructor
    public class ListNode {
        int val;
        ListNode next;
    }

    @Test
    public void leetCode002() {
        ListNode l1 = new ListNode(2, new ListNode(4, new ListNode(3, null)));
        ListNode l2 = new ListNode(5, new ListNode(6, new ListNode(4, null)));

        print(l1);
        print(l2);
        ListNode result = addTwoNumbers(l1, l2);
        print(result);
    }

    public void print(ListNode node) {
        while (node != null) {
            System.out.print(node.val);
            node = node.next;
            if (node != null) {
                System.out.print(" -> ");
            }
        }
        System.out.println();
    }

    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode dummy = new ListNode(0, null);
        while (l1 != null || l2 != null) {
            if(l1==null){

            }
        }

        return dummy.next;
    }
}
