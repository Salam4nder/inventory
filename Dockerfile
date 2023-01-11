FROM golang:1.19.0-alpine3.13 AS modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM golang:1.19.0-alpine3.13 AS builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/app ./cmd/app

from scratch
COPY --from=builder /bin/app /app
COPY --from=builder /app/config /config
COPY --from=builder /app/migration /migration
CMD ["/app"]
