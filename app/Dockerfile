FROM golang:1.15.2-alpine as build

RUN apk add --update --no-cache git

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o seat_info .

FROM alpine:3.10

RUN apk add --update --no-cache ca-certificates
WORKDIR /app
COPY --from=build /app/seat_info /app/seat_info

EXPOSE 80

ENTRYPOINT ["/app/seat_info"]