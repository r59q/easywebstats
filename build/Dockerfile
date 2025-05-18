FROM golang:1.24 AS builder

WORKDIR /app


COPY ../go.sum .
COPY ../go.mod .

RUN go mod download

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -o /app/bin/app

FROM scratch

WORKDIR /
RUN mkdir "data"
COPY --from=builder /app/bin/app .
ENTRYPOINT ["./app"]