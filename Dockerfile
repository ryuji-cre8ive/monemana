FROM golang:1.20.0 as builder

WORKDIR /go/src

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

# RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

COPY ./main.go  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/main \
    -ldflags '-s -w'

CMD ["air", "-c", ".air.toml"]


FROM scratch as runner

COPY --from=builder /go/bin/main /app/main

# COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=build /go/bin/app /app

ENTRYPOINT ["/app/main"]