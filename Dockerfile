FROM golang:1.9

# ENV HTTP_PROXY http://proxy.corproot.net:8079
# ENV HTTPS_PROXY http://proxy.corproot.net:8079
ENV GOPATH /go
ENV PATH "/go/bin:${PATH}"
ENV PORT 8080

RUN mkdir /go/src/github.com/gobeli/go-test

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/tockins/realize

WORKDIR /go/src/github.com/gobeli/go-test
COPY . .