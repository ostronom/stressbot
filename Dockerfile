# Build image.
FROM golang:latest AS builder
RUN apt-get update && apt-get install -y curl
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY . /go/src/github.com/dialogs/stressbot
WORKDIR /go/src/github.com/dialogs/stressbot
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/dialog-stress-bots

CMD ["/dialog-stress-bots"]


