FROM golang:alpine as builder

ENV GO111MODULE=on

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Add Maintainer Info
LABEL maintainer="Dao Mai <dao.mai@s3corp.com.vn>"

WORKDIR /app
#COPY env/env.env ./env/env.env
#COPY docs ./docs/
COPY go.mod .
#COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
#COPY --from=builder /app/env/env.env ./env/env.env
#COPY --from=builder /app/docs ./docs/

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]
