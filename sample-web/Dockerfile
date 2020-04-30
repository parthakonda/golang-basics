FROM golang:1.12-alpine as builder
RUN apk --no-cache add ca-certificates git
WORKDIR /build/myapp

# Fetch dependencies
COPY go.mod ./
RUN go mod download

# Build
COPY . ./
RUN CGO_ENABLED=0 go build

# Create final image
FROM alpine
WORKDIR /root
COPY --from=builder /build/myapp/sample-web .
EXPOSE 8080
CMD ["./sample-web"]