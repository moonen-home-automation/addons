# Build stage
FROM golang:1.22.4-bullseye as BuildStage

WORKDIR /app

COPY . .

EXPOSE 9091

RUN go build ./cmd/grpc_bridge/grpc_bridge.go

#Deploy stage
FROM ghcr.io/hassio-addons/ubuntu-base:10.0.1

WORKDIR /app

COPY --from=BuildStage /app/grpc_bridge /grpc_bridge

EXPOSE 9091

ENTRYPOINT [ "/grpc_bridge" ]