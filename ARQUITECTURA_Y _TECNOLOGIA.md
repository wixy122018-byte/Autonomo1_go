# 📚 ARQUITECTURA DEL SISTEMA DE GESTIÓN DE LIBROS ELECTRÓNICOS

## 📋 Tabla de Contenidos
1. [Introducción](#introducción)
2. [Flujo General de Arquitectura](#flujo-general)
3. [Estructura del Proyecto](#estructura)
4. [Usuarios y Roles](#usuarios-roles)
5. [Capas de la Arquitectura](#capas)
6. [Flujo Completo](#flujo-completo)

---

## 🎯 Introducción {#introducción}

El Sistema de Gestión de Libros Electrónicos será desarrollado en **Go**, implementando una arquitectura por capas que permite mantener el sistema organizado, escalable y fácil de mantener.

**Beneficios de esta arquitectura:**
- ✅ Separación de responsabilidades
- ✅ Código modular y reutilizable
- ✅ Fácil mantenimiento
- ✅ Mejor seguridad
- ✅ Escalabilidad

---

## 🔄 Flujo General de Arquitectura {#flujo-general}

```
┌─────────┐
│ Cliente │ (Interfaz Visual)
└────┬────┘
     │ HTTP
     ▼
┌─────────────┐
│ API REST    │ (Gin Framework)
└────┬────────┘
     │
     ▼
┌─────────────────┐
│ Controladores   │ (Recepción y Validación)
└────┬────────────┘
     │
     ▼
┌─────────────────┐
│ Servicios       │ (Lógica de Negocio)
└────┬────────────┘
     │
     ▼
┌─────────────────┐
│ Repositorios    │ (Acceso a Datos)
└────┬────────────┘
     │
     ▼
┌─────────────────┐
│ Base de Datos   │ (MySQL)
└─────────────────┘
```

---

## 🗂️ Estructura del Proyecto {#estructura}

```
/proyecto-libros
│
├── main.go                 # Punto de entrada de la aplicación
│
├── controllers/            # Controladores (Recepción de solicitudes)
│   ├── user_controller.go
│   ├── book_controller.go
│   ├── loan_controller.go
│   └── reservation_controller.go
│
├── services/               # Servicios (Lógica de negocio)
│   ├── user_service.go
│   ├── book_service.go
│   ├── loan_service.go
│   └── reservation_service.go
│
├── repositories/           # Repositorios (Acceso a datos)
│   ├── user_repository.go
│   ├── book_repository.go
│   ├── loan_repository.go
│   └── reservation_repository.go
│
├── models/                 # Modelos de datos
│   ├── user.go
│   ├── book.go
│   ├── loan.go
│   └── reservation.go
│
├── routes/                 # Rutas de la API
│   └── routes.go
│
├── middleware/             # Middleware (Autenticación, CORS)
│   ├── auth_middleware.go
│   └── cors_middleware.go
│
├── database/               # Configuración de BD
│   └── connection.go
│
├── utils/                  # Utilidades y funciones auxiliares
│   └── validators.go
│
├── config/                 # Configuración de la aplicación
│   └── config.go
│
└── README.md
```

---

## 👥 Usuarios y Roles {#usuarios-roles}

### Tabla de Roles y Permisos

| **Rol** | **Permisos Clave** | **Funcionalidades Principales** |
|---------|-------------------|--------------------------------|
| **Administrador** | Control total | • Gestiona usuarios<br>• Administra libros<br>• Visualiza reportes<br>• Configura parámetros del sistema |
| **Bibliotecario** | Gestión de contenido | • Registra libros<br>• Gestiona préstamos<br>• Controla reservas<br>• Actualiza información |
| **Lector/Estudiante** | Acceso limitado | • Busca libros<br>• Solicita préstamos<br>• Reserva libros<br>• Consulta historial |

---

## 🏗️ Capas de la Arquitectura {#capas}

### 1️⃣ CAPA CLIENTE

**Descripción:** Interfaz visual donde interactúan los usuarios directamente.

| **Aspecto** | **Descripción** |
|------------|-----------------|
| **Tecnologías** | HTML5, CSS3, JavaScript |
| **Responsabilidad** | Captura de datos y visualización |
| **Tipo de Interacción** | Formularios, botones, menús |

**Funcionalidades Principales:**

| **Funcionalidad** | **Descripción** |
|------------------|-----------------|
| **Inicio de Sesión** | Autenticación mediante correo y contraseña |
| **Visualización de Libros** | Ver catálogo, consultar autores, leer descripciones |
| **Búsqueda y Filtros** | Buscar por título, autor, categoría o palabras clave |
| **Solicitudes y Reservas** | Solicitar préstamos, reservar libros, consultar historial |
| **Panel Administrativo** | Gestión de usuarios, reportes, estadísticas y configuraciones |

**Flujo de Funcionamiento:**
1. Usuario realiza una acción en la interfaz
2. Se captura la información
3. Se genera una solicitud HTTP
4. La solicitud se envía a la API REST
5. El sistema procesa la operación
6. La respuesta retorna al cliente

---

### 2️⃣ CAPA API REST

**Framework:** Gin

**Descripción:** Puente de comunicación entre la interfaz y el backend.

| **Aspecto** | **Descripción** |
|------------|-----------------|
| **Framework** | Gin (Go) |
| **Formato de Datos** | JSON |
| **Responsabilidad** | Enrutamiento y coordinación |
| **Clientes Soportados** | Web, móvil, servicios externos |

**Métodos HTTP Soportados:**

| **Método** | **Propósito** | **Ejemplo** |
|-----------|--------------|-----------|
| **GET** | Obtener información | `GET /libros` |
| **POST** | Enviar información nueva | `POST /usuarios` |
| **PUT** | Actualizar información | `PUT /libros/15` |
| **DELETE** | Eliminar información | `DELETE /reservas/10` |

**Endpoints Principales:**

| **Endpoint** | **Método** | **Descripción** |
|-------------|----------|-----------------|
| `/login` | POST | Iniciar sesión |
| `/libros` | GET | Obtener lista de libros |
| `/libros/:id` | GET | Obtener detalle de un libro |
| `/usuarios` | POST | Registrar nuevo usuario |
| `/prestamos` | POST | Crear nuevo préstamo |
| `/prestamos/:id` | PUT | Actualizar préstamo |
| `/reservas` | POST | Crear reserva |
| `/reservas/:id` | DELETE | Eliminar reserva |

**Ejemplo de Respuesta JSON:**
```json
{
  "titulo": "Redes Informáticas",
  "autor": "Autor Ejemplo",
  "categoria": "Tecnología",
  "disponible": true,
  "isbn": "978-3-16-148410-0"
}
```

**Ventajas de la API REST:**
- ✅ Separa frontend y backend
- ✅ Facilita integración futura
- ✅ Mejora organización
- ✅ Centraliza servicios
- ✅ Permite escalabilidad

---

### 3️⃣ CAPA DE CONTROLADORES

**Descripción:** Intermediarios entre la API y la lógica de negocio.

| **Aspecto** | **Descripción** |
|------------|-----------------|
| **Responsabilidad** | Recepción y validación de solicitudes |
| **Función Principal** | Coordinar flujo entre API y servicios |
| **Ubicación** | `/controllers` |

**Funciones Principales:**

| **Función** | **Descripción** |
|------------|-----------------|
| **Recepción de Solicitudes** | Recibir peticiones HTTP de la API |
| **Validación de Datos** | Verificar campos, tipos y formatos |
| **Envío a Servicios** | Pasar datos validados a la capa de servicios |
| **Generación de Respuestas** | Retornar resultados al cliente |

**Validaciones que Realiza:**

| **Validación** | **Ejemplo** |
|---------------|-----------|
| Campos obligatorios | Verificar que email no esté vacío |
| Tipos de datos | Confirmar que edad sea numérica |
| Formatos correctos | Validar formato de correo electrónico |
| Parámetros válidos | Verificar ID de recurso existe |

**Ejemplo en Go:**
```go
func ObtenerLibros(c *gin.Context) {
    libros := services.ListarLibros()
    c.JSON(200, libros)
}

func CrearPrestamo(c *gin.Context) {
    var prestamo models.Prestamo
    if err := c.ShouldBindJSON(&prestamo); err != nil {
        c.JSON(400, gin.H{"error": "Datos inválidos"})
        return
    }
    resultado := services.CrearPrestamo(prestamo)
    c.JSON(201, resultado)
}
```

**Flujo de Ejemplo - Login:**
1. Cliente envía correo y contraseña → 
2. API recibe solicitud → 
3. Controlador valida datos → 
4. Envía al servicio de autenticación → 
5. Espera respuesta → 
6. Devuelve resultado al usuario

---

### 4️⃣ CAPA DE SERVICIOS

**Descripción:** Contiene la lógica de negocio del sistema.

| **Aspecto** | **Descripción** |
|------------|-----------------|
| **Responsabilidad** | Ejecutar reglas de negocio |
| **Importancia** | Parte más crítica del sistema |
| **Ubicación** | `/services` |

**Funciones Principales por Módulo:**

| **Módulo** | **Funcionalidades** |
|-----------|-------------------|
| **Gestión de Usuarios** | • Registrar usuarios<br>• Validar credenciales<br>• Gestionar roles<br>• Controlar permisos |
| **Gestión de Libros** | • Registrar libros<br>• Actualizar información<br>• Verificar disponibilidad<br>• Clasificar categorías |
| **Gestión de Préstamos** | • Procesar solicitudes<br>• Calcular fechas<br>• Verificar disponibilidad<br>• Mantener historiales |
| **Gestión de Reservas** | • Registrar reservas<br>• Notificar disponibilidad<br>• Gestionar lista de espera<br>• Cancelar reservas |
| **Seguridad** | • Encriptar contraseñas<br>• Validar sesiones<br>• Autenticar usuarios<br>• Proteger información |

**Ejemplo Detallado - Solicitud de Préstamo:**

| **Paso** | **Acción** |
|---------|-----------|
| 1 | Verificar que el usuario existe |
| 2 | Comprobar que el libro está disponible |
| 3 | Revisar límite de préstamos del usuario |
| 4 | Registrar la solicitud |
| 5 | Actualizar disponibilidad del libro |
| 6 | Generar respuesta final |

**Ejemplo en Go:**
```go
func ListarLibros() []models.Libro {
    return repositories.ObtenerLibros()
}

func CrearPrestamo(usuarioID int, libroID int) (models.Prestamo, error) {
    if !verificarDisponibilidad(libroID) {
        return models.Prestamo{}, errors.New("libro no disponible")
    }
    
    prestamo := models.Prestamo{
        UsuarioID: usuarioID,
        LibroID: libroID,
        FechaInicio: time.Now(),
        FechaVencimiento: time.Now().AddDate(0, 0, 14),
    }
    
    return repositories.CrearPrestamo(prestamo), nil
}
```

---

### 5️⃣ CAPA DE REPOSITORIOS

**Descripción:** Responsable de comunicarse directamente con la base de datos.

| **Aspecto** | **Descripción** |
|------------|-----------------|
| **Responsabilidad** | Acceso y manipulación de datos |
| **Ubicación** | `/repositories` |
| **Herramienta ORM** | GORM |

**Operaciones Principales:**

| **Operación** | **Descripción** | **Ejemplos** |
|--------------|-----------------|-------------|
| **Consultas (READ)** | Obtener datos de BD | • Buscar usuarios<br>• Consultar libros<br>• Obtener préstamos |
| **Inserción (CREATE)** | Guardar nuevos registros | • Nuevos usuarios<br>• Nuevos libros<br>• Préstamos registrados |
| **Actualización (UPDATE)** | Modificar registros existentes | • Datos de usuarios<br>• Info de libros<br>• Estados de préstamos |
| **Eliminación (DELETE)** | Borrar registros | • Reservas canceladas<br>• Registros innecesarios<br>• Datos obsoletos |

**Ejemplo Práctico - Registrar Préstamo:**

| **Paso** | **Acción** |
|---------|-----------|
| 1 | Servicio envía solicitud |
| 2 | Repositorio genera consulta SQL |
| 3 | Información se guarda en BD |
| 4 | Base de datos confirma operación |

**Ejemplo en Go:**
```go
func ObtenerLibros() []models.Libro {
    var libros []models.Libro
    database.DB.Find(&libros)
    return libros
}

func CrearPrestamo(prestamo models.Prestamo) models.Prestamo {
    database.DB.Create(&prestamo)
    return prestamo
}

func ActualizarLibro(id int, libro models.Libro) models.Libro {
    database.DB.Model(&models.Libro{}).Where("id = ?", id).Updates(libro)
    return libro
}
```

---

### 6️⃣ CAPA DE BASE DE DATOS

**Sistema:** MySQL

**Descripción:** Almacenamiento permanente de toda la información del sistema.

| **Aspecto** | **Descripción** |
|------------|-----------------|
| **DBMS** | MySQL 8.0+ |
| **ORM** | GORM |
| **Responsabilidad** | Persistencia de datos |

**Tablas Principales:**

| **Tabla** | **Información Almacenada** |
|----------|---------------------------|
| **Usuarios** | • IDs<br>• Nombres<br>• Correos<br>• Contraseñas cifradas<br>• Roles |
| **Libros** | • Títulos<br>• Autores<br>• Categorías<br>• Archivos digitales<br>• ISBN |
| **Préstamos** | • Usuario<br>• Libro<br>• Fechas<br>• Estados<br>• Historiales |
| **Reservas** | • Usuario<br>• Libro<br>• Fechas<br>• Disponibilidad<br>• Posición en cola |
| **Auditorías** | • Actividades del sistema<br>• Logs<br>• Estadísticas<br>• Cambios registrados |

**Ejemplo de Conexión en Go:**
```go
dsn := "usuario:password@tcp(localhost:3306)/biblioteca"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

if err != nil {
    panic("Error al conectar a la base de datos")
}
```

**Beneficios de la Base de Datos:**
- ✅ Persistencia de información
- ✅ Seguridad de datos
- ✅ Integridad referencial
- ✅ Disponibilidad 24/7
- ✅ Consultas rápidas y optimizadas

---

## 🔄 Flujo Completo {#flujo-completo}

### Ciclo de una Solicitud Completa:

| **Paso** | **Componente** | **Acción** |
|---------|---------------|-----------|
| 1 | 👤 Usuario | Interactúa con la interfaz web |
| 2 | 🌐 Cliente | Envía solicitud HTTP a la API REST |
| 3 | 🔀 Gin Framework | Recibe la petición HTTP |
| 4 | 🎮 Controlador | Procesa y valida los datos |
| 5 | ⚙️ Servicio | Ejecuta la lógica de negocio |
| 6 | 📦 Repositorio | Consulta la base de datos MySQL |
| 7 | 💾 Base de Datos | Devuelve los resultados solicitados |
| 8 | 📄 Servicio | Genera respuesta JSON con los datos |
| 9 | ↩️ Cliente | Recibe respuesta y actualiza interfaz |

### Ejemplo Práctico - Búsqueda de Libro:

```
Usuario: Escribe "Go Programming" en la barra de búsqueda
   ↓
Cliente: Genera solicitud GET /libros?titulo=Go%20Programming
   ↓
API/Gin: Enruta la solicitud al controlador
   ↓
Controlador: Valida el parámetro de búsqueda
   ↓
Servicio: Ejecuta lógica de búsqueda
   ↓
Repositorio: Ejecuta consulta: SELECT * FROM libros WHERE titulo LIKE '%Go%'
   ↓
BD MySQL: Retorna coincidencias encontradas
   ↓
Respuesta: JSON con libros encontrados
   ↓
Cliente: Muestra resultados en pantalla
```

---

## ✅ Conclusión

La arquitectura basada en **Cliente → API REST → Controladores → Servicios → Repositorios → Base de Datos** permite:

| **Beneficio** | **Descripción** |
|--------------|-----------------|
| **Organización** | Código estructurado y fácil de navegar |
| **Seguridad** | Validación en múltiples capas |
| **Escalabilidad** | Capacidad de crecer sin rediseño |
| **Mantenibilidad** | Cambios aislados por capas |
| **Reutilización** | Componentes reutilizables |
| **Testabilidad** | Fácil de probar unitariamente |

La separación por capas facilita la administración del código, mejora el mantenimiento y permite que el sistema pueda crecer en el futuro incorporando nuevas funcionalidades sin afectar su estructura actual.

---

**Última actualización:** Mayo 2026  
**Desarrollador:** Leonardo Sánchez  
**Rama:** Leonardo_Sanchez
