package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

/**
 * 860. 柠檬水找零 在柠檬水摊上，每一杯柠檬水的售价为 5 美元。顾客排队购买你的产品，（按账单 bills 支付的顺序）一次购买一杯。
 * 每位顾客只买一杯柠檬水，然后向你付 5 美元、10 美元或 20 美元。你必须给每个顾客正确找零，也就是说净交易是每位顾客向你支付 5 美元。
 * 注意，一开始你手头没有任何零钱。如果你能给每位顾客正确找零，返回 true ，否则返回 false 。
 * 
 * 示例 1：输入：[5,5,5,10,20] 输出：true 解释：前 3 位顾客那里，我们按顺序收取 3 张 5 美元的钞票。第 4
 * 位顾客那里，我们收取一张 10 美元的钞票，并返还 5 美元。 第 5 位顾客那里，我们找还一张 10 美元的钞票和一张 5
 * 美元的钞票。由于所有客户都得到了正确的找零，所以我们输出 true。
 * 
 * 示例 2：输入：[5,5,10] 输出：true
 * 
 * 示例 3：输入：[10,10] 输出：false
 * 
 * 示例 4：输入：[5,5,10,10,20] 输出：false 解释： 前 2 位顾客那里，我们按顺序收取 2 张 5 美元的钞票。对于接下来的 2
 * 位顾客，我们收取一张 10 美元的钞票，然后返还 5 美元。 对于最后一位顾客，我们无法退回 15 美元，因为我们现在只有两张 10
 * 美元的钞票。由于不是每位顾客都得到了正确的找零，所以答案是 false。
 * 
 * 提示：0 <= bills.length <= 10000 bills[i] 不是 5 就是 10 或是 20 
 * 
 * 思路：模拟 + 贪心
 */
public class LeetCode860 {

    @Test
    public void leetCode842() {
        Flow.executor(() -> lemonadeChange2(new int[] { 5, 5, 5, 10, 20 }));
        Flow.executor(() -> lemonadeChange2(new int[] { 5, 5, 10 }));
        Flow.executor(() -> lemonadeChange2(new int[] { 10, 10 }));
        Flow.executor(() -> lemonadeChange2(new int[] { 5, 5, 10, 10, 20 }));
    }

    public boolean lemonadeChange1(int[] bills) {
        return lemonadeChange(0, 0, 0, bills);
    }

    public boolean lemonadeChange(int m5, int m10, int startSub, int[] bills) {
        int length = bills.length;
        for (int i = startSub; i < length; i++) {
            if (bills[i] == 5) {
                m5 = m5 + 1;
            } else if (bills[i] == 10) {
                if (m5 < 1) {
                    return false;
                }
                m5 = m5 - 1;
                m10 = m10 + 1;
            } else {
                if (m5 == 0 || (m5 < 3 && m10 == 0)) {
                    return false;
                }
                if (m5 > 2) {
                    if (lemonadeChange(m5 - 3, m10, i + 1, bills)) {
                        return true;
                    }
                }
                if (m5 > 0 && m10 > 0) {
                    if (lemonadeChange(m5 - 1, m10 - 1, i + 1, bills)) {
                        return true;
                    }
                }
                return false;
            }
        }
        return true;
    }

    public boolean lemonadeChange2(int[] bills) {
        int m5 = 0, m10 = 0;
        for (int bill : bills) {
            if (bill == 5) {
                m5++;
            } else if (bill == 10) {
                if (m5 > 0) {
                    m5--;
                    m10++;
                } else {
                    return false;
                }
            } else {
                if (m5 > 0 && m10 > 0) {
                    m5--;
                    m10--;
                } else if (m5 > 2) {
                    m5 = m5 - 3;
                } else {
                    return false;
                }
            }
        }
        return true;
    }

    public boolean lemonadeChange3(int[] bills) {
        int m5 = 0, m10 = 0;
        for (int bill : bills) {
            switch (bill) {
                case 5:
                    m5++;
                    break;
                case 10:
                    if (m5 > 0) {
                        m5--;
                        m10++;
                    } else {
                        return false;
                    }
                    break;
                case 20:
                    if (m5 > 0 && m10 > 0) {
                        m5--;
                        m10--;
                    } else if (m5 > 2) {
                        m5 = m5 - 3;
                    } else {
                        return false;
                    }
                    break;
            }
        }
        return true;
    }

    public boolean lemonadeChange4(int[] bills) {
        int m5 = 0, m10 = 0;
        for (int bill : bills) {
            if (bill == 5) {
                m5++;
            } else if (bill == 10) {
                m5--;
                m10++;
            } else if (m10 > 0) {
                m5--;
                m10--;
            } else {
                m5 = m5 - 3;
            }
            if (m5 < 0 || m10 < 0) {
                return false;
            }
        }
        return true;
    }

    public boolean lemonadeChange5(int[] bills) {
        int[] array = { 0, 0, 0, 0, 0 };
        for (int bill : bills) {
            array[bill / 5]++;
            array[bill / 10]--;
            array[bill / 20]--;
            if (array[1] < 0 || array[1] + 2 * array[2] < 0) {
                return false;
            }
        }
        return true;
    }
}
