FROM golang:alpine as builder
ENV GIN_MODE=release
WORKDIR /app
COPY . /app/
RUN go build -o /app/build/todo-go .


FROM golang:alpine as app
WORKDIR /app
COPY --from=builder /app/build/todo-go .
EXPOSE 8080
CMD [ "todo-go" ]
