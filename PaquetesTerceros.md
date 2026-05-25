# Paquetes de Terceros Recomendados

## Descripción general

Para el desarrollo del Sistema de Gestión de Libros Electrónicos en Go, se recomienda utilizar paquetes de terceros que permitan construir un backend organizado, seguro y mantenible.

La selección de estos paquetes responde a las necesidades principales del proyecto: creación de una API REST, conexión con base de datos MySQL, autenticación de usuarios, protección de contraseñas, configuración del entorno y comunicación con el frontend.

No se recomienda agregar librerías innecesarias, ya que esto puede aumentar la complejidad del sistema sin aportar beneficios reales. Por esta razón, los paquetes seleccionados se enfocan únicamente en funciones necesarias para la arquitectura propuesta.

## Paquetes recomendados

| Paquete | Función principal |
|---|---|
| `github.com/gin-gonic/gin` | Creación de la API REST, manejo de rutas, solicitudes HTTP, middlewares y respuestas JSON. |
| `gorm.io/gorm` | ORM para manejar operaciones con la base de datos mediante modelos y repositorios. |
| `gorm.io/driver/mysql` | Driver necesario para conectar GORM con la base de datos MySQL. |
| `github.com/golang-jwt/jwt/v5` | Generación y validación de tokens JWT para autenticación y control de sesiones. |
| `golang.org/x/crypto/bcrypt` | Cifrado seguro de contraseñas antes de almacenarlas en la base de datos. |
| `github.com/spf13/viper` | Manejo centralizado de variables de configuración y entorno. |
| `github.com/gin-contrib/cors` | Configuración de permisos CORS para permitir la comunicación entre frontend y backend. |

## Justificación de los paquetes

### Gin Web Framework

Gin será utilizado para construir la API REST del sistema. Este paquete permite definir rutas, recibir solicitudes HTTP, procesar datos enviados por el cliente y devolver respuestas en formato JSON.

Dentro del proyecto, Gin se relaciona directamente con la capa de controladores, ya que permitirá manejar las operaciones asociadas a usuarios, libros, préstamos, reservas, descargas y reportes.

Su uso resulta adecuado porque permite estructurar el backend de forma clara y facilita la creación de endpoints para cada módulo del sistema.

### GORM

GORM será utilizado como ORM para facilitar la interacción entre el backend y la base de datos. Su función principal será permitir operaciones de creación, lectura, actualización y eliminación de registros sin depender completamente de consultas SQL manuales.

En el proyecto, GORM se aplicará principalmente en la capa de repositorios, donde se manejará el acceso a los datos de entidades como usuarios, libros, préstamos, reservas y descargas.

Este paquete ayuda a mantener una separación más clara entre la lógica de negocio y la lógica de acceso a datos.

### Driver MySQL para GORM

El paquete `gorm.io/driver/mysql` será necesario para conectar GORM con MySQL, que es el sistema de base de datos definido en la arquitectura del proyecto.

Este driver permite que el backend desarrollado en Go pueda comunicarse correctamente con la base de datos, ejecutar operaciones sobre las tablas y gestionar la persistencia de la información.

Su inclusión es importante porque GORM requiere un driver específico según el motor de base de datos utilizado.

### JWT

El paquete `github.com/golang-jwt/jwt/v5` será utilizado para implementar autenticación basada en tokens. Después de que un usuario inicie sesión correctamente, el sistema podrá generar un token JWT que será usado para validar futuras solicitudes.

Este paquete será importante para proteger rutas privadas del sistema y controlar el acceso según el tipo de usuario.

Dentro del proyecto, JWT puede aplicarse especialmente en funciones relacionadas con inicio de sesión, autorización de usuarios y validación de permisos.

### Bcrypt

El paquete `golang.org/x/crypto/bcrypt` será utilizado para cifrar las contraseñas de los usuarios antes de almacenarlas en la base de datos.

Este paquete es necesario porque las contraseñas no deben guardarse en texto plano. Al utilizar bcrypt, el sistema mejora la protección de las credenciales y reduce el riesgo en caso de acceso no autorizado a la base de datos.

Su uso complementa el proceso de autenticación junto con JWT.

### Viper

Viper será utilizado para manejar la configuración general del sistema. Este paquete permite centralizar valores como el puerto del servidor, la conexión a la base de datos, claves secretas, variables de entorno y otros parámetros necesarios para la ejecución del backend.

Su uso evita escribir datos sensibles directamente en el código fuente y facilita la administración del sistema en distintos entornos, como desarrollo, pruebas o producción.

### CORS

El paquete `github.com/gin-contrib/cors` será utilizado para configurar los permisos de comunicación entre el frontend y el backend.

Esto es importante porque el sistema contempla una interfaz cliente que deberá consumir los servicios expuestos por la API REST. CORS permitirá definir qué orígenes, métodos y encabezados estarán permitidos al realizar solicitudes al servidor.

Su uso ayuda a controlar de forma más segura la interacción entre la aplicación web y el backend desarrollado en Go.

## Relación con la arquitectura del proyecto

Los paquetes recomendados se relacionan con las diferentes capas del sistema de la siguiente manera:

| Capa del sistema | Paquetes relacionados |
|---|---|
| API REST | Gin, CORS |
| Controladores | Gin |
| Servicios | JWT, Bcrypt |
| Repositorios | GORM |
| Base de datos | GORM, Driver MySQL |
| Seguridad | JWT, Bcrypt, CORS |
| Configuración | Viper |

## Importancia dentro del sistema

La selección de estos paquetes permite cubrir las necesidades principales del backend sin sobrecargar el proyecto con dependencias innecesarias.

Gin facilita la construcción de la API REST. GORM y el driver de MySQL permiten trabajar con la base de datos de forma ordenada. JWT y bcrypt fortalecen la seguridad del sistema. Viper permite administrar la configuración de manera centralizada. CORS permite controlar la comunicación entre el frontend y el backend.

En conjunto, estos paquetes contribuyen a que el sistema sea más organizado, seguro, escalable y fácil de mantener.

## Conclusión

Los paquetes de terceros recomendados para el proyecto fueron seleccionados de acuerdo con las necesidades reales del Sistema de Gestión de Libros Electrónicos.

Cada paquete cumple una función específica dentro de la arquitectura: gestión de rutas, acceso a datos, autenticación, cifrado de contraseñas, configuración del entorno y comunicación entre cliente y servidor.

Esta selección permite construir un backend funcional y coherente con la arquitectura planteada, evitando dependencias innecesarias y manteniendo una estructura clara para el desarrollo del sistema.
