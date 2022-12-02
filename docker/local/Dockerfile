FROM golang:1.18-alpine AS builder

WORKDIR /dao-chain

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o dao-chain main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates jq

WORKDIR /dao-chain

COPY --from=builder /dao-chain/dao-chain ./
COPY ./docker/local/dao-chain.sh ./

# Expose json-rpc, libp2p and grpc ports
EXPOSE 8545 9632 1478

ENTRYPOINT ["./dao-chain.sh"]
