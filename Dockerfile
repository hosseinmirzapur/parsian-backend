FROM golang:1.20-alpine as build
COPY . ./
WORKDIR /app

RUN go mod download && GOOS=linux GOARCH=arm64 go build -o main && chmod +x ./main

FROM alpine:latest

RUN apk --no-cache add bash

WORKDIR /app

COPY --from=build /app/main .

CMD [ "./main" ]

EXPOSE 3000