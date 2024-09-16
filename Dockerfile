FROM golang:latest AS builder

RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY backend .

RUN go mod download

RUN ls

RUN swag init -g ./main.go

RUN CGO_ENABLED=0 GOOS=linux go build -o /build

FROM alpine:latest AS runner

#ENV SERVER_ADDRESS localhost:8080
    #POSTGRES_CONN="admin:12345678@localhost:5432/avito" \
ENV POSTGRES_JDBC_URL jdbc:postgresql://localhost:5432/avito
ENV POSTGRES_USERNAME "postgres"
ENV POSTGRES_PASSWORD "postgres"
ENV POSTGRES_HOST "localhost"
ENV POSTGRES_PORT 5432
ENV POSTGRES_DATABASE "avito"

COPY --from=builder build /bin/build

EXPOSE 8080

ENTRYPOINT [ "/bin/build" ]