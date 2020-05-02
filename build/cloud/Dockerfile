FROM golang AS builder

WORKDIR /go/src/github.com/davidchristie/cloud

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN mkdir /cloud

RUN CGO_ENABLED=0 go build -o /cloud ./...

FROM alpine

COPY --from=builder /cloud /cloud