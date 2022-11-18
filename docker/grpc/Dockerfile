FROM golang:1.19
RUN mkdir /app
WORKDIR /app
COPY ./ /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]
