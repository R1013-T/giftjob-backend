FROM golang:1.19-alpine as builder

WORKDIR /app

ENV DB_URL="postgres://takahashi:TestPostgres01@rt-test-sqlserver.postgres.database.azure.com:5432/postgres"
ENV JWE_SECRET="Ltf1BuN6W416gwmiLvM9xHxlDInxGQq13UzhzJLE65I="
ENV FRONTEND_URL="https://test-nextauth-apollo.vercel.app"

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