FROM golang:1.15.2-alpine as build

RUN apk add --update --no-cache git

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
CMD ls

RUN GOOS=linux GOARCH=amd64 go build -o seat_info .

FROM alpine:3.10

RUN apk add --update --no-cache ca-certificates && \
    apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=build /app /app

EXPOSE 8080

ENTRYPOINT ["/app/seat_info"]