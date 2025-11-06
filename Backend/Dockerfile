#build stage
FROM golang:1.25.3-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go mod download
RUN go build -o dynamic-qr-code-backend .


#final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/dynamic-qr-code-backend .
EXPOSE 3080
CMD [ "./dynamic-qr-code-backend" ]
