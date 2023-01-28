FROM golang:1.19.5-alpine AS build-env

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /tmp ./...

# Runtime image

FROM golang:1.19.5-alpine

WORKDIR /app

COPY --from=build-env /tmp/sqzsvc /app
COPY --from=build-env /usr/src/app/.env  /app


ENV GIN_MODE=release

CMD ["/app/sqzsvc"]
