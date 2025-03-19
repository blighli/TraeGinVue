FROM golang:1.21-alpine

# 设置 Go 模块代理
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

# 只复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 分层复制其他文件
COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]