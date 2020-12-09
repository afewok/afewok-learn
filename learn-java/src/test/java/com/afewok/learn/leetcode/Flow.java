package com.afewok.learn.leetcode;

import com.afewok.learn.util.Json;

public class Flow {

    public static <T> T executor(IExecutor<T> executor) {
        long timestamp = System.nanoTime();
        T result = executor.execute();
        System.out.println("结果：" + Json.toJSONString(result) + "  时间戳:" + (System.nanoTime() - timestamp) / 1000000D);
        return result;
    }

    public interface IExecutor<T> {
        T execute();
    }
}
