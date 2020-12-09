package com.afewok.learn.leetcode;

import java.util.ArrayList;
import java.util.List;

import com.afewok.learn.util.Json;

import org.testng.annotations.Test;

/**
 * 842. 将数组拆分成斐波那契序列
 * 
 * 给定一个数字字符串 S，比如 S = "123456579"，我们可以将它分成斐波那契式的序列 [123, 456,
 * 579]。形式上，斐波那契式序列是一个非负整数列表 F， 且满足：0 <= F[i] <= 2^31 - 1，（也就是说，每个整数都符合 32
 * 位有符号整数类型）；F.length >= 3； 对于所有的0 <= i < F.length - 2，都有 F[i] + F[i+1] = F[i+2]
 * 成立。 另外，请注意，将字符串拆分成小块时，每个块的数字一定不要以零开头，除非这个块是数字 0 本身。 返回从 S
 * 拆分出来的任意一组斐波那契式的序列块，如果不能拆分则返回 []。
 * 
 * 示例 1：输入："123456579" 输出：[123,456,579]
 * 
 * 示例 2：输入: "11235813" 输出: [1,1,2,3,5,8,13]
 * 
 * 示例 3：输入: "112358130" 输出: [] 解释: 这项任务无法完成。
 * 
 * 示例 4：输入："0123" 输出：[] 解释：每个块的数字不能以零开头，因此 "01"，"2"，"3" 不是有效答案。
 * 
 * 示例 5：输入: "1101111" 输出: [110, 1, 111] 解释: 输出 [11,0,11,11] 也同样被接受。
 * 
 * 提示：1 <= S.length <= 200 字符串 S 中只含有数字。
 * 
 */
public class LeetCode842 {

    @Test
    public void leetCode659() {
        System.out.println(Json.toJSONString(splitIntoFibonacci1("1")));
        System.out.println(Json.toJSONString(splitIntoFibonacci1("17522")));
        System.out.println(Json.toJSONString(splitIntoFibonacci1("214748364721474836422147483641")));
        System.out.println(Json.toJSONString(splitIntoFibonacci1("123456579")));
        System.out.println(Json.toJSONString(splitIntoFibonacci1("11235813")));
        System.out.println(Json.toJSONString(splitIntoFibonacci1("112358130")));
        System.out.println(Json.toJSONString(splitIntoFibonacci1("0123")));
        System.out.println(Json.toJSONString(splitIntoFibonacci1("1101111")));
    }

    public List<Integer> splitIntoFibonacci1(String S) {
        List<Integer> list = new ArrayList<>();

        int length = S.length(), half = length / 2;
        for (int first = 1; first <= half; first++) {
            String mStr = S.substring(0, first);
            if (mStr.length() > 1 && mStr.indexOf("0") == 0) {
                continue;
            }
            if(Long.valueOf(mStr)>Integer.MAX_VALUE){
                break;
            }
            for (int two = 1; two <= half; two++) {
                int sub = first;
                String nStr = S.substring(sub, sub + two);
                if (nStr.length() > 1 && nStr.indexOf("0") == 0) {
                    continue;
                }
                if(Long.valueOf(nStr)>Integer.MAX_VALUE){
                    break;
                }

                int mInt = Integer.parseInt(mStr);
                int nInt = Integer.parseInt(nStr);
                list.add(mInt);

                while (S.substring(sub).indexOf(nStr) == 0) {
                    sub = sub + nStr.length();
                    list.add(nInt);
                    nInt = nInt + mInt;
                    mInt = nInt - mInt;
                    nStr = String.valueOf(nInt);
                }

                if (length == sub && list.size() > 2) {
                    return list;
                }
                list.clear();
            }
        }
        return list;
    }
}