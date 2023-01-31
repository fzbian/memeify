FROM golang:latest
WORKDIR /app
COPY . /app
RUN go get -d -v ./...
RUN go build -o main .
CMD ["./main"]
EXPOSE 8080
