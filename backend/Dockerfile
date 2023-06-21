FROM golang:1.19-alpine AS build-stage
WORKDIR "/app"

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-hiep-be

FROM alpine:3.14 AS build-release-stage

WORKDIR /

COPY --from=build-stage /docker-hiep-be /docker-hiep-be

RUN mkdir "/config"
COPY config/*.json /config

ENV APP_PROFILE=dev

EXPOSE 8989


CMD ["/docker-hiep-be"]

