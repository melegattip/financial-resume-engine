# Etapa 1: Compilación
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Instalar dependencias del sistema
RUN apk add --no-cache gcc musl-dev

# Copiar archivos de dependencias
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# Etapa 2: Imagen final
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copiar el binario compilado
COPY --from=builder /app/main .

# Exponer el puerto
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./main"]
