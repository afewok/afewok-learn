package com.afewok.learn.leetcode;

import com.afewok.learn.util.Json;

public class Flow {

    public static <T> T executor(IExecutor<T> executor) {
        long timestamp = System.nanoTime();
        T result = executor.execute();
        System.out.printf("结果：%s  耗时：%f ns\n", Json.toJSONString(result), (System.nanoTime() - timestamp) / 1000000D);
        return result;
    }

    public interface IExecutor<T> {
        T execute();
    }
}
