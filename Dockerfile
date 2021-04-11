#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
#ENTRYPOINT ["tail", "-f", "/dev/null"]


#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/go-stock-portfolio /go-stock-portfolio
ENTRYPOINT ./go-stock-portfolio
LABEL Name=gostockportfolio Version=0.0.1