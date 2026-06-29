# 📚 Sistema de Gestión de Libros Electrónicos en Go

**Una Aplicación Backend Completa para la Administración Integral de una Biblioteca Digital**

![Go](https://img.shields.io/badge/Go-1.26.4+-00ADD8?logo=go&logoColor=white&style=flat-square)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-12+-336791?logo=postgresql&logoColor=white&style=flat-square)
![Gin](https://img.shields.io/badge/Gin-Web%20Framework-00ADD8?style=flat-square)
![GORM](https://img.shields.io/badge/GORM-ORM-336791?style=flat-square)
![JWT](https://img.shields.io/badge/JWT-Auth-FF6B6B?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![Status](https://img.shields.io/badge/Status-Production%20Ready-success?style=flat-square)

---

## 📋 Tabla de Contenidos

- [Presentación Académica](#presentación-académica)
- [Descripción General](#descripción-general)
- [Características Principales](#características-principales)
- [Tecnologías Utilizadas](#tecnologías-utilizadas)
- [Arquitectura del Sistema](#arquitectura-del-sistema)
- [Estructura de Directorios](#estructura-de-directorios)
- [Instalación y Configuración](#instalación-y-configuración)
- [Módulos del Sistema](#módulos-del-sistema)
- [Endpoints de la API](#endpoints-de-la-api)
- [Base de Datos](#base-de-datos)
- [Seguridad y Autenticación](#seguridad-y-autenticación)
- [Uso y Ejemplos](#uso-y-ejemplos)
- [Despliegue](#despliegue)
- [Equipo de Desarrollo](#equipo-de-desarrollo)
- [Limitaciones y Mejoras Futuras](#limitaciones-y-mejoras-futuras)
- [Conclusión](#conclusión)

---

## 🎓 Presentación Académica

### Contexto Académico

Este proyecto corresponde al **trabajo final integrador** de la asignatura de Programación en Go, desarrollado en el semestre académico de Junio de 2026. El proyecto integra de manera práctica los conocimientos adquiridos durante todas las unidades de estudio.

### Objetivos de Aprendizaje Alcanzados

El desarrollo de este sistema permite validar la adquisición de las siguientes competencias:

- ✅ **Dominio del lenguaje Go**: Sintaxis, estructuras, funciones, métodos y paquetes
- ✅ **Arquitectura de software**: Patrón MVC, separación de capas y arquitectura modular
- ✅ **Desarrollo de APIs REST**: Diseño, implementación y documentación de servicios web
- ✅ **Gestión de bases de datos**: Modelado relacional, migraciones y ORM
- ✅ **Seguridad informática**: Autenticación (JWT), encriptación (bcrypt) y control de acceso (RBAC)
- ✅ **Desarrollo colaborativo**: Trabajo en equipo mediante Git, ramas de desarrollo y control de versiones
- ✅ **Serialización de datos**: Transformación entre formatos (JSON, Go structs)
- ✅ **Manejo de errores**: Tratamiento exhaustivo de excepciones y validaciones

### Problemática

Los sistemas bibliotecarios tradicionales generalmente operan con tecnologías legadas que dificultan la modernización, escalabilidad y acceso digital. Este proyecto propone una **solución backend moderna** que permite:

1. Gestión integral de recursos digitales
2. Control de acceso basado en roles
3. Automatización de procesos (préstamos, devoluciones, renovaciones)
4. Generación de reportes y estadísticas
5. Escalabilidad horizontal mediante arquitectura de capas

---

## 🎯 Descripción General

El **Sistema de Gestión de Libros Electrónicos** es una aplicación backend robusta y escalable desarrollada en Go que implementa las operaciones principales de una biblioteca digital moderna. El sistema permite la gestión completa del ciclo de vida de los recursos digitales, desde el registro hasta el análisis de uso.

### Características Clave

**🔹 Modularidad**: Arquitectura desacoplada mediante inyección de dependencias  
**🔹 Escalabilidad**: Diseñada para crecer horizontalmente  
**🔹 Seguridad**: Autenticación de múltiples capas y validaciones exhaustivas  
**🔹 Performance**: Optimización de consultas y caché de datos  
**🔹 Mantenibilidad**: Código limpio, documentado y testeable

---

## ✨ Características Principales

### 👥 Gestión de Usuarios y Autenticación
- ✅ Registro de nuevos usuarios con validaciones exhaustivas
- ✅ Autenticación mediante JWT (JSON Web Tokens)
- ✅ Cifrado de contraseñas con bcrypt (salt rounds: 10)
- ✅ Control de acceso basado en roles (RBAC): ADMINISTRADOR, BIBLIOTECARIO, LECTOR
- ✅ Consulta de perfil de usuario autenticado
- ✅ Administración de usuarios (solo ADMIN)
- ✅ Prevención de acceso no autorizado mediante middleware

### 📖 Catálogo de Libros
- ✅ Registro y gestión completa del catálogo
- ✅ Búsqueda avanzada por múltiples criterios
- ✅ Filtrado por categoría, autor y disponibilidad
- ✅ Información bibliográfica completa (ISBN, editorial, año, descripción)
- ✅ Soporte de múltiples formatos (PDF, EPUB, MOBI)
- ✅ Control de disponibilidad y estado
- ✅ Gestión de permisos para crear/editar/desactivar libros

### 📥 Historial de Descargas
- ✅ Registro automático de descargas de usuarios
- ✅ Historial completo y auditable
- ✅ Estadísticas de uso por usuario y por libro
- ✅ Análisis de libros más descargados
- ✅ Rastreo de acceso a recursos

### 🔄 Sistema de Préstamos (Planificado)
- 📋 Crear y gestionar préstamos
- 📋 Registrar devoluciones
- 📋 Renovación de préstamos
- 📋 Control de fechas de vencimiento

### 📌 Gestión de Reservas (Planificado)
- 📋 Reservar libros no disponibles
- 📋 Cola de espera automática

### 📊 Reportes y Estadísticas (Planificado)
- 📋 Reportes generales del sistema
- 📋 Estadísticas de actividad

---

## 🔧 Tecnologías Utilizadas

### Stack Principal

| Componente | Tecnología | Versión | Propósito |
|-----------|-----------|---------|----------|
| **Lenguaje** | Go (Golang) | 1.26.4+ | Lenguaje de programación principal |
| **Framework Web** | Gin | v1.12.0 | Enrutamiento HTTP y manejo de solicitudes |
| **ORM** | GORM | v1.31.1 | Mapeo objeto-relacional con PostgreSQL |
| **Base de Datos** | PostgreSQL | 12+ | Sistema de gestión de base de datos relacional |
| **Autenticación** | JWT | v5.3.1 | Generación y validación de tokens |
| **Criptografía** | bcrypt | golang.org/x/crypto | Cifrado de contraseñas |
| **Configuración** | godotenv | v1.5.1 | Gestión de variables de entorno |

---

## 🏗️ Arquitectura del Sistema

### Patrón Arquitectónico: Arquitectura en Capas

```
┌─────────────────────────────────────────────────────────────┐
│         CLIENTE HTTP (Postman / Frontend / cURL)            │
├─────────────────────────────────────────────────────────────┤
│  ┌─────────────────────────────────────────────────────┐   │
│  │ 1. CAPA DE PRESENTACIÓN (Gin Router)               │   │
│  │    - Enrutamiento de solicitudes HTTP              │   │
│  │    - Validación inicial de parámetros              │   │
│  │    - Serialización de respuestas JSON              │   │
│  └─────────────────────────────────────────────────────┘   │
│           ↓                                                  │
│  ┌─────────────────────────────────────────────────────┐   │
│  │ 2. CAPA DE MIDDLEWARE (Autenticación)              │   │
│  │    - Validación de tokens JWT                      │   │
│  │    - Control de acceso (RBAC)                      │   │
│  │    - Manejo de errores global                      │   │
│  └─────────────────────────────────────────────────────┘   │
│           ↓                                                  │
│  ┌─────────────────────────────────────────────────────┐   │
│  │ 3. CAPA DE HANDLERS (Controladores)                │   │
│  │    - Recepción de solicitudes HTTP                 │   │
│  │    - Orquestación con servicios                    │   │
│  └─────────────────────────────────────────────────────┘   │
│           ↓                                                  │
│  ┌─────────────────────────────────────────────────────┐   │
│  │ 4. CAPA DE SERVICIOS (Lógica de Negocio)          │   │
│  │    - Validaciones de datos                         │   │
│  │    - Transformaciones de datos                     │   │
│  │    - Reglas de negocio                             │   │
│  └─────────────────────────────────────────────────────┘   │
│           ↓                                                  │
│  ┌─────────────────────────────────────────────────────┐   │
│  │ 5. CAPA DE REPOSITORIOS (Acceso a Datos)          │   │
│  │    - Operaciones CRUD                              │   │
│  │    - Consultas complejas                           │   │
│  └─────────────────────────────────────────────────────┘   │
│           ↓                                                  │
│  ┌─────────────────────────────────────────────────────┐   │
│  │ 6. CAPA DE MODELOS (Entidades)                     │   │
│  │    - Estructuras de datos (Go structs)             │   │
│  │    - Mappeo a tablas de BD                         │   │
│  └─────────────────────────────────────────────────────┘   │
│           ↓                                                  │
│  ┌─────────────────────────────────────────────────────┐   │
│  │ 7. CAPA DE PERSISTENCIA (PostgreSQL)               │   │
│  │    - Almacenamiento relacional                     │   │
│  │    - Integridad referencial                        │   │
│  └─────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
```

---

## 📂 Estructura de Directorios

```
GO_CODE_GESTION/
│
├── 📄 Main.go                              # Punto de entrada (main)
│   ├─ config.LoadEnv() - Carga variables
│   ├─ database.Connect() - Conexión a BD
│   ├─ routes.RegisterRoutes() - Registro de rutas
│   └─ router.Run() - Inicio del servidor
│
├── 📄 go.mod                               # Definición del módulo Go
├── 📄 go.sum                               # Checksums de dependencias
├── 📄 .env                                 # Variables de entorno
├── 📄 .env.example                         # Plantilla de variables
│
├── 📁 database/
│   ├── 📄 connection.go                    # Conexión GORM a PostgreSQL
│   │   └─ Connect() - Establece BD
│   ├── 📄 schema.sql                       # Esquema SQL
│   └── 📄 seed.sql                         # Datos de prueba
│
└── 📁 internal/
    │
    ├── 📁 config/
    │   └── 📄 config.go
    │       ├─ LoadEnv() - Carga .env
    │       ├─ GetServerPort() - Puerto
    │       ├─ GetJWTSecret() - Secret JWT
    │       └─ GetDatabaseDSN() - Conexión BD
    │
    ├── 📁 models/                          # CAPA DE MODELOS
    │   ├── 📄 user.go
    │   │   ├─ User struct (id, name, email, password, role)
    │   │   ├─ SetRole(role) - Validar rol
    │   │   ├─ GetRole() - Obtener rol
    │   │   └─ HasRole(role) - Verificar rol
    │   ├── 📄 book.go
    │   │   ├─ Book struct (title, author, category, isbn, available)
    │   │   └─ Métodos de validación
    │   └── 📄 download.go
    │       ├─ Download struct (user_id, book_id, created_at)
    │       └─ Relaciones User/Book
    │
    ├── 📁 repositories/                    # CAPA DE REPOSITORIOS
    │   ├── 📄 user_repository.go
    │   │   ├─ FindByID(id) - Buscar usuario
    │   │   ├─ FindAll() - Listar usuarios
    │   │   ├─ FindByEmail(email) - Buscar por email
    │   │   ├─ Save(user) - Guardar usuario
    │   │   └─ DeleteByID(id) - Eliminar usuario
    │   ├── 📄 book_repository.go
    │   │   ├─ FindByID(id) - Buscar libro
    │   │   ├─ FindAll() - Listar libros
    │   │   ├─ Search(query) - Búsqueda
    │   │   ├─ Save(book) - Guardar libro
    │   │   └─ DeleteByID(id) - Eliminar libro
    │   └── 📄 download_repository.go
    │       ├─ Save(download) - Registrar descarga
    │       ├─ FindByUserID(userID) - Historial
    │       └─ GetStats(bookID) - Estadísticas
    │
    ├── 📁 services/                        # CAPA DE SERVICIOS
    │   ├── 📄 auth_service.go
    │   │   ├─ Register(name, email, pwd, role) - Crear usuario
    │   │   ├─ Login(email, password) - Autenticar
    │   │   ├─ ValidateToken(token) - Validar JWT
    │   │   └─ generateToken(user) - Generar JWT
    │   ├── 📄 user_service.go
    │   │   ├─ GetUserByID(id) - Obtener usuario
    │   │   ├─ GetAllUsers() - Listar usuarios
    │   │   ├─ DeleteUser(id) - Eliminar usuario
    │   │   ├─ ValidateEmail(email) - Validar email
    │   │   └─ ValidatePassword(pwd) - Validar contraseña
    │   ├── 📄 book_service.go
    │   │   ├─ CreateBook(book) - Crear libro
    │   │   ├─ UpdateBook(id, updates) - Actualizar
    │   │   ├─ DeleteBook(id) - Desactivar
    │   │   ├─ ListBooks(page, size) - Listar paginado
    │   │   └─ SearchBooks(query) - Buscar
    │   └── 📄 download_service.go
    │       ├─ RegisterDownload(userID, bookID) - Registrar
    │       ├─ GetUserHistory(userID) - Historial
    │       └─ GetStatistics(bookID) - Estadísticas
    │
    ├── 📁 handlers/                        # CAPA DE HANDLERS
    │   ├── 📄 auth_handler.go
    │   │   ├─ Register(c) - POST /register
    │   │   └─ Login(c) - POST /login
    │   ├── 📄 user_handler.go
    │   │   ├─ GetUsers(c) - GET /users
    │   │   ├─ GetUserByID(c) - GET /users/:id
    │   │   ├─ Profile(c) - GET /profile
    │   │   └─ DeleteUser(c) - DELETE /users/:id
    │   ├── 📄 book_handler.go
    │   │   ├─ List(c) - GET /books
    │   │   ├─ FindByID(c) - GET /books/:id
    │   │   ├─ Search(c) - GET /books/search
    │   │   ├─ Create(c) - POST /books
    │   │   ├─ Update(c) - PUT /books/:id
    │   │   └─ Deactivate(c) - DELETE /books/:id
    │   ├── 📄 download_handler.go
    │   │   ├─ Register(c) - POST /downloads
    │   │   └─ History(c) - GET /downloads/history
    │   └── 📄 common.go
    │       └─ WebServicesCatalog(c) - GET /api/v1/services
    │
    ├── 📁 middleware/                      # CAPA DE MIDDLEWARE
    │   ├── 📄 auth.go
    │   │   ├─ RequireAuth() - Validar JWT
    │   │   ├─ RequireRoles(...roles) - Validar roles
    │   │   └─ extractToken(c) - Extraer token
    │   └── 📄 errors.go
    │       └─ Manejo centralizado de errores
    │
    ├── 📁 routes/                          # CAPA DE RUTAS
    │   └── 📄 routes.go
    │       ├─ RegisterRoutes(router) - Registro central
    │       ├─ Rutas públicas (register, login)
    │       ├─ Rutas protegidas (con middleware JWT)
    │       └─ Rutas por rol (ADMIN, BIBLIOTECARIO)
    │
    └── 📁 utils/                           # UTILIDADES
        ├── 📄 helpers.go - Funciones auxiliares
        ├── 📄 validators.go - Validadores
        └── 📄 errors.go - Tipos de errores personalizados
```

---

## 🚀 Instalación y Configuración

### Requisitos Previos

- **Go**: versión 1.26.4 o superior
- **PostgreSQL**: versión 12 o superior
- **Git**: para control de versiones
- **Postman** o **cURL**: para testing de APIs

### Paso 1: Clonar el Repositorio

```bash
git clone https://github.com/wixy122018-byte/Autonomo1_go.git
cd Autonomo1_go/GO_CODE_GESTION
```

### Paso 2: Instalar Dependencias

```bash
go mod tidy
```

### Paso 3: Configurar Base de Datos

```bash
# Crear base de datos
psql -U postgres -c "CREATE DATABASE biblioteca;"

# Crear tablas (automático con AutoMigrate o manual)
psql -U postgres -d biblioteca -f database/schema.sql

# Cargar datos de prueba
psql -U postgres -d biblioteca -f database/seed.sql
```

### Paso 4: Configurar Variables de Entorno

```bash
cp .env.example .env
# Editar .env con tus valores
```

**Contenido de `.env`:**

```env
PORT=8080
DATABASE_DSN=host=localhost user=postgres password=postgres dbname=biblioteca port=5432 sslmode=disable
JWT_SECRET=tu_clave_secreta_muy_segura_minimo_32_caracteres
JWT_EXPIRE_HOURS=24
ENV=development
LOG_LEVEL=debug
```

### Paso 5: Ejecutar la Aplicación

```bash
go run Main.go

# Salida esperada:
# Conexión exitosa a la base de datos
# [GIN-debug] Listening and serving HTTP on :8080
```

### Paso 6: Verificar el Servidor

```bash
curl http://localhost:8080/health
# {"status":"ok"}
```

---

## 📚 Módulos del Sistema

### 1️⃣ Módulo de Usuarios y Autenticación 👥

**Responsable**: Yandry Álvarez

**Archivos Clave**:
- `internal/models/user.go` - Modelo User con roles
- `internal/services/auth_service.go` - Lógica de registro/login
- `internal/services/user_service.go` - Gestión de usuarios
- `internal/handlers/auth_handler.go` - Endpoints de autenticación
- `internal/handlers/user_handler.go` - Endpoints de usuarios
- `internal/middleware/auth.go` - Validación JWT y RBAC

**Funcionalidades**:
- ✅ Registro con validaciones
- ✅ Login con generación JWT
- ✅ Autenticación por token
- ✅ Control de roles (ADMIN, BIBLIOTECARIO, LECTOR)
- ✅ Gestión de usuarios (solo ADMIN)

### 2️⃣ Módulo de Libros 📖

**Responsable**: Michael Lescano

**Archivos Clave**:
- `internal/models/book.go` - Modelo Book
- `internal/services/book_service.go` - Lógica de libros
- `internal/handlers/book_handler.go` - Endpoints de libros
- `internal/repositories/book_repository.go` - Acceso a datos

**Funcionalidades**:
- ✅ Crear, leer, actualizar, desactivar libros
- ✅ Búsqueda por título/autor
- ✅ Filtrado por categoría
- ✅ Control de disponibilidad

### 3️⃣ Módulo de Descargas 📥

**Responsable**: Michael Lescano

**Archivos Clave**:
- `internal/models/download.go` - Modelo Download
- `internal/services/download_service.go` - Lógica de descargas
- `internal/handlers/download_handler.go` - Endpoints
- `internal/repositories/download_repository.go` - Acceso a datos

**Funcionalidades**:
- ✅ Registrar descargas
- ✅ Historial de descargas
- ✅ Estadísticas de uso

---

## 🔌 Endpoints de la API

### Resumen

| Categoría | Cantidad | Estado |
|-----------|----------|--------|
| Autenticación | 2 | ✅ Implementado |
| Usuarios | 4 | ✅ Implementado |
| Libros | 6 | ✅ Implementado |
| Descargas | 2 | ✅ Implementado |
| Préstamos | 7 | 📋 Planificado |
| Reservas | 5 | 📋 Planificado |
| Reportes | 6 | 📋 Planificado |
| Utilidad | 3 | ✅ Implementado |
| **TOTAL** | **35+** | - |

### Autenticación (Públicos)

```
POST /register           - Registrar usuario
POST /login              - Iniciar sesión
```

**Ejemplo**:
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Juan","email":"juan@example.com","password":"Pass123","role":"LECTOR"}'
```

### Usuarios (Protegidos)

```
GET /profile             - Mi perfil (autenticado)
GET /users               - Listar usuarios (ADMIN)
GET /users/:id           - Obtener usuario (ADMIN)
DELETE /users/:id        - Eliminar usuario (ADMIN)
```

### Libros (Protegidos)

```
GET /books               - Listar libros (autenticado)
GET /books/:id           - Obtener libro (autenticado)
GET /books/search?q=     - Buscar libros (autenticado)
POST /books              - Crear libro (ADMIN/BIBLIOTECARIO)
PUT /books/:id           - Actualizar libro (ADMIN/BIBLIOTECARIO)
DELETE /books/:id        - Desactivar libro (ADMIN/BIBLIOTECARIO)
```

### Descargas (Protegidos)

```
POST /downloads          - Registrar descarga (autenticado)
GET /downloads/history   - Mi historial (autenticado)
```

### Utilidad (Públicos)

```
GET /                    - Ruta principal
GET /health              - Estado del servidor
GET /api/v1/services     - Catálogo de servicios
```

---

## 💾 Base de Datos

### Tablas Implementadas

```sql
-- Tabla: users
-- Campos: id, name, email, password, role, created_at, updated_at
-- Índices: email (UNIQUE), role

-- Tabla: books
-- Campos: id, title, author, category, publisher, year, description, 
--         format, isbn (UNIQUE), file_path, available, created_at, updated_at
-- Índices: title, author, category, available, isbn

-- Tabla: downloads
-- Campos: id, user_id (FK), book_id (FK), file_size, format, created_at
-- Índices: user_id, book_id, created_at

-- Tablas Planificadas:
-- - loans (préstamos)
-- - reservations (reservas)
```

### Configuración de Conexión

```env
# Desarrollo local
DATABASE_DSN=host=localhost user=postgres password=postgres dbname=biblioteca port=5432 sslmode=disable

# Producción (Railway)
DATABASE_URL=postgresql://usuario:password@host.railway.app:5432/biblioteca
```

---

## 🔒 Seguridad y Autenticación

### Estrategia Multicapa

```
NIVEL 1: RBAC (Control de Roles)
├─ ADMINISTRADOR: Acceso total
├─ BIBLIOTECARIO: Gestión de libros
└─ LECTOR: Acceso a recursos

NIVEL 2: JWT (Autenticación)
├─ Token generado en login
├─ Validado en cada solicitud protegida
└─ Expira en 24 horas

NIVEL 3: bcrypt (Criptografía)
├─ Salt rounds: 10
├─ Imposible recuperar contraseña original
└─ Resistente a ataques de fuerza bruta

NIVEL 4: Validaciones (Input Validation)
├─ Email: formato válido y único
├─ Contraseña: mínimo 6 caracteres
└─ Prevención de inyección SQL (GORM)
```

### Flujo JWT

1. Usuario hace login con credenciales
2. Se valida en BD y se hashea contraseña
3. Se genera JWT con claims: user_id, email, role, exp
4. Token se devuelve al cliente
5. Cliente incluye token en header Authorization
6. Middleware valida firma y expiración
7. Se extrae user_id y se verifica rol
8. Se ejecuta handler si autorización es correcta

---

## 💡 Uso y Ejemplos

### Escenario: Registro, Login y Descarga

```bash
# 1. Registrar usuario
TOKEN=$(curl -s -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"name":"María","email":"maria@ex.com","password":"Pass123","role":"LECTOR"}' \
  | jq -r '.token')

# 2. Listar libros (con token)
curl -X GET http://localhost:8080/books \
  -H "Authorization: Bearer $TOKEN"

# 3. Descargar libro
curl -X POST http://localhost:8080/downloads \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"book_id":1}'

# 4. Ver historial
curl -X GET http://localhost:8080/downloads/history \
  -H "Authorization: Bearer $TOKEN"
```

---

## 🌐 Despliegue

### Local

```bash
go run Main.go
# Servidor en http://localhost:8080
```

### Railway (Producción)

1. Push a GitHub
2. Conectar repositorio en railway.app
3. Configurar variables de entorno
4. Railway se despliega automáticamente
5. URL de producción: https://autonomo1go-production-eb86.up.railway.app/

---

## 👥 Equipo de Desarrollo

| # | Nombre | Rol | Responsabilidades |
|---|--------|-----|-------------------|
| 1 | Leonardo Sánchez | Integrante 1 | Estructura base, configuración, integración |
| 2 | Michael Lescano | Integrante 2 | Libros, búsqueda, descargas |
| 3 | Martín Gómez | Integrante 3 | Préstamos, reservas, reportes |
| 4 | Yandry Álvarez | Integrante 4 | Usuarios, autenticación, seguridad, documentación |

---

### Mejoras Futuras

**Corto Plazo**: Tests unitarios, Swagger, rate limiting  
**Mediano Plazo**: Frontend web, OAuth2, notificaciones por email, Redis  
**Largo Plazo**: ML/IA, microservicios, app móvil

---

## 🎓 Conclusión

El **Sistema de Gestión de Libros Electrónicos** demuestra la aplicación práctica de conceptos fundamentales de ingeniería de software en Go:

✅ Sistema completamente funcional en producción  
✅ Arquitectura modular y escalable  
✅ Autenticación segura con JWT + bcrypt  
✅ APIs REST completas y documentadas  
✅ Trabajo colaborativo exitoso con Git  
✅ Despliegue en Railway  

Competencias validadas: Programación en Go, Arquitectura de software, APIs REST, Seguridad, Trabajo en equipo.

---

## 📄 Licencia

MIT License - Copyright (c) 2026

---

<div align="center">


Junio 2026 | Programación en Go

Última actualización: 29 de Junio, 2026 | Versión: 2.0 | 🟢 Production Ready ✅

</div>
