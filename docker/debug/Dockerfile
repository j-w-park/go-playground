FROM golang:alpine
RUN go install github.com/go-delve/delve/cmd/dlv@latest
WORKDIR /app
EXPOSE 12345 8080
ENTRYPOINT [ "dlv" ]
