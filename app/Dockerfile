FROM golang:1.14-alpine as build

RUN apk update && apk add git

RUN mkdir -p /go/src/github.com/blacknikka/timestream-golang
WORKDIR /go/src/github.com/blacknikka/timestream-golang

# COPY ./src /go/src/github.com/blacknikka/timestream-golang
# RUN go get
# RUN go build -o main

# RUN chmod +x /go/src/github.com/blacknikka/timestream-golang/main

# EXPOSE 8080

# CMD [ "/go/src/github.com/blacknikka/timestream-golang/main" ]
