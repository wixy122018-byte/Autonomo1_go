# Sistema de Gestión de Libros Electrónicos en Go

## Descripción General del Sistema

El proyecto consiste en desarrollar una aplicación backend en Golang orientada a la administración de libros electrónicos dentro de una biblioteca virtual o plataforma digital de lectura.

El sistema permitirá gestionar información relacionada con libros virtuales, usuarios, préstamos digitales, descargas y reportes estadísticos mediante una API REST desarrollada en Go.

La plataforma automatizará procesos importantes como:

- Registro de libros electrónicos
- Organización por categorías
- Búsquedas dinámicas
- Gestión de usuarios
- Préstamos digitales
- Descargas de archivos electrónicos
- Generación de reportes
- Seguridad y autenticación

Además, el sistema utilizará una arquitectura modular y aplicará programación funcional mediante funciones puras, closures y funciones de orden superior.

---

## Objetivo General

Desarrollar un Sistema de Gestión de Libros Electrónicos en Golang que permita administrar libros virtuales, usuarios, préstamos y descargas digitales mediante una arquitectura modular, APIs REST y principios de programación funcional.

---

## Objetivos Específicos

- Diseñar una estructura modular organizada para el sistema.
- Implementar funcionalidades para la gestión de libros electrónicos.
- Gestionar usuarios y autenticación mediante JWT.
- Aplicar programación funcional en Golang.
- Implementar concurrencia mediante goroutines y channels.
- Gestionar persistencia de datos utilizando PostgreSQL y GORM.
- Crear APIs REST para la comunicación cliente-servidor.

---

## Gestión del Sistema

### ¿En qué consiste la gestión del sistema?

La gestión del sistema consiste en administrar de manera digital toda la información relacionada con los libros electrónicos y los usuarios que interactúan dentro de la plataforma virtual. El propósito principal es automatizar procesos que normalmente se realizan manualmente en una biblioteca tradicional, permitiendo un acceso más rápido, organizado y seguro a los recursos digitales.

A través del sistema, los usuarios podrán consultar libros virtuales, solicitar préstamos digitales, descargar archivos electrónicos y acceder a diferentes funcionalidades según sus permisos dentro de la plataforma. De igual manera, los administradores podrán mantener control sobre el catálogo de libros, usuarios registrados, estadísticas y actividades realizadas dentro del sistema.

La gestión del sistema no solamente se enfoca en almacenar información, sino también en optimizar la organización y administración de los libros electrónicos, facilitando el acceso al conocimiento mediante una plataforma moderna desarrollada en Golang.

Además, este sistema busca aprovechar las ventajas de las bibliotecas digitales, ya que los usuarios podrán acceder a los libros desde cualquier lugar y en cualquier momento, eliminando muchas de las limitaciones presentes en las bibliotecas físicas tradicionales.

---

### Procesos principales de la gestión del sistema

#### Gestión de libros electrónicos

La gestión de libros representa el núcleo principal del sistema, ya que permite administrar todo el catálogo digital disponible dentro de la biblioteca virtual.

Este módulo permitirá registrar nuevos libros electrónicos incluyendo información como:

- título
- autor
- categoría
- formato
- disponibilidad
- descripción del libro

También permitirá modificar información existente, eliminar registros y mantener actualizado el catálogo digital de la plataforma.

Además, los usuarios podrán realizar búsquedas dinámicas por:

- nombre del libro
- autor
- categoría
- disponibilidad

Esto facilitará el acceso rápido a los recursos digitales y mejorará la experiencia de navegación dentro de la biblioteca virtual.

---

#### Gestión de usuarios

La gestión de usuarios consiste en administrar toda la información relacionada con las personas registradas dentro de la plataforma.

Este proceso permitirá:

- registrar usuarios
- iniciar sesión
- editar perfiles
- recuperar información
- controlar permisos de acceso

El sistema podrá manejar diferentes tipos de usuarios como:

- administradores
- usuarios normales
- lectores registrados

Cada usuario tendrá permisos específicos según su rol dentro del sistema.

Además, la autenticación mediante JWT permitirá proteger el acceso a la plataforma y garantizar que únicamente usuarios autorizados puedan acceder a determinadas funcionalidades.

---

#### Gestión de préstamos digitales

La gestión de préstamos digitales permitirá controlar la asignación temporal de libros electrónicos a los usuarios registrados.

El sistema permitirá:

- solicitar préstamos
- validar disponibilidad
- registrar devoluciones
- consultar historial de préstamos
- controlar tiempos de acceso

Cada vez que un usuario solicite un préstamo, el sistema verificará si el libro se encuentra disponible antes de autorizar el acceso.

Posteriormente, el préstamo quedará registrado dentro de la base de datos para mantener un historial organizado de actividades.

---

#### Gestión de descargas

La gestión de descargas permitirá a los usuarios autorizados descargar libros electrónicos directamente desde la plataforma.

El sistema validará:

- permisos de acceso
- autenticación del usuario
- disponibilidad del archivo

Además, se podrá registrar un historial de descargas para llevar control sobre:

- libros descargados
- frecuencia de uso
- actividad de los usuarios

Este módulo resulta importante porque facilita el acceso rápido a los contenidos digitales y mejora la experiencia de los lectores dentro de la plataforma virtual.

---

#### Gestión de reportes

La gestión de reportes permitirá generar información estadística relacionada con el funcionamiento general de la plataforma.

Este módulo permitirá obtener datos como:

- libros más descargados
- usuarios más activos
- categorías más consultadas
- historial de préstamos
- estadísticas generales del sistema

Los reportes ayudarán a los administradores a analizar el comportamiento de la plataforma y tomar decisiones relacionadas con la administración de la biblioteca virtual.

---

### Importancia de la gestión del sistema

La gestión del sistema resulta importante porque permite modernizar la administración de bibliotecas mediante herramientas digitales y tecnologías actuales.

El sistema facilita:

- organización de información
- automatización de procesos
- acceso rápido a libros virtuales
- control de usuarios
- seguridad
- administración eficiente de recursos digitales

Además, al desarrollarse en Golang, el proyecto permitirá implementar una arquitectura moderna, modular y escalable, aplicando conceptos fundamentales estudiados durante la materia como:

- structs
- slices
- maps
- programación funcional
- concurrencia
- APIs REST
- manejo de JSON
- persistencia de datos

## Relación con temas estudiados en Go
### Para el desarrollo de la herramienta de gestión, acudimos a distintos fundamentos claves de GO impartidos en clases tales como:

| Tema en GO | <center>Función dentro del proyecto</center> |
| :--- | :--- |
| Structs | Los utilizaremos para el modelado de las entidades principales del sistema |
| Slices | Nos permitirán gestionar listas de datos dinámicas en memoria  |
| Maps | Los implementaremos para optimizar una indexación rápida y ágil  |
| Functions | Para el diseño de funciones modulares |
| Closures | Para validaciones dinámicas y filtros personalizados |
| Interfaces | Lo aplicaremos para definir contratos claros de distintas capas y repositorios |
| Goroutines y Channels | Nos permitirán el desarrollo de procesos y comunicación simultáneos en segundo plano |
| JSON Y HTTP | Funcionarán como base de nuestra API de comunicación (cliente - servidor) |
| Packages | Permiten la organización del código en paquetes estructurados |
| File Handling | Permitirán la administración, almacenamiento y lectura de archivos para el usuario |

## Arquitectura del Sistema de Gestión de Libros Electrónicos

### Tabla de Contenidos
1. [Introducción](#introducción)
2. [Flujo General de Arquitectura](#flujo-general)
3. [Estructura del Proyecto](#estructura)
4. [Usuarios y Roles](#usuarios-roles)
5. [Capas de la Arquitectura](#capas)
6. [Flujo Completo](#flujo-completo)

---

### Introducción

El Sistema de Gestión de Libros Electrónicos será desarrollado en **Go**, implementando una arquitectura por capas que permite mantener el sistema organizado, escalable y fácil de mantener.

**Beneficios de esta arquitectura:**
- Separación de responsabilidades
- Código modular y reutilizable
- Fácil mantenimiento
- Mejor seguridad
- Escalabilidad

---

### Flujo General de Arquitectura

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
│ Base de Datos   │ (PostgreSQL)
└─────────────────┘
```

---

### Estructura del Proyecto

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

### Usuarios y Roles

#### Tabla de Roles y Permisos

| **Rol** | **Permisos Clave** | **Funcionalidades Principales** |
|---------|-------------------|--------------------------------|
| **Administrador** | Control total | • Gestiona usuarios<br>• Administra libros<br>• Visualiza reportes<br>• Configura parámetros del sistema |
| **Bibliotecario** | Gestión de contenido | • Registra libros<br>• Gestiona préstamos<br>• Controla reservas<br>• Actualiza información |
| **Lector/Estudiante** | Acceso limitado | • Busca libros<br>• Solicita préstamos<br>• Reserva libros<br>• Consulta historial |

---

### Capas de la Arquitectura

#### 1 CAPA CLIENTE

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

#### 2 CAPA API REST

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
- Separa frontend y backend
- Facilita integración futura
- Mejora organización
- Centraliza servicios
- Permite escalabilidad

---

#### 3 CAPA DE CONTROLADORES

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

#### 4 CAPA DE SERVICIOS

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

#### 5 CAPA DE REPOSITORIOS

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

#### 6 CAPA DE BASE DE DATOS

**Sistema:** PostgreSQL

**Descripción:** Almacenamiento permanente de toda la información del sistema.

| **Aspecto** | **Descripción** |
|------------|-----------------|
| **DBMS** | PostgreSQL |
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
dsn := "host=localhost user=usuario password=password dbname=biblioteca port=5432 sslmode=disable"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

if err != nil {
    panic("Error al conectar a la base de datos")
}
```

**Beneficios de la Base de Datos:**
- Persistencia de información
- Seguridad de datos
- Integridad referencial
- Disponibilidad 24/7
- Consultas rápidas y optimizadas

---

### Flujo Completo

#### Ciclo de una Solicitud Completa:

| **Paso** | **Componente** | **Acción** |
|---------|---------------|-----------|
| 1 |  Usuario | Interactúa con la interfaz web |
| 2 |  Cliente | Envía solicitud HTTP a la API REST |
| 3 |  Gin Framework | Recibe la petición HTTP |
| 4 |  Controlador | Procesa y valida los datos |
| 5 |  Servicio | Ejecuta la lógica de negocio |
| 6 |  Repositorio | Consulta la base de datos PostgreSQL |
| 7 |  Base de Datos | Devuelve los resultados solicitados |
| 8 |  Servicio | Genera respuesta JSON con los datos |
| 9 | ↩ Cliente | Recibe respuesta y actualiza interfaz |

#### Ejemplo Práctico - Búsqueda de Libro:

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
BD PostgreSQL: Retorna coincidencias encontradas
   ↓
Respuesta: JSON con libros encontrados
   ↓
Cliente: Muestra resultados en pantalla
```

---

## Módulos del Sistema de Gestión de Libros Electrónicos en Go

### Introducción

El Sistema de Gestión de Libros Electrónicos será desarrollado utilizando **Go** bajo una arquitectura modular. Esta estructura permite dividir el sistema en módulos independientes que facilitan el desarrollo colaborativo, mantenimiento y escalabilidad.

---

### Tabla de Contenidos

1. [Módulo de Gestión de Usuarios](#1-módulo-de-gestión-de-usuarios)
2. [Módulo de Gestión de Libros](#2-módulo-de-gestión-de-libros)
3. [Módulo de Búsqueda de Libros](#3-módulo-de-búsqueda-de-libros)
4. [Módulo de Préstamos Electrónicos](#4-módulo-de-préstamos-electrónicos)
5. [Módulo de Reservas](#5-módulo-de-reservas)
6. [Módulo de Reportes y Estadísticas](#6-módulo-de-reportes-y-estadísticas)
7. [Módulo de Seguridad](#7-módulo-de-seguridad)

---

### 1. MÓDULO DE GESTIÓN DE USUARIOS

#### Descripción General
Administra toda la información relacionada con los usuarios del sistema. Es uno de los módulos más importantes porque controla el acceso, autenticación y permisos dentro de la plataforma.

#### Objetivos del Módulo
- Registrar nuevos usuarios
- Permitir inicio de sesión
- Gestionar perfiles de usuario
- Recuperar contraseñas perdidas
- Controlar acceso según permisos asignados

#### Arquitectura Interna

```
┌─────────────────────────┐
│  Usuario Controller     │ (Recibe solicitudes HTTP)
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│   Usuario Service       │ (Lógica de negocio)
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│  Usuario Repository     │ (Acceso a datos)
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│       PostgreSQL             │ (Base de datos)
└─────────────────────────┘
```

#### Componentes Internos

##### Usuario Controller
**Responsabilidades:**
- Recibir solicitudes HTTP
- Validar datos básicos
- Enviar respuestas JSON

**Operaciones principales:**
- Registro de usuario
- Login/Autenticación
- Actualización de perfil
- Obtener información de usuario

##### Usuario Service
**Responsabilidades:**
Contiene toda la lógica del negocio:
- Verificar si el correo ya existe
- Encriptar contraseñas
- Validar credenciales
- Generar tokens JWT
- Validar permisos de usuario

##### Usuario Repository
**Responsabilidades:**
- Guardar nuevos usuarios
- Consultar información de usuarios
- Actualizar registros existentes
- Eliminar usuarios
- Buscar usuarios por criterios

#### Procesos Principales

##### 1. Registro de Usuarios
```
1. Controlador recibe datos de registro
   ↓
2. Servicio valida información
   - Email válido
   - Contraseña fuerte
   - Datos completos
   ↓
3. Cifra la contraseña con bcrypt
   ↓
4. Repositorio guarda datos en PostgreSQL
   ↓
5. Retorna token JWT al cliente
```

##### 2. Inicio de Sesión
```
1. Búsqueda del usuario por email
   ↓
2. Verificación de contraseña
   ↓
3. Validación de estado (activo/inactivo)
   ↓
4. Generación de token de autenticación
   ↓
5. Devolución de acceso al cliente
```

#### Seguridad Implementada

**Librerías utilizadas:**
- **bcrypt** → Cifrado y hash de contraseñas
- **JWT (JSON Web Tokens)** → Autenticación y autorización

**Medidas de seguridad:**
- Contraseñas hasheadas (nunca en texto plano)
- Tokens JWT con expiración
- Validación de datos entrada
- Control de acceso por roles
- Auditoría de accesos

#### Importancia del Módulo
Controla aspectos críticos del sistema:
- Seguridad general
- Autenticación de usuarios
- Asignación de roles y permisos
- Protección de información personal

---

### 2. MÓDULO DE GESTIÓN DE LIBROS

#### Descripción General
Administra todo el catálogo digital de libros electrónicos. Es considerado el núcleo principal de la biblioteca virtual.

#### Objetivos del Módulo
- Registrar nuevos libros
- Actualizar información de libros
- Organizar por categorías
- Gestionar archivos digitales (PDF, EPUB)
- Controlar disponibilidad

#### Arquitectura Interna

```
┌─────────────────────────┐
│   Libro Controller      │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│    Libro Service        │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│   Libro Repository      │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│       PostgreSQL             │
└─────────────────────────┘
```

#### Información Gestionada por Libro

Cada registro de libro almacenará:
- **Identificador único** (ISBN)
- **Título** del libro
- **Autor/Autores**
- **Categoría/Subcategoría**
- **Editorial**
- **Año de publicación**
- **Archivo digital** (PDF o EPUB)
- **Estado de disponibilidad**
- **Número de copias disponibles**
- **Descripción/Resumen**
- **Fecha de registro**

#### Procesos Principales

##### 1. Registro de Libros
```
1. Controlador recibe información del libro
   ↓
2. Servicio valida datos
   - Campos obligatorios completos
   - Formato de archivo válido
   - ISBN único
   ↓
3. Verifica y procesa archivo digital
   ↓
4. Repositorio guarda información en PostgreSQL
   ↓
5. Archivo se almacena en servidor
   ↓
6. Retorna confirmación al bibliotecario
```

##### 2. Actualización de Libros
El sistema permite:
- Editar títulos y autores
- Cambiar categorías
- Reemplazar archivos digitales
- Modificar estado de disponibilidad
- Actualizar número de copias

##### 3. Eliminación de Libros
- Eliminar registros permanentemente
- Desactivar libros (mantener histórico)
- Limpiar información obsoleta
- Mantener referencia en préstamos históricos

#### Tecnologías Utilizadas en Go

- **Gin** → Framework para API REST
- **GORM** → ORM para conexión con PostgreSQL
- **FileSystem de Go** → Manejo y almacenamiento de archivos
- **Multer/Form parsing** → Procesamiento de uploads

#### Importancia del Módulo

Permite:
- Organizar biblioteca digital de forma eficiente
- Gestionar contenido digital
- Facilitar búsquedas rápidas
- Administrar archivos electrónicos
- Mantener integridad del catálogo

---

### 3. MÓDULO DE BÚSQUEDA DE LIBROS

#### Descripción General
Facilita búsquedas rápidas y eficientes dentro del sistema mediante filtros y consultas avanzadas.

#### Objetivos del Módulo
- Localizar libros por múltiples criterios
- Aplicar filtros avanzados
- Retornar resultados relevantes
- Optimizar tiempo de búsqueda

#### Arquitectura Interna

```
┌─────────────────────────┐
│  Busqueda Controller    │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│   Busqueda Service      │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│  Busqueda Repository    │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│       PostgreSQL             │
└─────────────────────────┘
```

#### Funciones Principales

##### 1. Búsqueda por Título
- Permite buscar libros mediante palabras clave
- Búsqueda parcial (LIKE)
- Orden por relevancia

##### 2. Búsqueda por Autor
- Localiza libros relacionados con autores específicos
- Soporte para múltiples autores
- Búsqueda flexible

##### 3. Búsqueda por Categorías
Ejemplos de categorías:
- Programación
- Redes y Telecomunicaciones
- Matemáticas
- Literatura
- Ciencias
- Historia

##### 4. Filtros Avanzados
Permite filtrar por:
- **Año de publicación** (rango)
- **Estado de disponibilidad** (disponible/prestado)
- **Popularidad** (más consultados)
- **Categoría/Subcategoría**
- **Editorial**
- **Tipo de archivo** (PDF, EPUB)

#### Procesamiento en Go

```
1. Controlador recibe parámetros de búsqueda
   ↓
2. Servicio valida y sanitiza criterios
   ↓
3. Repositorio construye consultas SQL dinámicas
   ↓
4. Ejecuta queries optimizadas
   ↓
5. Filtra y ordena resultados
   ↓
6. Retorna JSON con resultados paginados
```

#### Ventajas del Módulo

- Mejora significativa de la experiencia del usuario
- Reduce tiempo de búsqueda
- Optimiza navegación en el catálogo
- Facilita descubrimiento de libros
- Paginación eficiente de resultados

---

### 4. MÓDULO DE PRÉSTAMOS ELECTRÓNICOS

#### Descripción General
Controla todo el proceso de préstamos digitales. Es uno de los procesos más importantes del sistema.

#### Objetivos del Módulo
- Administrar solicitudes de préstamo
- Verificar disponibilidad
- Gestionar devoluciones
- Mantener historiales completos
- Controlar límites por usuario

#### Arquitectura Interna

```
┌─────────────────────────┐
│  Prestamo Controller    │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│   Prestamo Service      │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│  Prestamo Repository    │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│       PostgreSQL             │
└─────────────────────────┘
```

#### Flujo de Préstamo en Go

##### Cuando un usuario solicita un libro:

```
PASO 1: Recepción de Solicitud
└─→ Controlador recibe solicitud de préstamo
   └─→ Extrae ID usuario e ID libro

PASO 2: Validación
└─→ Servicio verifica:
   ├─→ Existencia del usuario (activo)
   ├─→ Disponibilidad del libro
   ├─→ Límite de préstamos por usuario
   ├─→ Historial de usuario (morosos)
   └─→ Formato de solicitud

PASO 3: Registro
└─→ Repositorio registra información en PostgreSQL
   ├─→ Guarda registro en tabla PRESTAMOS
   ├─→ Registra fecha de inicio
   ├─→ Establece fecha de vencimiento

PASO 4: Actualización
└─→ Sistema actualiza disponibilidad del libro
   ├─→ Reduce cantidad disponible
   ├─→ Marca libro como parcialmente prestado
   └─→ Retorna confirmación al usuario
```

#### Funciones Principales

##### 1. Solicitud de Préstamos
- Permite solicitar libros digitales
- Validación automática de criterios
- Generación de comprobante digital
- Notificación al usuario

##### 2. Gestión de Devoluciones
- Procesar devoluciones de libros
- Calcular multas si es necesario
- Actualizar disponibilidad
- Generar comprobante de devolución

##### 3. Historial de Préstamos
- Guarda todos los préstamos realizados
- Mantiene registro de devoluciones
- Información sobre multas generadas
- Análisis de hábitos de lectura

##### 4. Control de Restricciones
- Límite de libros por usuario (ej: máximo 5)
- Duración máxima de préstamo (ej: 30 días)
- Penalizaciones por retraso
- Bloqueo temporal por incumplimiento

#### Datos del Préstamo

```
Tabla: PRESTAMOS
├─ id_prestamo (PK)
├─ id_usuario (FK)
├─ id_libro (FK)
├─ fecha_inicio
├─ fecha_vencimiento
├─ fecha_devolucion
├─ estado (activo/completado/vencido)
├─ multa_generada
└─ observaciones
```

#### Importancia del Módulo

Permite:
- Automatizar procesos bibliotecarios
- Mantener control sobre circulación digital
- Generar estadísticas de uso
- Prevenir conflictos entre usuarios
- Facilitar auditoría de préstamos

---

### 5. MÓDULO DE RESERVAS

#### Descripción General
Administra reservas de libros temporalmente no disponibles.

#### Objetivos del Módulo
- Permitir reservas de libros ocupados
- Mantener lista de espera
- Notificar disponibilidad
- Gestionar cancelaciones

#### Arquitectura Interna

```
┌─────────────────────────┐
│  Reserva Controller     │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│   Reserva Service       │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│  Reserva Repository     │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│       PostgreSQL             │
└─────────────────────────┘
```

#### Proceso de Reserva

```
1. Usuario solicita reserva de libro no disponible
   ↓
2. Sistema verifica:
   - Disponibilidad actual del libro
   - Posición en lista de espera
   - Máximo de reservas por usuario
   ↓
3. Registra la reserva
   - Asigna posición en cola
   - Establece fecha de solicitud
   ↓
4. Guarda orden de espera
   - Ordena automáticamente
   - Mantiene histórico
   ↓
5. Notifica disponibilidad futura
   - Email al usuario
   - Notificación en el sistema
```

#### Funciones Principales

##### 1. Registro de Reservas
- Guardar solicitudes de reserva
- Validar datos de entrada
- Asignar prioridad automática

##### 2. Lista de Espera
- Organizar usuarios automáticamente
- Respetar orden FIFO (First In, First Out)
- Mantener histórico de esperas

##### 3. Notificaciones
- Informar cuando libro está disponible
- Enviar recordatorios
- Confirmar recepción de notificación

##### 4. Cancelaciones
- Permitir eliminar reservas
- Liberar posición en lista
- Notificar cancelación al usuario

#### Datos de Reserva

```
Tabla: RESERVAS
├─ id_reserva (PK)
├─ id_usuario (FK)
├─ id_libro (FK)
├─ posicion_cola
├─ fecha_solicitud
├─ estado (pendiente/disponible/cancelada)
├─ fecha_disponibilidad
└─ fecha_cancelacion
```

#### Importancia del Módulo

- Evita conflictos entre usuarios
- Mejora administración de recursos digitales
- Garantiza acceso equitativo
- Reduce pérdida de demanda

---

### 6. MÓDULO DE REPORTES Y ESTADÍSTICAS

#### Descripción General
Genera información administrativa y estadística del sistema.

#### Objetivos del Módulo
- Permitir análisis de funcionamiento
- Generar reportes administrativos
- Crear estadísticas de uso
- Facilitar toma de decisiones

#### Arquitectura Interna

```
┌─────────────────────────┐
│  Reporte Controller     │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│   Reporte Service       │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│  Reporte Repository     │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│       PostgreSQL             │
└─────────────────────────┘
```

#### Funciones Principales

##### 1. Reportes de Préstamos
- Muestra historial de préstamos realizados
- Análisis de tendencias
- Identificación de períodos de alta demanda
- Libros más prestados

##### 2. Reportes de Usuarios
- Identifica usuarios más activos
- Análisis de perfiles de lectura
- Detecta patrones de uso
- Usuarios con retrasos o multas

##### 3. Reportes de Libros
- Muestra libros más consultados
- Libros menos utilizados
- Análisis de categorías populares
- Disponibilidad vs. demanda

##### 4. Estadísticas Generales
- Métricas generales del sistema
- Número total de usuarios
- Cantidad de libros en catálogo
- Préstamos activos
- Reservas pendientes
- Tasa de ocupación

#### Exportación de Información

El sistema puede generar reportes en formatos:
- **PDF** → Para impresión y distribución formal
- **Excel** → Para análisis adicional
- **CSV** → Para integración con otras herramientas

#### Procesamiento en Go

```
1. Controlador recibe solicitud de reporte
   ↓
2. Servicio procesa parámetros (rango fechas, filtros)
   ↓
3. Repositorio ejecuta queries complejas
   ↓
4. Agrega datos y calcula estadísticas
   ↓
5. Genera archivo en formato solicitado
   ↓
6. Retorna archivo comprimido al usuario
```

**Ventajas de usar Go:**
- Generación rápida de archivos
- Procesa grandes volúmenes de datos eficientemente
- Optimiza consultas complejas
- Manejo concurrente de múltiples reportes

#### Importancia del Módulo

Facilita:
- Toma de decisiones administrativas
- Optimización del catálogo
- Mejora de servicios
- Supervisión del sistema
- Evaluación de desempeño

---

### 7. MÓDULO DE SEGURIDAD

#### Descripción General
Protege toda la información del sistema. Fundamental debido a que maneja datos críticos.

#### Datos Sensibles Protegidos

- Información de usuarios
- Contraseñas
- Archivos digitales
- Información académica
- Registros de préstamos

#### Objetivos del Módulo

Garantizar:
- **Confidencialidad** → Solo usuarios autorizados acceden
- **Integridad** → Datos no alterados
- **Disponibilidad** → Servicio siempre accesible
- **Autenticidad** → Verificar identidad de usuarios

#### Arquitectura Interna

```
┌──────────────────────────────────┐
│   Middleware de Seguridad        │
│  - Validación de tokens JWT      │
│  - Protección CORS               │
│  - Rate limiting                 │
└──────────────────┬───────────────┘
                   │
┌──────────────────▼───────────────┐
│  Servicios de Autenticación      │
│  - Manejo de JWT                 │
│  - Encriptación bcrypt           │
│  - Validación de permisos        │
└──────────────────┬───────────────┘
                   │
┌──────────────────▼───────────────┐
│   Control de Acceso              │
│  - Roles y permisos              │
│  - Auditoría de accesos          │
│  - Registro de eventos críticos  │
└──────────────────────────────────┘
```

#### Funciones Principales

##### 1. Autenticación JWT
Go utiliza JWT para:
- **Validar sesiones** → Verificar token activo
- **Generar tokens** → Crear token al login
- **Proteger rutas** → Requerir token válido
- **Expiración automática** → Token vence después de tiempo

```
Estructura del Token JWT:
{
  "sub": "id_usuario",
  "email": "usuario@ejemplo.com",
  "rol": "usuario",
  "iat": 1234567890,
  "exp": 1234571490
}
```

##### 2. Cifrado de Contraseñas
- Utiliza **bcrypt** para hash seguro
- Nunca almacena contraseña en texto plano
- Salt automático incluido
- Algoritmo resistente a ataques

```
Flujo de contraseña:
1. Usuario ingresa: "MiContraseña123"
   ↓
2. bcrypt genera hash: $2a$12$R9h7cIPz0gi...
   ↓
3. Se almacena hash en BD
   ↓
4. En login: bcrypt compara contra hash
   ↓
5. Valida o rechaza acceso
```

##### 3. Middleware de Autenticación
Permite:
- Proteger endpoints específicos
- Validar permisos por rol
- Bloquear accesos indebidos
- Registrar intentos fallidos

##### 4. Auditoría y Logging
El sistema registra:
- **Inicios de sesión** → Quién, cuándo, desde dónde
- **Errores de autenticación** → Intentos fallidos
- **Actividades críticas** → Cambios importantes
- **Acceso a datos sensibles** → Quién consultó qué

#### Medidas de Seguridad Adicionales

- HTTPS/TLS para comunicación cifrada
- CORS configurado
- Rate limiting para prevenir fuerza bruta
- Validación de entrada en todos los endpoints
- Sanitización de queries SQL
- Passwords policy (complejidad mínima)
- Bloqueo temporal tras intentos fallidos

#### Importancia del Módulo

Protege:
- Información personal de usuarios
- Integridad del sistema
- Archivos digitales
- Cumplimiento normativo
- Confianza de usuarios

---

## Programación Funcional en Go

### Descripción general

Go es un lenguaje principalmente imperativo, estructurado y concurrente. Aunque no es un lenguaje funcional puro, permite aplicar ciertos principios de programación funcional mediante el uso de funciones como valores, closures, funciones puras y composición de comportamientos.

Dentro del Sistema de Gestión de Libros Electrónicos, la programación funcional se plantea como un complemento para mejorar la organización de la lógica interna del sistema. Su aplicación se enfoca principalmente en el procesamiento de datos, validaciones, filtros, transformación de información y generación de reportes.

Este enfoque no reemplaza la arquitectura principal del proyecto, sino que ayuda a que ciertas operaciones sean más limpias, reutilizables y fáciles de mantener.

### Aplicación dentro del proyecto

La programación funcional se aplicará principalmente en la capa de servicios y en funciones auxiliares del sistema. Estas áreas concentran la lógica de negocio y permiten trabajar con datos como usuarios, libros, préstamos, descargas, reservas y reportes.

Su uso será útil en procesos donde se necesite aplicar reglas específicas sobre colecciones de datos, transformar información antes de enviarla como respuesta o validar condiciones del sistema de manera ordenada.

Entre los principales usos dentro del proyecto se consideran:

- Filtrado de libros por disponibilidad, categoría, autor o popularidad.
- Validación de datos de usuarios, libros y préstamos.
- Transformación de modelos internos en respuestas para la API.
- Procesamiento de listas de libros, usuarios o préstamos.
- Cálculo de estadísticas para reportes.
- Aplicación de reglas dinámicas según el tipo de usuario.
- Separación de la lógica general de procesamiento y las condiciones específicas.

### Funciones de orden superior

Las funciones de orden superior permiten recibir otras funciones como parámetros o devolver funciones como resultado. En el proyecto, este principio puede utilizarse para construir procesos reutilizables, especialmente en filtros, validaciones y reportes.

Este enfoque evita repetir la misma estructura lógica en diferentes módulos. Por ejemplo, en lugar de crear una función distinta para cada tipo de filtro, se puede tener una función general que reciba la condición que debe aplicarse.

Dentro del sistema, esto resulta útil para módulos como búsqueda de libros, gestión de préstamos y generación de reportes.

### Transformación de datos

La transformación de datos permite convertir información interna del sistema en estructuras más adecuadas para su uso o presentación. En el proyecto, este principio puede aplicarse cuando los datos obtenidos desde la base de datos necesitan ser adaptados antes de enviarse al cliente mediante la API REST.

Esto es importante porque no siempre se debe exponer toda la información interna de los modelos. La transformación permite seleccionar únicamente los datos necesarios, mejorar la presentación de las respuestas y mantener una separación clara entre la lógica interna y la información mostrada al usuario.

Este enfoque puede aplicarse especialmente en el módulo de libros, usuarios, préstamos y reportes.

### Uso de closures

Los closures permiten crear funciones que conservan información del contexto donde fueron definidas. En el proyecto, pueden utilizarse para generar reglas dinámicas de búsqueda, validación o configuración sin depender de variables globales.

Este recurso es útil cuando el sistema necesita aplicar condiciones que cambian según los parámetros recibidos. Por ejemplo, filtros por categoría, disponibilidad, año, autor o rol del usuario.

El uso de closures permite que el sistema sea más flexible y que las reglas puedan adaptarse sin modificar de forma extensa la lógica principal.

### Funciones puras

Las funciones puras son aquellas que producen siempre el mismo resultado cuando reciben los mismos datos y no modifican elementos externos. En el proyecto, este tipo de funciones es importante para mantener una lógica más predecible y fácil de probar.

Pueden aplicarse en cálculos, validaciones y reglas de negocio que no dependan directamente de la base de datos ni de estados externos.

Dentro del sistema, las funciones puras pueden utilizarse para:

- Validar si un libro está disponible.
- Calcular días de préstamo.
- Verificar condiciones de usuario.
- Calcular estadísticas básicas.
- Determinar estados de préstamos.
- Evaluar reglas simples del sistema.

Su principal ventaja es que facilitan las pruebas unitarias y reducen errores causados por modificaciones inesperadas de datos externos.

### Relación con los módulos del sistema

La programación funcional puede aportar valor en varios módulos del Sistema de Gestión de Libros Electrónicos.

En el módulo de libros, puede utilizarse para filtrar, ordenar y transformar información antes de presentarla al usuario.

En el módulo de usuarios, puede apoyar la validación de datos, roles y permisos.

En el módulo de préstamos, puede servir para evaluar estados, calcular condiciones y procesar historiales.

En el módulo de reportes, puede facilitar el conteo, clasificación y análisis de datos relacionados con préstamos, descargas y actividad de usuarios.

En el módulo de seguridad, puede apoyar validaciones específicas sin mezclar reglas de negocio con lógica de controladores.

### Ventajas para el proyecto

La aplicación de programación funcional en Go aporta varias ventajas al desarrollo del sistema.

Permite reducir la duplicación de código, ya que una misma estructura de procesamiento puede reutilizarse con diferentes condiciones.

Mejora la organización interna del proyecto, separando la lógica general de las reglas específicas.

Facilita las pruebas unitarias, especialmente cuando se emplean funciones puras.

Ayuda a mantener el código más predecible, ya que reduce los efectos secundarios y las modificaciones innecesarias de datos externos.

Favorece la escalabilidad del sistema, porque permite agregar nuevas reglas, filtros o validaciones sin reestructurar completamente el código existente.

### Limitaciones

Aunque Go permite aplicar principios funcionales, no debe tratarse como un lenguaje funcional puro. Forzar demasiada abstracción puede hacer que el código sea más difícil de leer y mantener.

Por esta razón, la programación funcional debe utilizarse únicamente donde aporte claridad y reutilización. En el proyecto, su uso debe mantenerse principalmente en servicios, validaciones, filtros, transformaciones y reportes.

Los controladores deben conservar una estructura simple, enfocada en recibir solicitudes y devolver respuestas. Los repositorios deben mantenerse orientados al acceso a datos. La lógica funcional debe ubicarse donde realmente mejore la organización del sistema.

## Estructura de la Gestión del Sistema

#### Organización de paquetes del sistema

El Sistema de Gestión de Libros Electrónicos utilizará una arquitectura modular organizada mediante paquetes en Golang con el objetivo de mantener el código limpio, escalable y fácil de mantener.

La estructura del proyecto estará dividida en diferentes capas, donde cada paquete tendrá una responsabilidad específica dentro del sistema.

La organización modular permitirá separar correctamente:

- lógica de negocio
- acceso a datos
- rutas
- autenticación
- configuración
- modelos
- servicios
- controladores

Esto facilitará el trabajo colaborativo y permitirá realizar mantenimiento o ampliaciones futuras de manera más eficiente.

---

### Estructura general del proyecto

```text
sistema-libros-electronicos/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── config/
│   ├── models/
│   ├── handlers/
│   ├── services/
│   ├── repositories/
│   ├── middleware/
│   ├── routes/
│   └── database/
│
├── pkg/
│
├── docs/
│
├── .env
├── go.mod
└── go.sum
```

### Explicación de Paquetes

#### cmd/

Este paquete contiene el punto de entrada principal del sistema mediante el archivo `main.go`.

Su función es:

- iniciar el servidor
- cargar configuraciones
- conectar la base de datos
- ejecutar las rutas principales del sistema

---

#### internal/config/

Este paquete administra toda la configuración del sistema.

Permite:

- cargar variables de entorno
- gestionar configuraciones del servidor
- administrar claves JWT
- configurar conexiones con PostgreSQL

---

#### internal/models/

Contiene las estructuras principales del sistema utilizando structs de Go.

Aquí se modelan entidades como:

- usuarios
- libros
- préstamos
- descargas
- reportes

Estas estructuras representan las tablas de la base de datos.

---

#### internal/handlers/

Este paquete gestiona las solicitudes HTTP de la API REST.

Sus funciones son:

- recibir peticiones del cliente
- validar información
- enviar respuestas JSON
- controlar errores del sistema

---

#### internal/services/

El paquete services contiene la lógica de negocio del sistema.

Aquí se implementan procesos como:

- validación de préstamos
- control de disponibilidad de libros
- generación de reportes
- autenticación de usuarios

Esta capa conecta los handlers con los repositories.

---

#### internal/repositories/

Este paquete maneja el acceso a la base de datos.

Permite realizar:

- operaciones CRUD
- consultas SQL
- inserción de registros
- actualización de datos
- eliminación de información

Se utilizará GORM para facilitar la comunicación con PostgreSQL.

---

#### internal/middleware/

Contiene middlewares para la seguridad y control de acceso del sistema.

Permite implementar:

- autenticación JWT
- validación de permisos
- manejo de errores
- logging
- control de sesiones

---

#### internal/routes/

Este paquete organiza las rutas de la API REST.

Aquí se definirán endpoints como:

- /login
- /usuarios
- /libros
- /prestamos
- /descargas

Las rutas estarán separadas por módulos para mantener orden en el proyecto.

---

#### internal/database/

Este paquete administra la conexión principal con PostgreSQL.

Sus funciones son:

- abrir conexiones
- verificar estado de la base de datos
- ejecutar migraciones
- gestionar transacciones

---

#### pkg/

El paquete pkg almacenará utilidades reutilizables para distintos módulos del sistema.

Ejemplos:

- validadores
- generadores de tokens
- manejo de fechas
- funciones auxiliares
- herramientas de cifrado

---

#### docs/

Este directorio contendrá documentación técnica del proyecto como:

- diagramas
- documentación API
- manuales
- reportes técnicos

## Base de Datos Recomendada

Para el desarrollo del Sistema de Gestión de Libros Electrónicos se recomienda utilizar PostgreSQL como sistema gestor de base de datos principal.

PostgreSQL es una base de datos relacional moderna, segura y altamente utilizada en aplicaciones empresariales debido a su estabilidad, rendimiento y capacidad para manejar grandes cantidades de información.

La base de datos será una de las partes más importantes del sistema, ya que permitirá almacenar, organizar y administrar toda la información relacionada con la biblioteca virtual.

Dentro del sistema se almacenarán datos como:

- usuarios registrados
- libros electrónicos
- categorías
- préstamos digitales
- historial de descargas
- autenticación
- reportes y estadísticas

El uso de PostgreSQL permitirá que toda esta información se gestione de manera rápida, segura y eficiente.

---

### ¿Por qué se eligió PostgreSQL?

Se eligió PostgreSQL porque ofrece múltiples ventajas para proyectos backend desarrollados en Golang.

Entre las principales razones podemos destacar:

#### Seguridad

PostgreSQL posee sistemas avanzados de protección de datos y control de accesos.

Esto permitirá:

- proteger la información de los usuarios
- evitar accesos no autorizados
- administrar permisos
- garantizar integridad de datos

Además, se integrará con autenticación JWT para reforzar la seguridad del sistema.

---

#### Escalabilidad

La plataforma podrá crecer progresivamente sin afectar el rendimiento del sistema.

PostgreSQL permite trabajar con:

- múltiples usuarios simultáneos
- grandes volúmenes de información
- consultas complejas
- procesamiento eficiente de datos

Esto resulta importante para una biblioteca virtual donde constantemente existirán búsquedas, préstamos y descargas.

---

#### Compatibilidad con Golang

PostgreSQL tiene excelente integración con Golang mediante librerías modernas como:

- GORM
- pq
- pgx

Estas herramientas facilitan:

- conexión con la base de datos
- consultas automáticas
- operaciones CRUD
- migraciones automáticas
- validación de estructuras

Gracias a esto, el desarrollo backend será más organizado y eficiente.

---

### Relación entre las tablas

La base de datos trabajará mediante relaciones entre distintas entidades del sistema.

Por ejemplo:

- un usuario puede realizar múltiples préstamos
- un libro puede pertenecer a varias categorías
- un usuario puede descargar varios libros
- cada descarga quedará registrada en el historial

Estas relaciones permitirán mantener una estructura organizada y evitar duplicación de información.

---

### Tablas principales del sistema

#### Tabla Usuarios

Esta tabla almacenará toda la información relacionada con las personas registradas dentro de la plataforma.

Campos principales:

- id_usuario
- nombre
- correo electrónico
- contraseña cifrada
- rol del usuario
- fecha de registro

Los roles permitirán diferenciar:

- administradores
- lectores
- usuarios normales

---

#### Tabla Libros

Contendrá el catálogo digital de libros electrónicos disponibles dentro del sistema.

Campos principales:

- id_libro
- título
- autor
- categoría
- descripción
- formato del archivo
- disponibilidad

Esta tabla permitirá organizar correctamente la biblioteca virtual.

---

#### Tabla Préstamos

Registrará los préstamos digitales realizados por los usuarios.

Campos principales:

- id_prestamo
- id_usuario
- id_libro
- fecha de préstamo
- fecha de devolución
- estado del préstamo

Esta tabla permitirá llevar control sobre la disponibilidad de los libros electrónicos.

---

#### Tabla Descargas

Permitirá almacenar el historial de descargas realizadas dentro de la plataforma.

Campos principales:

- id_descarga
- id_usuario
- id_libro
- fecha de descarga

Esto ayudará a generar estadísticas y reportes administrativos.

---

### Beneficios de la base de datos en el sistema

La implementación de PostgreSQL permitirá:

- mejor organización de información
- acceso rápido a los datos
- control eficiente de usuarios
- seguridad de información
- generación de reportes
- optimización de búsquedas
- estabilidad del sistema
- soporte para múltiples usuarios

## Paquetes de Terceros Recomendados

### Descripción general

Para el desarrollo del Sistema de Gestión de Libros Electrónicos en Go, se recomienda utilizar paquetes de terceros que permitan construir un backend organizado, seguro y mantenible.

La selección de estos paquetes responde a las necesidades principales del proyecto: creación de una API REST, conexión con base de datos PostgreSQL, autenticación de usuarios, protección de contraseñas, configuración del entorno y comunicación con el frontend.

No se recomienda agregar librerías innecesarias, ya que esto puede aumentar la complejidad del sistema sin aportar beneficios reales. Por esta razón, los paquetes seleccionados se enfocan únicamente en funciones necesarias para la arquitectura propuesta.

### Paquetes recomendados

| Paquete | Función principal |
|---|---|
| `github.com/gin-gonic/gin` | Creación de la API REST, manejo de rutas, solicitudes HTTP, middlewares y respuestas JSON. |
| `gorm.io/gorm` | ORM para manejar operaciones con la base de datos mediante modelos y repositorios. |
| `gorm.io/driver/postgres` | Driver necesario para conectar GORM con la base de datos PostgreSQL. |
| `github.com/golang-jwt/jwt/v5` | Generación y validación de tokens JWT para autenticación y control de sesiones. |
| `golang.org/x/crypto/bcrypt` | Cifrado seguro de contraseñas antes de almacenarlas en la base de datos. |
| `github.com/spf13/viper` | Manejo centralizado de variables de configuración y entorno. |
| `github.com/gin-contrib/cors` | Configuración de permisos CORS para permitir la comunicación entre frontend y backend. |

### Justificación de los paquetes

#### Gin Web Framework

Gin será utilizado para construir la API REST del sistema. Este paquete permite definir rutas, recibir solicitudes HTTP, procesar datos enviados por el cliente y devolver respuestas en formato JSON.

Dentro del proyecto, Gin se relaciona directamente con la capa de controladores, ya que permitirá manejar las operaciones asociadas a usuarios, libros, préstamos, reservas, descargas y reportes.

Su uso resulta adecuado porque permite estructurar el backend de forma clara y facilita la creación de endpoints para cada módulo del sistema.

#### GORM

GORM será utilizado como ORM para facilitar la interacción entre el backend y la base de datos. Su función principal será permitir operaciones de creación, lectura, actualización y eliminación de registros sin depender completamente de consultas SQL manuales.

En el proyecto, GORM se aplicará principalmente en la capa de repositorios, donde se manejará el acceso a los datos de entidades como usuarios, libros, préstamos, reservas y descargas.

Este paquete ayuda a mantener una separación más clara entre la lógica de negocio y la lógica de acceso a datos.

#### Driver PostgreSQL para GORM

El paquete `gorm.io/driver/postgres` será necesario para conectar GORM con PostgreSQL, que es el sistema de base de datos definido en la arquitectura del proyecto.

Este driver permite que el backend desarrollado en Go pueda comunicarse correctamente con la base de datos, ejecutar operaciones sobre las tablas y gestionar la persistencia de la información.

Su inclusión es importante porque GORM requiere un driver específico según el motor de base de datos utilizado.

#### JWT

El paquete `github.com/golang-jwt/jwt/v5` será utilizado para implementar autenticación basada en tokens. Después de que un usuario inicie sesión correctamente, el sistema podrá generar un token JWT que será usado para validar futuras solicitudes.

Este paquete será importante para proteger rutas privadas del sistema y controlar el acceso según el tipo de usuario.

Dentro del proyecto, JWT puede aplicarse especialmente en funciones relacionadas con inicio de sesión, autorización de usuarios y validación de permisos.

#### Bcrypt

El paquete `golang.org/x/crypto/bcrypt` será utilizado para cifrar las contraseñas de los usuarios antes de almacenarlas en la base de datos.

Este paquete es necesario porque las contraseñas no deben guardarse en texto plano. Al utilizar bcrypt, el sistema mejora la protección de las credenciales y reduce el riesgo en caso de acceso no autorizado a la base de datos.

Su uso complementa el proceso de autenticación junto con JWT.

#### Viper

Viper será utilizado para manejar la configuración general del sistema. Este paquete permite centralizar valores como el puerto del servidor, la conexión a la base de datos, claves secretas, variables de entorno y otros parámetros necesarios para la ejecución del backend.

Su uso evita escribir datos sensibles directamente en el código fuente y facilita la administración del sistema en distintos entornos, como desarrollo, pruebas o producción.

#### CORS

El paquete `github.com/gin-contrib/cors` será utilizado para configurar los permisos de comunicación entre el frontend y el backend.

Esto es importante porque el sistema contempla una interfaz cliente que deberá consumir los servicios expuestos por la API REST. CORS permitirá definir qué orígenes, métodos y encabezados estarán permitidos al realizar solicitudes al servidor.

Su uso ayuda a controlar de forma más segura la interacción entre la aplicación web y el backend desarrollado en Go.

### Relación con la arquitectura del proyecto

Los paquetes recomendados se relacionan con las diferentes capas del sistema de la siguiente manera:

| Capa del sistema | Paquetes relacionados |
|---|---|
| API REST | Gin, CORS |
| Controladores | Gin |
| Servicios | JWT, Bcrypt |
| Repositorios | GORM |
| Base de datos | GORM, Driver PostgreSQL |
| Seguridad | JWT, Bcrypt, CORS |
| Configuración | Viper |

## Concurrencia en el Sistema

La concurrencia es uno de los conceptos más importantes implementados dentro del Sistema de Gestión de Libros Electrónicos.

La concurrencia permite que múltiples procesos se ejecuten al mismo tiempo dentro del sistema sin afectar el rendimiento general de la aplicación.

En este proyecto, la concurrencia será implementada utilizando herramientas nativas de Golang como:

- goroutines
- channels
- sincronización de procesos
- procesamiento simultáneo

Gracias a estas herramientas, el sistema podrá atender varios usuarios al mismo tiempo de manera rápida, eficiente y segura.

---

### ¿Por qué es importante la concurrencia?

Dentro de una plataforma digital existirán múltiples acciones ejecutándose simultáneamente.

Por ejemplo:

- usuarios iniciando sesión
- consultas de libros electrónicos
- búsquedas dinámicas
- préstamos digitales
- descargas de archivos
- generación de reportes
- validaciones de seguridad

Si el sistema ejecutara todas estas tareas una por una, el rendimiento sería lento y la experiencia del usuario se vería afectada.

Por esta razón, la concurrencia resulta fundamental para mejorar:

- velocidad
- eficiencia
- rendimiento
- estabilidad
- capacidad de respuesta

---

### Implementación de concurrencia en Golang

Golang posee un sistema de concurrencia moderno y eficiente que permite ejecutar múltiples tareas al mismo tiempo utilizando pocos recursos del servidor.

La concurrencia será implementada principalmente mediante:

#### Goroutines

Las goroutines son funciones ligeras que pueden ejecutarse en segundo plano de manera simultánea.

A diferencia de los hilos tradicionales, las goroutines consumen menos memoria y permiten procesar múltiples tareas de forma eficiente.

Dentro del sistema serán utilizadas para procesos como:

- procesamiento de descargas
- validación de préstamos
- generación automática de reportes
- actualización de estadísticas
- envío de notificaciones
- registro de actividades
- validación de usuarios

Gracias a las goroutines, el sistema podrá realizar varias operaciones simultáneamente sin bloquear el funcionamiento general de la plataforma.

Por ejemplo:

Mientras un usuario descarga un libro electrónico, otro usuario podrá iniciar sesión o realizar una búsqueda sin interrupciones.

---

### Uso de Channels

Los channels son mecanismos de comunicación utilizados en Golang para permitir el intercambio seguro de información entre goroutines.

Los channels ayudarán a:

- sincronizar procesos
- enviar información entre tareas concurrentes
- evitar conflictos de datos
- controlar operaciones simultáneas
- mantener estabilidad del sistema

Esto resulta importante porque múltiples usuarios podrán acceder al sistema al mismo tiempo.

Los channels permitirán coordinar correctamente la información para evitar errores o pérdidas de datos.

---

### Procesos concurrentes dentro del sistema

La concurrencia será aplicada en diferentes módulos del proyecto.

#### Gestión de préstamos digitales

Cuando varios usuarios soliciten préstamos simultáneamente, el sistema verificará la disponibilidad de los libros en tiempo real sin generar bloqueos.

Esto permitirá:

- validar préstamos rápidamente
- evitar duplicación de préstamos
- mantener disponibilidad actualizada

---

#### Gestión de descargas

Las descargas de libros electrónicos podrán ejecutarse en segundo plano mediante goroutines.

Esto permitirá:

- descargas simultáneas
- mejor velocidad de procesamiento
- menor tiempo de espera
- mejor experiencia del usuario

---

#### Generación de reportes

Los reportes estadísticos podrán generarse sin interrumpir el funcionamiento del sistema.

Por ejemplo:

- libros más descargados
- usuarios más activos
- historial de préstamos
- estadísticas generales

Todo esto podrá ejecutarse concurrentemente mientras otros usuarios continúan utilizando la plataforma.

---

### Beneficios de implementar concurrencia

La implementación de concurrencia dentro del Sistema de Gestión de Libros Electrónicos permitirá múltiples ventajas.

#### Mejor rendimiento

El sistema responderá más rápido incluso cuando existan muchos usuarios conectados simultáneamente.

---

#### Optimización de recursos

Las goroutines utilizan pocos recursos del servidor, permitiendo mayor eficiencia en el procesamiento.

---

#### Escalabilidad

La plataforma podrá crecer progresivamente soportando más usuarios y operaciones sin afectar significativamente el rendimiento.

---

#### Procesamiento simultáneo

Múltiples tareas podrán ejecutarse al mismo tiempo:

- búsquedas
- descargas
- préstamos
- autenticación
- reportes

---

#### Mejor experiencia de usuario

Los usuarios tendrán respuestas más rápidas y una navegación más fluida dentro de la plataforma.

---

### Relación de la concurrencia con Golang

La concurrencia representa una de las características más importantes de Golang y uno de los motivos principales por los cuales fue elegido para el desarrollo del proyecto.

Golang fue diseñado específicamente para:

- aplicaciones concurrentes
- servidores backend
- APIs REST
- procesamiento simultáneo
- sistemas escalables

Por esta razón, implementar concurrencia en el Sistema de Gestión de Libros Electrónicos permitirá aprovechar una de las mayores fortalezas del lenguaje.

---

## Flujo del Sistema

El flujo del sistema representa el proceso general de funcionamiento del Sistema de Gestión de Libros Electrónicos desde que el usuario ingresa a la plataforma hasta que finaliza sus actividades dentro del sistema.

El objetivo principal del flujo es organizar correctamente cada proceso que realiza el usuario y garantizar que todas las operaciones se ejecuten de manera segura, ordenada y eficiente.

Dentro del sistema existirán diferentes módulos conectados entre sí, permitiendo automatizar procesos relacionados con:

- autenticación
- gestión de usuarios
- búsqueda de libros
- préstamos digitales
- descargas
- reportes
- administración general

El flujo permitirá controlar paso a paso el funcionamiento de la plataforma virtual.

---

### Inicio del sistema

El flujo inicia cuando el usuario accede a la plataforma digital mediante la aplicación o navegador web.

Al ingresar al sistema, el usuario tendrá dos opciones principales:

- registrarse
- iniciar sesión

Esto permitirá controlar el acceso y garantizar que únicamente usuarios autorizados utilicen la plataforma.

---

### Registro de usuario

Si el usuario aún no posee una cuenta, deberá completar un proceso de registro.

El sistema solicitará información básica como:

- nombre
- correo electrónico
- contraseña
- tipo de usuario

Posteriormente, la información será validada y almacenada dentro de la base de datos PostgreSQL.

Durante este proceso el sistema verificará:

- correos duplicados
- campos obligatorios
- seguridad de contraseña
- formato correcto de datos

Una vez completado el registro, el usuario podrá acceder a la plataforma.

---

### Inicio de sesión y autenticación

Cuando el usuario inicia sesión, el sistema validará sus credenciales mediante autenticación JWT.

En esta etapa se verificará:

- correo electrónico
- contraseña
- permisos de acceso
- rol del usuario

Si las credenciales son correctas, el sistema generará un token JWT que permitirá mantener la sesión activa y segura.

Esto ayudará a proteger la información del sistema y evitar accesos no autorizados.

---

### Acceso al sistema

Después de la autenticación, el usuario ingresará al panel principal de la plataforma.

Dependiendo del tipo de usuario, el sistema mostrará diferentes funcionalidades.

Por ejemplo:

#### Administradores

Podrán:

- gestionar libros
- editar información
- eliminar registros
- generar reportes
- controlar usuarios

#### Usuarios lectores

Podrán:

- consultar libros
- realizar búsquedas
- solicitar préstamos
- descargar archivos electrónicos

Esto permitirá mantener control de permisos y seguridad dentro del sistema.

---

### Gestión y búsqueda de libros electrónicos

Una de las funciones principales del sistema será la consulta del catálogo digital.

Los usuarios podrán realizar búsquedas utilizando filtros como:

- nombre del libro
- autor
- categoría
- disponibilidad
- formato

El sistema procesará las búsquedas dinámicamente y mostrará resultados organizados.

Esto facilitará el acceso rápido a la información y mejorará la experiencia de navegación.

---

### Selección del libro

Después de encontrar un libro, el usuario podrá visualizar información detallada relacionada con:

- título
- autor
- descripción
- categoría
- formato del archivo
- disponibilidad

El sistema permitirá decidir si el usuario desea:

- solicitar préstamo
- descargar el libro
- guardar información para futuras consultas

---

### Validación de disponibilidad

Antes de autorizar un préstamo digital, el sistema verificará automáticamente si el libro se encuentra disponible.

En esta etapa el sistema validará:

- estado del libro
- permisos del usuario
- historial de préstamos
- límites de acceso

Si el libro está disponible, el préstamo será aprobado.

En caso contrario, el sistema mostrará un mensaje indicando que el recurso no se encuentra disponible temporalmente.

---

### Registro del préstamo digital

Una vez aprobado el préstamo, el sistema almacenará automáticamente la información dentro de la base de datos.

Se registrarán datos como:

- usuario solicitante
- libro seleccionado
- fecha de préstamo
- fecha de devolución
- estado del préstamo

Esto permitirá mantener un historial organizado de actividades realizadas dentro de la plataforma.

---

### Descarga de libros electrónicos

Después del préstamo, el usuario podrá descargar el archivo digital.

El sistema verificará:

- autenticación del usuario
- permisos de descarga
- disponibilidad del archivo

Posteriormente iniciará la descarga del libro electrónico.

Además, el sistema registrará automáticamente el historial de descargas para generar estadísticas futuras.

---

### Gestión de reportes

El sistema también permitirá generar reportes administrativos relacionados con el funcionamiento de la plataforma.

Entre los reportes se podrán obtener:

- libros más descargados
- usuarios más activos
- historial de préstamos
- categorías más consultadas
- estadísticas generales

Estos reportes ayudarán a mejorar la administración y control de la biblioteca virtual.

---

### Cierre de sesión

Finalmente, el usuario podrá cerrar sesión de forma segura.

Al cerrar sesión:

- el token JWT será invalidado
- finalizará el acceso del usuario
- se protegerá la información del sistema

Esto permitirá mantener seguridad y control de accesos dentro de la plataforma.

---

### Importancia del flujo del sistema

El flujo general del sistema permitirá:

- organizar correctamente los procesos
- automatizar actividades
- mejorar la seguridad
- optimizar tiempos de respuesta
- facilitar navegación de usuarios
- mantener control de información

Además, ayudará a que el sistema funcione de manera estructurada, moderna y eficiente.

---

### Resumen general del flujo

El funcionamiento general del sistema seguirá el siguiente proceso:

Inicio
↓
Registro o inicio de sesión
↓
Autenticación JWT
↓
Acceso al sistema
↓
Consulta de libros
↓
Búsquedas dinámicas
↓
Selección del libro
↓
Validación de disponibilidad
↓
Préstamo digital
↓
Descarga del archivo
↓
Registro de actividades
↓
Generación de reportes
↓
Cierre de sesión
↓
Fin

## Metodología a Utilizar

Se utilizará una metodología incremental y modular, permitiendo desarrollar el sistema por etapas funcionales independientes para posteriormente integrarlas en una sola plataforma.

Las etapas principales serán:

1. Planeación y análisis
2. Diseño de arquitectura
3. Desarrollo de modelos y base de datos
4. Implementación de repositorios y servicios
5. Desarrollo de APIs REST
6. Seguridad y autenticación
7. Aplicación de programación funcional
8. Integración y pruebas


