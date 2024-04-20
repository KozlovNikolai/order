FROM golang:1.21.7-alpine AS builder

COPY . /github.com/KozlovNikolai/order/source/
WORKDIR /github.com/KozlovNikolai/order/source/

RUN go mod download
RUN go build -o ./bin/order cmd/order/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/KozlovNikolai/order/source/bin/order .

CMD ["./order"]