FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build -o bcwserver
USER default
CMD ["./bcwserver"]