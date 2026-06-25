# Sistema de Gestión de Libros Electrónicos en Go

## Datos del Proyecto

**Nombre del proyecto:** Sistema de Gestión de Libros Electrónicos
**Lenguaje:** Go / Golang
**Base de datos:** PostgreSQL
**Framework web:** Gin
**ORM:** GORM
**Despliegue:** Railway
**Fecha:** Junio 2026

## Integrantes

* Integrante 1: Estructura base, almacenamiento e integración.
* Integrante 2: Gestión de libros, búsquedas y descargas.
* Integrante 3: Préstamos, reservas y reportes.
* Integrante 4: Usuarios, seguridad y documentación.

## Objetivo del Programa

Desarrollar un sistema de gestión de libros electrónicos en Go que permita administrar usuarios, libros, préstamos, reservas, descargas y reportes mediante servicios web conectados a una base de datos PostgreSQL.

El sistema busca simular el funcionamiento de una biblioteca digital, permitiendo que los usuarios puedan registrarse, iniciar sesión, consultar libros, gestionar préstamos y acceder a funcionalidades administrativas mediante una API REST.

## Introducción

El proyecto consiste en una aplicación backend desarrollada en Golang para la administración de libros electrónicos dentro de una biblioteca virtual.

La finalidad del sistema es automatizar procesos relacionados con la gestión de libros digitales, usuarios, préstamos, reservas y descargas. A través de servicios web, el aplicativo permite intercambiar información en formato JSON y responder a las solicitudes realizadas por los usuarios o administradores del sistema.

Este proyecto integra conocimientos adquiridos durante las unidades de la asignatura, como estructuras de datos, funciones, paquetes, métodos, encapsulación, interfaces, manejo de errores, conexión a base de datos, servicios web y serialización JSON.

## Tecnologías Utilizadas

* **Go:** lenguaje principal del proyecto.
* **Gin:** framework utilizado para crear servicios web.
* **GORM:** ORM utilizado para conectar Go con PostgreSQL.
* **PostgreSQL:** base de datos relacional.
* **Railway:** plataforma utilizada para el despliegue.
* **JSON:** formato utilizado para la serialización de datos.
* **GitHub:** repositorio utilizado para el trabajo colaborativo.
* **bcrypt:** paquete utilizado para cifrar contraseñas.

## Estructura General del Proyecto

```text
Autonomo1_go/
│
├── README.md
│
└── GO_CODE_GESTION/
    ├── Main.go
    ├── go.mod
    ├── go.sum
    │
    └── internal/
        ├── config/
        ├── database/
        ├── handlers/
        ├── models/
        ├── repositories/
        ├── routes/
        └── services/
```

## Módulos Principales

### Usuarios y Seguridad

Este módulo permite administrar los usuarios del sistema. Incluye funcionalidades para registrar, consultar, actualizar y eliminar usuarios. También permite validar credenciales mediante el inicio de sesión.

Funcionalidades principales:

* Registro de usuarios.
* Inicio de sesión.
* Consulta de usuarios.
* Actualización de datos.
* Eliminación de usuarios.
* Validación de campos obligatorios.
* Validación de correo duplicado.
* Cifrado de contraseñas mediante bcrypt.

### Libros

Este módulo permite gestionar el catálogo de libros electrónicos dentro de la biblioteca digital.

Funcionalidades principales:

* Registrar libros.
* Consultar libros.
* Buscar libros.
* Actualizar información.
* Eliminar registros.
* Controlar disponibilidad.

### Préstamos y Reservas

Este módulo permite controlar el préstamo de libros electrónicos y las reservas realizadas por los usuarios.

Funcionalidades principales:

* Crear préstamos.
* Registrar devoluciones.
* Consultar historial.
* Crear reservas.
* Cancelar reservas.
* Consultar reservas pendientes.

### Descargas

Este módulo permite registrar y controlar las descargas de libros electrónicos realizadas por los usuarios.

Funcionalidades principales:

* Registrar descargas.
* Consultar historial de descargas.
* Validar disponibilidad del libro.
* Relacionar descargas con usuarios y libros.

### Reportes

Este módulo permite generar información estadística sobre el uso del sistema.

Funcionalidades principales:

* Consultar préstamos activos.
* Consultar reservas pendientes.
* Revisar libros más utilizados.
* Obtener información general del sistema.

## Servicios Web Implementados

El sistema expone sus funcionalidades mediante servicios web REST. Las respuestas se serializan en formato JSON.

### Rutas generales

| Método | Endpoint | Descripción                     |
| ------ | -------- | ------------------------------- |
| GET    | /        | Ruta principal del sistema      |
| GET    | /health  | Verifica el estado del servidor |

### Usuarios y Seguridad

| Método | Endpoint   | Descripción                   |
| ------ | ---------- | ----------------------------- |
| GET    | /users     | Lista todos los usuarios      |
| GET    | /users/:id | Consulta un usuario por ID    |
| POST   | /users     | Crea un nuevo usuario         |
| PUT    | /users/:id | Actualiza un usuario          |
| DELETE | /users/:id | Elimina un usuario            |
| POST   | /register  | Registra un usuario           |
| POST   | /login     | Valida credenciales de acceso |

## Ejemplo de JSON para Registro

```json
{
  "name": "Martin Gomez",
  "email": "martin@example.com",
  "password": "123456",
  "role": "lector"
}
```

## Ejemplo de JSON para Login

```json
{
  "email": "martin@example.com",
  "password": "123456"
}
```

## Instalación y Ejecución

### 1. Clonar el repositorio

```bash
git clone https://github.com/wixy122018-byte/Autonomo1_go.git
```

### 2. Ingresar al proyecto

```bash
cd Autonomo1_go/GO_CODE_GESTION
```

### 3. Instalar dependencias

```bash
go mod tidy
```

### 4. Ejecutar el servidor

```bash
go run Main.go
```

### 5. Verificar el servidor

```text
http://localhost:8080/health
```

## Base de Datos

El sistema utiliza PostgreSQL como base de datos. La conexión se configura mediante variables de entorno.

Tablas principales:

* users
* books
* loans
* reservations
* downloads

## Aporte del Responsable D: Usuarios, Seguridad y Documentación

Dentro del proyecto se desarrolló el módulo de usuarios y seguridad, implementando servicios web para registrar, consultar, actualizar y eliminar usuarios.

También se agregó la lógica de validación de credenciales para el inicio de sesión, verificación de campos obligatorios, validación de correo duplicado, manejo de errores y cifrado de contraseñas mediante bcrypt.

Además, se elaboró la documentación base del proyecto mediante este archivo README, explicando el objetivo del sistema, tecnologías utilizadas, estructura, endpoints principales e instrucciones de ejecución.

## Relación con los Temas de la Asignatura

El proyecto integra conocimientos de las cuatro unidades de la materia:

* Sintaxis básica de Go.
* Funciones.
* Paquetes.
* Structs.
* Métodos.
* Encapsulación.
* Interfaces.
* Manejo de errores.
* Base de datos.
* Servicios web.
* Serialización JSON.
* Organización modular del código.

## Visualización del Futuro

Este sistema puede evolucionar hacia una biblioteca digital inteligente, en la cual los usuarios reciban recomendaciones personalizadas de libros según sus intereses, historial de préstamos y categorías consultadas.

También podría integrarse con inteligencia artificial para sugerir lecturas, generar resúmenes automáticos, clasificar libros por temática y mejorar la experiencia de los usuarios dentro de la plataforma.

En el futuro, una biblioteca digital de este tipo podría ser utilizada en instituciones educativas, universidades, bibliotecas públicas y plataformas de aprendizaje virtual, facilitando el acceso al conocimiento desde cualquier lugar.

## Limitaciones

Aunque el sistema presenta una estructura funcional, todavía puede mejorar en aspectos como:

* Implementación completa de autenticación con JWT.
* Protección de rutas privadas.
* Validaciones más avanzadas.
* Interfaz gráfica para usuarios finales.
* Pruebas unitarias automatizadas.
* Mejora de reportes estadísticos.

## Conclusión

El Sistema de Gestión de Libros Electrónicos en Go permite aplicar los conocimientos adquiridos durante la asignatura mediante el desarrollo de un backend funcional basado en servicios web.

El proyecto integra programación en Go, estructuras de datos, manejo de errores, interfaces, conexión a base de datos, serialización JSON y documentación técnica. Además, representa una solución práctica para la administración de una biblioteca digital moderna, con posibilidades de crecimiento hacia tecnologías futuras como inteligencia artificial, automatización y recomendación personalizada de contenidos.
