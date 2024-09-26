package prived.medved;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication(scanBasePackages = {"prived.medved.controllers"})
public class App {

  public static void main(final String[] args) {
    SpringApplication.run(App.class, args);
  }
}
