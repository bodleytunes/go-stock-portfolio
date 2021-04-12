#build stage
FROM golang:alpine 
RUN apk add --no-cache git && apk add build-base
# important: use delv for debugging
RUN go get github.com/go-delve/delve/cmd/dlv 
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
#ENTRYPOINT ./dlv
#ENTRYPOINT ./go-stock-portfolio
LABEL Name=gostockportfolio Version=0.0.1
#ENTRYPOINT ["tail", "-f", "/dev/null"]


#final stage
#FROM golang:latest
##RUN apk --no-cache add ca-certificates
#COPY --from=builder /go/bin/go-stock-portfolio /go-stock-portfolio
## important: use delv for debugging
#COPY --from=builder /go/bin/dlv /dlv
