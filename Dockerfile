
#build stage
FROM golang:1.14.6 AS builder
ENV GOPROXY=https://goproxy.io,https://proxy.golang.org,direct
ENV GO111MODULE=on
WORKDIR /go/src/app
COPY go.mod . 
COPY go.sum .
VOLUME [ "go-modules:/go/pkg/mod" ]
RUN go mod download -x
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM debian:10.4-slim
#RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin /app
COPY ./wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh
ENTRYPOINT ["/bin/sh","-c","/app/wait-for-it.sh db:5432 -t 30 -- /app/Grasindo.API.Products"]
LABEL Name=grasindo.api.products Version=0.0.1
EXPOSE 1337
