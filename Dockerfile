FROM golang:1.15

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 

RUN go get github.com/libp2p/go-libp2p-core
RUN go build -o main .

ENV TYPE rsa
ENV BITSIZE 2048

CMD go run main.go -type=$TYPE -bitsize=$BITSIZE
