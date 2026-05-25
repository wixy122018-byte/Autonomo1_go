# 📚 Sistema de Gestión de Libros Electrónicos

> Un sistema moderno y escalable desarrollado en **Go** para la gestión integral de bibliotecas digitales

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![MySQL](https://img.shields.io/badge/MySQL-8.0+-4479A1?style=flat-square&logo=mysql)](https://www.mysql.com)
[![Gin](https://img.shields.io/badge/Gin-v1-00A86B?style=flat-square)](https://gin-gonic.com)
[![License](https://img.shields.io/badge/License-MIT-blue?style=flat-square)](LICENSE)

---

## 📑 Tabla de Contenidos

- [🎯 Descripción del Proyecto](#descripción)
- [✨ Características Principales](#características)
- [🏗️ Arquitectura](#arquitectura)
- [👥 Roles y Usuarios](#roles)
- [⚙️ Tecnologías](#tecnologías)
- [���� Quick Start](#quick-start)
- [📂 Estructura del Proyecto](#estructura)
- [🔄 Flujo del Sistema](#flujo)
- [🛠️ Instalación](#instalación)
- [📖 Documentación Adicional](#documentación)
- [👨‍💻 Desarrollador](#desarrollador)

---

## 🎯 Descripción del Proyecto {#descripción}

El **Sistema de Gestión de Libros Electrónicos** es una plataforma web integral que permite a bibliotecas y instituciones educativas administrar colecciones de libros digitales de forma eficiente.

**Características clave:**
- 📱 Interfaz responsiva e intuitiva
- 🔐 Autenticación y autorización por roles
- 📖 Gestión completa de catálogo de libros
- 🎫 Sistema de préstamos y reservas
- 📊 Reportes y estadísticas
- ⚡ Alto rendimiento con Go

---

## ✨ Características Principales {#características}

### 👤 Gestión de Usuarios
- ✅ Registro y autenticación de usuarios
- ✅ Gestión de roles (Administrador, Bibliotecario, Lector)
- ✅ Control de permisos granulares
- ✅ Historial de actividades

### 📚 Administración de Libros
- ✅ Catálogo digital completo
- ✅ Búsqueda avanzada por título, autor, categoría
- ✅ Control de disponibilidad
- ✅ Gestión de metadatos (ISBN, descripción, portada)

### 🎫 Sistema de Préstamos
- ✅ Solicitud de préstamos
- ✅ Cálculo automático de fechas de devolución
- ✅ Notificaciones de vencimiento
- ✅ Historial de préstamos

### 📌 Reservas y Disponibilidad
- ✅ Sistema de reservas para libros ocupados
- ✅ Cola de espera inteligente
- ✅ Notificaciones automáticas
- ✅ Gestión de reservas caducadas

### 📊 Reportes y Análisis
- ✅ Estadísticas de uso
- ✅ Reportes de préstamos
- ✅ Análisis de popularidad
- ✅ Dashboard administrativo

---

## 🏗️ Arquitectura {#arquitectura}

### Arquitectura por Capas

El sistema implementa una **arquitectura por capas** que garantiza separación de responsabilidades, escalabilidad y fácil mantenimiento.

```
┌─────────────────────────────────────┐
│  🌐 Cliente (HTML5/CSS3/JavaScript) │
└──────────────────┬──────────────────┘
                   │ HTTP/JSON
┌──────────────────▼──────────────────┐
│      🔀 API REST (Gin Framework)    │
└──────────────────┬──────────────────┘
                   │
┌──────────────────▼──────────────────┐
│   🎮 Controladores (Validación)     │
└──────────────────┬──────────────────┘
                   │
┌──────────────────▼──────────────────┐
│    ⚙️ Servicios (Lógica Negocio)    │
└──────────────────┬──────────────────┘
                   │
┌──────────────────▼──────────────────┐
│   📦 Repositorios (Acceso a Datos)  │
└──────────────────┬──────────────────┘
                   │
┌──────────────────▼──────────────────┐
│      💾 Base de Datos (MySQL)       │
└─────────────────────────────────────┘
```

### Ventajas de esta Arquitectura

| Aspecto | Beneficio |
|---------|-----------|
| **Separación de Responsabilidades** | Cada capa tiene un propósito específico |
| **Mantenibilidad** | Cambios aislados por capas, sin afectar otras |
| **Escalabilidad** | Fácil de extender sin rediseño |
| **Reutilización** | Componentes modulares y reutilizables |
| **Testabilidad** | Pruebas unitarias independientes por capa |
| **Seguridad** | Validación en múltiples puntos |

---

## 👥 Roles y Usuarios {#roles}

### Tabla de Roles y Permisos

| Rol | Permisos Clave | Funcionalidades |
|-----|-----------------|-----------------|
| **👨‍💼 Administrador** | Control total del sistema | • Gestionar usuarios<br>• Administrar libros<br>• Visualizar reportes completos<br>• Configurar parámetros del sistema<br>• Auditoría y logs |
| **📚 Bibliotecario** | Gestión de contenido | • Registrar nuevos libros<br>• Gestionar préstamos<br>• Controlar reservas<br>• Actualizar información<br>• Generar reportes básicos |
| **👨‍🎓 Lector/Estudiante** | Acceso limitado | • Buscar libros<br>• Solicitar préstamos<br>• Reservar libros<br>• Consultar historial<br>• Renovar préstamos |

---

## ⚙️ Tecnologías {#tecnologías}

### Backend
- **Go 1.21+** - Lenguaje principal
- **Gin Framework** - Framework web rápido
- **GORM** - ORM para Go
- **MySQL 8.0+** - Base de datos relacional

### Frontend
- **HTML5** - Estructura
- **CSS3** - Estilos
- **JavaScript (ES6+)** - Interactividad

### Herramientas de Desarrollo
- **Docker** - Containerización
- **Postman** - Testing de APIs
- **Git** - Control de versiones

---

## 🚀 Quick Start {#quick-start}

### Requisitos Previos
```bash
Go >= 1.21
MySQL >= 8.0
Git
```

### Instalación Rápida

```bash
# 1. Clonar el repositorio
git clone https://github.com/wixy122018-byte/Autonomo1_go.git
cd Autonomo1_go

# 2. Instalar dependencias
go mod download

# 3. Configurar base de datos
# Editar config/config.go con tus credenciales MySQL

# 4. Ejecutar la aplicación
go run main.go

# 5. Acceder a la aplicación
# Cliente: http://localhost:3000
# API: http://localhost:8080
```

---

## 📂 Estructura del Proyecto {#estructura}

```
Autonomo1_go/
│
├── main.go                          # Punto de entrada
│
├── controllers/                     # 🎮 Captura de solicitudes
│   ├── user_controller.go
│   ├── book_controller.go
│   ├── loan_controller.go
│   └── reservation_controller.go
│
├── services/                        # ⚙️ Lógica de negocio
│   ├── user_service.go
│   ├── book_service.go
│   ├── loan_service.go
│   └── reservation_service.go
│
├── repositories/                    # 📦 Acceso a datos
│   ├── user_repository.go
│   ├── book_repository.go
│   ├── loan_repository.go
│   └── reservation_repository.go
│
├── models/                          # 📊 Estructuras de datos
│   ├── user.go
│   ├── book.go
│   ├── loan.go
│   └── reservation.go
│
├── routes/                          # 🛣️ Rutas de la API
│   └── routes.go
│
├── middleware/                      # 🔐 Middlewares
│   ├── auth_middleware.go
│   └── cors_middleware.go
│
├── database/                        # 💾 Conexión BD
│   └── connection.go
│
├── utils/                           # 🛠️ Funciones auxiliares
│   ├── validators.go
│   └── helpers.go
│
├── config/                          # ⚙️ Configuración
│   └── config.go
│
├── frontend/                        # 🌐 Interfaz usuario
│   ├── index.html
│   ├── css/
│   └── js/
│
├── go.mod                           # Dependencias Go
├── go.sum                           # Hash de dependencias
├── .env.example                     # Variables de entorno
├── ARQUITECTURA_Y_TECNOLOGIA.md    # Documentación técnica
└── README.md                        # Este archivo
```

---

## 🔄 Flujo del Sistema {#flujo}

### Ciclo Completo de una Solicitud

```
1️⃣  Usuario interactúa con la interfaz
                    ↓
2️⃣  Cliente envía solicitud HTTP a la API REST
                    ↓
3️⃣  Gin Router recibe y enruta la petición
                    ↓
4️⃣  Controlador procesa y valida datos
                    ↓
5️⃣  Servicio ejecuta la lógica de negocio
                    ↓
6️⃣  Repositorio genera consultas SQL
                    ↓
7️⃣  Base de datos MySQL procesa la operación
                    ↓
8️⃣  Datos fluyen de regreso por todas las capas
                    ↓
9️⃣  Cliente recibe respuesta JSON
                    ↓
🔟 Interfaz se actualiza con los resultados
```

### Ejemplo Práctico: Búsqueda de Libro

```
Usuario escribe: "Clean Code"
        ↓
GET /api/libros?titulo=Clean%20Code
        ↓
Controlador valida parámetro de búsqueda
        ↓
Servicio ejecuta búsqueda
        ↓
Repositorio: SELECT * FROM libros WHERE titulo LIKE '%Clean Code%'
        ↓
MySQL retorna resultados
        ↓
JSON con libros encontrados
        ↓
Interfaz muestra resultados al usuario
```

---

## 🛠️ Instalación {#instalación}

### Opción 1: Instalación Local

```bash
# Clonar repositorio
git clone https://github.com/wixy122018-byte/Autonomo1_go.git
cd Autonomo1_go

# Descargar dependencias
go mod download

# Configurar variables de entorno
cp .env.example .env
# Editar .env con tus valores

# Crear base de datos
mysql -u root -p < database/init.sql

# Ejecutar aplicación
go run main.go

# La aplicación estará disponible en http://localhost:8080
```

### Opción 2: Usando Docker

```bash
# Construir imagen
docker build -t autonomo1_go .

# Ejecutar contenedor
docker run -p 8080:8080 -p 3000:3000 autonomo1_go

# Acceder a la aplicación
open http://localhost:8080
```

### Variables de Entorno (.env)

```env
# Base de Datos
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=tu_password
DB_NAME=biblioteca

# Server
SERVER_PORT=8080
GIN_MODE=debug

# JWT
JWT_SECRET=tu_secret_key_aqui
JWT_EXPIRATION=24
```

---

## 📖 Documentación Adicional {#documentación}

- **[ARQUITECTURA_Y_TECNOLOGIA.md](ARQUITECTURA_Y_TECNOLOGIA.md)** - Documentación técnica detallada
- **[API_ENDPOINTS.md](API_ENDPOINTS.md)** - Lista completa de endpoints
- **[GUIA_DESARROLLO.md](GUIA_DESARROLLO.md)** - Guía para desarrolladores
- **[MODELO_DATOS.md](MODELO_DATOS.md)** - Esquema de base de datos

---

## 📊 Endpoints Principales de la API

### Autenticación
```
POST   /api/auth/login          - Iniciar sesión
POST   /api/auth/register       - Registrarse
POST   /api/auth/logout         - Cerrar sesión
```

### Libros
```
GET    /api/libros              - Obtener catálogo
GET    /api/libros/:id          - Obtener detalle
POST   /api/libros              - Crear libro (Admin)
PUT    /api/libros/:id          - Actualizar libro (Admin)
DELETE /api/libros/:id          - Eliminar libro (Admin)
```

### Préstamos
```
GET    /api/prestamos           - Listar préstamos del usuario
POST   /api/prestamos           - Crear préstamo
PUT    /api/prestamos/:id       - Renovar préstamo
DELETE /api/prestamos/:id       - Devolver préstamo
```

### Reservas
```
GET    /api/reservas            - Listar reservas del usuario
POST   /api/reservas            - Crear reserva
DELETE /api/reservas/:id        - Cancelar reserva
```

---

## 🔒 Seguridad

- ✅ **Encriptación de contraseñas** con bcrypt
- ✅ **Autenticación JWT** para API
- ✅ **CORS configurado** para desarrollo/producción
- ✅ **Validación de entrada** en múltiples capas
- ✅ **SQL Injection prevention** con GORM parametrizado
- ✅ **Rate limiting** en endpoints críticos
- ✅ **Logs de auditoría** de operaciones sensibles

---

## 📈 Rendimiento

- ⚡ Desarrollado en **Go** para máximo rendimiento
- 🚀 Consultas optimizadas en base de datos
- 💨 Caché inteligente de datos frecuentes
- 📊 Índices en tablas principales

---

## 🤝 Contribuciones

Las contribuciones son bienvenidas. Para cambios significativos:

1. Fork el repositorio
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

---

## 📝 Licencia

Este proyecto está bajo la Licencia MIT. Ver archivo [LICENSE](LICENSE) para más detalles.

---

## 👨‍💻 Desarrollador {#desarrollador}

**Leonardo Sánchez**
- 🐱 GitHub: [@wixy122018-byte](https://github.com/wixy122018-byte)
- 📧 Email: leonardo.sanchez@example.com

---

## 🙋 Soporte

¿Preguntas o problemas?

- 📖 Consulta la [documentación completa](ARQUITECTURA_Y_TECNOLOGIA.md)
- 🐛 Abre un [issue](https://github.com/wixy122018-byte/Autonomo1_go/issues)
- 💬 Inicia una [discusión](https://github.com/wixy122018-byte/Autonomo1_go/discussions)

---

## 📅 Información del Proyecto

- **Versión**: 1.0.0
- **Última actualización**: Mayo 2026
- **Rama principal**: Leonardo_Sanchez
- **Estado**: 🟢 En desarrollo activo

---

<div align="center">

**⭐ Si te gusta este proyecto, considera darle una estrella en GitHub**

[Ver Repositorio](https://github.com/wixy122018-byte/Autonomo1_go) | [Reportar Bug](https://github.com/wixy122018-byte/Autonomo1_go/issues) | [Solicitar Feature](https://github.com/wixy122018-byte/Autonomo1_go/issues)

</div>
