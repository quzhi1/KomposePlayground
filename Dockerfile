FROM golang:latest

WORKDIR /opt/app

ADD . /opt/app

EXPOSE 8090

RUN go test

CMD ["go", "run", "main.go"]