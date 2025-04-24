FROM golang:1.24.0-alpine as base
WORKDIR /work
RUN apk update && apk add --no-cache git
COPY services/auth/go.mod services/auth/go.sum ./
RUN go mod download

FROM base as dev
CMD ["go", "run", "./cmd/server/main.go"]

FROM base as build
COPY cmd/ cmd/
COPY internal/ internal/
RUN apk add --no-cache swag
RUN swag init -g cmd/server/main.go -o cmd/server/docs
RUN mkdir -p /app
RUN go build -o /app/server ./cmd/server/main.go
CMD ["/app/server"]
