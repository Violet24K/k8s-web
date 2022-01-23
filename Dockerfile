FROM golang:alpine
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]