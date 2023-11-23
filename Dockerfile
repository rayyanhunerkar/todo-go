FROM golang:1.20-alpine as build-binary
WORKDIR /go/src/code
COPY . .
RUN go mod vendor
RUN go build -v -o todo-go

FROM alpine:3.12
WORKDIR /usr/bin
RUN apk add --no-cache tzdata
COPY --from=build-binary /go/src/code/todo-go .
COPY --from=build-binary /go/src/code/src/config /usr/bin/src/config
EXPOSE 8000
ENTRYPOINT ./todo-go