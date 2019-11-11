FROM golang:1.12-alpine as builder
WORKDIR /go/src/github.com/proskehy/flux-webhook-receiver
COPY cmd cmd/
COPY pkg pkg/
RUN cd cmd && go build

FROM alpine:latest
COPY --from=builder /go/src/github.com/proskehy/flux-webhook-receiver/cmd/cmd /bin/flux-webhook-receiver
EXPOSE 3031
ENTRYPOINT ["flux-webhook-receiver"]

