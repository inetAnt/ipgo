# build executable binary
FROM golang:alpine AS builder
# install git
# RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/inetAnt/ipgo/
COPY . .
# build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/ipgo

# build a small image
FROM scratch
# copy our static executable
COPY --from=builder /go/bin/ipgo /go/bin/ipgo
# run the ipgo binary on port 8080
ENTRYPOINT ["/go/bin/ipgo", "-l", ":8080"]
