# Build stage
FROM golang:1.22.4-bullseye as BuildStage

WORKDIR /app

COPY . .

ARG SSH_PRIVATE_KEY
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

RUN mkdir -p /root/.ssh && \
    chmod 0700 /root/.ssh && \
    ssh-keyscan github.com > /root/.ssh/known_hosts

RUN echo "$SSH_PRIVATE_KEY" > /root/.ssh/id_rsa && \
    chmod 600 /root/.ssh/id_rsa && \
    go mod download && \
    rm -rf /root/.ssh/

EXPOSE 9091

RUN go build ./cmd/grpc_bridge/grpc_bridge.go

#Deploy stage
FROM ghcr.io/hassio-addons/ubuntu-base:10.0.1

WORKDIR /app

COPY --from=BuildStage /app/grpc_bridge /grpc_bridge

EXPOSE 9091

ENTRYPOINT [ "/grpc_bridge" ]