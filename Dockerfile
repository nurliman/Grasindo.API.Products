
#build stage
FROM golang:1.14.6 AS builder
ENV GOPROXY=https://gocenter.io,https://goproxy.io,direct
ENV GO111MODULE=on
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -a -installsuffix cgo -v ./...

#final stage
FROM debian:10.4-slim

COPY --from=builder /go/bin /app
COPY ./wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh
ENTRYPOINT ["/bin/sh","-c","/app/wait-for-it.sh db:5432 -t 30 -- /app/Grasindo.API.Products"]
LABEL Name=grasindo.api.products Version=0.0.1
EXPOSE 1337
