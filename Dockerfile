FROM golang:1.20.1-bullseye AS builder

RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

COPY . /app

WORKDIR /app

RUN go build -o bin/publicip main.go

FROM debian:bullseye

COPY --from=builder /app/bin/publicip /usr/local/bin/publicip
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD [ "publicip" ]
