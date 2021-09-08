# syntax=docker/dockerfile:1

# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
RUN apk add --no-cache gcc musl-dev

RUN mkdir -p /apps
WORKDIR /apps

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./app

EXPOSE 8181

ENTRYPOINT [ "./app" ]