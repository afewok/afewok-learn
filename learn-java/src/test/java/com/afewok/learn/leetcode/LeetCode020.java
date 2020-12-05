package com.afewok.learn.leetcode;

import java.util.HashMap;
import java.util.Map;
import java.util.Stack;

import org.testng.annotations.Test;

/**
 * 20. 有效的括号
 * 
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 * 有效字符串需满足：左括号必须用相同类型的右括号闭合。左括号必须以正确的顺序闭合。注意空字符串可被认为是有效字符串。
 * 
 * 示例 1: 输入: "()" 输出: true
 * 
 * 示例 2: 输入: "()[]{}" 输出: true
 * 
 * 示例 3: 输入: "(]" 输出: false
 * 
 * 示例 4: 输入: "([)]" 输出: false
 * 
 * 示例 5: 输入: "{[]}" 输出: true
 */
public class LeetCode020 {

    @Test
    public void leetCode020() {
        System.out.println(isValid("()"));
        System.out.println(isValid("()[]{}"));
        System.out.println(isValid("(]"));
        System.out.println(isValid("([)]"));
        System.out.println(isValid("{[]}"));

    }

    private static final Map<Character, Character> mapKv = new HashMap<>(3);
    private static final Map<Character, Character> mapVk = new HashMap<>(3);
    static {
        mapKv.put('(', ')');
        mapKv.put('{', '}');
        mapKv.put('[', ']');

        mapVk.put(')', '(');
        mapVk.put('}', '{');
        mapVk.put(']', '[');
    }

    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();
        int length = s.length();
        for (int i = 0; i < length; i++) {
            char c = s.charAt(i);

            if (mapVk.get(c) != null) {
                if (stack.size() == 0) {
                    return false;
                }
                char key = mapVk.get(c);
                if (stack.peek() != key) {
                    return false;
                }
                stack.pop();
            } else if (mapKv.get(c) != null) {
                stack.push(c);
            }
        }
        return stack.size() == 0 ? true : false;
    }
}
