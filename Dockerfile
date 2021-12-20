FROM golang:latest

WORKDIR /opt/app

ADD . /opt/app

EXPOSE 8090

RUN go build -mod vendor -o bin/hello-go main.go

CMD ["bin/hello-go"]