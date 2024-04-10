FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build
USER default
CMD ["./bcwserver"]