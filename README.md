# Financial Resume Engine

Motor de resumen financiero que permite gestionar transacciones, categorías y generar reportes.

## Requisitos Previos

- Go 1.23 o superior
- Docker y Docker Compose
- Sistema operativo compatible:
  - Windows 10/11 Pro, Enterprise o Education (64-bit)
  - Linux (Ubuntu 20.04 LTS o superior)
  - macOS 10.15 o superior

## Instalación

1. Clonar el repositorio:
```bash
git clone https://github.com/melegattip/financial-resume-engine.git
cd financial-resume-engine
```

2. Configurar variables de entorno:
```bash
# Windows (PowerShell)
Copy-Item .env.example .env

# Linux/macOS
cp .env.example .env

# Editar .env con tus credenciales seguras
```

3. Configurar Docker según tu sistema operativo:

### Windows
1. Instalar Docker Desktop desde [https://www.docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop)
2. Iniciar Docker Desktop
3. Esperar a que el ícono de Docker en la bandeja del sistema indique que está listo

### Linux (Ubuntu)
```bash
# Actualizar repositorios
sudo apt-get update

# Instalar dependencias
sudo apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

# Agregar la clave GPG oficial de Docker
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Configurar el repositorio estable
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Instalar Docker Engine
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io

# Agregar tu usuario al grupo docker
sudo usermod -aG docker $USER
```

### macOS
1. Opción 1 - Docker Desktop:
   - Instalar Docker Desktop desde [https://www.docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop)
   - Iniciar Docker Desktop

2. Opción 2 - Colima (alternativa más ligera):
```bash
# Instalar Colima
brew install colima

# Iniciar Colima
colima start

# Configurar el socket de Docker
export DOCKER_HOST="unix://${HOME}/.colima/default/docker.sock"
```

## Ejecución

1. Iniciar los contenedores:
```bash
# Windows (PowerShell)
docker compose up --build

# Linux/macOS
docker compose up --build
```

2. Acceder a pgAdmin:
- URL: http://localhost:5050
- Email: Configurado en PGADMIN_DEFAULT_EMAIL
- Password: Configurado en PGADMIN_DEFAULT_PASSWORD

3. Configurar el servidor en pgAdmin:
- Host: Configurado en DB_HOST
- Port: Configurado en DB_PORT
- Database: Configurado en DB_NAME
- Username: Configurado en DB_USER
- Password: Configurado en DB_PASSWORD

## Endpoints Disponibles

### Transacciones

#### Crear Transacción
```bash
# Windows (PowerShell)
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/transactions" `
  -Method Post `
  -Headers @{
    "Content-Type" = "application/json"
    "x-caller-id" = "test-user"
  } `
  -Body '{
    "type_id": 1,
    "description": "Test Transaction",
    "amount": 100.00,
    "payed": false,
    "expiry_date": "2024-12-31",
    "category": "Test"
  }'

# Linux/macOS
curl --location 'http://localhost:8080/api/v1/transactions' \
--header 'Content-Type: application/json' \
--header 'x-caller-id: test-user' \
--data '{
    "type_id": 1,
    "description": "Test Transaction",
    "amount": 100.00,
    "payed": false,
    "expiry_date": "2024-12-31",
    "category": "Test"
}'
```

#### Listar Transacciones
```bash
# Windows (PowerShell)
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/transactions" `
  -Method Get `
  -Headers @{
    "x-caller-id" = "test-user"
  }

# Linux/macOS
curl --location 'http://localhost:8080/api/v1/transactions' \
--header 'x-caller-id: test-user'
```

### Categorías

#### Crear Categoría
```bash
# Windows (PowerShell)
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/categories" `
  -Method Post `
  -Headers @{
    "Content-Type" = "application/json"
    "x-caller-id" = "test-user"
  } `
  -Body '{
    "name": "Test Category",
    "description": "Test Description"
  }'

# Linux/macOS
curl --location 'http://localhost:8080/api/v1/categories' \
--header 'Content-Type: application/json' \
--header 'x-caller-id: test-user' \
--data '{
    "name": "Test Category",
    "description": "Test Description"
}'
```

#### Listar Categorías
```bash
# Windows (PowerShell)
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/categories" `
  -Method Get `
  -Headers @{
    "x-caller-id" = "test-user"
  }

# Linux/macOS
curl --location 'http://localhost:8080/api/v1/categories' \
--header 'x-caller-id: test-user'
```

### Reportes

#### Obtener Resumen por Categoría
```bash
# Windows (PowerShell)
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/reports/category-summary" `
  -Method Get `
  -Headers @{
    "x-caller-id" = "test-user"
  }

# Linux/macOS
curl --location 'http://localhost:8080/api/v1/reports/category-summary' \
--header 'x-caller-id: test-user'
```

## Desarrollo

### Configuración de VS Code

El proyecto incluye configuración para debugging en VS Code:

1. Abrir el proyecto en VS Code
2. Presionar F5 para iniciar el debugging
3. Seleccionar "Start Docker & Debug Go"

### Estructura del Proyecto

```
.
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── categories/
│   ├── config/
│   ├── models/
│   ├── reports/
│   └── transactions/
├── Dockerfile
├── docker-compose.yml
├── .env.example
└── go.mod
```

## Detener la Aplicación

Para detener los contenedores:
```bash
# Windows (PowerShell)
docker compose down

# Linux/macOS
docker compose down
```

## Notas Importantes

- El header `x-caller-id` es requerido en todas las peticiones
- Las transacciones están asociadas a un usuario específico
- Los reportes se generan por usuario
- Nunca compartas o comitees el archivo `.env` con tus credenciales reales
- En Windows, asegúrate de que Docker Desktop esté ejecutándose antes de usar los comandos
- En Linux, puede ser necesario reiniciar la sesión después de agregar el usuario al grupo docker