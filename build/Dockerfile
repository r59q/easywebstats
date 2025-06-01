FROM golang:1.24.3 AS builder

WORKDIR /app


COPY ../go.sum .
COPY ../go.mod .

RUN go mod download

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -o /app/bin/app

FROM scratch

WORKDIR /
ENV GIN_MODE=release
COPY --from=builder /app/bin/app .
ENTRYPOINT ["./app"]