# Etapa 1: Compilación
FROM golang:1.23-bullseye AS builder
WORKDIR /app

# Instalar dependencias del sistema
RUN apt-get update && apt-get install -y gcc && rm -rf /var/lib/apt/lists/*

# Copiar archivos de dependencias
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación con optimizaciones
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/app

# Etapa 2: Imagen final
FROM alpine:latest
WORKDIR /root/

# Copiar el binario compilado
COPY --from=builder /app/main .

# Exponer el puerto
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./main"]
