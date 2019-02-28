# Start from golang v1.11 base image
FROM golang:1.11

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/beaverGO

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Run the executable
CMD ["beaverGO/cmd/getData/main.go"]