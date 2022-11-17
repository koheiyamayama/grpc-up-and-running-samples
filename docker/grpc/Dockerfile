FROM golang:1.19
RUN mkdir /app
WORKDIR /app
COPY ./ /app
RUN go build -v -o /usr/local/bin/grpc_up_and_running_samples cmd/server/main.go
CMD ["grpc_up_and_running_samples"]
