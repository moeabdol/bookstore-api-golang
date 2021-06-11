# Build stage
FROM golang:1.16-alpine3.13 AS build
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=build /app/.env .
COPY --from=build /app/main .

EXPOSE 3000
CMD ["/app/main"]
