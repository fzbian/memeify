FROM golang:latest
WORKDIR /app
COPY meme-generator /app
RUN go get -d -v ./...
RUN go build -o main .
CMD ["./main"]
EXPOSE 8080
