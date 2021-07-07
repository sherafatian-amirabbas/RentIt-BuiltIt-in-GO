FROM golang:alpine as api
WORKDIR /src
COPY . .
RUN go mod download
EXPOSE 8081
ENTRYPOINT ["go", "run", "main.go"]

FROM golang:alpine as test
WORKDIR /src
COPY . .
RUN chmod +x wait-for
RUN apk update && apk add build-base
RUN go mod download
ENTRYPOINT ["go", "test", "http_test.go"]