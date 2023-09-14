FROM golang:1.21.0 as builder

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN GOPROXY=https://goproxy.cn,direct CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# RUN GO111MODULE=on GOPROXY=https://goproxy.cn,direct CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/etc  etc

EXPOSE 80

CMD ["/app/main"]