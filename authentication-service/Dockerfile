FROM golang:1.18.3-alpine3.16
RUN apk update && \
    apk add --no-cache \
    netcat-openbsd \
    iproute2 \
    nmap \
    curl

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY src/ ./
RUN CGO_ENABLED=0 go build -o auth ./cmd/api

CMD [ "/app/auth" ]


EXPOSE 8080
