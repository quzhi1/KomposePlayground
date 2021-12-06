FROM golang:latest

WORKDIR /opt/app

ADD . /opt/app

EXPOSE 80

RUN go test

CMD ["go", "run", "main.go"]