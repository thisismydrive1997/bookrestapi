FROM golang:1.18-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o /docker-book-rest-api

CMD ["/docker-book-rest-api"]

#Buid: sudo docker build --tag docker-book-rest-api .