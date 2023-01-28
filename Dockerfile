FROM golang:1.19.5

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify


COPY . .
RUN go build -v -o /usr/local/bin ./...

ENV GIN_MODE=release
# ENV DB_CONNECTION=host=localhost user=postgres password=secret dbname=sqz-data port=5432 sslmode=disable
# ENV TOKEN_SECRET=superDuperSecret
# ENV TOKEN_HOUR_LIFESPAN=1


EXPOSE 8010
CMD ["sqzsvc"]
