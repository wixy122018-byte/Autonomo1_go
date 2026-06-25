# Informe Final

## Sistema de Gestión de Libros Electrónicos en Go

## 1. Introducción

El presente proyecto corresponde al desarrollo de un Sistema de Gestión de Libros Electrónicos utilizando el lenguaje de programación Go. La finalidad del sistema es simular el funcionamiento de una biblioteca digital, permitiendo administrar usuarios, libros, préstamos, reservas, descargas y reportes mediante servicios web.

El sistema fue desarrollado como parte de la evaluación final de la asignatura, integrando los conocimientos adquiridos durante las diferentes unidades de estudio. Entre los principales temas aplicados se encuentran la sintaxis de Go, el uso de funciones, paquetes, estructuras, métodos, interfaces, manejo de errores, conexión a base de datos, servicios web y serialización de datos en formato JSON.

El proyecto se desarrolló de manera colaborativa mediante GitHub, utilizando ramas de trabajo para organizar los aportes de cada integrante. Además, se utilizó PostgreSQL como base de datos, Gin como framework web, GORM como herramienta ORM y Railway como plataforma de despliegue.

## 2. Objetivo del Proyecto

Desarrollar un sistema backend funcional en Go para la gestión de libros electrónicos, permitiendo el manejo de usuarios, libros, préstamos, reservas, descargas y reportes a través de servicios web conectados a una base de datos PostgreSQL.

## 3. Desarrollo del Proyecto

El desarrollo del sistema se organizó en diferentes módulos, cada uno relacionado con una funcionalidad principal del aplicativo. La estructura del proyecto permite separar responsabilidades mediante carpetas internas como modelos, rutas, controladores, servicios, repositorios, configuración y conexión a base de datos.

La aplicación utiliza una arquitectura organizada por capas. Los modelos representan las entidades principales del sistema, los repositorios permiten interactuar con la base de datos, los servicios contienen la lógica de negocio, los handlers reciben las peticiones HTTP y las rutas conectan los endpoints con cada funcionalidad.

Durante el desarrollo se implementaron servicios web tipo REST, los cuales permiten enviar y recibir información en formato JSON. Esto permite que el sistema pueda ser consumido por otras aplicaciones, interfaces web o herramientas de prueba como Postman.

## 4. Tecnologías Utilizadas

Las tecnologías utilizadas en el proyecto fueron:

* Go / Golang como lenguaje principal de programación.
* Gin como framework para la creación de servicios web.
* GORM como ORM para la conexión entre Go y PostgreSQL.
* PostgreSQL como sistema de base de datos.
* Railway como plataforma de despliegue.
* GitHub como repositorio colaborativo.
* JSON como formato de intercambio de información.
* bcrypt para el cifrado de contraseñas.
* JWT para el manejo de autenticación y seguridad.

## 5. Módulos del Sistema

### 5.1 Módulo de Usuarios y Seguridad

Este módulo permite gestionar los usuarios registrados en el sistema. Incluye funcionalidades como registro, inicio de sesión, consulta de perfil y consulta de usuarios según permisos.

También se aplican validaciones para evitar campos vacíos, correos duplicados y credenciales incorrectas. Las contraseñas se almacenan de forma cifrada mediante bcrypt, aumentando la seguridad del sistema.

Además, el sistema implementa autenticación mediante JWT, permitiendo proteger rutas privadas y validar que solo usuarios autorizados puedan acceder a ciertas funcionalidades.

### 5.2 Módulo de Libros

El módulo de libros permite administrar el catálogo de libros electrónicos. Incluye información como título, autor, categoría, editorial, año, descripción, formato, ISBN, disponibilidad y ruta del archivo.

Este módulo es fundamental porque representa el recurso principal de la biblioteca digital.

### 5.3 Módulo de Préstamos

El módulo de préstamos permite registrar cuándo un usuario solicita un libro electrónico. También permite manejar fechas de inicio, fechas de vencimiento, fechas de devolución y estado del préstamo.

### 5.4 Módulo de Reservas

El módulo de reservas permite que los usuarios puedan reservar libros electrónicos. Cada reserva se relaciona con un usuario y un libro, permitiendo controlar su fecha y estado.

### 5.5 Módulo de Descargas

El módulo de descargas permite registrar el acceso de los usuarios a los libros electrónicos disponibles. Esto permite llevar un control del uso del sistema y del historial de descargas.

### 5.6 Módulo de Reportes

El módulo de reportes permite generar información útil sobre el funcionamiento del sistema, como préstamos activos, libros disponibles, reservas pendientes y datos generales de uso.

## 6. Servicios Web Implementados

El sistema expone sus funcionalidades mediante servicios web REST. Algunos de los endpoints principales son:

| Método | Endpoint         | Descripción                                |
| ------ | ---------------- | ------------------------------------------ |
| GET    | /                | Ruta principal del sistema                 |
| GET    | /health          | Verifica el estado del servidor            |
| POST   | /register        | Registra un nuevo usuario                  |
| POST   | /login           | Permite iniciar sesión                     |
| GET    | /profile         | Consulta el perfil del usuario autenticado |
| GET    | /users           | Lista usuarios registrados                 |
| GET    | /users/:id       | Consulta un usuario por ID                 |
| POST   | /books           | Registra libros electrónicos               |
| GET    | /books           | Consulta libros electrónicos               |
| POST   | /loans           | Registra préstamos                         |
| POST   | /reservations    | Registra reservas                          |
| GET    | /reports/general | Consulta reportes generales                |

Estos servicios utilizan JSON como formato de respuesta, cumpliendo con el requisito de serialización de datos.

## 7. Resultados Obtenidos

Como resultado del proyecto se obtuvo un backend funcional desarrollado en Go, con conexión a una base de datos PostgreSQL y servicios web organizados mediante Gin.

El sistema permite gestionar diferentes entidades relacionadas con una biblioteca digital, como usuarios, libros, préstamos, reservas y descargas. Además, se implementó seguridad mediante cifrado de contraseñas y autenticación con JWT.

También se logró trabajar de manera colaborativa mediante GitHub, utilizando ramas de desarrollo para integrar los aportes de cada integrante. Se realizaron pruebas de compilación mediante el comando `go test ./...`, verificando que el proyecto compile correctamente y que no existan errores de estructura.

## 8. Implicaciones del Proyecto

El sistema desarrollado puede ser aplicado en instituciones educativas, bibliotecas universitarias o plataformas de aprendizaje digital que requieran organizar y controlar el acceso a libros electrónicos.

Además, permite comprender cómo las tecnologías actuales pueden facilitar la administración de recursos digitales, mejorar el acceso a la información y optimizar procesos que antes podían realizarse de forma manual.

Desde el punto de vista académico, el proyecto permitió aplicar conocimientos de programación, bases de datos, servicios web, seguridad informática y trabajo colaborativo.

## 9. Limitaciones del Proyecto

Aunque el sistema presenta una estructura funcional, todavía existen aspectos que pueden mejorarse en futuras versiones, como:

* Implementar una interfaz gráfica para usuarios finales.
* Mejorar los reportes estadísticos.
* Agregar pruebas unitarias más completas.
* Implementar recuperación de contraseña.
* Mejorar el control de roles y permisos.
* Agregar validaciones más avanzadas en todos los módulos.
* Integrar recomendaciones inteligentes de libros.
* Mejorar la documentación técnica del API.

## 10. Visualización del Futuro

En el futuro, este sistema podría evolucionar hacia una biblioteca digital inteligente. Mediante el uso de inteligencia artificial, el sistema podría recomendar libros según los intereses, historial de préstamos y comportamiento de cada usuario.

También podría incluir funciones como resúmenes automáticos, clasificación inteligente de libros, búsqueda avanzada por temas, análisis de lectura y generación de reportes personalizados.

Este tipo de sistema podría ser útil en universidades, colegios, bibliotecas públicas y plataformas virtuales de aprendizaje, facilitando el acceso al conocimiento desde cualquier lugar y reduciendo la dependencia de procesos manuales.

## 11. Conclusión

El desarrollo del Sistema de Gestión de Libros Electrónicos permitió aplicar de forma práctica los conocimientos adquiridos durante la asignatura. El proyecto integró conceptos fundamentales de Go, estructuras, funciones, métodos, paquetes, manejo de errores, conexión a base de datos, servicios web, JSON y seguridad.

Además, el trabajo permitió fortalecer habilidades de organización, solución de errores, uso de GitHub y colaboración en equipo. A pesar de las dificultades encontradas durante la integración de ramas y resolución de conflictos, se logró consolidar un sistema funcional y correctamente estructurado.

En conclusión, el proyecto representa una solución tecnológica aplicable al contexto educativo y bibliotecario, con posibilidades de crecimiento hacia nuevas tecnologías como inteligencia artificial, automatización y análisis de datos.
