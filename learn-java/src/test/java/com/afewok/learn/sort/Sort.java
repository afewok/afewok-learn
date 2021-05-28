package com.afewok.learn.sort;

import com.afewok.learn.util.Json;
import org.testng.annotations.AfterMethod;
import org.testng.annotations.BeforeMethod;

import java.util.List;

public abstract class Sort {

    public int[][] arrs = {
            {5, 7, 1, 8, 4},
            {1, 9, 5, 3, 8, 7, 6, 2, 4},
            {3, 7, 8, 5, 2, 1, 9, 5, 4},
            {47, 29, 71, 99, 78, 19, 24, 47},
            {4, 7, 6, 5, 3, 2, 8, 1},
            {0, 5, 1, 9, 6, 7, 19, 11, 3, 8, 4},
    };

    @BeforeMethod
    public void beforeMethod() {
        System.out.println("BeforeMethod");
    }

    @AfterMethod
    public void afterMethod() {
        System.out.println("AfterMethod");
    }

    public interface IExecutor<T> {
        T execute();
    }

    public static <T> T executor(IExecutor<T> executor) {
        long timestamp = System.nanoTime();
        T result = executor.execute();
        boolean b = true;
        String string = Json.toJSONString(result);
        List<Integer> list = Json.parseArray(string, Integer.class);
        for (int i = 1; i < list.size(); i++) {
            if (list.get(i - 1) > list.get(i)) {
                b = false;
                break;
            }
        }
        System.out.printf("结果：%s  耗时：%f ns\n", b, (System.nanoTime() - timestamp) / 1000000D);
        return result;
    }
}
