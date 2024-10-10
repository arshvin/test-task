FROM docker.io/gradle:8.8-jdk21 AS builder
WORKDIR /home/gradle/project
COPY ./ ./
RUN gradle build --info

FROM docker.io/openjdk:21-jdk as webapp
WORKDIR /opt/webapp
COPY --from=builder /home/gradle/project/build/libs/webapp-spring-boot.jar ./
CMD ["/bin/java", "-jar", "/opt/webapp/webapp-spring-boot.jar"]
