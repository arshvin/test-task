FROM gradle:6.6.1-jdk11-openj9 AS builder
WORKDIR /home/gradle/project
COPY ./ ./
RUN gradle build --info

FROM openjdk:11 as webapp
WORKDIR /opt/webapp
COPY --from=builder /home/gradle/project/build/libs/webapp-spring-boot.jar ./
CMD ["/bin/java", "-jar", "/opt/webapp/webapp-spring-boot.jar"]
