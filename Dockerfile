# Etapa 1: Compilación
FROM --platform=linux/arm64 golang:1.23-rc-alpine
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
RUN go build -o main ./cmd/app

# Exponer el puerto
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./main"]
