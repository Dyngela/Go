##
## Build
##
FROM golang:1.18-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/api/*.go cmd/api/
COPY data/*.go data/

RUN set GOOS=linux && set GOARCH=amd64 && set CGO_ENABLED=0 && go build -o /app/authApp ./cmd/api

##
## Deploy
##
FROM alpine:latest

COPY --from=build /app/authApp /authApp

EXPOSE 80
USER 1337

ENTRYPOINT ["/authApp"]
