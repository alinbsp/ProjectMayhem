package org.mayhem;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
@RestController
@EnableAutoConfiguration
public class MayhemApplication {

    @RequestMapping("/")
    String home() {
        return "Mayhem";
    }

    public static void main(String[] args) {
        SpringApplication.run(MayhemApplication.class, args);
    }
}
