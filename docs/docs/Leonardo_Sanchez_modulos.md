# MÓDULOS DEL SISTEMA DE GESTIÓN DE LIBROS ELECTRÓNICOS EN GO

## Introducción

El Sistema de Gestión de Libros Electrónicos será desarrollado utilizando Go bajo una **arquitectura modular**. Esto significa que el sistema estará dividido en diferentes módulos independientes, cada uno responsable de funcionalidades específicas.

En Go, los módulos permiten mantener el proyecto estructurado, reutilizable y fácil de mantener. Gracias a la organización por paquetes que ofrece Go, cada módulo podrá separarse en carpetas independientes, facilitando el desarrollo colaborativo y el mantenimiento del código.

## Estructura de cada módulo

Cada módulo contará con una arquitectura en capas:

- **Controladores** - Manejan solicitudes HTTP y respuestas
- **Servicios** - Contienen la lógica de negocio
- **Repositorios** - Interactúan con la base de datos
- **Modelos** - Definen las estructuras de datos
- **Rutas** - Definen los endpoints
- **Validaciones** - Validan datos de entrada

Esta separación permite mantener un código limpio, modular y profesional.

## Organización del proyecto

```
/proyecto-libros
│
├── controllers/
│   ├── usuario_controller.go
│   ├── libro_controller.go
│   ├── prestamo_controller.go
│   ├── reserva_controller.go
│   ├── busqueda_controller.go
│   └── reporte_controller.go
│
├── services/
├── repositories/
├── models/
├── routes/
├── middleware/
├── database/
├── config/
└── utils/
```

---

# MÓDULOS DEL SISTEMA

## 1️⃣ MÓDULO DE GESTIÓN DE USUARIOS

### Descripción

Responsable de administrar toda la información relacionada con los usuarios del sistema. Es uno de los módulos más críticos porque controla el acceso, autenticación y permisos dentro de la plataforma.

### Objetivos

- ✅ Registrar nuevos usuarios
- ✅ Permitir inicio de sesión
- ✅ Gestionar perfiles de usuario
- ✅ Recuperar contraseñas
- ✅ Controlar acceso según permisos asignados

### Arquitectura del módulo

```
┌─────────────────────┐
│ Usuario Controller  │
└──────────┬──────────┘
           │ Recibe solicitudes HTTP
           │
┌──────────▼──────────┐
│ Usuario Service     │
└──────────┬──────────┘
           │ Lógica de negocio
           │
┌──────────▼──────────┐
│ Usuario Repository  │
└──────────┬──────────┘
           │ Acceso a datos
           │
┌──────────▼──────────┐
│ MySQL Database      │
└─────────────────────┘
```

### Componentes

#### 🎮 Usuario Controller

**Responsabilidades:**
- Recibir y procesar solicitudes HTTP
- Validar datos de entrada básicos
- Enviar respuestas en formato JSON
- Manejar errores de forma controlada

**Endpoints principales:**
- `POST /usuarios/registro` - Registrar usuario
- `POST /usuarios/login` - Iniciar sesión
- `PUT /usuarios/{id}` - Actualizar perfil
- `POST /usuarios/recuperar-password` - Recuperar contraseña

#### 🔧 Usuario Service

**Responsabilidades:**
- Implementar toda la lógica de negocio
- Verificar si el correo ya existe
- Encriptar contraseñas
- Validar credenciales
- Generar tokens JWT
- Validar reglas de negocio

#### 💾 Usuario Repository

**Responsabilidades:**
- Guardar nuevos usuarios en la BD
- Consultar información de usuarios
- Actualizar registros existentes
- Eliminar usuarios (soft delete)

### Flujo de Registro

```
1. Usuario envía formulario de registro
   ↓
2. Controlador recibe y valida datos básicos
   ↓
3. Servicio verifica que el email no exista
   ↓
4. Servicio encripta la contraseña con bcrypt
   ↓
5. Repositorio guarda datos en MySQL
   ↓
6. Sistema retorna confirmación con token JWT
```

### Flujo de Inicio de Sesión

```
1. Usuario envía credenciales
   ↓
2. Repositorio busca el usuario por email
   ↓
3. Servicio verifica la contraseña con bcrypt
   ↓
4. Servicio genera token JWT
   ↓
5. Sistema retorna token de autenticación
```

### Seguridad implementada

**Tecnologías utilizadas:**
- **bcrypt** - Cifrado de contraseñas (hashing seguro)
- **JWT (JSON Web Tokens)** - Autenticación y autorización

**Garantías:**
- ✅ Contraseñas encriptadas y almacenadas de forma segura
- ✅ Protección de datos sensibles
- ✅ Control granular de acceso
- ✅ Tokens de sesión seguros

---

## 2️⃣ MÓDULO DE GESTIÓN DE LIBROS

### Descripción

Administra el catálogo completo de libros electrónicos. Es el núcleo principal de la biblioteca virtual donde se almacena y gestiona toda la información bibliográfica.

### Objetivos

- 📚 Registrar libros en el sistema
- 📚 Actualizar información bibliográfica
- 📚 Organizar libros por categorías
- 📚 Gestionar archivos digitales (PDF, EPUB)

### Arquitectura del módulo

```
┌────────────────────┐
│ Libro Controller   │
└─────────┬──────────┘
          │
┌─────────▼──────────┐
│ Libro Service      │
└─────────┬──────────┘
          │
┌─────────▼──────────┐
│ Libro Repository   │
└─────────┬──────────┘
          │
┌─────────▼──────────┐
│ MySQL Database     │
└────────────────────┘
```

### Información almacenada por libro

```
Título
Autor(es)
Categoría/Género
Editorial
ISBN
Año de publicación
Archivo (PDF/EPUB)
Disponibilidad
Fecha de ingreso
```

### Funcionalidades principales

#### 📥 Registro de libros

```
1. Bibliotecario carga información del libro
   ↓
2. Controlador recibe y valida datos
   ↓
3. Servicio verifica ISBN único
   ↓
4. Servicio valida formato del archivo
   ↓
5. Repositorio guarda información
   ↓
6. Archivo se almacena en el servidor
```

#### ✏️ Actualización de libros

- Editar títulos y metadata
- Cambiar categorías
- Reemplazar archivos
- Modificar disponibilidad
- Agregar/actualizar descripciones

#### 🗑️ Eliminación de libros

- Eliminar registros (soft delete)
- Desactivar libros obsoletos
- Mantener historial de eliminaciones

### Tecnologías utilizadas

- **Gin** - Framework para API REST de alto rendimiento
- **GORM** - ORM para conexión y consultas a MySQL
- **Go FileSystem** - Manejo seguro de archivos

### Importancia

Este módulo es crítico porque:
- Organiza la biblioteca digital
- Gestiona el contenido principal del sistema
- Facilita búsquedas y consultas
- Administra archivos electrónicos grandes

---

## 3️⃣ MÓDULO DE BÚSQUEDA DE LIBROS

### Descripción

Permite búsquedas rápidas y eficientes dentro del sistema mediante filtros avanzados y consultas optimizadas.

### Objetivo

Facilitar a los usuarios la localización de libros electrónicos de forma intuitiva y rápida.

### Arquitectura del módulo

```
┌─────────────────────────┐
│ Búsqueda Controller     │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│ Búsqueda Service        │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│ Búsqueda Repository     │
└────────────┬────────────┘
             │
┌────────────▼────────────┐
│ MySQL Database          │
└─────────────────────────┘
```

### Tipos de búsqueda

#### 🔍 Búsqueda por título

- Buscar mediante palabras clave
- Búsqueda parcial (LIKE)
- Auto-completar sugerencias

#### 👤 Búsqueda por autor

- Localizar libros de autores específicos
- Búsqueda por apellido o nombre completo
- Listar todos los libros de un autor

#### 📂 Búsqueda por categorías

**Categorías disponibles:**
- Programación
- Redes y Telecomunicaciones
- Matemáticas
- Literatura
- Ciencias
- Historia
- Educación

#### 🎚️ Filtros avanzados

El sistema permite filtrar por:

| Filtro | Descripción |
|--------|-------------|
| Año | Año de publicación |
| Disponibilidad | Libros disponibles/prestados |
| Popularidad | Libros más consultados |
| Categoría | Género o tema |
| Editorial | Casa editorial |

### Procesamiento en Go

```
1. Usuario envía parámetros de búsqueda
   ↓
2. Controlador valida parámetros
   ↓
3. Servicio construye consulta SQL optimizada
   ↓
4. Repositorio ejecuta query en MySQL
   ↓
5. Servicio filtra y ordena resultados
   ↓
6. Controlador retorna JSON con resultados
```

### Ventajas

- ⚡ Mejora experiencia del usuario
- ⚡ Reduce tiempo de búsqueda
- ⚡ Optimiza navegación del catálogo
- ⚡ Consultas optimizadas y cacheadas

---

## 4️⃣ MÓDULO DE PRÉSTAMOS ELECTRÓNICOS

### Descripción

Controla todo el proceso de préstamos digitales. Es uno de los procesos más importantes del sistema y automatiza la circulación de libros.

### Objetivo

Administrar de forma integral:
- 📤 Solicitudes de préstamo
- 📊 Disponibilidad de libros
- 📥 Devoluciones
- 📋 Historiales de transacciones

### Arquitectura del módulo

```
┌──────────────────────┐
│ Préstamo Controller  │
└─────────┬────────────┘
          │
┌─────────▼────────────┐
│ Préstamo Service     │
└─────────┬────────────┘
          │
┌─────────▼────────────┐
│ Préstamo Repository  │
└─────────┬────────────┘
          │
┌─────────▼────────────┐
│ MySQL Database       │
└──────────────────────┘
```

### Flujo del préstamo

#### Paso 1: Solicitud del usuario

```
Usuario solicita un libro
  ↓
Controlador recibe la solicitud
```

#### Paso 2: Validación del servicio

El servicio verifica:

- ✅ Usuario existe y está activo
- ✅ Libro existe en el catálogo
- ✅ Libro está disponible
- ✅ Usuario no ha excedido límite de préstamos
- ✅ Usuario no tiene deuda pendiente

#### Paso 3: Registro en base de datos

```
Repositorio registra información:
  - ID usuario
  - ID libro
  - Fecha préstamo
  - Fecha devolución esperada
  - Estado: ACTIVO
```

#### Paso 4: Actualización de disponibilidad

```
Sistema actualiza:
  - Estado del libro a NO DISPONIBLE
  - Contador de préstamos
  - Fecha última transacción
```

### Funcionalidades principales

#### 🎫 Solicitud de préstamos

- Permite solicitar libros electrónicos
- Genera confirmación de préstamo
- Define automáticamente fecha de devolución (ej: 15 días)

#### 🔄 Gestión de devoluciones

- Registra devolución del libro
- Calcula multas por retraso (si aplica)
- Libera el libro para nuevos préstamos

#### 📚 Historial de préstamos

- Guarda todos los préstamos realizados
- Muestra estado actual de cada préstamo
- Permite generar reportes

#### ⚠️ Restricciones

- Límite máximo de préstamos simultáneos por usuario (ej: 5 libros)
- Máximo de renovaciones permitidas (ej: 2 veces)
- Control de deudas y multas

### Importancia

- Automatiza procesos bibliotecarios manuales
- Mantiene control sobre circulación digital
- Previene fraudes y accesos no autorizados
- Facilita estadísticas de uso

---

## 5️⃣ MÓDULO DE RESERVAS

### Descripción

Administra reservas de libros que están temporalmente no disponibles. Permite que usuarios esperen su turno de forma ordenada.

### Objetivo

Permitir que usuarios puedan **reservar libros no disponibles** y ser notificados automáticamente cuando estén disponibles.

### Arquitectura del módulo

```
┌────────────────────┐
│ Reserva Controller │
└────────┬───────────┘
         │
┌────────▼───────────┐
│ Reserva Service    │
└────────┬───────────┘
         │
┌────────▼───────────┐
│ Reserva Repository │
└────────┬───────────┘
         │
┌────────▼───────────┐
│ MySQL Database     │
└────────────────────┘
```

### Proceso de reserva

```
1. Usuario solicita reservar libro (no disponible)
   ↓
2. Sistema verifica disponibilidad actual
   ↓
3. Registra la reserva en cola de espera
   ↓
4. Guarda posición en lista de espera
   ↓
5. Usuario recibe confirmación de reserva
   ↓
6. Al devolver libro, se notifica a usuario
```

### Funcionalidades

#### 📝 Registro de reservas

- Guardar solicitud de reserva
- Asignar número de posición en cola
- Registrar fecha de solicitud

#### ⏳ Lista de espera

- Organizar usuarios automáticamente por fecha
- FIFO (First In, First Out)
- Mostrar posición actual en cola

#### 🔔 Notificaciones

- Email cuando libro está disponible
- SMS (opcional)
- Ventana de disponibilidad (ej: 48 horas)

#### ❌ Cancelaciones

- Permite cancelar reservas
- Libera espacio en cola
- Mueve usuarios hacia arriba automáticamente

### Importancia

- Evita conflictos entre usuarios
- Mejora administración de recursos digitales
- Aumenta equidad en acceso a contenido popular
- Reduce frustraciones de usuarios

---

## 6️⃣ MÓDULO DE REPORTES Y ESTADÍSTICAS

### Descripción

Genera información administrativa y estadística para análisis del funcionamiento del sistema. Proporciona insights sobre uso y tendencias.

### Objetivo

Permitir **análisis detallado y control** del funcionamiento del sistema mediante reportes y gráficos.

### Arquitectura del módulo

```
┌──────────────────────┐
│ Reporte Controller   │
└─────────┬────────────┘
          │
┌─────────▼────────────┐
│ Reporte Service      │
└─────────┬────────────┘
          │
┌─────────▼────────────┐
│ Reporte Repository   │
└─────────┬────────────┘
          │
┌─────────▼────────────┐
│ MySQL Database       │
└──────────────────────┘
```

### Reportes disponibles

#### 📊 Reportes de préstamos

- Historial completo de préstamos realizados
- Libros más prestados
- Períodos de mayor demanda
- Estadísticas de renovaciones

#### 👥 Reportes de usuarios

- Usuarios más activos
- Patrones de comportamiento
- Usuarios inactivos
- Deudores

#### 📚 Reportes de libros

- Libros más consultados/descargados
- Libros nunca consultados
- Desempeño por categoría
- Tasa de disponibilidad promedio

#### 📈 Estadísticas generales

- Métricas del sistema completo
- Crecimiento de colección
- Índice de circulación
- Satisfacción estimada

### Formatos de exportación

El sistema puede generar reportes en múltiples formatos:

- **PDF** - Documentos profesionales imprimibles
- **Excel** - Hojas de cálculo para análisis adicional
- **CSV** - Datos crudos para importación

### Procesamiento en Go

Go permite:

- ✨ Generar archivos rápidamente
- ✨ Procesar grandes cantidades de datos
- ✨ Optimizar consultas complejas
- ✨ Manejar múltiples exportaciones simultáneas

### Importancia

Facilita:
- 🎯 Toma de decisiones informadas
- 🎯 Administración estratégica
- 🎯 Supervisión del sistema
- 🎯 Identificación de mejoras necesarias

---

## 7️⃣ MÓDULO DE SEGURIDAD

### Descripción

Protege toda la información del sistema. Es fundamental porque el sistema maneja datos sensibles de usuarios, contraseñas y acceso a contenido protegido.

### Información sensible gestionada

- 🔐 Datos de usuarios
- 🔐 Contraseñas
- 🔐 Archivos digitales
- 🔐 Información académica/personal
- 🔐 Tokens de sesión

### Objetivos de seguridad

- 🛡️ **Confidencialidad** - Información no accesible a no autorizados
- 🛡️ **Integridad** - Datos no modificados sin autorización
- 🛡️ **Disponibilidad** - Sistema accesible para usuarios autorizados

### Arquitectura del módulo

```
┌──────────────────────────┐
│ Middleware de Seguridad  │
└────────────┬─────────────┘
             │ Valida cada solicitud
             │
┌────────────▼─────────────┐
│ Servicios de Autenticación│
└────────────┬─────────────┘
             │ Genera/valida tokens
             │
┌────────────▼─────────────┐
│ Control de Acceso        │
└──────────────────────────┘
             │ Verifica permisos
```

### Componentes de seguridad

#### 🔑 Autenticación JWT

**Propósito:** Validar identidad de usuarios

**Características:**
- Validar sesiones activas
- Generar tokens seguros y únicos
- Proteger rutas/endpoints
- Tokens con expiración configurable

**Flujo:**

```
1. Usuario inicia sesión
   ↓
2. Sistema verifica credenciales
   ↓
3. Genera JWT con información del usuario
   ↓
4. Usuario envía token en cada solicitud
   ↓
5. Middleware valida token antes de procesar
```

#### 🔒 Cifrado de contraseñas

**Algoritmo:** bcrypt (hashing seguro y adaptativo)

**Características:**
- Contraseñas nunca se almacenan en texto plano
- Salt automático incorporado
- Resistente a ataques de fuerza bruta
- Adaptativo a aumento de poder computacional

#### 🚪 Middleware de autenticación

**Responsabilidades:**
- Proteger endpoints sensibles
- Validar tokens JWT en cada solicitud
- Verificar permisos de usuario
- Bloquear accesos indebidos
- Registrar intentos fallidos

#### 📝 Auditoría

**Registro de:**
- Inicios de sesión (exitosos y fallidos)
- Cambios de contraseña
- Modificaciones de datos sensibles
- Errores y excepciones
- Actividades críticas del sistema

**Almacenamiento:** Se guardan con:
- Timestamp
- ID usuario
- Acción realizada
- Resultado (éxito/fallo)

### Buenas prácticas implementadas

- ✅ Validación de entrada en todos los endpoints
- ✅ Rate limiting para prevenir ataques
- ✅ HTTPS/TLS obligatorio
- ✅ Cookies de sesión con HttpOnly
- ✅ CORS configurado restrictivamente
- ✅ SQL Injection prevention (GORM con parametrización)
- ✅ XSS protection en respuestas

### Importancia crítica

Protege:
- 🎯 Información personal de usuarios
- 🎯 Contenido digital de valor
- 🎯 Integridad del sistema
- 🎯 Acceso seguro y controlado
- 🎯 Cumplimiento normativo

---

# CONCLUSIÓN

## Resumen de la arquitectura modular

La implementación modular utilizando Go permite desarrollar un **Sistema de Gestión de Libros Electrónicos** altamente:

- ✅ **Organizado** - Cada módulo con responsabilidad específica
- ✅ **Eficiente** - Rendimiento optimizado con Go
- ✅ **Escalable** - Fácil de expandir con nuevos módulos
- ✅ **Mantenible** - Código limpio y separación de concerns

## Beneficios de la separación modular

- 👥 **Desarrollo colaborativo** - Equipos pueden trabajar en módulos independientes
- 🔧 **Mantenimiento** - Bugs aislados a módulos específicos
- 🔐 **Seguridad** - Cada módulo puede tener validaciones propias
- ♻️ **Reutilización** - Componentes reutilizables entre módulos
- 📈 **Escalabilidad** - Fácil agregar nuevos módulos o funcionalidades

## Stack tecnológico

| Tecnología | Propósito |
|------------|-----------|
| **Go** | Lenguaje base - Alto rendimiento y concurrencia |
| **Gin** | Framework HTTP REST - Rápido y minimalista |
| **GORM** | ORM - Abstracción de BD con seguridad |
| **MySQL** | Base de datos - Confiable y escalable |
| **bcrypt** | Seguridad - Hash de contraseñas |
| **JWT** | Autenticación - Tokens seguros |

## Resultado final

Este diseño modular utilizando Go permitirá construir una **aplicación moderna, rápida y segura** para la gestión de bibliotecas digitales, con capacidad de crecimiento y fácil mantenimiento a largo plazo.
