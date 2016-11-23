FROM buildpack-deps:jessie-scm

# gcc for cgo
RUN apt-get update && apt-get install -y --no-install-recommends \
        git \
        wget \
        ca-certificates \
        g++ \
        gcc \
        libc6-dev \
        make \
        pkg-config \
    && rm -rf /var/lib/apt/lists/*


#install go 1.4 to compile lest go version
RUN wget --no-check-certificate https://storage.googleapis.com/golang/go1.4.1.linux-amd64.tar.gz
RUN tar -xvf go1.4.1.linux-amd64.tar.gz
RUN mv go /tmp/go

#NV PATH $PATH:/usr/local/go/bin

#UN chmod -R 777 /usr/local/go

#installing golang from git/master
WORKDIR /app

# copying go project with tests skipped
# tests are failing
COPY ./go /app/go

#RUN git clone https://go.googlesource.com/go 

WORKDIR /app/go/src

ENV GOROOT_BOOTSTRAP /tmp/go

RUN ./all.bash

RUN cp -rf /app/go /usr/local/go

ENV PATH $PATH:/usr/local/go/bin:/go/bin

RUN chmod -R 777 /usr/local/go

RUN mkdir /go

ENV GOPATH /go

WORKDIR /go
