package com.afewok.learn;

import lombok.extern.slf4j.Slf4j;
import org.testng.annotations.BeforeMethod;
import org.testng.annotations.Test;

/**
 * 默认unit
 */
@Slf4j
public class DefaultUnitTest extends AbstractUnitTest {

    @BeforeMethod
    public void unitTestBefore() {
        System.out.println(applicationContext.getBeanDefinitionCount() + "个Bean的名字如下：");
        String[] beanDefinitionNames = applicationContext.getBeanDefinitionNames();
        for (String beanName : beanDefinitionNames) {
            System.out.println(beanName);
        }
    }

    @Test
    public void helloWorld() {
        log.info("hello unit test!");
    }
}
