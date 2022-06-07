##
## Build
##
FROM golang:1.18-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/api/*.go cmd/api/

RUN set GOOS=linux && set GOARCH=amd64 && set CGO_ENABLED=0 && go build -o /app/brokerApp ./cmd/api

##
## Deploy
##
FROM alpine:latest

COPY --from=build /app/brokerApp /

USER 1337
EXPOSE 80

ENTRYPOINT ["/brokerApp"]
