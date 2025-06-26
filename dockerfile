FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-sample-microservice

# Run
FROM gcr.io/distroless/base-debian11 AS runner

WORKDIR /app

COPY --from=builder /go-sample-microservice ./go-sample-microservice

ENTRYPOINT ["./go-sample-microservice"]