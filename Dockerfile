FROM ubuntu:latest

RUN apt-get update && apt-get install -y git

WORKDIR /app

RUN git clone https://go.googlesource.com/go

WORKDIR /app/go/src

RUN ./all.bash

COPY ./main.go /app/main.go

WORKDIR /app

RUN go run main.go --proxy 173.44.253.143:60099 --user neoway --password AxafhhhCEvvU5xg --dest http://www.google.com.br
