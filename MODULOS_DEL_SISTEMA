# MÓDULOS DEL SISTEMA DE GESTIÓN DE LIBROS ELECTRÓNICOS EN GO

## Introducción

El Sistema de Gestión de Libros Electrónicos será desarrollado utilizando **Go** bajo una arquitectura modular. Esta estructura permite dividir el sistema en módulos independientes que facilitan el desarrollo colaborativo, mantenimiento y escalabilidad.

---

## Tabla de Contenidos

1. [Módulo de Gestión de Usuarios](#1-módulo-de-gestión-de-usuarios)
2. [Módulo de Gestión de Libros](#2-módulo-de-gestión-de-libros)
3. [Módulo de Búsqueda de Libros](#3-módulo-de-búsqueda-de-libros)
4. [Módulo de Préstamos Electrónicos](#4-módulo-de-préstamos-electrónicos)
5. [Módulo de Reservas](#5-módulo-de-reservas)
6. [Módulo de Reportes y Estadísticas](#6-módulo-de-reportes-y-estadísticas)
7. [Módulo de Seguridad](#7-módulo-de-seguridad)

---

## 1. MÓDULO DE GESTIÓN DE USUARIOS

### Descripción General
Administra toda la información relacionada con los usuarios del sistema. Es uno de los módulos más importantes porque controla el acceso, autenticación y permisos dentro de la plataforma.

### Objetivos del Módulo
- ✓ Registrar nuevos usuarios
- ✓ Permitir inicio de sesión
- ✓ Gestionar perfiles de usuario
- ✓ Recuperar contraseñas perdidas
- ✓ Controlar acceso según permisos asignados

### Arquitectura Interna

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
│       MySQL             │ (Base de datos)
└─────────────────────────┘
```

### Componentes Internos

#### Usuario Controller
**Responsabilidades:**
- Recibir solicitudes HTTP
- Validar datos básicos
- Enviar respuestas JSON

**Operaciones principales:**
- Registro de usuario
- Login/Autenticación
- Actualización de perfil
- Obtener información de usuario

#### Usuario Service
**Responsabilidades:**
Contiene toda la lógica del negocio:
- Verificar si el correo ya existe
- Encriptar contraseñas
- Validar credenciales
- Generar tokens JWT
- Validar permisos de usuario

#### Usuario Repository
**Responsabilidades:**
- Guardar nuevos usuarios
- Consultar información de usuarios
- Actualizar registros existentes
- Eliminar usuarios
- Buscar usuarios por criterios

### Procesos Principales

#### 1. Registro de Usuarios
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
4. Repositorio guarda datos en MySQL
   ↓
5. Retorna token JWT al cliente
```

#### 2. Inicio de Sesión
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

### Seguridad Implementada

**Librerías utilizadas:**
- **bcrypt** → Cifrado y hash de contraseñas
- **JWT (JSON Web Tokens)** → Autenticación y autorización

**Medidas de seguridad:**
- ✓ Contraseñas hasheadas (nunca en texto plano)
- ✓ Tokens JWT con expiración
- ✓ Validación de datos entrada
- ✓ Control de acceso por roles
- ✓ Auditoría de accesos

### Importancia del Módulo
Controla aspectos críticos del sistema:
- Seguridad general
- Autenticación de usuarios
- Asignación de roles y permisos
- Protección de información personal

---

## 2. MÓDULO DE GESTIÓN DE LIBROS

### Descripción General
Administra todo el catálogo digital de libros electrónicos. Es considerado el núcleo principal de la biblioteca virtual.

### Objetivos del Módulo
- ✓ Registrar nuevos libros
- ✓ Actualizar información de libros
- ✓ Organizar por categorías
- ✓ Gestionar archivos digitales (PDF, EPUB)
- ✓ Controlar disponibilidad

### Arquitectura Interna

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
│       MySQL             │
└─────────────────────────┘
```

### Información Gestionada por Libro

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

### Procesos Principales

#### 1. Registro de Libros
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
4. Repositorio guarda información en MySQL
   ↓
5. Archivo se almacena en servidor
   ↓
6. Retorna confirmación al bibliotecario
```

#### 2. Actualización de Libros
El sistema permite:
- Editar títulos y autores
- Cambiar categorías
- Reemplazar archivos digitales
- Modificar estado de disponibilidad
- Actualizar número de copias

#### 3. Eliminación de Libros
- Eliminar registros permanentemente
- Desactivar libros (mantener histórico)
- Limpiar información obsoleta
- Mantener referencia en préstamos históricos

### Tecnologías Utilizadas en Go

- **Gin** → Framework para API REST
- **GORM** → ORM para conexión con MySQL
- **FileSystem de Go** → Manejo y almacenamiento de archivos
- **Multer/Form parsing** → Procesamiento de uploads

### Importancia del Módulo

Permite:
- Organizar biblioteca digital de forma eficiente
- Gestionar contenido digital
- Facilitar búsquedas rápidas
- Administrar archivos electrónicos
- Mantener integridad del catálogo

---

## 3. MÓDULO DE BÚSQUEDA DE LIBROS

### Descripción General
Facilita búsquedas rápidas y eficientes dentro del sistema mediante filtros y consultas avanzadas.

### Objetivos del Módulo
- ✓ Localizar libros por múltiples criterios
- ✓ Aplicar filtros avanzados
- ✓ Retornar resultados relevantes
- ✓ Optimizar tiempo de búsqueda

### Arquitectura Interna

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
│       MySQL             │
└─────────────────────────┘
```

### Funciones Principales

#### 1. Búsqueda por Título
- Permite buscar libros mediante palabras clave
- Búsqueda parcial (LIKE)
- Orden por relevancia

#### 2. Búsqueda por Autor
- Localiza libros relacionados con autores específicos
- Soporte para múltiples autores
- Búsqueda flexible

#### 3. Búsqueda por Categorías
Ejemplos de categorías:
- Programación
- Redes y Telecomunicaciones
- Matemáticas
- Literatura
- Ciencias
- Historia

#### 4. Filtros Avanzados
Permite filtrar por:
- **Año de publicación** (rango)
- **Estado de disponibilidad** (disponible/prestado)
- **Popularidad** (más consultados)
- **Categoría/Subcategoría**
- **Editorial**
- **Tipo de archivo** (PDF, EPUB)

### Procesamiento en Go

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

### Ventajas del Módulo

- Mejora significativa de la experiencia del usuario
- Reduce tiempo de búsqueda
- Optimiza navegación en el catálogo
- Facilita descubrimiento de libros
- Paginación eficiente de resultados

---

## 4. MÓDULO DE PRÉSTAMOS ELECTRÓNICOS

### Descripción General
Controla todo el proceso de préstamos digitales. Es uno de los procesos más importantes del sistema.

### Objetivos del Módulo
- ✓ Administrar solicitudes de préstamo
- ✓ Verificar disponibilidad
- ✓ Gestionar devoluciones
- ✓ Mantener historiales completos
- ✓ Controlar límites por usuario

### Arquitectura Interna

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
│       MySQL             │
└─────────────────────────┘
```

### Flujo de Préstamo en Go

#### Cuando un usuario solicita un libro:

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
└─→ Repositorio registra información en MySQL
   ├─→ Guarda registro en tabla PRESTAMOS
   ├─→ Registra fecha de inicio
   ├─→ Establece fecha de vencimiento

PASO 4: Actualización
└─→ Sistema actualiza disponibilidad del libro
   ├─→ Reduce cantidad disponible
   ├─→ Marca libro como parcialmente prestado
   └─→ Retorna confirmación al usuario
```

### Funciones Principales

#### 1. Solicitud de Préstamos
- Permite solicitar libros digitales
- Validación automática de criterios
- Generación de comprobante digital
- Notificación al usuario

#### 2. Gestión de Devoluciones
- Procesar devoluciones de libros
- Calcular multas si es necesario
- Actualizar disponibilidad
- Generar comprobante de devolución

#### 3. Historial de Préstamos
- Guarda todos los préstamos realizados
- Mantiene registro de devoluciones
- Información sobre multas generadas
- Análisis de hábitos de lectura

#### 4. Control de Restricciones
- Límite de libros por usuario (ej: máximo 5)
- Duración máxima de préstamo (ej: 30 días)
- Penalizaciones por retraso
- Bloqueo temporal por incumplimiento

### Datos del Préstamo

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

### Importancia del Módulo

Permite:
- Automatizar procesos bibliotecarios
- Mantener control sobre circulación digital
- Generar estadísticas de uso
- Prevenir conflictos entre usuarios
- Facilitar auditoría de préstamos

---

## 5. MÓDULO DE RESERVAS

### Descripción General
Administra reservas de libros temporalmente no disponibles.

### Objetivos del Módulo
- ✓ Permitir reservas de libros ocupados
- ✓ Mantener lista de espera
- ✓ Notificar disponibilidad
- ✓ Gestionar cancelaciones

### Arquitectura Interna

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
│       MySQL             │
└─────────────────────────┘
```

### Proceso de Reserva

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

### Funciones Principales

#### 1. Registro de Reservas
- Guardar solicitudes de reserva
- Validar datos de entrada
- Asignar prioridad automática

#### 2. Lista de Espera
- Organizar usuarios automáticamente
- Respetar orden FIFO (First In, First Out)
- Mantener histórico de esperas

#### 3. Notificaciones
- Informar cuando libro está disponible
- Enviar recordatorios
- Confirmar recepción de notificación

#### 4. Cancelaciones
- Permitir eliminar reservas
- Liberar posición en lista
- Notificar cancelación al usuario

### Datos de Reserva

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

### Importancia del Módulo

- Evita conflictos entre usuarios
- Mejora administración de recursos digitales
- Garantiza acceso equitativo
- Reduce pérdida de demanda

---

## 6. MÓDULO DE REPORTES Y ESTADÍSTICAS

### Descripción General
Genera información administrativa y estadística del sistema.

### Objetivos del Módulo
- ✓ Permitir análisis de funcionamiento
- ✓ Generar reportes administrativos
- ✓ Crear estadísticas de uso
- ✓ Facilitar toma de decisiones

### Arquitectura Interna

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
│       MySQL             │
└─────────────────────────┘
```

### Funciones Principales

#### 1. Reportes de Préstamos
- Muestra historial de préstamos realizados
- Análisis de tendencias
- Identificación de períodos de alta demanda
- Libros más prestados

#### 2. Reportes de Usuarios
- Identifica usuarios más activos
- Análisis de perfiles de lectura
- Detecta patrones de uso
- Usuarios con retrasos o multas

#### 3. Reportes de Libros
- Muestra libros más consultados
- Libros menos utilizados
- Análisis de categorías populares
- Disponibilidad vs. demanda

#### 4. Estadísticas Generales
- Métricas generales del sistema
- Número total de usuarios
- Cantidad de libros en catálogo
- Préstamos activos
- Reservas pendientes
- Tasa de ocupación

### Exportación de Información

El sistema puede generar reportes en formatos:
- **PDF** → Para impresión y distribución formal
- **Excel** → Para análisis adicional
- **CSV** → Para integración con otras herramientas

### Procesamiento en Go

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
- ✓ Generación rápida de archivos
- ✓ Procesa grandes volúmenes de datos eficientemente
- ✓ Optimiza consultas complejas
- ✓ Manejo concurrente de múltiples reportes

### Importancia del Módulo

Facilita:
- Toma de decisiones administrativas
- Optimización del catálogo
- Mejora de servicios
- Supervisión del sistema
- Evaluación de desempeño

---

## 7. MÓDULO DE SEGURIDAD

### Descripción General
Protege toda la información del sistema. Fundamental debido a que maneja datos críticos.

### Datos Sensibles Protegidos

- ✓ Información de usuarios
- ✓ Contraseñas
- ✓ Archivos digitales
- ✓ Información académica
- ✓ Registros de préstamos

### Objetivos del Módulo

Garantizar:
- **Confidencialidad** → Solo usuarios autorizados acceden
- **Integridad** → Datos no alterados
- **Disponibilidad** → Servicio siempre accesible
- **Autenticidad** → Verificar identidad de usuarios

### Arquitectura Interna

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

### Funciones Principales

#### 1. Autenticación JWT
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

#### 2. Cifrado de Contraseñas
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

#### 3. Middleware de Autenticación
Permite:
- Proteger endpoints específicos
- Validar permisos por rol
- Bloquear accesos indebidos
- Registrar intentos fallidos

#### 4. Auditoría y Logging
El sistema registra:
- **Inicios de sesión** → Quién, cuándo, desde dónde
- **Errores de autenticación** → Intentos fallidos
- **Actividades críticas** → Cambios importantes
- **Acceso a datos sensibles** → Quién consultó qué

### Medidas de Seguridad Adicionales

- ✓ HTTPS/TLS para comunicación cifrada
- ✓ CORS configurado
- ✓ Rate limiting para prevenir fuerza bruta
- ✓ Validación de entrada en todos los endpoints
- ✓ Sanitización de queries SQL
- ✓ Passwords policy (complejidad mínima)
- ✓ Bloqueo temporal tras intentos fallidos

### Importancia del Módulo

Protege:
- Información personal de usuarios
- Integridad del sistema
- Archivos digitales
- Cumplimiento normativo
- Confianza de usuarios

---

## CONCLUSIÓN

La implementación modular utilizando **Go** permite desarrollar un Sistema de Gestión de Libros Electrónicos altamente organizado, eficiente y escalable.

### Beneficios de la Arquitectura Modular

#### Desarrollo
- ✓ Equipos pueden trabajar en módulos independientes
- ✓ Reduce conflictos en versionado
- ✓ Desarrollo más rápido y eficiente

#### Mantenimiento
- ✓ Módulos aislados facilitan identificación de problemas
- ✓ Actualizaciones sin afectar otros módulos
- ✓ Código más legible y documentado

#### Seguridad
- ✓ Control de acceso granular
- ✓ Aislamiento de responsabilidades
- ✓ Auditoría por módulo

#### Escalabilidad
- ✓ Fácil agregar nuevos módulos
- ✓ Preparado para crecimiento futuro
- ✓ Microservicios posibles

### Stack Tecnológico

| Componente | Tecnología | Función |
|-----------|------------|---------|
| API REST | Gin | Framework web ligero y rápido |
| ORM | GORM | Manejo de base de datos |
| BD | MySQL | Almacenamiento de datos |
| Autenticación | JWT + bcrypt | Seguridad y autenticación |
| Gestión de Archivos | FileSystem Go | Almacenamiento digital |
| Reportes | Exportación múltiples formatos | PDF, Excel, CSV |

### Próximos Pasos

1. **Configuración del entorno** de desarrollo
2. **Creación de estructura** de carpetas por módulo
3. **Implementación de modelos** (structs en Go)
4. **Desarrollo de Controllers** y Servicios
5. **Testing** unitario por módulo
6. **Integración y testing** de módulos
7. **Deployment** en producción

---

**Documento preparado para:** Sistema de Gestión de Libros Electrónicos  
**Lenguaje:** Go  
**Fecha de actualización:** 2026  
**Estado:** Documentación de diseño arquitectónico
