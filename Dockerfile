FROM golang:1.22 AS builder

WORKDIR /app

# Only download Go modules (improves build caching)
COPY go.mod ./
RUN go mod download

# Copy our source code over
COPY . .

# Build the binary
RUN go mod vendor
RUN go build gotu-bookstore/cmd/gotu-bookstore

FROM debian:stable-slim as final
WORKDIR /app

# Copy over the binary artifact
COPY --from=builder /app/gotu-bookstore /app/gotu-bookstore

# Copy the resources directory
COPY --from=builder /app/cmd/gotu-bookstore/resources /app/resources

USER nobody

EXPOSE 4041/tcp

ENTRYPOINT ["/app/gotu-bookstore"]
