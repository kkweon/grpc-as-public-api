FROM golang:1.14 as builder

RUN GRPC_HEALTH_PROBE_VERSION=v0.3.2 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

WORKDIR /app

COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .


#########################

FROM alpine:latest
# Copy our static executable.
COPY --from=builder /app/server /app/server
COPY --from=builder /bin/grpc_health_probe /bin/grpc_health_probe

# Run the hello binary.
ENTRYPOINT ["/app/server"]
