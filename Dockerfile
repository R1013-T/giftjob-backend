FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine
RUN apk update && apk upgrade
RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/main ./main

CMD ["./main"]