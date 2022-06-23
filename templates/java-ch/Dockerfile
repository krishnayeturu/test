FROM openjdk:17-alpine

ARG JARFILE=REPLACE_ME_USING_BUILD_ARG
ENV JARFILE=$JARFILE
ENV LIBSTDCpp_VERSION 10.3

RUN apk add --update --no-cache \
    libstdc++~${LIBSTDCpp_VERSION}

COPY build/libs/$JARFILE /app/$JARFILE

WORKDIR /app

CMD java -jar ${JARFILE}