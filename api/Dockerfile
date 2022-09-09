ARG BASE_IMAGE=gcr.io/distroless/static:latest-amd64
FROM $BASE_IMAGE

COPY build/main /service/main
COPY .env /service/.env

CMD ["/service/main"]