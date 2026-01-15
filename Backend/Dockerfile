FROM golang:1.25-alpine

WORKDIR /app

# Install Air for hot reload
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# We do NOT run "go build" here. We let Air handle that.
EXPOSE 3080

CMD ["air"]