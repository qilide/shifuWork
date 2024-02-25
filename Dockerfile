FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 运行应用
CMD ["./main"]