# ARQUITECTURA DEL SISTEMA DE GESTIÓN DE LIBROS ELECTRÓNICOS

El Sistema de Gestión de Libros Electrónicos será desarrollado utilizando el lenguaje de programación Go, implementando una arquitectura por capas que permita mantener el sistema organizado, seguro y escalable.

## Flujo General de la Arquitectura

```
Cliente
   ↓
API REST
   ↓
Controladores
   ↓
Servicios
   ↓
Repositorios
   ↓
Base de Datos
```

Este tipo de arquitectura permite dividir el sistema en componentes organizados, donde cada capa tiene responsabilidades específicas. Gracias a esta separación, el sistema será más fácil de desarrollar y mantener.

La arquitectura también ayuda a mejorar la seguridad, reutilización del código y organización interna del proyecto, permitiendo que cada componente trabaje de forma independiente pero conectada con los demás.

---

## ARQUITECTURA GENERAL EN GO

En Go el proyecto se organizará mediante paquetes y carpetas estructuradas para mantener un código limpio y modular.

### Estructura del Proyecto

```
/proyecto-libros
│
├── main.go
├── controllers/
├── services/
├── repositories/
├── models/
├── routes/
├── database/
├── middleware/
├── utils/
└── config/
```

Cada carpeta tendrá una función específica dentro de la arquitectura.

---

## 1. CAPA CLIENTE

La capa cliente corresponde a la interfaz visual del sistema, es decir, la parte con la que interactúan directamente los usuarios. Esta capa es responsable de mostrar información y permitir que las personas realicen acciones dentro de la aplicación.

Aquí se encuentran todas las pantallas, formularios, botones y menús que permiten navegar dentro del sistema.

### Tipos de Usuarios

El sistema tendrá distintos tipos de usuarios con permisos y funcionalidades diferentes:

#### Administrador
- Tiene control total del sistema
- Gestiona usuarios
- Administra libros
- Visualiza reportes
- Configura parámetros del sistema

#### Bibliotecario
- Se encarga de administrar el contenido bibliográfico
- Registra libros
- Gestiona préstamos
- Controla reservas
- Actualiza información

#### Lector o Estudiante
- Es el usuario final que utiliza la biblioteca digital
- Busca libros
- Solicita préstamos
- Reserva libros
- Consulta historial

### Funciones Principales de la Capa Cliente

#### Inicio de Sesión
Permite autenticar usuarios mediante:
- Correo electrónico
- Contraseña

El sistema validará la información antes de enviarla al servidor.

#### Visualización de Libros
Los usuarios podrán:
- Ver catálogo digital
- Consultar autores
- Leer descripciones
- Revisar disponibilidad

#### Búsqueda y Filtros
La interfaz permitirá buscar libros mediante:
- Título
- Autor
- Categoría
- Palabras clave

#### Solicitudes y Reservas
Los usuarios podrán:
- Solicitar préstamos
- Reservar libros ocupados
- Consultar historial

#### Panel Administrativo
Los administradores tendrán acceso a:
- Gestión de usuarios
- Reportes
- Estadísticas
- Configuraciones

### Tecnologías Utilizadas en la Capa Cliente

La interfaz será desarrollada utilizando:
- **HTML5**
- **CSS3**
- **JavaScript**

Estas tecnologías permitirán construir una plataforma:
- Dinámica
- Interactiva
- Responsiva
- Compatible con dispositivos modernos

### Funcionamiento de la Capa Cliente

Cuando el usuario realiza una acción:

1. La interfaz captura la información
2. Se genera una solicitud HTTP
3. La solicitud es enviada hacia la API REST desarrollada en Go
4. El sistema procesa la operación
5. La respuesta regresa al cliente

La capa cliente únicamente se encarga de interacción visual y no procesa lógica compleja del sistema.

### Ejemplo Práctico

Si un usuario busca un libro:
1. Escribe el nombre del libro
2. La interfaz envía una petición GET
3. La API procesa la solicitud
4. Los resultados se muestran en pantalla

---

## 2. CAPA API REST

### ¿Qué es una API REST?

Una API REST es un conjunto de rutas o endpoints que permiten enviar y recibir información. Cada funcionalidad del sistema tendrá una ruta específica.

### Gin Framework

Gin es un framework para Go que permite construir APIs rápidas y eficientes. La API REST actuará como puente entre la interfaz y el backend.

La API permite que diferentes aplicaciones puedan comunicarse con el sistema, incluyendo:
- Aplicaciones web
- Aplicaciones móviles
- Servicios externos

### Ejemplos de Endpoints

```
POST   /login
GET    /libros
POST   /prestamos
PUT    /usuarios/:id
DELETE /reservas/:id
```

Cada endpoint representa una funcionalidad concreta del sistema.

### Funciones Principales de la API REST

#### Recepción de Solicitudes

La API recibe solicitudes enviadas desde el cliente.

Ejemplos:
- Iniciar sesión
- Registrar usuarios
- Buscar libros
- Solicitar préstamos

#### Procesamiento de Métodos HTTP

| Método | Descripción | Ejemplo |
|--------|-------------|---------|
| **GET** | Obtiene información | `GET /libros` - Sirve para consultar libros registrados |
| **POST** | Envía información nueva | `POST /usuarios` - Sirve para registrar nuevos usuarios |
| **PUT** | Actualiza información existente | `PUT /libros/15` - Actualiza datos de un libro |
| **DELETE** | Elimina información | `DELETE /reservas/10` - Elimina una reserva registrada |

### Respuestas de la API

La información será enviada generalmente en formato JSON.

Ejemplo:
```json
{
  "titulo": "Redes Informáticas",
  "autor": "Autor Ejemplo",
  "categoria": "Tecnología",
  "disponible": true
}
```

### Importancia de la API REST

La API permite:
- Separar frontend y backend
- Facilitar integración futura
- Mejorar organización
- Centralizar servicios
- Permitir escalabilidad

### Implementación en Go

En Go las rutas se configurarán mediante Gin.

Ejemplo conceptual:
```go
router.GET("/libros", controllers.ObtenerLibros)
```

Esto significa que cuando el usuario consulte `/libros`, la solicitud será enviada al controlador correspondiente.

---

## 3. CAPA DE CONTROLADORES

Los controladores son responsables de recibir las solicitudes provenientes de la API REST y coordinar el flujo interno del sistema. Funcionan como intermediarios entre la API y la lógica de negocio.

### Funciones Principales de los Controladores

#### Recepción de Solicitudes

Los controladores reciben peticiones enviadas desde la API.

Por ejemplo:
- Solicitudes de login
- Registro de libros
- Préstamos
- Consultas

#### Validación Inicial de Datos

Antes de procesar una operación, el controlador verifica:
- Campos obligatorios
- Tipos de datos
- Formatos correctos
- Parámetros válidos

Ejemplo:
- Validar que el correo tenga formato correcto
- Confirmar que la contraseña no esté vacía

#### Ejemplo en Go

```go
func ObtenerLibros(c *gin.Context) {
    libros := services.ListarLibros()
    c.JSON(200, libros)
}
```

#### Envío hacia Servicios

Una vez validada la información, el controlador envía los datos hacia la capa de servicios.

### Ejemplo Práctico: Inicio de Sesión

Cuando un usuario inicia sesión:

1. El cliente envía correo y contraseña
2. La API recibe la solicitud
3. El controlador valida los datos
4. Envía la información al servicio de autenticación
5. Espera respuesta
6. Devuelve resultado al usuario

### Importancia de los Controladores

Permiten:
- Organizar solicitudes
- Validar entradas
- Controlar respuestas
- Mantener código limpio

---

## 4. CAPA DE SERVICIOS

La capa de servicios contiene toda la lógica de negocio del sistema. Es considerada la parte más importante porque aquí se ejecutan las reglas y procesos internos que permiten el funcionamiento de la aplicación.

Los servicios toman decisiones y controlan cómo debe comportarse el sistema.

### Funciones Principales de los Servicios

#### Gestión de Usuarios
Permite:
- Registrar usuarios
- Validar credenciales
- Gestionar roles
- Controlar permisos

#### Gestión de Libros
Se encarga de:
- Registrar libros
- Actualizar información
- Verificar disponibilidad
- Clasificar categorías

#### Gestión de Préstamos
Controla:
- Solicitudes de préstamo
- Fechas de devolución
- Disponibilidad
- Historiales

#### Gestión de Reservas
Permite:
- Registrar reservas
- Notificar disponibilidad
- Gestionar listas de espera

#### Seguridad
Implementa:
- Encriptación de contraseñas
- Validación de sesiones
- Autenticación
- Protección de información

### Ejemplo Detallado de Funcionamiento

Cuando un estudiante solicita un libro:

1. El servicio verifica si el usuario existe
2. Comprueba que el libro esté disponible
3. Revisa límite de préstamos
4. Registra la solicitud
5. Actualiza disponibilidad
6. Genera respuesta final

Toda esta lógica ocurre dentro de la capa de servicios.

### Importancia de la Capa de Servicios

Permite:
- Centralizar lógica del negocio
- Reutilizar procesos
- Mantener organización
- Facilitar mantenimiento

### Ejemplo en Go

```go
func ListarLibros() []models.Libro {
    return repositories.ObtenerLibros()
}
```

---

## 5. CAPA DE REPOSITORIOS

La capa de repositorios es responsable de comunicarse directamente con la base de datos. Su función es ejecutar operaciones relacionadas con almacenamiento y recuperación de información.

### Funciones Principales de los Repositorios

| Operación | Descripción |
|-----------|-------------|
| **Consultas de Datos** | Permite: Buscar usuarios, Consultar libros, Obtener préstamos, Revisar reservas |
| **Inserción de Información** | Registra: Nuevos usuarios, Nuevos libros, Préstamos, Reportes |
| **Actualización de Registros** | Permite modificar: Datos de usuarios, Información de libros, Estados de préstamos |
| **Eliminación de Información** | Elimina: Reservas, Registros innecesarios, Datos obsoletos |

### Ejemplo Práctico

Cuando se registra un préstamo:

1. El servicio envía la solicitud
2. El repositorio genera la consulta SQL
3. La información se guarda en la base de datos

### Importancia de los Repositorios

Permiten:
- Separar lógica y datos
- Mejorar organización
- Facilitar mantenimiento
- Optimizar consultas

### Ejemplo en Go

```go
func ObtenerLibros() []models.Libro {
    var libros []models.Libro
    database.DB.Find(&libros)
    return libros
}
```

---

## 6. CAPA DE BASE DE DATOS

La base de datos es la capa encargada de almacenar permanentemente toda la información del sistema.

### Sistema de Base de Datos: MySQL

MySQL permitirá administrar la información de forma segura, rápida y estructurada.

### Información Almacenada

| Categoría | Contenido |
|-----------|-----------|
| **Usuarios** | Nombres, Correos, Contraseñas cifradas, Roles |
| **Libros** | Títulos, Autores, Categorías, Archivos digitales |
| **Préstamos** | Usuarios, Fechas, Estados |
| **Reservas** | Fechas, Disponibilidad, Historiales |
| **Reportes y Auditorías** | Actividades del sistema, Logs, Estadísticas |

### Importancia de la Base de Datos

Garantiza:
- Persistencia de información
- Seguridad de datos
- Integridad
- Disponibilidad
- Consultas rápidas

### Ejemplo de Conexión

```go
dsn := "usuario:password@tcp(localhost:3306)/biblioteca"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
```

---

## FLUJO COMPLETO DEL SISTEMA

```
1. El usuario interactúa con la interfaz web
                    ↓
2. El cliente envía solicitud HTTP a la API REST
                    ↓
3. Gin recibe la petición
                    ↓
4. El controlador procesa y valida datos
                    ↓
5. El servicio ejecuta la lógica de negocio
                    ↓
6. El repositorio consulta MySQL
                    ↓
7. La base de datos devuelve resultados
                    ↓
8. El sistema genera respuesta JSON
                    ↓
9. La respuesta retorna al cliente
```

---

## CONCLUSIÓN

La arquitectura basada en Cliente, API REST, Controladores, Servicios, Repositorios y Base de Datos permite desarrollar un Sistema de Gestión de Libros Electrónicos organizado, seguro y escalable.

La separación por capas facilita la administración del código, mejora el mantenimiento y permite que el sistema pueda crecer en el futuro incorporando nuevas funcionalidades sin afectar su estructura fundamental.

Este modelo de arquitectura es flexible, mantenible y sigue las mejores prácticas de desarrollo de software moderno.
