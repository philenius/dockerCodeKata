FROM golang:stretch

WORKDIR /go/src/hello-world

COPY main.go /go/src/hello-world/

RUN go get
RUN go build

ENV NAME=opitz
EXPOSE 8080

ENTRYPOINT ["hello-world"]