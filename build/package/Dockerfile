FROM golang:alpine3.16

COPY web/build /app/
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

WORKDIR /app

RUN apk add --no-cache git
RUN go install github.com/shurcooL/goexec@latest
RUN go get github.com/shurcooL/go-goon

EXPOSE 80

CMD ["goexec", "http.ListenAndServe(`:80`, http.FileServer(http.Dir(`.`)))"]