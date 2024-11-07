# * Stage for building
FROM golang:1.22-bullseye AS builder

# * Enable CGO
ENV CGO_ENABLED=1

# * Install necessary dependencies
RUN apt-get update && \
    apt-get install -y \
    gdal-data \
    libgdal-dev \
    gdal-bin \
    pkg-config \
    gcc

# * Set working directory
WORKDIR /app-build

# * Copy project files to the container
COPY . .

# * Update go module (may pull compatible version of godal)
RUN go mod tidy

# * Build application with CGO
RUN go build -o ./bins/service ./cmd/main.go

# * Stage for final image
FROM debian:bullseye-slim

# * Install minimal necessary dependencies
RUN apt-get update && \
    apt-get install -y \
    libgdal-dev \
    gdal-bin \
    gdal-data \
    libproj-dev \
    pkg-config

# * Set working directory
WORKDIR /app-prod


# * Copy binary and configuration
COPY --from=builder /app-build/bins/service /app-build/cmd/conf.yaml ./

# * Set command for running application
ENTRYPOINT ["./service", "./conf.yaml"]

# * Expose port
EXPOSE 8086

# !------------------------------!
# *------------NOTES-------------*
# * 1) Contain configs in /app-prod, pass configpath argument direct to ENV
# * 2) Any modifications via additions of the current image
#
#
#