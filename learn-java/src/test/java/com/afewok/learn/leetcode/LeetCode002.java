package com.afewok.learn.leetcode;

import java.util.List;

import org.testng.annotations.Test;

import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;

/**
 * 2. 两数相加
 * 
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 * 
 * 示例： 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4) 输出：7 -> 0 -> 8 原因：342 + 465 = 807
 * 
 * 思路：虚拟头结点遍历、递归维护一个节点
 */
public class LeetCode002 {

    @NoArgsConstructor
    @AllArgsConstructor
    public class ListNode {
        int val;
        ListNode next;
    }

    @Test
    public void leetCode002() {
        ListNode[] list1 = preDate();

        print(list1[0]);
        print(list1[1]);
        ListNode result1 = addTwoNumbers1(list1[0], list1[1]);
        print(result1);

        ListNode[] list2 = preDate();

        print(list2[0]);
        print(list2[1]);
        ListNode result2 = addTwoNumbers2(list2[0], list2[1]);
        print(result2);
    }

    private ListNode[] preDate() {
        ListNode l1 = new ListNode(9, new ListNode(9,
                new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, null)))))));
        ListNode l2 = new ListNode(9, new ListNode(9, new ListNode(9, null)));
        return new ListNode[] { l1, l2 };
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

    public ListNode addTwoNumbers1(ListNode l1, ListNode l2) {
        boolean addOne = false;
        ListNode dummy = new ListNode();
        ListNode p = dummy;
        while (addOne || l1 != null || l2 != null) {
            p.next = new ListNode(addOne ? 1 : 0, null);
            if (l1 != null) {
                p.next.val += l1.val;
                l1 = l1.next;
            }
            if (l2 != null) {
                p.next.val += l2.val;
                l2 = l2.next;
            }
            addOne = p.next.val > 9 ? true : false;
            p.next.val = p.next.val % 10;
            p = p.next;
        }
        return dummy.next;
    }

    public ListNode addTwoNumbers2(ListNode l1, ListNode l2) {
        l1.val = l1.val + (l2 == null ? 0 : l2.val);
        boolean addOne = l1.val > 9 ? true : false;
        l1.val = l1.val % 10;

        ListNode l2next = (l2 == null ? null : l2.next);
        if (l1.next == null && (addOne || l2next != null)) {
            l1.next = new ListNode();
        }
        if (addOne) {
            l1.next.val++;
        }

        if (l1.next != null || l2next != null) {
            addTwoNumbers2(l1.next, l2next);
        }
        return l1;
    }
}
