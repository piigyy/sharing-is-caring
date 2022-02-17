FROM golang:1.17-alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
RUN mkdir /app
COPY . /app
WORKDIR /app/cmd/auth/http

# Fetch dependencies.
# Using go get.
RUN go get -d -v

# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/authservice

FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/authservice /go/bin/authservice
COPY --from=builder /app/config.local.yaml /go/bin/
# Run the hello binary.
EXPOSE 8088
ENTRYPOINT ["/go/bin/authservice"]
