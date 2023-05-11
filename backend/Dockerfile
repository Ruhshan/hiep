FROM golang:1.19-alpine
WORKDIR "/app"

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-hiep-be

ENV APP_PROFILE=dev

EXPOSE 8989

CMD ["/docker-hiep-be"]

