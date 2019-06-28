FROM golang:1.12

RUN apt-get update

RUN mkdir $GOPATH/src/states-webservice

ADD . $GOPATH/src/states-webservice

WORKDIR $GOPATH/src/states-webservice

RUN go get

RUN go build -o server .

EXPOSE 8001

CMD ["./server", ":8001"]