FROM golang:1.19-alpine AS build

# Move to working directory (/build).
WORKDIR /app

# Copy and download dependency using go mod.
COPY . .

# RUN go mod tidy
RUN go mod tidy

# Generate swagger documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init  --output swagger/

# Build
RUN go build -o server

FROM alpine:latest

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=build /app/server /server

EXPOSE 5000

# Command to run when starting the container.
ENTRYPOINT [ "/server" ]
