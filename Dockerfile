FROM golang:alpine
WORKDIR /go/src/app
RUN apk add --no-cache git
COPY . .
RUN go get -d -v ./...
RUN go install -v ./cmd/rhc-manager

ENTRYPOINT ["rhc-manager"]