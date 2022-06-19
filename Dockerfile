# stage 1: building application binary file
FROM golang:alpine3.16 AS builder

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN go build -mod vendor -o main main.go

# stage 2: copy only the application binary file and necessary files to the alpine container
FROM alpine:3.16 AS production
RUN apk --update add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

ENTRYPOINT [ "/app/main" ]