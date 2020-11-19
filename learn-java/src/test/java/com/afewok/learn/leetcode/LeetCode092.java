package com.afewok.learn.leetcode;

import lombok.AllArgsConstructor;
import org.testng.annotations.Test;

/**
 * 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
 * <p>
 * 说明:
 * 1 <= m <= n <= 链表长度。
 * <p>
 * 示例:
 * 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 * 输出: 1->4->3->2->5->NULL
 */
public class LeetCode092 {

    @AllArgsConstructor
    public class ListNode {
        int val;
        ListNode next;
    }

    @Test
    public void leetCode092() {
        int m = 2, n = 4;
        ListNode head1 = new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(4, new ListNode(5, null)))));
        print(head1);
        head1 = reverseBetween2(head1, m, n);
        print(head1);

        ListNode head2 = new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(4, new ListNode(5, null)))));
        print(head2);
        head2 = reverseBetween2(head2, m, n);
        print(head2);
    }

    public void print(ListNode head) {
        while (head != null) {
            System.out.print(head.val + "->");
            head = head.next;
        }
        System.out.println("NULL");
    }

    public ListNode reverseBetween1(ListNode head, int m, int n) {
        if (m > n || m < 1) {
            throw new RuntimeException("异常");
        }

        ListNode first = null;
        ListNode two = null;
        ListNode three = head;
        ListNode l = null;
        ListNode ml = null;

        int sub = 1;

        while (three != null) {
            if (sub == m) {
                l = two;
                ml = three;
            } else if (sub - 1 > m && sub <= n) {
                two.next = first;
            }

            if (sub == n) {
                if (l != null) {
                    l.next = three;
                }
                ml.next = three.next;
                if (m != n) {
                    three.next = two;
                }
                break;
            }

            sub++;
            first = two;
            two = three;
            three = three.next;
        }
        return m == 1 ? three : head;
    }


    public ListNode reverseBetween2(ListNode head, int m, int n) {
        ListNode virtual = new ListNode(0, head.next);
        ListNode p = virtual;
        ListNode temp = null;
        for (int i = 1; i < m; i++) {
            p = p.next;
        }
        head = p.next;
        for (int i = m; i < n; i++) {
            temp =head.next;


            p.next=temp;


        }
        return virtual.next;
    }
}