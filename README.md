# poc-proxy-https

POC to test https proxying using golang

Cool connect diagram:

https://blog.thousandeyes.com/wp-content/uploads/2013/09/image002.png


## build

    go build main.go

## run

    ./main --proxy IP:PORT --user USER --password PASSWORD --dest https://www.google.com.br
