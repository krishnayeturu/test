ARG BASE_IMAGE=openjdk:17-slim
FROM $BASE_IMAGE

COPY build/libs/main.jar /service/main.jar
COPY .env /service/.env

WORKDIR /service

CMD java -jar main.jar