FROM golang:1.16.5-alpine3.13

COPY ../app/src /go/src/api/

WORKDIR /go/src/api/

COPY ../go.mod /go/src/

COPY ../go.sum /go/src/

RUN go mod tidy
