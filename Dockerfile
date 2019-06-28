FROM golang:1.12

RUN apt-get update

RUN mkdir $GOPATH/src/states-webservice

ADD . $GOPATH/src/states-webservice

WORKDIR $GOPATH/src/states-webservice

RUN go get

EXPOSE 8001