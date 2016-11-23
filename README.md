# poc-proxy-https
POC to test https proxying with authentication using golang

The last golang version (1.7.3) don't pass header parameters to method CONNECT, then parameter Proxy-Authorization don't sent to proxy server.

To resolve this, field ProxyConnectHeader are added to struct Transport.

[link](https://github.com/golang/go/commit/b06c93e45b7b03a5d670250ff35e42d62aface82) to path with this change

For turn proccess more easy, golang from master with last change are installed in one docker container.

See Dockerfile for more information.

## build

    git clone https://go.googlesource.com/go
    docker build -t leocbs/golang-devel .

## run
    
    docker run --rm leocbs/golang-devel go run main.go --proxy IP:PORT --user USER --password PASSWORD --dest https://www.google.com.br
