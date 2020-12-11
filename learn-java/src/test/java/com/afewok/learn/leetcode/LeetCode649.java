package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

/**
 * 649. Dota2 参议院
 * 
 * 思路：
 */
public class LeetCode649 {

    @Test
    public void leetCode() {
        Flow.executor(() -> predictPartyVictory("RD"));
        Flow.executor(() -> predictPartyVictory("RDD"));
        Flow.executor(() -> predictPartyVictory("DRRDRDRDRDDRDRDR"));
    }

    public String predictPartyVictory1(String senate) {
        char[] ch = senate.toCharArray();
        int i = 0, j = 0, length = ch.length;
        while (true) {
            char temp = ch[i];
            if (ch[i] != '.') {
                j = i + 1;
                while (true) {
                    j = j % length;
                    if (ch[j] != temp && ch[j] != '.') {
                        ch[j] = '.';
                        break;
                    } else if (i == j) {
                        break;
                    }
                    j++;

                }
                if (i == j) {
                    return temp == 'R' ? "Radiant" : "Dire";
                }
            }
            i++;
            i = i % length;
        }
    }

    public String predictPartyVictory2(String senate) {
        char[] ch = senate.toCharArray();
        int i = -1, j = 0, length = ch.length;
        while (true) {
            i = (++i) % length;
            if (ch[i] == '.') {
                continue;
            }
            j = i;
            while (true) {
                j = (++j) % length;
                if (ch[j] == '.') {
                    continue;
                } else if (ch[j] != ch[i]) {
                    ch[j] = '.';
                    break;
                } else if (j == i) {
                    return ch[i] == 'R' ? "Radiant" : "Dire";
                }
            }
        }
    }

    public String predictPartyVictory3(String senate) {
        char[] ch = senate.toCharArray();
        int i = -1, j = 0, length = ch.length;
        while (true) {
            if (++i == length) {
                i = 0;
            }
            if (ch[i] == '.') {
                continue;
            }
            j = i;
            while (true) {
                if (++j == length) {
                    j = 0;
                }
                if (ch[j] == '.') {
                    continue;
                } else if (ch[j] != ch[i]) {
                    ch[j] = '.';
                    break;
                } else if (j == i) {
                    return ch[i] == 'R' ? "Radiant" : "Dire";
                }
            }
        }
    }

    public class ListNode {
        char val;
        ListNode next;

        public ListNode(char val, ListNode next) {
            this.val = val;
            this.next = next;
        }
    }

    public String predictPartyVictory(String senate) {
        char[] ch = senate.toCharArray();
        int length = ch.length;
        ListNode p = new ListNode(ch[0], null), m = p, n = null;
        for (int i = 1; i < length; i++) {
            m.next = new ListNode(ch[i], null);
            m = m.next;
        }
        m.next = p;

        while (true) {
            m = p;
            n = p.next;
            while (true) {
                if (n == p) {
                    return p.val == 'R' ? "Radiant" : "Dire";
                } else if (n.val != p.val) {
                    m.next = n.next;
                    n = n.next;
                    break;
                } else {
                    m = m.next;
                    n = n.next;
                }
            }
            p = p.next;
        }
    }

//"DRRDRDRDRDDRDRDR"
//     public String predictPartyVictory(String senate) {
//         int r = 0, d = 0, fr = 0, fd = 0,hr=-1;
//         char[] ch = senate.toCharArray();
//         for (char c : ch) {
//             if (c == 'R') {
//                 if (fr > 0) {
//                     fr--;
//                 } else {
//                     r++;
//                     fd++;
//                 }
//             } else {
//                 if (fd > 0) {
//                     fd--;
//                 } else {
//                     d++;
//                     fr++;
//                 }
//             }
//         }
//         r = r - Math.abs(fr);
//         d = d - Math.abs(fd);
//         if (r > d) {
//             return "Radiant";
//         } else if (r < d) {
//             return "Dire";
//         } else if (ch[0] == 'R') {
//             return "Radiant";
//         }
//         return "Dire";
//     }
}