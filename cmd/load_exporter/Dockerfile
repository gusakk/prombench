FROM golang:latest

ARG BINARY_PATH

WORKDIR /app
COPY ${BINARY_PATH} /app/

# Run incremental Prometheus exporter when the container starts.
ENTRYPOINT ["./load_exporter"]