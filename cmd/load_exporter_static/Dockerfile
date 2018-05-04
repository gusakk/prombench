FROM golang:latest

ARG BINARY_PATH

WORKDIR /app
COPY ${BINARY_PATH} /app/

# Run static Prometheus exporter when the container starts.
ENTRYPOINT ["./load_exporter_static"]