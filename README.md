# 📚 Sistema de Gestión de Libros Electrónicos en Go

> Una aplicación backend completa desarrollada en Go para la administración integral de una biblioteca digital moderna.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white&style=flat-square)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-12+-336791?logo=postgresql&logoColor=white&style=flat-square)
![Gin](https://img.shields.io/badge/Gin-Web%20Framework-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![Status](https://img.shields.io/badge/Status-Production%20Ready-success?style=flat-square)

---

## 📋 Tabla de Contenidos

- [Descripción General](#descripción-general)
- [Características Principales](#características-principales)
- [Datos del Proyecto](#datos-del-proyecto)
- [Tecnologías Utilizadas](#tecnologías-utilizadas)
- [Arquitectura del Sistema](#arquitectura-del-sistema)
- [Instalación y Configuración](#instalación-y-configuración)
- [Uso y Ejemplos](#uso-y-ejemplos)
- [Módulos Disponibles](#módulos-disponibles)
- [Endpoints de la API](#endpoints-de-la-api)
- [Base de Datos](#base-de-datos)
- [Seguridad](#seguridad)
- [Despliegue](#despliegue)
- [Equipo de Desarrollo](#equipo-de-desarrollo)
- [Contribuciones](#contribuciones)
- [Licencia](#licencia)

---

## 🎯 Descripción General

El **Sistema de Gestión de Libros Electrónicos** es una aplicación backend robusta y escalable desarrollada en Go que simula el funcionamiento de una biblioteca digital moderna. El sistema permite la administración completa de usuarios, libros, préstamos, reservas, descargas y generación de reportes analíticos.

Este proyecto integra conocimientos avanzados de:
- 🔹 Lenguaje Go (Golang)
- 🔹 Arquitectura en capas (MVC + Repository Pattern)
- 🔹 APIs REST
- 🔹 Bases de datos relacionales (PostgreSQL)
- 🔹 Autenticación y seguridad (JWT + bcrypt)
- 🔹 Trabajo colaborativo con Git

---

## ✨ Características Principales

### 👥 Gestión de Usuarios y Autenticación
- ✅ Registro de nuevos usuarios
- ✅ Inicio de sesión con generación de JWT
- ✅ Autenticación segura con tokens
- ✅ Cifrado de contraseñas con bcrypt
- ✅ Control de acceso basado en roles (RBAC)
- ✅ Validaciones exhaustivas de entrada

### 📖 Catálogo de Libros
- ✅ Registro y gestión de libros electrónicos
- ✅ Búsqueda y filtrado avanzado (título, autor, categoría)
- ✅ Control de disponibilidad
- ✅ Información bibliográfica completa
- ✅ Soporte de múltiples formatos (PDF, EPUB, MOBI)

### 🔄 Sistema de Préstamos
- ✅ Crear y gestionar préstamos
- ✅ Registrar devoluciones
- ✅ Renovación de préstamos
- ✅ Control de fechas de vencimiento
- ✅ Historial de préstamos
- ✅ Detección de préstamos vencidos

### 📌 Gestión de Reservas
- ✅ Reservar libros no disponibles
- ✅ Cola de espera automática
- ✅ Cancelación de reservas
- ✅ Notificación cuando libro está disponible

### 📥 Historial de Descargas
- ✅ Registro de descargas de usuarios
- ✅ Historial completo y auditable
- ✅ Estadísticas de uso
- ✅ Análisis de libros más descargados

### 📊 Reportes y Estadísticas
- ✅ Reportes generales del sistema
- ✅ Estadísticas de préstamos
- ✅ Análisis de actividad de usuarios
- ✅ Libros más consultados
- ✅ Métricas de rendimiento

---

## 📊 Datos del Proyecto

| Aspecto | Descripción |
|---------|------------|
| **Nombre** | Sistema de Gestión de Libros Electrónicos |
| **Lenguaje Principal** | Go / Golang 1.21+ |
| **Base de Datos** | PostgreSQL 12+ |
| **Framework Web** | Gin Web Framework |
| **ORM** | GORM v1+ |
| **Autenticación** | JWT (JSON Web Tokens) |
| **Encriptación** | bcrypt |
| **API Format** | REST con JSON |
| **Plataforma de Despliegue** | Railway (Cloud) |
| **Control de Versiones** | Git / GitHub |
| **Fecha de Desarrollo** | Junio 2026 |
| **Estado** | Production Ready ✅ |
| **Total de Endpoints** | 35+ |

---

## 🔧 Tecnologías Utilizadas

### Backend Core
```
Go 1.21+              → Lenguaje de programación principal
Gin Web Framework     → Framework HTTP de alto rendimiento
GORM v1+              → ORM para PostgreSQL
```

### Base de Datos
```
PostgreSQL 12+        → Sistema de gestión de base de datos
pgAdmin               → Administración de PostgreSQL
```

### Seguridad
```
JWT (golang-jwt)      → Autenticación basada en tokens
bcrypt                → Cifrado de contraseñas
```

### Serialización
```
JSON                  → Formato de intercambio de datos
encoding/json         → Librería estándar de Go
```

### Desarrollo y Despliegue
```
GitHub                → Control de versiones colaborativo
Railway               → Plataforma de despliegue en la nube
Docker (opcional)     → Contenedorización
```

### Herramientas de Desarrollo
```
Postman               → Testing de APIs
Git                   → Sistema de control de versiones
Go Modules            → Gestión de dependencias
```

---

## 🏗️ Arquitectura del Sistema

### Diseño en Capas

```
┌─────────────────────────────────────────────────────────┐
│              HTTP CLIENT / POSTMAN / FRONTEND           │
├─────────────────────────────────────────────────────────┤
│                   GIN ROUTER (HTTP)                      │
│            Enrutamiento y validación de rutas            │
├─────────────────────────────────────────────────────────┤
│              MIDDLEWARE (Auth, Logging, etc)             │
│        Validación JWT, Control de acceso, Errores       │
├─────────────────────────────────────────────────────────┤
│                  HANDLERS (Controladores)                │
│          Recepción de solicitudes, Validación inicial    │
├─────────────────────────────────────────────────────────┤
│               SERVICES (Lógica de Negocio)              │
│    Validaciones, Transformaciones, Operaciones complejas │
├─────────────────────────────────────────────────────────┤
│              REPOSITORIES (Acceso a Datos)               │
│          Operaciones CRUD, Consultas SQL (GORM)         │
├─────────────────────────────────────────────────────────┤
│                   DATABASE LAYER (GORM)                  │
│              Mapeo relacional, Migraciones               │
├─────────────────────────────────────────────────────────┤
│                 PostgreSQL DATABASE                      │
│        Almacenamiento persistente, Transacciones        │
└─────────────────────────────────────────────────────────┘
```

### Estructura de Directorios

```
Autonomo1_go/
│
├── 📄 README.md                        # Documentación principal
├── 📄 InformeFinal.md                  # Informe técnico completo
├── 📄 VIDEO_PRESENTACION.md            # Referencias de presentación
│
└── GO_CODE_GESTION/                    # Código fuente principal
    │
    ├── 📄 Main.go                      # Punto de entrada del sistema
    ├── 📄 go.mod                       # Definición del módulo Go
    ├── 📄 go.sum                       # Checksums de dependencias
    ├── 📄 .env                         # Variables de entorno
    ├── 📄 .env.example                 # Template de variables
    │
    ├── 📁 database/
    │   ├── 📄 README.md                # Documentación de BD
    │   ├── 📄 connection.go            # Conexión a PostgreSQL
    │   ├── 📄 schema.sql               # Esquema de base de datos
    │   └── 📄 seed.sql                 # Datos de prueba
    │
    ├── 📁 internal/
    │   │
    │   ├── 📁 config/
    │   │   └── 📄 config.go            # Configuración general
    │   │
    │   ├── 📁 models/
    │   │   ├── 📄 user.go              # Modelo de Usuario
    │   │   ├── 📄 book.go              # Modelo de Libro
    │   │   ├── 📄 loan.go              # Modelo de Préstamo
    │   │   ├── 📄 reservation.go       # Modelo de Reserva
    │   │   └── 📄 download.go          # Modelo de Descarga
    │   │
    │   ├── 📁 handlers/
    │   │   ├── 📄 user_handler.go      # Handlers de Usuarios
    │   │   ├── 📄 book_handler.go      # Handlers de Libros
    │   │   ├── 📄 loan_handler.go      # Handlers de Préstamos
    │   │   ├── 📄 reservation_handler.go # Handlers de Reservas
    │   │   ├── 📄 download_handler.go  # Handlers de Descargas
    │   │   └── 📄 report_handler.go    # Handlers de Reportes
    │   │
    │   ├── 📁 services/
    │   │   ├── 📄 user_service.go      # Lógica de Usuarios
    │   │   ├── 📄 book_service.go      # Lógica de Libros
    │   │   ├── 📄 loan_service.go      # Lógica de Préstamos
    │   │   ├── 📄 reservation_service.go # Lógica de Reservas
    │   │   ├── 📄 download_service.go  # Lógica de Descargas
    │   │   └── 📄 report_service.go    # Lógica de Reportes
    │   │
    │   ├── 📁 repositories/
    │   │   ├── 📄 user_repository.go   # Repositorio de Usuarios
    │   │   ├── 📄 book_repository.go   # Repositorio de Libros
    │   │   ├── 📄 loan_repository.go   # Repositorio de Préstamos
    │   │   ├── 📄 reservation_repository.go # Repositorio de Reservas
    │   │   ├── 📄 download_repository.go # Repositorio de Descargas
    │   │   └── 📄 report_repository.go # Repositorio de Reportes
    │   │
    │   ├── 📁 middleware/
    │   │   ├── 📄 auth.go              # Middleware de Autenticación JWT
    │   │   ├── 📄 error_handler.go     # Manejo de errores global
    │   │   └── 📄 cors.go              # CORS (si aplica)
    │   │
    │   ├── 📁 routes/
    │   │   ├── 📄 routes.go            # Configuración central de rutas
    │   │   ├── 📄 user_routes.go       # Rutas de Usuarios
    │   │   ├── 📄 book_routes.go       # Rutas de Libros
    │   │   ├── 📄 loan_routes.go       # Rutas de Préstamos
    │   │   ├── 📄 reservation_routes.go # Rutas de Reservas
    │   │   ├── 📄 download_routes.go   # Rutas de Descargas
    │   │   └── 📄 report_routes.go     # Rutas de Reportes
    │   │
    │   └── 📁 utils/
    │       ├── 📄 helpers.go           # Funciones auxiliares
    │       ├── 📄 validators.go        # Validadores
    │       └── 📄 errors.go            # Manejo de errores personalizado

```

---

## 🚀 Instalación y Configuración

### Requisitos Previos

- **Go** 1.21 o superior
- **PostgreSQL** 12 o superior
- **Git**
- **Postman** o similar (opcional, para testing)

### Pasos de Instalación

#### 1️⃣ Clonar el Repositorio

```bash
git clone https://github.com/wixy122018-byte/Autonomo1_go.git
cd Autonomo1_go
```

#### 2️⃣ Navegar al Directorio del Código

```bash
cd GO_CODE_GESTION
```

#### 3️⃣ Instalar Dependencias

```bash
go mod tidy
```

Esto descargará e instalará todas las dependencias especificadas en `go.mod`.

#### 4️⃣ Configurar Base de Datos

**Crear la base de datos:**

```bash
psql -U postgres
CREATE DATABASE biblioteca;
\q
```

**Crear las tablas (automático con AutoMigrate):**

```bash
# Las tablas se crearán automáticamente cuando ejecutes Main.go
# O manualmente:
psql -U postgres -d biblioteca -f database/schema.sql
```

**Cargar datos de prueba (opcional):**

```bash
psql -U postgres -d biblioteca -f database/seed.sql
```

#### 5️⃣ Configurar Variables de Entorno

Crear archivo `.env` en la raíz de `GO_CODE_GESTION/`:

```env
# Puerto del servidor
PORT=8080

# Configuración de la Base de Datos
DATABASE_DSN=host=localhost user=postgres password=postgres dbname=biblioteca port=5432 sslmode=disable

# Configuración JWT
JWT_SECRET=clave_secreta_proyecto_go_muy_segura_min_32_caracteres
JWT_EXPIRE_HOURS=24

# Entorno
ENV=development

# Logging (opcional)
LOG_LEVEL=debug
```

**Para Railway (Producción):**

```env
DATABASE_URL=postgresql://usuario:password@host:5432/biblioteca
PORT=8080
JWT_SECRET=clave_secreta_produccion
ENV=production
```

#### 6️⃣ Ejecutar la Aplicación

```bash
go run Main.go
```

**Salida esperada:**

```
[GIN-debug] Listening and serving HTTP on :8080
[GIN-debug] GET    /health              --> github.com/.../handlers.Health (1 handler)
[GIN-debug] POST   /register            --> github.com/.../handlers.Register (1 handler)
[GIN-debug] POST   /login               --> github.com/.../handlers.Login (1 handler)
...
```

#### 7️⃣ Verificar que el Servidor está Corriendo

```bash
curl http://localhost:8080/health
```

**Respuesta esperada:**

```json
{
  "status": "ok",
  "database": "connected"
}
```

---

## 💡 Uso y Ejemplos

### Ejemplo 1: Registro de Usuario

**Solicitud:**

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Martín Gómez",
    "email": "martin@example.com",
    "password": "SecurePass123",
    "role": "LECTOR"
  }'
```

**Respuesta exitosa (201):**

```json
{
  "id": 1,
  "name": "Martín Gómez",
  "email": "martin@example.com",
  "role": "LECTOR",
  "created_at": "2026-06-29T10:30:00Z",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Ejemplo 2: Inicio de Sesión

**Solicitud:**

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "martin@example.com",
    "password": "SecurePass123"
  }'
```

**Respuesta exitosa (200):**

```json
{
  "id": 1,
  "name": "Martín Gómez",
  "email": "martin@example.com",
  "role": "LECTOR",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_in": 3600
}
```

### Ejemplo 3: Listar Libros

**Solicitud:**

```bash
curl -X GET http://localhost:8080/books \
  -H "Authorization: Bearer <TOKEN_JWT>"
```

**Respuesta exitosa (200):**

```json
{
  "total": 5,
  "limit": 10,
  "offset": 0,
  "books": [
    {
      "id": 1,
      "title": "Cien años de soledad",
      "author": "Gabriel García Márquez",
      "category": "Ficción",
      "publisher": "Penguin",
      "year": 1967,
      "available": true,
      "isbn": "978-8495618528",
      "format": "PDF",
      "description": "Una novela maestra de la literatura..."
    }
  ]
}
```

### Ejemplo 4: Crear Préstamo

**Solicitud:**

```bash
curl -X POST http://localhost:8080/loans \
  -H "Authorization: Bearer <TOKEN_JWT>" \
  -H "Content-Type: application/json" \
  -d '{
    "book_id": 1,
    "duration_days": 14
  }'
```

**Respuesta exitosa (201):**

```json
{
  "id": 1,
  "user_id": 1,
  "book_id": 1,
  "status": "ACTIVO",
  "start_date": "2026-06-29T10:30:00Z",
  "due_date": "2026-07-13T10:30:00Z",
  "return_date": null,
  "created_at": "2026-06-29T10:30:00Z"
}
```

### Ejemplo 5: Buscar Libros por Categoría

**Solicitud:**

```bash
curl -X GET "http://localhost:8080/books/category/Ficción" \
  -H "Authorization: Bearer <TOKEN_JWT>"
```

---

## 📚 Módulos Disponibles

### 1. **Módulo de Usuarios y Seguridad** 👥

**Responsable**: Yandry Álvarez

**Funcionalidades:**
- Registro de usuarios con validaciones
- Autenticación con JWT
- Control de acceso basado en roles
- Consulta y eliminación de usuarios
- Cifrado de contraseñas con bcrypt

**Endpoints:**
```
POST   /register              - Registrar usuario
POST   /login                 - Iniciar sesión
GET    /profile               - Obtener mi perfil
GET    /users                 - Listar usuarios (ADMIN)
GET    /users/:id             - Obtener usuario por ID (ADMIN)
PUT    /users/:id             - Actualizar usuario (ADMIN)
DELETE /users/:id             - Eliminar usuario (ADMIN)
```

### 2. **Módulo de Libros** 📖

**Responsable**: Michael Lescano

**Funcionalidades:**
- Gestión del catálogo de libros
- Búsqueda y filtrado avanzado
- Control de disponibilidad
- Información bibliográfica completa
- Soporte de múltiples formatos

**Endpoints:**
```
GET    /books                 - Listar todos los libros
GET    /books/:id             - Obtener detalles de un libro
GET    /books/search?q=:term  - Buscar libros
GET    /books/category/:cat   - Filtrar por categoría
POST   /books                 - Crear libro (ADMIN)
PUT    /books/:id             - Actualizar libro (ADMIN)
DELETE /books/:id             - Eliminar libro (ADMIN)
```

### 3. **Módulo de Préstamos** 🔄

**Responsable**: Martín Gómez

**Funcionalidades:**
- Crear y gestionar préstamos
- Registrar devoluciones
- Renovación de préstamos
- Control de vencimientos
- Historial completo

**Endpoints:**
```
POST   /loans                 - Crear préstamo
GET    /loans                 - Mis préstamos
GET    /loans/:id             - Detalles de préstamo
GET    /loans/active          - Préstamos activos
GET    /loans/overdue         - Préstamos vencidos
PUT    /loans/:id/return      - Registrar devolución
PUT    /loans/:id/renew       - Renovar préstamo
```

### 4. **Módulo de Reservas** 📌

**Responsable**: Martín Gómez

**Funcionalidades:**
- Reservar libros no disponibles
- Cola de espera automática
- Cancelación de reservas
- Notificaciones de disponibilidad

**Endpoints:**
```
POST   /reservations          - Crear reserva
GET    /reservations          - Mis reservas
GET    /reservations/:id      - Detalles de reserva
DELETE /reservations/:id      - Cancelar reserva
GET    /reservations/book/:id - Reservas pendientes
```

### 5. **Módulo de Descargas** 📥

**Responsable**: Michael Lescano

**Funcionalidades:**
- Registrar descargas
- Historial de descargas
- Estadísticas de uso
- Auditoría de acceso

**Endpoints:**
```
POST   /downloads             - Registrar descarga
GET    /downloads             - Mi historial
GET    /downloads/book/:id    - Descargas del libro
GET    /downloads/stats       - Estadísticas
```

### 6. **Módulo de Reportes** 📊

**Responsable**: Martín Gómez

**Funcionalidades:**
- Reportes generales del sistema
- Estadísticas de actividad
- Análisis de libros más usados
- Métricas de rendimiento

**Endpoints:**
```
GET    /reports/general       - Estadísticas generales
GET    /reports/loans         - Reporte de préstamos
GET    /reports/books         - Reporte de libros
GET    /reports/downloads     - Reporte de descargas
GET    /reports/users         - Reporte de usuarios
GET    /reports/analytics     - Análisis avanzado
```

---

## 🔌 Endpoints de la API

### Resumen de Endpoints

| Categoría | Total | Métodos |
|-----------|-------|---------|
| Autenticación | 2 | POST |
| Usuarios | 5 | GET, POST, PUT, DELETE |
| Libros | 7 | GET, POST, PUT, DELETE |
| Préstamos | 7 | GET, POST, PUT |
| Reservas | 5 | GET, POST, DELETE |
| Descargas | 4 | GET, POST |
| Reportes | 6 | GET |
| Utilidad | 2 | GET |
| **TOTAL** | **38** | - |

### Tabla Completa de Endpoints

| Método | Endpoint | Descripción | Autenticación | Rol |
|--------|----------|-------------|---------------|-----|
| **AUTENTICACIÓN** | | | | |
| POST | /register | Registrar usuario | No | - |
| POST | /login | Iniciar sesión | No | - |
| **USUARIOS** | | | | |
| GET | /profile | Mi perfil | Sí | - |
| GET | /users | Listar usuarios | Sí | ADMIN |
| GET | /users/:id | Obtener usuario | Sí | ADMIN |
| PUT | /users/:id | Actualizar usuario | Sí | ADMIN |
| DELETE | /users/:id | Eliminar usuario | Sí | ADMIN |
| **LIBROS** | | | | |
| GET | /books | Listar libros | Sí | - |
| GET | /books/:id | Obtener libro | Sí | - |
| GET | /books/search | Buscar libros | Sí | - |
| GET | /books/category/:cat | Filtrar por categoría | Sí | - |
| POST | /books | Crear libro | Sí | ADMIN |
| PUT | /books/:id | Actualizar libro | Sí | ADMIN |
| DELETE | /books/:id | Eliminar libro | Sí | ADMIN |
| **PRÉSTAMOS** | | | | |
| POST | /loans | Crear préstamo | Sí | - |
| GET | /loans | Mis préstamos | Sí | - |
| GET | /loans/:id | Obtener préstamo | Sí | - |
| GET | /loans/active | Préstamos activos | Sí | - |
| GET | /loans/overdue | Préstamos vencidos | Sí | ADMIN |
| PUT | /loans/:id/return | Registrar devolución | Sí | - |
| PUT | /loans/:id/renew | Renovar préstamo | Sí | - |
| **RESERVAS** | | | | |
| POST | /reservations | Crear reserva | Sí | - |
| GET | /reservations | Mis reservas | Sí | - |
| GET | /reservations/:id | Obtener reserva | Sí | - |
| DELETE | /reservations/:id | Cancelar reserva | Sí | - |
| GET | /reservations/book/:id | Reservas del libro | Sí | ADMIN |
| **DESCARGAS** | | | | |
| POST | /downloads | Registrar descarga | Sí | - |
| GET | /downloads | Mi historial | Sí | - |
| GET | /downloads/book/:id | Descargas del libro | Sí | ADMIN |
| GET | /downloads/stats | Estadísticas | Sí | ADMIN |
| **REPORTES** | | | | |
| GET | /reports/general | Estadísticas generales | Sí | ADMIN |
| GET | /reports/loans | Reporte de préstamos | Sí | ADMIN |
| GET | /reports/books | Reporte de libros | Sí | ADMIN |
| GET | /reports/downloads | Reporte de descargas | Sí | ADMIN |
| GET | /reports/users | Reporte de usuarios | Sí | ADMIN |
| GET | /reports/analytics | Análisis avanzado | Sí | ADMIN |
| **UTILIDAD** | | | | |
| GET | / | Ruta principal | No | - |
| GET | /health | Estado del servidor | No | - |

---

## 💾 Base de Datos

### Configuración de PostgreSQL

#### Variables de Entorno

```env
# En .env
DATABASE_DSN=host=localhost user=postgres password=postgres dbname=biblioteca port=5432 sslmode=disable

# En Railway
DATABASE_URL=postgresql://usuario:password@host:5432/biblioteca
```

#### Tablas Principales

**Tabla: users**
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'LECTOR',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Tabla: books**
```sql
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    category VARCHAR(100),
    publisher VARCHAR(255),
    year INTEGER,
    description TEXT,
    format VARCHAR(50),
    isbn VARCHAR(20) UNIQUE,
    file_path VARCHAR(500),
    available BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Tabla: loans**
```sql
CREATE TABLE loans (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    book_id INTEGER NOT NULL REFERENCES books(id),
    status VARCHAR(50) NOT NULL DEFAULT 'ACTIVO',
    start_date TIMESTAMP NOT NULL,
    due_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Tabla: reservations**
```sql
CREATE TABLE reservations (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    book_id INTEGER NOT NULL REFERENCES books(id),
    status VARCHAR(50) NOT NULL DEFAULT 'PENDIENTE',
    queue_position INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Tabla: downloads**
```sql
CREATE TABLE downloads (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    book_id INTEGER NOT NULL REFERENCES books(id),
    file_size INTEGER,
    format VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Usuarios de Prueba

Después de ejecutar `database/seed.sql`:

```
Email: estudiante@example.com
Password: secreto123
Role: LECTOR

Email: admin@example.com
Password: secreto123
Role: ADMINISTRADOR
```

---

## 🔒 Seguridad

### Autenticación JWT

El sistema utiliza **JSON Web Tokens (JWT)** para autenticación:

```
1. Usuario se registra o inicia sesión
2. Sistema genera un JWT con claims personalizados
3. Cliente incluye el token en el header Authorization
4. Middleware valida el token en rutas protegidas
5. Token expira después de 24 horas (configurable)
```

**Header del Token:**
```bash
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### Cifrado de Contraseñas

Se utiliza **bcrypt** con salt rounds de 10:

```
- Las contraseñas NUNCA se almacenan en texto plano
- Imposible recuperar la contraseña original
- Resistente a ataques de fuerza bruta
- Validación mediante comparación de hashes
```

### Control de Acceso (RBAC)

Roles disponibles:
- **LECTOR**: Acceso a libros, préstamos personales, descargas
- **ADMINISTRADOR**: Acceso completo a todas las funciones

### Validaciones de Entrada

- Email: Formato válido y único
- Contraseña: Mínimo 8 caracteres
- Nombre: No vacío, 3-100 caracteres
- Prevención de inyección SQL mediante GORM/parametrización

### Variables de Entorno Sensibles

```env
JWT_SECRET=clave_secreta_min_32_caracteres
DATABASE_DSN=postgresql://user:password@host/db
```

**Nunca commits secrets o tokens en el repositorio**

---

## 🌐 Despliegue

### Despliegue Local

```bash
# 1. Clonar y configurar
git clone https://github.com/wixy122018-byte/Autonomo1_go.git
cd Autonomo1_go/GO_CODE_GESTION

# 2. Instalar dependencias
go mod tidy

# 3. Configurar .env
cp .env.example .env
# Editar .env con tus valores

# 4. Ejecutar
go run Main.go
```

### Despliegue en Railway

Railway es una plataforma moderna para desplegar aplicaciones:

#### Paso 1: Conectar GitHub

1. Ir a [railway.app](https://railway.app)
2. Iniciar sesión con GitHub
3. Crear nuevo proyecto
4. Conectar repositorio `Autonomo1_go`

#### Paso 2: Configurar Variables de Entorno

En el dashboard de Railway:

```env
PORT=8080
DATABASE_URL=postgresql://...
JWT_SECRET=tu_clave_secreta_aqui
ENV=production
```

#### Paso 3: Desplegar

Railway se despliega automáticamente en cada push a la rama principal.

#### Verificar en Producción

```bash
curl https://tu-app.railway.app/health
```

---

## 👥 Equipo de Desarrollo

| Nombre | Rol | Responsabilidades |
|--------|-----|------------------|
| **Leonardo Sánchez** | Integrante 1 | Estructura base, almacenamiento e integración |
| **Michael Lescano** | Integrante 2 | Gestión de libros, búsquedas y descargas |
| **Martín Gómez** | Integrante 3 | Préstamos, reservas y reportes |
| **Yandry Álvarez** | Integrante 4 | Usuarios, seguridad y documentación |

---

## 🤝 Contribuciones

### Cómo Contribuir

1. **Fork el repositorio**

```bash
git clone https://github.com/wixy122018-byte/Autonomo1_go.git
```

2. **Crear rama de feature**

```bash
git checkout -b feature/nueva-funcionalidad
```

3. **Hacer cambios y commit**

```bash
git add .
git commit -m "Agregar nueva funcionalidad"
```

4. **Push a la rama**

```bash
git push origin feature/nueva-funcionalidad
```

5. **Crear Pull Request**

Describe los cambios en el PR y solicita revisión.

### Estándares de Código

- Seguir convenciones de Go
- Nombres descriptivos en español/inglés
- Documentar funciones públicas
- Incluir tests unitarios
- Validar con `go fmt` y `go vet`

---

## 📝 Documentación Completa

Para documentación más detallada, ver:

- **[GO_CODE_GESTION/database/README.md](./GO_CODE_GESTION/database/README.md)** - Documentación de base de datos
- **Código fuente** - Comentarios detallados en cada archivo

---

## 🎓 Relación con los Temas de la Asignatura

El proyecto integra conocimientos de todas las unidades:

- ✅ **Sintaxis básica de Go**: Estructuras, funciones, métodos
- ✅ **Funciones y paquetes**: Modularización y reutilización
- ✅ **Structs y métodos**: Modelado de datos
- ✅ **Encapsulación**: Paquetes internos (private/public)
- ✅ **Interfaces**: Uso de interfaces de Go
- ✅ **Manejo de errores**: Error handling exhaustivo
- ✅ **Base de datos**: PostgreSQL con GORM
- ✅ **Servicios web**: REST APIs con Gin
- ✅ **Serialización JSON**: Codificación/decodificación de datos
- ✅ **Organización modular**: Arquitectura en capas

---

## 🚀 Mejoras Futuras

### Corto Plazo
- [ ] Implementar pruebas unitarias completas
- [ ] Agregar Swagger/OpenAPI documentation
- [ ] Mejorar validaciones de entrada
- [ ] Implementar rate limiting

### Mediano Plazo
- [ ] Interfaz web (Frontend React/Vue)
- [ ] Autenticación OAuth2
- [ ] Sistema de notificaciones por email
- [ ] Caché con Redis
- [ ] WebSockets para actualizaciones en tiempo real

### Largo Plazo
- [ ] Recomendaciones basadas en IA
- [ ] Análisis predictivo de demanda
- [ ] Arquitectura de microservicios
- [ ] Integración con sistemas bibliotecarios externos
- [ ] Sincronización con otras bibliotecas

---

## 📊 Estadísticas del Proyecto

- **Lenguajes**: Go (95%), SQL (5%)
- **Líneas de código**: ~2500+ líneas
- **Endpoints**: 38+ endpoints REST
- **Módulos**: 6 módulos principales
- **Tablas BD**: 5 tablas relacionales
- **Commits**: 50+ commits en Git
- **Ramas**: 9 ramas de desarrollo
- **Status**: ✅ Production Ready


## ✨ Logros Alcanzados

✅ Sistema completamente funcional en Go  
✅ Arquitectura modular y escalable  
✅ Autenticación segura con JWT + bcrypt  
✅ Base de datos robusta en PostgreSQL  
✅ APIs REST completas y documentadas  
✅ Trabajo colaborativo exitoso  
✅ Despliegue en producción  
✅ Documentación técnica integral  

---

## Conclusión

El **Sistema de Gestión de Libros Electrónicos** demuestra la aplicación práctica de conceptos avanzados de programación backend en Go. El proyecto integra seguridad, modularidad, escalabilidad y mejores prácticas de desarrollo de software.

**Estado**: 🟢 Production Ready - Junio 2026

---

<div align="center">

</div>

---

**Última actualización**: Junio 29, 2026  
**Versión**: 2.0  
**Estado**: Active & Maintained ✅
