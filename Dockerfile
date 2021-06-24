FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

EXPOSE 7000
ADD simple-go-app /
CMD ["/simple-go-app"]