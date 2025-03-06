FROM golang:1.24.0
WORKDIR /site-monitoring
COPY . .
RUN go build -o main main.go
ENTRYPOINT ./main