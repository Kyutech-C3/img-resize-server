FROM golang:1.20.5-alpine as base

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

RUN apk upgrade --update && apk --no-cache add git
RUN apk add --no-cache gcc g++ libgcc make

EXPOSE 8000

FROM base as dev
RUN go get -u github.com/cosmtrek/air@v1.44.0
RUN go build -o /go/bin/air github.com/cosmtrek/air
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12

ARG CGO_ENABLED=0
ARG GOOS=linux
# アーキテクチャを指定する
# ARG GOARCH=amd64
ARG GOARCH=arm64


CMD [ "air", "-c", ".air.toml" ]

FROM base as prod

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

CMD ["GIN_MODE=release", "go", "run", "main.go"]