# first stage
FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o interviewServer main.go


# second stage
FROM ubuntu:latest

COPY --from=build /app/interviewServer .

EXPOSE 8080

CMD ["./interviewServer"]
