ARG OPENJDK_TAG=17-alpine

FROM bash as builder

ARG MINECRAFT_JAR_PATH="https://launcher.mojang.com/v1/objects/125e5adf40c659fd3bce3e66e67a16bb49ecc1b9/server.jar"

WORKDIR /app/tmp

RUN apt-get update && apt-get install -y curl

RUN curl ${MINECRAFT_JAR_PATH} -o /app/tmp/server.jar

FROM openjdk:${OPENJDK_TAG}

COPY --from=builder /app/tmp/server.jar /server/

WORKDIR server

RUN echo "eula=true" > eula.txt

ENTRYPOINT ["java", "-jar", "server.jar"]
