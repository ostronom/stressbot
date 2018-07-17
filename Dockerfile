# Build image.
FROM golang:latest AS builder
COPY . /go/src/dialog-stress-bots
WORKDIR /go/src/dialog-stress-bots
#RUN make deps
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/dialog-stress-bots



# Final image.
FROM scratch as production
COPY --from=builder /go/src/dialog-stress-bots/bin/dialog-stress-bots .
EXPOSE 8081

CMD ["/dialog-stress-bots"]


