package com.afewok.learn.leetcode;

import org.testng.annotations.Test;

/**
 * 5629. 重新格式化电话号码
 */
public class LeetCodeWeek5629 {
    
    @Test
    public void leetCode() {
        Flow.executor(() -> reformatNumber("1-23-45 6"));
        Flow.executor(() -> reformatNumber("123 4-567"));
        Flow.executor(() -> reformatNumber("123 4-5678"));
        Flow.executor(() -> reformatNumber("12"));
        Flow.executor(() -> reformatNumber("--17-5 229 35-39475 "));
    }

    public String reformatNumber(String number) {
        StringBuilder sb=new StringBuilder();
        for (int i = 0; i < number.length(); i++) {
            char temp=number.charAt(i);
            if(temp =='-'||temp ==' '){
                continue;
            }
            sb.append(temp);
            if(sb.length()%4==3){
                sb.append('-');
            }
        }
        if(sb.charAt(sb.length()-1) =='-'){
            sb.deleteCharAt(sb.length()-1);
        }else if(sb.charAt(sb.length()-2) =='-'){
            sb.deleteCharAt(sb.length()-2);
            sb.insert(sb.length()-2, '-');
        }
        return sb.toString();
    }
}
