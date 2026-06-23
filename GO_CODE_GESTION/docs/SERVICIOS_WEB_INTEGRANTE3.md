# Servicios web - Integrante 3

El modulo expone servicios web REST en formato JSON para libros, busquedas, descargas y autenticacion.

## URL local

```text
http://localhost:8081
```

## Catalogo de servicios

```http
GET /api/v1/services
```

Devuelve en JSON la lista de rutas del modulo.

## Rutas versionadas

| Metodo | Ruta | Descripcion |
|---|---|---|
| POST | `/api/v1/books` | Registrar libro |
| GET | `/api/v1/books` | Listar libros |
| GET | `/api/v1/books/:id` | Consultar libro por ID |
| PUT | `/api/v1/books/:id` | Actualizar libro |
| DELETE | `/api/v1/books/:id` | Desactivar libro |
| GET | `/api/v1/books/search` | Buscar y filtrar libros |
| POST | `/api/v1/downloads` | Registrar descarga |
| GET | `/api/v1/downloads/history` | Consultar historial |
| POST | `/api/v1/users/register` | Registrar usuario |
| POST | `/api/v1/login` | Iniciar sesion y obtener JWT |
| GET | `/api/v1/users/me` | Consultar usuario autenticado |

Tambien se mantienen las rutas originales sin `/api/v1` para compatibilidad.

## Ejemplos de prueba

Listar libros:

```bash
curl http://localhost:8081/api/v1/books
```

Buscar libros de informatica:

```bash
curl "http://localhost:8081/api/v1/books/search?category=Informatica"
```

Login:

```bash
curl -X POST http://localhost:8081/api/v1/login \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"estudiante@example.com\",\"password\":\"secreto123\"}"
```

Registrar descarga:

```bash
curl -X POST http://localhost:8081/api/v1/downloads \
  -H "Content-Type: application/json" \
  -d "{\"user_id\":1,\"book_id\":1}"
```
