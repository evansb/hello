FROM golang:1.10.0-stretch

RUN go get github.com/evansb/hello
RUN cd /go/src/github.com/evansb/hello && go build -o hello

WORKDIR /go/src/github.com/evansb/hello

ENTRYPOINT [ "./hello" ]

EXPOSE 8080
