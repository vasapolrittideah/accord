FROM alpine:latest
WORKDIR /app
COPY .env tmp/ ./
EXPOSE 8080
CMD ["./server"]
