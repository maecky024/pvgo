# Start from the official Go image
FROM golang:bullseye
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy the source from the current directory to the Working Directory inside the container
COPY ./ .
# Build the Go app
RUN wget -O /app/liquibase/liquibase.tar.gz https://github.com/liquibase/liquibase/releases/download/v4.30.0/liquibase-4.30.0.tar.gz
RUN mkdir /app/java/
RUN wget -O /app/java/java.tar.gz https://download.oracle.com/java/17/archive/jdk-17.0.12_linux-aarch64_bin.tar.gz
RUN tar -xzvf /app/java/java.tar.gz -C /app/java/
ENV JAVA_HOME=/app/java/jdk-17.0.12
RUN tar -xzvf /app/liquibase/liquibase.tar.gz -C /app/liquibase/
ENV LIQUIBASE_HOME=/app/liquibase
RUN rm -rf /app/liquibase/examples
RUN rm -rf /app/liquibase/liquibase.tar.gz
RUN rm -rf /app/java/java.tar.gz
RUN go build -o /app/cmd/pvgo/main /app/cmd/pvgo/main.go
# Expose port 8080 to the outside world
EXPOSE 8000
# Set environment variable for Gin mode
ENV GIN_MODE=release
# Run the executable
CMD ["/app/cmd/pvgo/main"]