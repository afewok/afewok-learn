package com.afewok.learn;

import com.google.common.collect.ImmutableMap;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Date;
import java.util.Map;

@RestController
@RequestMapping("/monitor")
@SpringBootApplication
public class StartApp {

    @GetMapping("/alive")
    public Map<String, Object> monitorAlive() {
        return ImmutableMap.<String, Object>builder()
                .put("data", "Hello World!")
                .put("date", new Date())
                .build();
    }

    public static void main(String[] args) {
        SpringApplication.run(StartApp.class, args);
    }

}
