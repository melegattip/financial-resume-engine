#!/bin/bash

# Verificar si Colima está en ejecución
if ! colima status > /dev/null 2>&1; then
    echo "Iniciando Colima..."
    colima start --cpu 2 --memory 4 --disk 100
fi

# Construir la imagen Docker
echo "Construyendo la imagen Docker..."
docker build -t financial-resume-engine .

# Detener y eliminar el contenedor si existe
docker stop financial-resume-engine 2>/dev/null || true
docker rm financial-resume-engine 2>/dev/null || true

# Iniciar el contenedor
echo "Iniciando la aplicación..."
docker run -d \
    --name financial-resume-engine \
    -p 8080:8080 \
    --env-file .env \
    financial-resume-engine

echo "La aplicación está corriendo en http://localhost:8080" 