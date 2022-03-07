FROM docker.io/library/golang:1.16 as builder
WORKDIR /tmd
COPY . .
RUN go mod tidy; \
    go build
FROM docker.io/library/golang:1.16
EXPOSE 14000
WORKDIR /tmd
COPY --from=builder /tmd/local-api-encrypted-variables ./local-api-encrypted-variables
RUN chmod +x local-api-encrypted-variables
ENTRYPOINT ["./local-api-encrypted-variables"]
