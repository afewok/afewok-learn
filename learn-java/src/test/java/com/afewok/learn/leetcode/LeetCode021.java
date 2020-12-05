package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * 21. 合并两个有序链表
 * 
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 * 
 * 示例：输入：1->2->4, 1->3->4 输出：1->1->2->3->4->4
 * 
 * 思路：递归、迭代
 */
public class LeetCode021 {
    @Test
    public void leetCode021() {
        ListNode l11 = new ListNode(1, new ListNode(2, new ListNode(4, null)));
        ListNode l12 = new ListNode(1, new ListNode(3, new ListNode(4, null)));
        ListNode result1 = mergeTwoLists1(l11, l12);

        ListNode l21 = new ListNode(1, new ListNode(2, new ListNode(4, null)));
        ListNode l22 = new ListNode(1, new ListNode(3, new ListNode(4, null)));
        ListNode result2 = mergeTwoLists1(l21, l22);
        print(result1);
        print(result2);
    }

    public void print(ListNode head) {
        while (head != null) {
            System.out.print(head.val + "->");
            head = head.next;
        }
        System.out.println("NULL");
    }

    @Data
    @NoArgsConstructor
    @AllArgsConstructor
    public class ListNode {
        int val;
        ListNode next;
    }

    public ListNode mergeTwoLists1(ListNode l1, ListNode l2) {
        ListNode result = new ListNode(), p = result;
        while (l1 != null || l2 != null) {
            if (l1 != null && l2 != null) {
                if (l1.val < l2.val) {
                    p.next = l1;
                    p = p.next;
                    l1 = l1.next;
                } else {
                    p.next = l2;
                    p = p.next;
                    l2 = l2.next;
                }
            } else if (l1 != null) {
                p.next = l1;
                p = p.next;
                l1 = l1.next;
            } else if (l2 != null) {
                p.next = l2;
                p = p.next;
                l2 = l2.next;
            }
        }
        return result.next;
    }

    public ListNode mergeTwoLists2(ListNode l1, ListNode l2) {
        if (l1 == null) {
            return l2;
        } else if (l2 == null) {
            return l1;
        } else if (l1.val < l2.val) {
            l1.next = mergeTwoLists2(l1.next, l2);
            return l1;
        }
        l2.next = mergeTwoLists2(l1, l2.next);
        return l2;
    }
}
