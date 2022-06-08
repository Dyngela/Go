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

ARG BUILD_ARGS
ENV BUILD_ARGS=${BUILD_ARGS}
ARG GO_MAIN
ENV GO_MAIN=${GO_MAIN}
RUN CGO_ENABLED=0 go build -o entry ${BUILD_ARGS} ${GO_MAIN}

CMD [ "/app/entry" ]


EXPOSE 8080
