ARG TAG=8-jre-slim

FROM alpine/curl as builder

ARG ARTIFACT_URL="https://launcher.mojang.com/v1/objects/0a269b5f2c5b93b1712d0f5dc43b6182b9ab254e/server.jar"

WORKDIR /app/tmp

RUN curl ${ARTIFACT_URL} -o /app/tmp/server.jar

FROM openjdk:${TAG}

COPY --from=builder /app/tmp/server.jar /server/

WORKDIR server

RUN echo "eula=true" > eula.txt

LABEL org.opencontainers.image.description='Minecraft Java Edition version ${VERSION}. See Changelog: ${VERSION_URL}.'
LABEL org.opencontainers.image.url='ghcr.io/hostfactor/minecraft-server'
LABEL org.opencontainers.image.version=${VERSION}
LABEL org.opencontainers.image.authors='eddie@hostfactor.io'

ENTRYPOINT ["java", "-jar", "server.jar"]
