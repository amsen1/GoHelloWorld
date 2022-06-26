# stage 1: building application binary file
FROM golang:bullseye AS builder

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .
# RUN ls -laRt

RUN go build -o main ./src/main.go

# stage 2: copy only the application binary file and necessary files to the alpine container
FROM debian:bullseye-slim AS production
# RUN apk --update add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

ENTRYPOINT [ "/app/main" ]