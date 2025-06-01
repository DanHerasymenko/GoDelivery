FROM golang:1.24.0-alpine as base
WORKDIR /service
RUN apk update && apk add --no-cache git
COPY services/auth/go.mod services/auth/go.sum ./
RUN go mod download

FROM base as dev
CMD ["go", "run", "./cmd/server/main.go"]

FROM base as build
COPY services/auth/cmd ./cmd
COPY services/auth/internal ./internal
COPY services/auth/migrations ./migrations
RUN apk add --no-cache swag
RUN swag init -g cmd/server/main.go -o cmd/server/docs
RUN mkdir -p /build
RUN go build -o /build/auth_service_server ./cmd/server/main.go
CMD ["/build/auth_service_server"]
