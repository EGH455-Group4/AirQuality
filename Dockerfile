FROM golang:1.15.6-alpine3.12

WORKDIR /app

ADD go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/api .

EXPOSE 8050

ENTRYPOINT [ "./bin/api" ]
