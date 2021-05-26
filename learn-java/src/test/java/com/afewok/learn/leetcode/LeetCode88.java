package com.afewok.learn.leetcode;

import com.afewok.learn.util.Json;

import org.testng.annotations.Test;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * 88. 合并两个有序数组
 */
public class LeetCode88 {

    @Test
    public void leetCode() {
        int[] a=new int[] {1};
        int[] b=new int[] {};
        merge(a,a.length,b,b.length);
        System.out.println(Json.toJSONString(a));

        a=new int[] {1};
        b=new int[] {1};
        merge(a,0,b,b.length);
        System.out.println(Json.toJSONString(a));
    }

    public void merge(int A[], int m, int B[], int n) {
        int sub=m+n;
        while(sub>0){
            sub--;
            if(n==0){
                A[sub]=A[--m];
            }else if(m==0){
                A[sub]=B[--n];
            }else{
                if(A[m-1]>B[n-1]){
                    A[sub]=A[--m];
                }else{
                    A[sub]=B[--n];
                }
            }
        }
    }
}
