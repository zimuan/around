FROM golang:1.18
MAINTAINER cuzimuan@gmail.com


WORKDIR /go/src/around
ADD . / /go/src/around/


RUN go get cloud.google.com/go/storage
RUN go get github.com/auth0/go-jwt-middleware
RUN go get github.com/form3tech-oss/jwt-go
RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux
RUN go get github.com/olivere/elastic/v7
RUN go get github.com/pborman/uuid


EXPOSE 8080
CMD ["/usr/local/go/bin/go", "run", "main.go"]
