# Informe técnico: servicios web de libros, búsquedas y descargas

## Datos generales

- **Proyecto:** Sistema de Gestión de Libros Electrónicos
- **Integrante responsable:** Leonardo Sánchez
- **Módulo:** Catálogo de libros, búsquedas y descargas
- **Tecnologías:** Go, Gin, GORM, PostgreSQL y JSON
- **Rama:** `Leonardo_Sanchez`
- **Fecha:** 27 de junio de 2026

## 1. Introducción

Los servicios web permiten que diferentes aplicaciones intercambien información mediante solicitudes HTTP. En este proyecto se implementó una API REST que utiliza JSON para enviar y recibir datos.

La parte desarrollada corresponde al catálogo de libros electrónicos. El módulo permite registrar libros, consultarlos, buscarlos mediante filtros, actualizarlos, desactivarlos y registrar las descargas realizadas por los usuarios.

## 2. Objetivo

Implementar servicios web REST funcionales para administrar libros electrónicos, aplicar filtros de búsqueda y registrar descargas almacenadas en PostgreSQL.

### Objetivos específicos

- Implementar el CRUD de libros.
- Buscar libros por título, autor, categoría y disponibilidad.
- Validar los campos obligatorios.
- Comprobar que el libro exista antes de actualizarlo o desactivarlo.
- Verificar la disponibilidad antes de registrar una descarga.
- Devolver códigos HTTP y errores claros en JSON.
- Organizar el código mediante modelos, repositorios, servicios, handlers y rutas.

## 3. Arquitectura del módulo

El módulo está organizado por capas para separar responsabilidades:

| Capa | Archivos | Responsabilidad |
|---|---|---|
| Modelos | `book.go`, `download.go` | Definen los structs, constructores, validaciones y métodos. |
| Repositorios | `book_repository.go`, `download_repository.go` | Guardan y consultan información mediante GORM. |
| Servicios | `book_service.go`, `download_service.go` | Contienen las reglas de negocio, filtros y validaciones. |
| Handlers | `book_handler.go`, `download_handler.go` | Reciben solicitudes HTTP y devuelven respuestas JSON. |
| Rutas | `routes.go` | Registra los endpoints y los relaciona con los handlers. |

### Flujo de una solicitud

1. El cliente envía una solicitud HTTP.
2. Gin dirige la solicitud al handler correspondiente.
3. El handler interpreta los parámetros o el cuerpo JSON.
4. El servicio aplica las validaciones y reglas de negocio.
5. El repositorio consulta o modifica PostgreSQL mediante GORM.
6. El handler devuelve una respuesta JSON y un código HTTP.

## 4. Servicios web implementados

La URL utilizada localmente es:

```text
http://localhost:8081
```

Los endpoints versionados utilizan el prefijo `/api/v1`.

| Método | Endpoint | Funcionalidad | Respuesta exitosa |
|---|---|---|---|
| `POST` | `/api/v1/books` | Registrar un libro | `201 Created` |
| `GET` | `/api/v1/books` | Listar libros activos | `200 OK` |
| `GET` | `/api/v1/books/:id` | Consultar un libro por ID | `200 OK` |
| `PUT` | `/api/v1/books/:id` | Actualizar un libro | `200 OK` |
| `DELETE` | `/api/v1/books/:id` | Desactivar un libro | `204 No Content` |
| `GET` | `/api/v1/books/search` | Buscar y filtrar libros | `200 OK` |
| `POST` | `/api/v1/downloads` | Registrar una descarga | `201 Created` |
| `GET` | `/api/v1/downloads/history` | Consultar historial | `200 OK` |

Con estos endpoints, esta parte aporta ocho servicios web de diferentes funcionalidades al proyecto.

## 5. Gestión de libros

El struct `Book` representa los libros del catálogo. Incluye título, autor, categoría, editorial, año, descripción, formato, ISBN, disponibilidad, estado activo, ruta del archivo y fecha de creación.

El constructor `NewBook()` limpia los espacios y ejecuta la validación antes de guardar un libro.

El método `IsAvailable()` encapsula la regla de disponibilidad:

```go
func (b Book) IsAvailable() bool {
    return b.Active && b.Available
}
```

Un libro solamente se considera disponible cuando está activo y su campo `available` es verdadero.

### Operaciones CRUD

- **Crear:** valida los datos y guarda el libro en PostgreSQL.
- **Listar:** devuelve únicamente libros activos.
- **Consultar:** busca un libro activo mediante su ID.
- **Actualizar:** valida los nuevos datos y conserva el ID y la fecha de creación.
- **Eliminar:** realiza un borrado lógico cambiando `active` y `available` a `false`.

El borrado lógico evita eliminar físicamente los registros y permite conservar trazabilidad.

## 6. Búsqueda de libros

El endpoint de búsqueda permite filtrar por:

| Criterio | Parámetro | Ejemplo |
|---|---|---|
| Título | `title` | `/api/v1/books/search?title=go` |
| Autor | `author` | `/api/v1/books/search?author=Ana` |
| Categoría | `category` | `/api/v1/books/search?category=Informatica` |
| Disponibilidad | `available` | `/api/v1/books/search?available=true` |

La función pura `FilterBooks()` recibe un slice de libros y una estructura de filtros. Utiliza:

- Un `map` para organizar los criterios.
- Un bucle `for` para recorrer los libros.
- Condicionales para comprobar cada filtro.
- Comparaciones sin distinguir mayúsculas y minúsculas.
- El método `IsAvailable()` para validar disponibilidad.

La función no modifica la lista original.

## 7. Registro de descargas

Para registrar una descarga se envía:

```json
{
  "user_id": 23,
  "book_id": 1
}
```

El servicio realiza los siguientes pasos:

1. Busca el libro mediante `BookRepository`.
2. Devuelve error si el libro no existe.
3. Comprueba `Book.IsAvailable()`.
4. Devuelve error si el libro no está disponible.
5. Crea el registro con la fecha actual.
6. Guarda la descarga en PostgreSQL.

El historial se ordena por fecha descendente para mostrar primero las descargas recientes.

## 8. Serialización JSON

Gin deserializa las solicitudes mediante `ShouldBindJSON()` y serializa las respuestas mediante `Context.JSON()`.

Ejemplo para crear un libro:

```http
POST http://localhost:8081/api/v1/books
Content-Type: application/json
```

```json
{
  "title": "Libro de Prueba",
  "author": "Leonardo Sanchez",
  "category": "Informatica",
  "editorial": "Editorial Demo",
  "year": 2026,
  "description": "Libro creado desde Postman",
  "format": "PDF",
  "isbn": "978-1-999",
  "available": true,
  "file_path": "/books/libro-prueba.pdf"
}
```

## 9. Validaciones y manejo de errores

| Situación | Código HTTP | Respuesta |
|---|---|---|
| JSON incorrecto | `400` | `{"error":"json invalido"}` |
| Título vacío | `400` | `{"error":"el titulo es obligatorio"}` |
| Autor vacío | `400` | `{"error":"el autor es obligatorio"}` |
| Categoría vacía | `400` | `{"error":"la categoria es obligatoria"}` |
| ID incorrecto | `400` | `{"error":"id invalido"}` |
| Libro inexistente | `404` | `{"error":"libro no encontrado"}` |
| Libro no disponible | `400` | `{"error":"el libro no esta disponible para descarga"}` |

Los errores propios `ErrBookNotFound` y `ErrBookUnavailable` permiten identificar claramente cada situación.

## 10. Interfaces y encapsulación

La persistencia se abstrae mediante las interfaces:

```go
type BookRepository interface {
    Create(book models.Book) (models.Book, error)
    FindAll() ([]models.Book, error)
    FindByID(id uint) (models.Book, error)
    Update(id uint, book models.Book) (models.Book, error)
    Deactivate(id uint) error
}
```

```go
type DownloadRepository interface {
    Create(download models.Download) (models.Download, error)
    History() ([]models.Download, error)
}
```

Los servicios dependen de estas interfaces y no directamente de GORM. Esto reduce el acoplamiento y permite utilizar otra implementación de almacenamiento o repositorios simulados para pruebas.

## 11. Persistencia en PostgreSQL

Las principales tablas del módulo son:

| Tabla | Función |
|---|---|
| `books` | Almacena el catálogo y el estado de los libros. |
| `downloads` | Registra el usuario, el libro y la fecha de descarga. |

Consultas utilizadas para comprobar los datos:

```sql
SELECT id, title, author, category, available
FROM books
ORDER BY id;
```

```sql
SELECT downloads.id,
       users.name AS usuario,
       books.title AS libro,
       downloads.download_date
FROM downloads
JOIN users ON users.id = downloads.user_id
JOIN books ON books.id = downloads.book_id
ORDER BY downloads.id;
```

La aplicación utiliza GORM y `AutoMigrate` para crear o actualizar las tablas al iniciar.

## 12. Pruebas realizadas

Las pruebas incluyeron:

- Ejecución de `go test ./...`.
- Creación de libros desde Postman.
- Listado y consulta por ID.
- Búsqueda por categoría y otros filtros.
- Actualización y desactivación de libros.
- Registro de descargas.
- Consulta del historial.
- Comprobación de datos mediante pgAdmin.

El comando de pruebas terminó correctamente:

```text
ok  sistema-libros-electronicos/internal/services
```

## 13. Contenidos de Go evidenciados

| Contenido | Evidencia |
|---|---|
| Structs | `Book`, `Download`, `BookInput`, `BookFilters` |
| Slices | Listas de libros usadas en filtros |
| Array | `FixedCategories` |
| Maps | `filterMap`, `CountBooksByCategory` |
| Condicionales | Validaciones y disponibilidad |
| Bucles | Recorrido y conteo de libros |
| Funciones puras | `FilterBooks`, `CountBooksByCategory` |
| Métodos | `Book.IsAvailable()`, `Book.Validate()` |
| Constructores | `NewBook`, `NewDownload` |
| Errores | `ErrBookNotFound`, `ErrBookUnavailable` |
| Interfaces | `BookRepository`, `DownloadRepository` |
| Polimorfismo | Servicios independientes de la implementación de persistencia |

## 14. Relación con la rúbrica

| Criterio | Evidencia en el módulo |
|---|---|
| Desarrollo del tema | Módulo enfocado en la gestión de libros electrónicos. |
| Análisis de software | Separación en modelos, repositorios, servicios, handlers y rutas. |
| Resultados | CRUD, búsquedas, filtros y descargas funcionales. |
| Implicaciones y limitaciones | Se identifican mejoras pendientes y usos futuros. |
| Toma de decisiones | Selección razonada de Go, Gin, GORM, PostgreSQL y JSON. |
| Argumento principal | Los servicios web centralizan y facilitan la gestión del catálogo. |
| Encapsulación | Reglas protegidas en modelos y servicios. |
| Errores e interfaces | Errores propios, códigos HTTP e interfaces de repositorio. |
| Comentarios | Funciones complejas de filtros, conteo y descarga documentadas. |

## 15. Limitaciones y trabajo futuro

### Limitaciones actuales

- Se registra la ruta del archivo, pero aún no se carga físicamente el PDF o EPUB.
- La validación de ISBN comprueba unicidad, pero no su algoritmo formal.
- No existe paginación para catálogos grandes.
- Los filtros se aplican en memoria después de consultar los libros activos.

### Mejoras futuras

- Almacenamiento de archivos en la nube.
- Lectura de libros en línea.
- Búsqueda de texto completo.
- Recomendaciones personalizadas.
- Estadísticas de descargas.
- Paginación y filtros avanzados.
- Interfaz web y aplicación móvil.
- Protección de rutas mediante JWT.

## 16. Conclusión

El módulo cumple el objetivo de ofrecer servicios web para administrar libros, realizar búsquedas y registrar descargas. La solución integra programación en Go, encapsulación, interfaces, manejo de errores, funciones puras, persistencia en PostgreSQL y serialización JSON.

Las pruebas demuestran que las reglas principales funcionan y que la arquitectura puede integrarse con el resto del proyecto mediante la rama `development` y el despliegue configurado por el equipo.

Este trabajo permitió comprender que un servicio web no consiste únicamente en crear una ruta. También requiere validaciones, lógica de negocio, persistencia, respuestas HTTP coherentes y pruebas que demuestren su funcionamiento.

## 17. Evidencias recomendadas para el video

1. Mostrar la rama `Leonardo_Sanchez` en GitHub.
2. Ejecutar `go run .` desde `GO_CODE_GESTION`.
3. Crear un libro con `POST /api/v1/books`.
4. Listar libros con `GET /api/v1/books`.
5. Buscar con `GET /api/v1/books/search`.
6. Consultar y actualizar un libro por ID.
7. Registrar una descarga con `POST /api/v1/downloads`.
8. Consultar `GET /api/v1/downloads/history`.
9. Mostrar las tablas `books` y `downloads` en pgAdmin.
10. Mostrar el resultado de `go test ./...`.

