FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build
RUN apk add --no-cache git mercurial \
    && go get github.com/gorilla/mux \
    && go get github.com/docker/docker/pkg/namesgenerator \
    && apk del git mercurial
RUN go build -o main ./...
FROM alpine
EXPOSE 80
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]
