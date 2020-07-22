FROM alpine:edge AS builder
RUN apk update && apk upgrade
RUN apk add --update go=1.14.6-r0 gcc=9.3.0-r4 g++=9.3.0-r4
RUN apk add --no-cache git
ENV GOPATH /go
WORKDIR /go/src/github.com/mpinta/goshort/backend
COPY . .
RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-contrib/cors
RUN go get github.com/jinzhu/gorm
RUN go get github.com/mattn/go-sqlite3
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

FROM alpine:latest
ENV GOPATH /go
WORKDIR /go/src/github.com/mpinta/goshort/backend
COPY --from=builder /go/src/github.com/mpinta/goshort/backend/config.yml .
COPY --from=builder /go/src/github.com/mpinta/goshort/backend/main .
EXPOSE 8080
CMD ["./main"]
