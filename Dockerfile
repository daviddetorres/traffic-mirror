FROM golang:1.14 as builder
WORKDIR /traffic-mirror
COPY go.mod go.sum ./
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s" -o traffic-mirror main.go
RUN strip traffic-mirror

FROM alpine:3.10
COPY --from=builder /traffic-mirror /bin
EXPOSE 8080
CMD ["/bin/traffic-mirror"]
