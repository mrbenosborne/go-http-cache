FROM golang:1.15.3 as builder
WORKDIR /app/
COPY . .
RUN CGO_ENABLED=0 go build -o go-http-cache .

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/go-http-cache .
RUN ls -la
CMD ./go-http-cache