FROM golang AS builder

ENV GO111MODULE=on

RUN useradd -u 10001 lecturer

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o lecturer httpd/main.go

FROM scratch
COPY --from=builder /app/lecturer /app/
COPY --from=builder /etc/passwd /etc/passwd
USER lecturer
ENV GIN_MODE=release
EXPOSE 8080
ENTRYPOINT ["/app/lecturer"]
