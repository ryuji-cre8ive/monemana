FROM golang:1.20.0 as builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/main \
    -ldflags '-s -w'

FROM scratch as runner

COPY --from=builder /go/bin/main /app/main

# Set the PORT environment variable to 8080
ENV PORT=8080

# Make sure the server listens on the correct port
CMD ["/app/main"]

# Make sure to expose the port
EXPOSE 8080
