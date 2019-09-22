# Build the manager binary
FROM registry.cn-hangzhou.aliyuncs.com/knative-sample/golang:1.12.9 as builder

# Copy in the go src
WORKDIR /go/src/github.com/knative-sample/dingtalk-service
COPY pkg/    pkg/
COPY cmd/    cmd/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o dingtalk-service-receive github.com/knative-sample/dingtalk-service/cmd/receive

# Copy the dingtalk-service-receive into a thin image
FROM alpine:3.7
WORKDIR /
COPY --from=builder /go/src/github.com/knative-sample/dingtalk-service/dingtalk-service-receive app/
ENTRYPOINT ["/app/dingtalk-service-receive"]