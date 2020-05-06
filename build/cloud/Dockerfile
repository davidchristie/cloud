FROM golang AS builder

WORKDIR /go/src/github.com/davidchristie/cloud

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg

RUN mkdir /cloud

RUN CGO_ENABLED=0 go build -o /cloud ./...
