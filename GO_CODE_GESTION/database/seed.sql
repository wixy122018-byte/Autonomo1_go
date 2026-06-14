-- Datos iniciales para probar usuarios, catalogo y descargas.
-- Contrasena para todos los usuarios de prueba: secreto123
-- El hash fue generado con bcrypt.

INSERT INTO users (name, email, password, role)
VALUES
    ('Estudiante Demo', 'estudiante@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Administrador Demo', 'admin@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'admin'),
    ('Leonardo Sanchez', 'leonardo.sanchez@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Ana Perez', 'ana.perez@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Carlos Mendoza', 'carlos.mendoza@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Maria Zambrano', 'maria.zambrano@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Luis Mora', 'luis.mora@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Carla Ruiz', 'carla.ruiz@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Daniel Torres', 'daniel.torres@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Sofia Castillo', 'sofia.castillo@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Jorge Andrade', 'jorge.andrade@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Paula Cevallos', 'paula.cevallos@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Mateo Guerrero', 'mateo.guerrero@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Valeria Ponce', 'valeria.ponce@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Kevin Lopez', 'kevin.lopez@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Gabriela Vera', 'gabriela.vera@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Diego Ramirez', 'diego.ramirez@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Nicole Delgado', 'nicole.delgado@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Andres Molina', 'andres.molina@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Camila Flores', 'camila.flores@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user')
ON CONFLICT (email) DO NOTHING;

INSERT INTO books (title, author, category, editorial, year, description, format, isbn, available, active, file_path)
VALUES
    ('Go Basico', 'Ana Perez', 'Informatica', 'Editorial UTM', 2026, 'Introduccion a programacion con Go.', 'PDF', '978-1-001', TRUE, TRUE, '/books/go-basico.pdf'),
    ('Algoritmos desde Cero', 'Carlos Mendoza', 'Informatica', 'Editorial Codigo', 2025, 'Fundamentos de algoritmos y pensamiento computacional.', 'PDF', '978-1-002', TRUE, TRUE, '/books/algoritmos-desde-cero.pdf'),
    ('Estructuras de Datos en Go', 'Valeria Ponce', 'Informatica', 'Editorial Codigo', 2026, 'Listas, pilas, colas, mapas, arboles y grafos.', 'PDF', '978-1-003', TRUE, TRUE, '/books/estructuras-datos-go.pdf'),
    ('Bases de Datos PostgreSQL', 'Daniel Torres', 'Informatica', 'Data Press', 2025, 'Diseno, consultas SQL e indices en PostgreSQL.', 'PDF', '978-1-004', TRUE, TRUE, '/books/postgresql.pdf'),
    ('SQL Practico', 'Sofia Castillo', 'Informatica', 'Data Press', 2024, 'Consultas, filtros, joins y agrupaciones.', 'PDF', '978-1-005', TRUE, TRUE, '/books/sql-practico.pdf'),
    ('Desarrollo Web con HTML CSS y JavaScript', 'Jorge Andrade', 'Informatica', 'Web Academy', 2025, 'Bases para construir sitios web interactivos.', 'PDF', '978-1-006', TRUE, TRUE, '/books/desarrollo-web.pdf'),
    ('APIs REST con Gin', 'Paula Cevallos', 'Informatica', 'Backend Lab', 2026, 'Creacion de servicios REST con Go y Gin.', 'PDF', '978-1-007', TRUE, TRUE, '/books/apis-rest-gin.pdf'),
    ('GORM y Persistencia de Datos', 'Mateo Guerrero', 'Informatica', 'Backend Lab', 2026, 'Uso de ORM para conectar Go con bases de datos.', 'PDF', '978-1-008', TRUE, TRUE, '/books/gorm-persistencia.pdf'),
    ('Ciberseguridad Basica', 'Kevin Lopez', 'Informatica', 'Security House', 2025, 'Conceptos esenciales de seguridad informatica.', 'PDF', '978-1-009', TRUE, TRUE, '/books/ciberseguridad-basica.pdf'),
    ('Redes de Computadoras', 'Gabriela Vera', 'Informatica', 'Net Press', 2024, 'Modelos OSI, TCP/IP, direccionamiento y servicios.', 'PDF', '978-1-010', TRUE, TRUE, '/books/redes-computadoras.pdf'),
    ('Sistemas Operativos', 'Diego Ramirez', 'Informatica', 'Tech Books', 2023, 'Procesos, memoria, archivos y administracion del sistema.', 'PDF', '978-1-011', TRUE, TRUE, '/books/sistemas-operativos.pdf'),
    ('Ingenieria de Software', 'Nicole Delgado', 'Informatica', 'Tech Books', 2025, 'Analisis, diseno, pruebas y mantenimiento de software.', 'PDF', '978-1-012', TRUE, TRUE, '/books/ingenieria-software.pdf'),
    ('Arquitectura de Software', 'Andres Molina', 'Informatica', 'Tech Books', 2026, 'Patrones, capas, servicios y buenas practicas.', 'PDF', '978-1-013', TRUE, TRUE, '/books/arquitectura-software.pdf'),
    ('Git y GitHub para Equipos', 'Camila Flores', 'Informatica', 'Dev Tools', 2024, 'Control de versiones, ramas, commits y pull requests.', 'PDF', '978-1-014', TRUE, TRUE, '/books/git-github.pdf'),
    ('Docker para Desarrolladores', 'Luis Mora', 'Informatica', 'DevOps Press', 2025, 'Contenedores, imagenes, volumenes y despliegue.', 'PDF', '978-1-015', TRUE, TRUE, '/books/docker-desarrolladores.pdf'),
    ('Introduccion a Linux', 'Carla Ruiz', 'Informatica', 'Open Source Press', 2024, 'Comandos, permisos, procesos y administracion basica.', 'PDF', '978-1-016', TRUE, TRUE, '/books/linux.pdf'),
    ('Python Aplicado', 'Ana Perez', 'Informatica', 'Editorial UTM', 2025, 'Automatizacion, datos y scripts con Python.', 'PDF', '978-1-017', TRUE, TRUE, '/books/python-aplicado.pdf'),
    ('Java Fundamentos', 'Carlos Mendoza', 'Informatica', 'Editorial Codigo', 2023, 'Programacion orientada a objetos con Java.', 'PDF', '978-1-018', TRUE, TRUE, '/books/java-fundamentos.pdf'),
    ('Programacion Orientada a Objetos', 'Valeria Ponce', 'Informatica', 'Editorial Codigo', 2024, 'Clases, objetos, herencia, interfaces y polimorfismo.', 'PDF', '978-1-019', TRUE, TRUE, '/books/poo.pdf'),
    ('Inteligencia Artificial Basica', 'Daniel Torres', 'Informatica', 'AI Press', 2026, 'Conceptos iniciales de IA y aprendizaje automatico.', 'PDF', '978-1-020', TRUE, TRUE, '/books/ia-basica.pdf'),
    ('Machine Learning Practico', 'Sofia Castillo', 'Informatica', 'AI Press', 2025, 'Modelos, entrenamiento, evaluacion y prediccion.', 'PDF', '978-1-021', TRUE, TRUE, '/books/machine-learning.pdf'),
    ('Analisis de Datos', 'Jorge Andrade', 'Informatica', 'Data Press', 2024, 'Limpieza, exploracion y visualizacion de datos.', 'PDF', '978-1-022', TRUE, TRUE, '/books/analisis-datos.pdf'),
    ('Computacion en la Nube', 'Paula Cevallos', 'Informatica', 'Cloud Books', 2025, 'Servicios cloud, escalabilidad y despliegue.', 'PDF', '978-1-023', TRUE, TRUE, '/books/cloud.pdf'),
    ('Fundamentos de DevOps', 'Mateo Guerrero', 'Informatica', 'DevOps Press', 2026, 'Integracion continua, entrega continua y automatizacion.', 'PDF', '978-1-024', TRUE, TRUE, '/books/devops.pdf'),
    ('Pruebas de Software', 'Kevin Lopez', 'Informatica', 'QA Press', 2025, 'Unit testing, pruebas de integracion y calidad.', 'PDF', '978-1-025', TRUE, TRUE, '/books/pruebas-software.pdf'),
    ('Seguridad en Aplicaciones Web', 'Gabriela Vera', 'Informatica', 'Security House', 2025, 'Riesgos comunes y buenas practicas de seguridad.', 'PDF', '978-1-026', TRUE, TRUE, '/books/seguridad-web.pdf'),
    ('Administracion de Servidores', 'Diego Ramirez', 'Informatica', 'Net Press', 2024, 'Configuracion y mantenimiento de servidores.', 'PDF', '978-1-027', TRUE, TRUE, '/books/administracion-servidores.pdf'),
    ('Aplicaciones Moviles', 'Nicole Delgado', 'Informatica', 'Mobile Press', 2025, 'Diseno y desarrollo de aplicaciones para moviles.', 'PDF', '978-1-028', TRUE, TRUE, '/books/apps-moviles.pdf'),
    ('UX UI para Sistemas', 'Andres Molina', 'Informatica', 'Design Tech', 2024, 'Experiencia de usuario aplicada a software.', 'PDF', '978-1-029', TRUE, TRUE, '/books/ux-ui.pdf'),
    ('Matematicas para Computacion', 'Camila Flores', 'Informatica', 'Academia Press', 2023, 'Logica, conjuntos, matrices y grafos.', 'PDF', '978-1-030', TRUE, TRUE, '/books/matematicas-computacion.pdf'),
    ('Historia Andina', 'Luis Mora', 'Historia', 'Editorial Saber', 2024, 'Texto de historia regional.', 'PDF', '978-1-031', FALSE, TRUE, '/books/historia-andina.pdf'),
    ('Ciencia para Todos', 'Carla Ruiz', 'Ciencia', 'Editorial Aula', 2025, 'Libro introductorio de ciencias.', 'EPUB', '978-1-032', TRUE, TRUE, '/books/ciencia-para-todos.epub'),
    ('Fisica General', 'Daniel Torres', 'Ciencia', 'Editorial Aula', 2022, 'Mecanica, energia y ondas.', 'PDF', '978-1-033', TRUE, TRUE, '/books/fisica-general.pdf'),
    ('Quimica Basica', 'Sofia Castillo', 'Ciencia', 'Editorial Aula', 2023, 'Materia, atomos, enlaces y reacciones.', 'PDF', '978-1-034', TRUE, TRUE, '/books/quimica-basica.pdf'),
    ('Biologia Moderna', 'Jorge Andrade', 'Ciencia', 'Editorial Aula', 2024, 'Celulas, genetica, evolucion y ecosistemas.', 'PDF', '978-1-035', TRUE, TRUE, '/books/biologia-moderna.pdf'),
    ('Estadistica Aplicada', 'Paula Cevallos', 'Ciencia', 'Academia Press', 2025, 'Probabilidad, muestras, graficos e interpretacion.', 'PDF', '978-1-036', TRUE, TRUE, '/books/estadistica-aplicada.pdf'),
    ('Matematica Discreta', 'Mateo Guerrero', 'Educacion', 'Academia Press', 2024, 'Logica, combinatoria y relaciones.', 'PDF', '978-1-037', TRUE, TRUE, '/books/matematica-discreta.pdf'),
    ('Comunicacion Academica', 'Kevin Lopez', 'Educacion', 'Editorial Saber', 2023, 'Lectura critica, escritura y presentaciones.', 'PDF', '978-1-038', TRUE, TRUE, '/books/comunicacion-academica.pdf'),
    ('Metodologia de Investigacion', 'Gabriela Vera', 'Educacion', 'Editorial Saber', 2024, 'Diseno de proyectos, hipotesis y referencias.', 'PDF', '978-1-039', TRUE, TRUE, '/books/metodologia-investigacion.pdf'),
    ('Administracion General', 'Diego Ramirez', 'Administracion', 'Empresa Press', 2023, 'Planeacion, organizacion, direccion y control.', 'PDF', '978-1-040', TRUE, TRUE, '/books/administracion-general.pdf'),
    ('Contabilidad Basica', 'Nicole Delgado', 'Administracion', 'Empresa Press', 2024, 'Principios contables y registros financieros.', 'PDF', '978-1-041', TRUE, TRUE, '/books/contabilidad-basica.pdf'),
    ('Marketing Digital', 'Andres Molina', 'Administracion', 'Empresa Press', 2025, 'Estrategias digitales, redes y analitica.', 'PDF', '978-1-042', TRUE, TRUE, '/books/marketing-digital.pdf'),
    ('Emprendimiento', 'Camila Flores', 'Administracion', 'Empresa Press', 2023, 'Ideas de negocio, validacion y plan empresarial.', 'PDF', '978-1-043', TRUE, TRUE, '/books/emprendimiento.pdf'),
    ('Psicologia Educativa', 'Ana Perez', 'Educacion', 'Editorial Saber', 2024, 'Aprendizaje, motivacion y procesos educativos.', 'PDF', '978-1-044', TRUE, TRUE, '/books/psicologia-educativa.pdf'),
    ('Derecho Informatico', 'Carlos Mendoza', 'Derecho', 'Legal Tech', 2025, 'Normas, delitos informaticos y proteccion de datos.', 'PDF', '978-1-045', TRUE, TRUE, '/books/derecho-informatico.pdf'),
    ('Etica Profesional', 'Valeria Ponce', 'Educacion', 'Editorial Saber', 2023, 'Responsabilidad profesional y valores.', 'PDF', '978-1-046', TRUE, TRUE, '/books/etica-profesional.pdf'),
    ('Literatura Latinoamericana', 'Luis Mora', 'Literatura', 'Letras Press', 2022, 'Autores y obras representativas de Latinoamerica.', 'PDF', '978-1-047', TRUE, TRUE, '/books/literatura-latinoamericana.pdf'),
    ('Ingles Tecnico', 'Carla Ruiz', 'Idiomas', 'Language Press', 2024, 'Vocabulario tecnico para carreras tecnologicas.', 'PDF', '978-1-048', TRUE, TRUE, '/books/ingles-tecnico.pdf'),
    ('Gestion de Proyectos', 'Daniel Torres', 'Administracion', 'Empresa Press', 2025, 'Planificacion, alcance, cronograma y riesgos.', 'PDF', '978-1-049', TRUE, TRUE, '/books/gestion-proyectos.pdf'),
    ('Innovacion Tecnologica', 'Sofia Castillo', 'Informatica', 'Tech Books', 2026, 'Transformacion digital, creatividad y nuevas tecnologias.', 'PDF', '978-1-050', TRUE, TRUE, '/books/innovacion-tecnologica.pdf')
ON CONFLICT (isbn) DO NOTHING;

INSERT INTO downloads (user_id, book_id, download_date)
SELECT users.id, books.id, NOW()
FROM users
CROSS JOIN books
WHERE users.email = 'estudiante@example.com'
  AND books.isbn = '978-1-001'
  AND NOT EXISTS (
      SELECT 1
      FROM downloads
      WHERE downloads.user_id = users.id
        AND downloads.book_id = books.id
  );

INSERT INTO downloads (user_id, book_id, download_date)
SELECT users.id, books.id, NOW()
FROM users
JOIN books ON books.isbn IN ('978-1-004', '978-1-007', '978-1-020', '978-1-023')
WHERE users.email IN ('leonardo.sanchez@example.com', 'ana.perez@example.com', 'carlos.mendoza@example.com')
  AND NOT EXISTS (
      SELECT 1
      FROM downloads
      WHERE downloads.user_id = users.id
        AND downloads.book_id = books.id
  );
