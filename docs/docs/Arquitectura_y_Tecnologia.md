ARQUITECTURA DEL SISTEMA DE GESTIÓN DE LIBROS ELECTRÓNICOS

El Sistema de Gestión de Libros Electrónicos será desarrollado utilizando el lenguaje de programación Go, implementando una arquitectura por capas que permita mantener el sistema organizado, seguro, escalable y fácil de mantener.
La arquitectura estará basada en el siguiente flujo:
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

Este tipo de arquitectura permite dividir el sistema en componentes organizados, donde cada capa tiene responsabilidades específicas. Gracias a esta separación, el sistema será más fácil de desarrollar, mantener, escalar y actualizar.
La arquitectura también ayuda a mejorar la seguridad, reutilización del código y organización interna del proyecto, permitiendo que cada componente trabaje de forma independiente pero conectada con los demás.

ARQUITECTURA GENERAL EN GO
En Go el proyecto se organizará mediante paquetes y carpetas estructuradas para mantener un código limpio y modular.
La estructura del proyecto será similar a la siguiente:
/proyecto-libros
│
├── main.go
│
├── controllers/
├── services/
├── repositories/
├── models/
├── routes/
├── database/
├── middleware/
├── utils/
└── config/

Cada carpeta tendrá una función específica dentro de la arquitectura.


1. CAPA CLIENTE
La capa cliente corresponde a la interfaz visual del sistema, es decir, la parte con la que interactúan directamente los usuarios. Esta capa es responsable de mostrar información y permitir que las personas utilicen las funcionalidades disponibles dentro de la plataforma.
Aquí se encuentran todas las pantallas, formularios, botones y menús que permiten navegar dentro del sistema.
Usuarios que interactúan con esta capa
El sistema tendrá distintos tipos de usuarios:
Administrador	Bibliotecario	Lector o estudiante
Tiene control total del sistema:
•	Gestiona usuarios. 
•	Administra libros. 
•	Visualiza reportes. 
•	Configura parámetros del sistema. 
	Se encarga de administrar el contenido bibliográfico:
•	Registra libros. 
•	Gestiona préstamos. 
•	Controla reservas. 
•	Actualiza información. 
	Es el usuario final que utiliza la biblioteca digital:
•	Busca libros. 
•	Solicita préstamos. 
•	Reserva libros. 
•	Consulta historial. 

Funciones principales de la capa cliente
Inicio de sesión
Permite autenticar usuarios mediante:
•	Correo electrónico. 
•	Contraseña. 
El sistema validará la información antes de enviarla al servidor.







Inicio de sesión	Visualización de libros	Búsqueda y filtros	Solicitudes y reservas	Panel administrativo
Permite autenticar usuarios mediante:
Correo electrónico.
Contraseña. 
El sistema validará la información antes de enviarla al servidor.
	Los usuarios podrán:
Ver catálogo digital. 
Consultar autores. 
Leer descripciones. 
Revisar disponibilidad. 
	La interfaz permitirá buscar libros mediante:
Título. 
Autor. 
Categoría. 
Palabras clave. 

	Los usuarios podrán:
Solicitar préstamos. 
Reservar libros ocupados. 
Consultar historial. 
	Los administradores tendrán acceso a:
Gestión de usuarios. 
Reportes. 
Estadísticas. 
Configuraciones. 


Tecnologías utilizadas en la capa cliente
La interfaz será desarrollada utilizando:
•	HTML5 
•	CSS3 
•	JavaScript 
Estas tecnologías permitirán construir una plataforma:
•	Dinámica. 
•	Interactiva. 
•	Responsiva. 
•	Compatible con dispositivos modernos. 
Funcionamiento de la capa cliente
Cuando el usuario realiza una acción:
1.	La interfaz captura la información. 
2.	Se genera una solicitud HTTP. 
3.	La solicitud es enviada hacia la API REST desarrollada en Go. 
4.	El sistema procesa la operación. 
5.	La respuesta regresa al cliente. 
La capa cliente únicamente se encarga de interacción visual y no procesa lógica compleja del sistema.


Ejemplo práctico
Si un usuario busca un libro:
•	Escribe el nombre del libro. 
•	La interfaz envía una petición GET. 
•	La API procesa la solicitud. 
•	Los resultados se muestran en pantalla.

2. CAPA API REST
Gin Framework
Gin es un framework para Go que permite construir APIs rápidas y eficientes.
La API REST actuará como puente entre la interfaz y el backend.
La API permite que diferentes aplicaciones puedan comunicarse con el sistema, incluyendo:
•	Aplicaciones web. 
•	Aplicaciones móviles. 
•	Servicios externos. 

¿Qué es una API REST?
Una API REST es un conjunto de rutas o endpoints que permiten enviar y recibir información.
Cada funcionalidad del sistema tendrá una ruta específica.
Ejemplos de endpoints
POST /login
GET /libros
POST /prestamos
PUT /usuarios/:id
DELETE /reservas/:id
Cada endpoint representa una funcionalidad concreta del sistema.
Funciones principales de la API REST
Recepción de solicitudes
La API recibe solicitudes enviadas desde el cliente.

Ejemplos:
•	Iniciar sesión. 
•	Registrar usuarios. 
•	Buscar libros. 
•	Solicitar préstamos. 

Procesamiento de métodos HTTP
La API trabaja con métodos como:
GET	POST	PUT
	DELETE
Obtiene información.
Ejemplo:
GET /libros
Sirve para consultar libros registrados.
	Envía información nueva.
Ejemplo:
POST /usuarios
Sirve para registrar nuevos usuarios.
	Actualiza información existente.
Ejemplo:
PUT /libros/15
Actualiza datos de un libro.	Elimina información.
Ejemplo:
DELETE /reservas/10
Elimina una reserva registrada.



Respuestas de la API
La información será enviada generalmente en formato JSON.
Ejemplo:
{
  "titulo": "Redes Informáticas",
  "autor": "Autor Ejemplo",
  "categoria": "Tecnología",
  "disponible": true
}

Importancia de la API REST
La API permite:
•	Separar frontend y backend. 
•	Facilitar integración futura. 
•	Mejorar organización. 
•	Centralizar servicios. 
•	Permitir escalabilidad. 
Implementación en Go
En Go las rutas se configurarán mediante Gin.
Ejemplo conceptual:
router.GET("/libros", controllers.ObtenerLibros)
Esto significa que cuando el usuario consulte /libros, la solicitud será enviada al controlador correspondiente.

3. CAPA DE CONTROLADORES
Los controladores son responsables de recibir las solicitudes provenientes de la API REST y coordinar el flujo interno del sistema.
Funcionan como intermediarios entre la API y la lógica de negocio.
Funciones principales de los controladores
Recepción de solicitudes
Los controladores reciben peticiones enviadas desde la API.
Por ejemplo:
•	Solicitudes de login. 
•	Registro de libros. 
•	Préstamos. 
•	Consultas. 

Validación inicial de datos
Antes de procesar una operación, el controlador verifica:
•	Campos obligatorios. 
•	Tipos de datos. 
•	Formatos correctos. 
•	Parámetros válidos. 
Ejemplo:
•	Validar que el correo tenga formato correcto. 
•	Confirmar que la contraseña no esté vacía. 


Ejemplo en Go
func ObtenerLibros(c *gin.Context) {
    libros := services.ListarLibros()
    c.JSON(200, libros)

Envío hacia servicios
Una vez validada la información, el controlador envía los datos hacia la capa de servicios.
Ejemplo práctico
Cuando un usuario inicia sesión:
1.	El cliente envía correo y contraseña. 
2.	La API recibe la solicitud. 
3.	El controlador valida los datos. 
4.	Envía la información al servicio de autenticación. 
5.	Espera respuesta. 
6.	Devuelve resultado al usuario. 

Importancia de los controladores
Permiten:
•	Organizar solicitudes. 
•	Validar entradas. 
•	Controlar respuestas. 
•	Mantener código limpio. 


4. CAPA DE SERVICIOS /services
La capa de servicios contiene toda la lógica de negocio del sistema.
Es considerada la parte más importante porque aquí se ejecutan las reglas y procesos internos que permiten el funcionamiento de la aplicación.
Los servicios toman decisiones y controlan cómo debe comportarse el sistema.



Funciones principales de los servicios
Gestión de usuarios	Gestión de libros	Gestión de préstamos	Gestión de reservas
	Seguridad
Permite:
Registrar usuarios. 
Validar credenciales. 
Gestionar roles. 
Controlar permisos. 
	Se encarga de:
Registrar libros. 
Actualizar información. 
Verificar disponibilidad. 
Clasificar categorías.	Controla:
Solicitudes de préstamo. 
Fechas de devolución. 
Disponibilidad. 
Historiales. 
	Permite:
Registrar reservas. 
Notificar disponibilidad. 
Gestionar listas de espera. 
	Implementa:
Encriptación de contraseñas. 
Validación de sesiones. 
Autenticación. 
Protección de información
Ejemplo detallado de funcionamiento
Cuando un estudiante solicita un libro:
1.	El servicio verifica si el usuario existe. 
2.	Comprueba que el libro esté disponible. 
3.	Revisa límite de préstamos. 
4.	Registra la solicitud. 
5.	Actualiza disponibilidad. 
6.	Genera respuesta final. 
Toda esta lógica ocurre dentro de la capa de servicios.

Importancia de la capa de servicios
Permite:
•	Centralizar lógica del negocio. 
•	Reutilizar procesos. 
•	Mantener organización. 
•	Facilitar mantenimiento. 

Ejemplo  en Go
func ListarLibros() []models.Libro {
    return repositories.ObtenerLibros()


5. CAPA DE REPOSITORIOS
La capa de repositorios es responsable de comunicarse directamente con la base de datos.
Su función es ejecutar operaciones relacionadas con almacenamiento y recuperación de información.

Funciones principales de los repositorios
Consultas de datos
	Inserción de información	Actualización de registros
	Eliminación de información

Permite:
•	Buscar usuarios. 
•	Consultar libros. 
•	Obtener préstamos. 
•	Revisar reservas. 
	Registra:
•	Nuevos usuarios. 
•	Nuevos libros. 
•	Préstamos. 
•	Reportes. 	Permite modificar:
•	Datos de usuarios. 
•	Información de libros. 
•	Estados de préstamos. 
	Elimina:
•	Reservas. 
•	Registros innecesarios. 
•	Datos obsoletos. 


Ejemplo práctico
Cuando se registra un préstamo:
1.	El servicio envía la solicitud. 
2.	El repositorio genera la consulta SQL. 
3.	La información se guarda en la base de datos. 




Importancia de los repositorios
Permiten:
•	Separar lógica y datos. 
•	Mejorar organización. 
•	Facilitar mantenimiento. 
•	Optimizar consultas. 
Ejemplo  en Go
func ObtenerLibros() []models.Libro {
    var libros []models.Libro
    database.DB.Find(&libros)
    return libros

6. CAPA DE BASE DE DATOS
La base de datos es la capa encargada de almacenar permanentemente toda la información del sistema.
Se utilizará:
MySQL
MySQL permitirá administrar la información de forma segura, rápida y estructurada.

Información almacenada
Usuarios
	Libros
	Préstamos
	Reservas
	Reportes y auditorías

Nombres. 
Correos. 
Contraseñas cifradas. 
Roles. 
	Títulos. 
Autores. 
Categorías. 
Archivos digitales. 
	Usuarios. 
Fechas. 
Estados. 
	•	Fechas. 
•	Disponibilidad. 
•	Historiales. 	Actividades del sistema. 
Logs. 
Estadísticas. 






Importancia de la base de datos
Garantiza:
•	Persistencia de información. 
•	Seguridad de datos. 
•	Integridad. 
•	Disponibilidad. 
•	Consultas rápidas. 
Ejemplo de conexión
dsn := "usuario:password@tcp(localhost:3306)/biblioteca"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

FLUJO COMPLETO DEL SISTEMA
1. El usuario interactúa con la interfaz web.
                    ↓
2. El cliente envía solicitud HTTP a la API REST.
                    ↓
3. Gin recibe la petición.
                    ↓
4. El controlador procesa y valida datos.
                    ↓
5. El servicio ejecuta la lógica de negocio.
                    ↓
6. El repositorio consulta MySQL.
                    ↓
7. La base de datos devuelve resultados.
                    ↓
8. El sistema genera respuesta JSON.
                    ↓
9. La respuesta retorna al cliente.
CONCLUSIÓN
La arquitectura basada en Cliente, API REST, Controladores, Servicios, Repositorios y Base de Datos permite desarrollar un Sistema de Gestión de Libros Electrónicos organizado, seguro y escalable.
La separación por capas facilita la administración del código, mejora el mantenimiento y permite que el sistema pueda crecer en el futuro incorporando nuevas funcionalidades sin afectar su estructura principal.

