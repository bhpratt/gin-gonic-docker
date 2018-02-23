FROM golang:latest
RUN go get github.com/gin-gonic/gin
RUN mkdir /app
COPY main.go /app
WORKDIR /app
RUN go build -o main .
EXPOSE 8080
CMD ["/app/main"]
