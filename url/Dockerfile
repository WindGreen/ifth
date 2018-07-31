FROM golang:1.10.3-alpine as builder
WORKDIR /go/src/ifth
COPY . .
WORKDIR /go/src/ifth/url
#RUN go get -d
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/ifth/www/app ./
EXPOSE 80
ENTRYPOINT ["./app"]
