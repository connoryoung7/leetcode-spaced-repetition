FROM golang:1.25.2-alpine3.22
RUN go install github.com/air-verse/air@latest
WORKDIR /app
COPY . .
ENTRYPOINT ["air"]
