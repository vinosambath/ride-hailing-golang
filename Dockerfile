FROM golang:1.23

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o /ride-hailing-golang
EXPOSE 8080
CMD ["/ride-hailing-golang"]