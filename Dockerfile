FROM golang:1.20-alpine as build
WORKDIR /app
COPY . ./

RUN go mod init github.com/hosseinmirzapur/arbitrage && \
    GOPROXY=https://goproxy.io,direct go mod tidy && \
    go build -o main ./main.go && \
    chmod +x ./main

FROM debian:latest

RUN apt-get update && \ 
    apt-get -y install bash
    
WORKDIR /app

COPY --from=build /app/main .

CMD [ "./main" ]

EXPOSE 3000