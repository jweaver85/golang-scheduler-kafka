FROM golang:1.21-alpine

WORKDIR /app/

# Turning off Cgo is required to install Delve on Alpine.
#ENV CGO_ENABLED=0
# Download and install the Delve debugger.
#RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/cosmtrek/air@latest

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go mod tidy

# compile for Devlve
#RUN go build -gcflags "all=-N -l" -o ./scheduler scheduler.go
#CMD dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./scheduler
RUN go mod vendor
RUN go build -o bin/scheduler src/scheduler/main.go
RUN go build -o bin/producer src/producer/main.go
RUN go build -o bin/consumer src/consumer/main.go
CMD ./scheduler
