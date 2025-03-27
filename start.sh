#!/bin/bash

# Verificar si Colima está en ejecución
if ! colima status > /dev/null 2>&1; then
    echo "Iniciando Colima..."
    colima start --cpu 2 --memory 4 --disk 100
fi

# Detener y eliminar contenedores existentes
echo "Deteniendo contenedores existentes..."
docker-compose down -v

# Construir la imagen Docker
echo "Construyendo la imagen Docker..."
docker build -t financial-resume-engine .

# Iniciar los servicios
echo "Iniciando los servicios..."
docker-compose up -d

echo "La aplicación está corriendo en http://localhost:8080" 