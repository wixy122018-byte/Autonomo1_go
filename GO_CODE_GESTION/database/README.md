# Base de datos PostgreSQL

El backend usa PostgreSQL mediante GORM. Al iniciar la aplicacion, `AutoMigrate` crea las tablas de acuerdo con los modelos de Go, pero tambien se incluye SQL para crear la base manualmente.

## Crear base de datos

En PostgreSQL ejecuta:

```sql
CREATE DATABASE digital_library;
```

Luego entra a esa base:

```bash
psql -U postgres -d digital_library
```

## Crear tablas

Desde la carpeta `GO_CODE_GESTION`:

```bash
psql -U postgres -d digital_library -f database/schema.sql
```

## Cargar datos de prueba

```bash
psql -U postgres -d digital_library -f database/seed.sql
```

Usuario de prueba:

```text
email: estudiante@example.com
password: secreto123
```

## Variables de entorno

El archivo `.env` debe tener una conexion como esta:

```env
PORT=8081
DATABASE_DSN=host=localhost user=postgres password=qwerty dbname=digital_library port=5432 sslmode=disable
JWT_SECRET=clave_secreta_proyecto_go
```

Si se usa Railway u otra plataforma, tambien se puede configurar:

```env
DATABASE_URL=postgresql://usuario:password@host:5432/digital_library
```

## Tablas incluidas

- `users`: usuarios, roles y contrasenas encriptadas.
- `books`: catalogo de libros.
- `downloads`: historial de descargas.
- `loans`: prestamos.
- `reservations`: reservas.
