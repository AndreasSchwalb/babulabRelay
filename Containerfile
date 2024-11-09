FROM golang:1.23 AS build-stage

WORKDIR /app

COPY go.mod main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /bambu-relay

FROM scratch AS build-release-stage

WORKDIR /

COPY --from=build-stage /bambu-relay /bambu-relay

EXPOSE 2021

ENV SERVER_ADDRESS=:2021

ENTRYPOINT ["/bambu-relay"]