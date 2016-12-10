FROM golang:1.7.4
MAINTAINER kamoljan@gmail.com

ADD . /go/src/github.com/microservices-today/aws-sushi

COPY sushi.conf /etc/sushi/sushi.conf

RUN go get github.com/microservices-today/aws-sushi
RUN go install github.com/microservices-today/aws-sushi

ENTRYPOINT /go/bin/aws-sushi

EXPOSE 9090
