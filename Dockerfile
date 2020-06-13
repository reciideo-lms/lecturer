FROM golang AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o lecturer httpd/main.go

FROM scratch
COPY --from=builder /app/lecturer /app/
ENV GIN_MODE=release
EXPOSE 8080
ENTRYPOINT ["/app/lecturer"]
