FROM golang:1.19-bullseye

# Set working directory for source
WORKDIR /usr/src/app

# Pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Compile the app
COPY . .
RUN go build -v -o /usr/local/bin ./...

# Switch to a low-privileged user
RUN adduser --no-create-home go --shell /bin/sh && chown -R go:go /usr/local/bin
USER go

# Expose the port
ENV CG_PORT="8080"
EXPOSE 8080

# Run the app
CMD [ "tic-tac-toe-simple" ]
