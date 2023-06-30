FROM alpine:latest

WORKDIR /workspace

COPY .env dist/apps/api ./

EXPOSE 8080

CMD [ "./api" ]
