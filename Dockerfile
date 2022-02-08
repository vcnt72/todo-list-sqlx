FROM golang as builder

ENV GO111MODULE=on

WORKDIR /src

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/app/main.go

FROM debian:buster-slim

RUN apt-get update

COPY --from=builder /src/main /app/main
# COPY --from=builder /src/config/config.yaml .
COPY --from=builder /src/config/config.yaml config/
# COPY --from=builder /src/config/config.yaml /app/config

ENTRYPOINT [ "/app/main" ]