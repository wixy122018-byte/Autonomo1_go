# Integrante 3: Libros, busqueda y descargas

Esta parte implementa el catalogo de libros, filtros y descargas dentro del proyecto `GO_CODE_GESTION`.

## Endpoints

| Metodo | Ruta | Descripcion |
|---|---|---|
| POST | `/books` | Registra un libro |
| GET | `/books` | Lista libros activos |
| GET | `/books/:id` | Consulta libro por ID |
| PUT | `/books/:id` | Actualiza un libro |
| DELETE | `/books/:id` | Desactiva un libro |
| GET | `/books/search` | Busca por titulo, autor, categoria o disponibilidad |
| POST | `/downloads` | Registra descarga |
| GET | `/downloads/history` | Consulta historial de descargas |

## Ejemplos

Crear libro:

```bash
curl -X POST http://localhost:8080/books ^
  -H "Content-Type: application/json" ^
  -d "{\"title\":\"Go Basico\",\"author\":\"Ana Perez\",\"category\":\"Tecnologia\",\"editorial\":\"UTM\",\"year\":2026,\"description\":\"Libro introductorio\",\"format\":\"PDF\",\"isbn\":\"978-1-001\",\"available\":true,\"file_path\":\"/books/go-basico.pdf\"}"
```

Buscar:

```bash
curl "http://localhost:8080/books/search?title=go"
curl "http://localhost:8080/books/search?author=ana"
curl "http://localhost:8080/books?category=Tecnologia"
curl "http://localhost:8080/books?available=true"
```

Registrar descarga:

```bash
curl -X POST http://localhost:8080/downloads ^
  -H "Content-Type: application/json" ^
  -d "{\"user_id\":1,\"book_id\":1}"
```

## Autenticacion agregada

Tambien se agrego un modulo basico de usuarios, login, contraseña encriptada y JWT.

| Metodo | Ruta | Descripcion |
|---|---|---|
| POST | `/users/register` | Registra usuario con contraseña encriptada |
| POST | `/login` | Valida credenciales y devuelve token JWT |
| GET | `/users/me` | Consulta el usuario autenticado con `Authorization: Bearer TOKEN` |

Registrar usuario:

```bash
curl -X POST http://localhost:8080/users/register ^
  -H "Content-Type: application/json" ^
  -d "{\"name\":\"Estudiante\",\"email\":\"estudiante@example.com\",\"password\":\"secreto123\",\"role\":\"user\"}"
```

Login:

```bash
curl -X POST http://localhost:8080/login ^
  -H "Content-Type: application/json" ^
  -d "{\"email\":\"estudiante@example.com\",\"password\":\"secreto123\"}"
```

Consultar usuario autenticado:

```bash
curl http://localhost:8080/users/me ^
  -H "Authorization: Bearer TU_TOKEN"
```

## Base de datos PostgreSQL

Se agregaron scripts en la carpeta `database/`:

- `database/schema.sql`: crea las tablas `users`, `books`, `downloads`, `loans` y `reservations`.
- `database/seed.sql`: agrega usuarios y libros de prueba.
- `database/README.md`: explica como crear la base `digital_library`.

El proyecto tambien mantiene `AutoMigrate` en `internal/database/connection.go`, por lo que puede crear las tablas automaticamente al iniciar la aplicacion.

## Contenidos de Go evidenciados

- Structs: `Book`, `Download`, inputs y filtros.
- Slices: filtros puros sobre listas de libros.
- Array: `FixedCategories`.
- Maps: `filterMap` y conteo por categoria.
- Condicionales: validacion de disponibilidad.
- Bucles `for`: filtros y conteo.
- Funciones puras: `FilterBooks`, `CountBooksByCategory`.
- Metodo sobre struct: `Book.IsAvailable()`.
- Constructor: `NewBook`, `NewDownload`.
- Manejo de errores: `ErrBookNotFound`, `ErrBookUnavailable`.
- Interfaces: `BookRepository`, `DownloadRepository`.
- Polimorfismo: los servicios dependen de interfaces y pueden usar GORM/PostgreSQL u otra implementacion.
- Encriptado de contraseña: `bcrypt`.
- JWT: token firmado con HMAC SHA256 usando `JWT_SECRET`.
