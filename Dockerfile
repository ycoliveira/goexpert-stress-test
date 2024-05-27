# Use a imagem base do Go para a construção do binário
FROM golang:1.18 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o stresstest

# Use a imagem base do Alpine para a imagem final
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/stresstest /stresstest
ENTRYPOINT ["/stresstest"]
