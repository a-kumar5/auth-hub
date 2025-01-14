FROM golang:1.22-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o auth-hub cmd/main.go

CMD ["/app/auth-hub"]