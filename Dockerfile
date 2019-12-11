FROM instrumentisto/dep:0.5-alpine as builder

LABEL stage=builder

WORKDIR /go/src/github.com/proskehy/flux-webhook-receiver

COPY Gopkg.lock ./
COPY Gopkg.toml ./
RUN dep ensure --vendor-only

COPY cmd ./cmd
COPY pkg ./pkg
RUN cd cmd && go build

FROM alpine:latest
COPY --from=builder /go/src/github.com/proskehy/flux-webhook-receiver/cmd/cmd /bin/flux-webhook-receiver
EXPOSE 3033
ENTRYPOINT ["flux-webhook-receiver"]

