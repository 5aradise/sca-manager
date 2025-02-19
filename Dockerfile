FROM golang:1.23-alpine AS builder

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -C cmd/manager/ -ldflags="-w -s" -o /go/bin/

FROM alpine

COPY --from=builder /go/bin/manager /go/bin/app

CMD ["/go/bin/app"]