FROM golang:alpine AS dev-env

WORKDIR /usr/local/go/src/github.com/yuleihua/trade
COPY . /usr/local/go/src/github.com/yuleihua/trade

RUN apk update && apk upgrade && \
    apk add --no-cache bash git

RUN go get ./...

RUN go build -o dist/trade &&\
    cp -f dist/trade /usr/local/bin/ &&\
    cp -f dist/trade.yaml /usr/local/etc/ &&\

RUN ls -l && ls -l dist

CMD ["/usr/local/bin/trade", "-c", "/usr/local/etc/trade.yaml" ]