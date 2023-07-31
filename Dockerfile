FROM golang as builder

ENV GO111MODULE=on

WORKDIR /go-modules-docker

COPY . .

WORKDIR /go-modules-docker/go-modules

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-modules-docker/app

FROM alpine:3.8

WORKDIR /root/

COPY --from=builder /go-modules-docker/app .

CMD ["./app"]