# Build image.
FROM golang:latest AS builder
COPY . /go/src/stressbot
WORKDIR /go/src/stressbot
#RUN make deps
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/stressbot


# Final image.
FROM scratch as production
COPY --from=builder /go/src/stressbot/bin/stressbot .
EXPOSE 8081

CMD ["/stressbot"]
