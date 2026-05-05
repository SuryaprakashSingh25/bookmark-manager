FROM golang:1.26-alpine

WORKDIR /app

COPY . .

# Build API
WORKDIR /app/api
RUN go mod download
RUN go build -o /app/api-main .

# Build gRPC service
WORKDIR /app/preview-service
RUN go mod download
RUN go build -o /app/preview-main .

WORKDIR /app

EXPOSE 8080
EXPOSE 50051

CMD sh -c "./preview-main & ./api-main"