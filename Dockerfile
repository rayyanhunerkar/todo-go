FROM golang:1.21.0-alpine3.18
WORKDIR /app
COPY . /app
RUN go install
CMD [ "go", "run", "main.go" ]
