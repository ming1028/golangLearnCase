#FROM alpine:latest

#WORKDIR /root/

#COPY http-server .

#CMD ["./http-server"]

#方式2
FROM golang:1.18

WORKDIR /go/src/http-server/

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o http-server .

FROM alpine:latest

WORKDIR /go/src

## --from 从第一阶段拷贝编译后的文件
COPY --from=0 /go/src/http-server/http-server .

CMD ["./http-server"]

## docker build -t http-server:latest .