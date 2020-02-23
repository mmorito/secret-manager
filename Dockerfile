FROM golang:1.13 as go
FROM gcr.io/distroless/base-debian10 as run

FROM go as build
WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o /go/bin/app

FROM run
COPY --from=build /go/bin/app /usr/local/bin/app
CMD ["app"]
