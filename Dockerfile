FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Copy the .env file
COPY .env ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Run
FROM scratch

COPY --from=builder /app/main /main
COPY --from=builder /app/.env ./

EXPOSE 8000

CMD ["/main"]