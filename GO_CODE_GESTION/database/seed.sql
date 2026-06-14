-- Datos iniciales para probar el catalogo y las descargas.
-- La contrasena del usuario de ejemplo es: secreto123
-- El hash fue generado con bcrypt.

INSERT INTO users (name, email, password, role)
VALUES
    ('Estudiante Demo', 'estudiante@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'user'),
    ('Administrador Demo', 'admin@example.com', '$2a$10$lU5WgO4NMZnb1PvXjDZWXONjq7hrLAmHwpwABlSkJrzUtVwyFeuW2', 'admin')
ON CONFLICT (email) DO NOTHING;

INSERT INTO books (title, author, category, editorial, year, description, format, isbn, available, active, file_path)
VALUES
    ('Go Basico', 'Ana Perez', 'Tecnologia', 'Editorial UTM', 2026, 'Introduccion a programacion con Go.', 'PDF', '978-1-001', TRUE, TRUE, '/books/go-basico.pdf'),
    ('Historia Andina', 'Luis Mora', 'Historia', 'Editorial Saber', 2024, 'Texto de historia regional.', 'PDF', '978-1-002', FALSE, TRUE, '/books/historia-andina.pdf'),
    ('Ciencia para Todos', 'Carla Ruiz', 'Ciencia', 'Editorial Aula', 2025, 'Libro introductorio de ciencias.', 'EPUB', '978-1-003', TRUE, TRUE, '/books/ciencia-para-todos.epub')
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
