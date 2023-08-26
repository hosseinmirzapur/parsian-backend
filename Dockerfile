FROM golang:1.20-alpine as build
WORKDIR /app
COPY . ./

RUN go mod init github.com/hosseinmirzapur/parsian-backend && \
    GOPROXY=https://goproxy.io,direct go mod tidy && \
    go build -o main ./cmd/main.go && \
    chmod +x ./main

FROM debian:latest

RUN apk --no-cache add bash

WORKDIR /app

COPY --from=build /app/main .

CMD [ "./main" ]

EXPOSE 3000