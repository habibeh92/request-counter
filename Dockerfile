FROM golang:latest

WORKDIR /app

COPY . /app

RUN mkdir "data"

ENV APP_HTTP_HOST="localhost"
ENV APP_HTTP_PORT=8080

RUN go build -o main ./cmd/server/

EXPOSE 8080

CMD ["./main"]
