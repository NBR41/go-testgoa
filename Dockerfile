FROM golang:1.11-alpine3.7
RUN apk add --no-cache git
COPY entrypoint.sh /bin/
WORKDIR /go/src/go-testgoa
COPY . .
RUN cd cmd && CGO_ENABLED=0 GO111MODULE=on go build -o /go/bin/myinventory && rm -rf /go/src/go-testgoa
EXPOSE 8089
ENTRYPOINT ["/bin/entrypoint.sh"]
