package com.afewok.learn;

import org.springframework.test.context.testng.AbstractTestNGSpringContextTests;
import org.springframework.test.context.web.WebAppConfiguration;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;

@WebAppConfiguration
@EnableWebMvc
//@EnableTransactionManagement
public abstract class AbstractUnitTest extends AbstractTestNGSpringContextTests {
//public abstract class AbstractUnitTest extends AbstractTransactionalTestNGSpringContextTests {

}
