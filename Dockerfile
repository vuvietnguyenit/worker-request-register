FROM golang:1.21-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
# Copy data file
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /worker-request-register
# Run
CMD ["/worker-request-login"]