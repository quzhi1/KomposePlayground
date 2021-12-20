FROM golang:latest

WORKDIR /opt/app

# ADD . /opt/app
ADD ./bin /opt/app/bin

EXPOSE 8090

# RUN go build -mod vendor -o bin/hello-go main.go

ENTRYPOINT bin/hello-go