# Paquetes de Terceros Recomendados
---
### Con el fin de asegurar la robustez, seguridad y escalabilidad del backend, hemos seleccionado un conjunto de librerías del ecosistema de Go que se alinean con los estándares de la industria y agilizan nuestro desarrollo incremental:

* **Gin Web Framework**
  * **Sitio Oficial / Repositorio:** [gin-gonic.com](https://gin-gonic.com/) o `github.com/gin-gonic/gin`
  * **Funciones Principales en el Proyecto:**
    * **Enrutamiento rápido y flexible:** Nos permite definir de manera estructurada los endpoints de la API REST para libros, usuarios, préstamos y descargas mediante un árbol de prefijos altamente eficiente.
    * **Manejo de Middlewares:** Permite interceptar las peticiones HTTP para ejecutar procesos globales automatizados, tales como el control de registros (logging), la recuperación ante errores (recovery) y la validación de seguridad.
    * **Validación y renderizado de JSON:** Facilita la deserialización automática de los datos que envían los usuarios (binding) y la respuesta estructurada en formato JSON hacia el cliente con un rendimiento óptimo.

* **GORM (Go Object Relational Mapping)**
  * **Sitio Oficial / Repositorio:** [gorm.io](https://gorm.io/) o `github.com/go-gorm/gorm`
  * **Funciones Principales en el Proyecto:**
    * **Abstracción del CRUD:** Automatiza las operaciones de creación, lectura, actualización y eliminación de registros en la base de datos PostgreSQL, evitando la escritura manual de consultas SQL complejas.
    * **Gestión de Relaciones (Eager Loading):** Permite modelar y consultar de forma nativa las asociaciones entre nuestras estructuras, tales como vincular un registro de la tabla `Prestamos` con un `Usuario` y un `Libro` específico.
    * **Migraciones Automáticas (AutoMigrate):** Sincroniza los esquemas de nuestras estructuras en Go con las tablas físicas de PostgreSQL de manera automática durante el arranque del servidor, asegurando la consistencia de los datos.

* **JWT-Go (golang-jwt)**
  * **Sitio Oficial / Repositorio:** `github.com/golang-jwt/jwt`
  * **Funciones Principales en el Proyecto:**
    * **Generación de Tokens Criptográficos:** Permite emitir tokens JSON Web Tokens firmados digitalmente tras un inicio de sesión exitoso, almacenando de forma segura los datos de identidad del usuario.
    * **Autenticación sin Estado (Stateless):** Facilita la validación de la identidad del usuario en cada petición HTTP a través de un middleware de seguridad, eliminando la necesidad de consultar constantemente la base de datos para verificar la sesión.
    * **Control de Acceso Basado en Roles (RBAC):** Permite leer los permisos (claims) incrustados en el token para restringir el acceso a módulos delicados, garantizando que solo los administradores gestionen el catálogo o los reportes estadísticos.

* **Viper**
  * **Sitio Oficial / Repositorio:** `github.com/spf13/viper`
  * **Funciones Principales en el Proyecto:**
    * **Inyección de Variables de Entorno:** Permite centralizar la lectura de credenciales sensibles (como la contraseña de PostgreSQL o la clave secreta del JWT) desde un archivo externo `.env`.
    * **Soporte Multiformato:** Ofrece la capacidad de leer archivos de configuración en múltiples extensiones de forma agnóstica, facilitando la modularidad del entorno de desarrollo.
    * **Valores por Defecto:** Permite establecer configuraciones base para el puerto del servidor o rutas de almacenamiento de archivos por si estas no se especifican explícitamente en el entorno del sistema.
