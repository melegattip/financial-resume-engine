# Etapa 1: Compilación
FROM golang:1.21-bullseye
WORKDIR /app

# Instalar dependencias del sistema
RUN apt-get update && apt-get install -y gcc && rm -rf /var/lib/apt/lists/*

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
