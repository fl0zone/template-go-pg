FROM golang:alpine AS development
ENV GO111MODULE=on \
    CGO_ENABLED=0
WORKDIR /app
COPY . .
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
CMD CompileDaemon -log-prefix=false -build="go build -o main" -command="./main"

FROM golang:alpine AS production
ENV GO111MODULE=on \
    CGO_ENABLED=0
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=production /app/main .
CMD ["./main"]
