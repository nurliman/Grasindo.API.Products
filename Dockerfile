
#build stage
FROM golang:1.14.6 AS builder
WORKDIR /go/src/app
COPY . .
#RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM debian:10.4-slim
#RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin /app
ENTRYPOINT ./app
LABEL Name=grasindo.api.products Version=0.0.1
EXPOSE 8080
RUN /app/Grasindo.API.Products
