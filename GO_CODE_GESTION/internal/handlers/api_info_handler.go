package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type WebServiceInfo struct {
	Method      string   `json:"method"`
	Path        string   `json:"path"`
	Description string   `json:"description"`
	RequestBody string   `json:"request_body,omitempty"`
	RequiresJWT bool     `json:"requires_jwt"`
	Roles       []string `json:"roles,omitempty"`
}

func WebServicesCatalog(c *gin.Context) {
	services := []WebServiceInfo{
		{Method: "POST", Path: "/register", Description: "Registra un nuevo usuario", RequestBody: "RegisterRequest", RequiresJWT: false},
		{Method: "POST", Path: "/login", Description: "Valida credenciales y devuelve JWT", RequestBody: "LoginRequest", RequiresJWT: false},
		{Method: "GET", Path: "/api/v1/services", Description: "Catálogo interactivo de servicios web", RequiresJWT: false},
		{Method: "GET", Path: "/profile", Description: "Perfil del usuario autenticado", RequiresJWT: true},
		{Method: "GET", Path: "/books", Description: "Lista libros activos con filtros (título, autor, categoría, disponibilidad)", RequiresJWT: true},
		{Method: "GET", Path: "/books/search", Description: "Busca libros por múltiples criterios", RequiresJWT: true},
		{Method: "GET", Path: "/books/:id", Description: "Consulta un libro por ID", RequiresJWT: true},
		{Method: "POST", Path: "/books", Description: "Registra un libro nuevo", RequestBody: "BookInput", RequiresJWT: true, Roles: []string{"ADMINISTRADOR", "BIBLIOTECARIO"}},
		{Method: "PUT", Path: "/books/:id", Description: "Actualiza los datos de un libro", RequestBody: "BookInput", RequiresJWT: true, Roles: []string{"ADMINISTRADOR", "BIBLIOTECARIO"}},
		{Method: "DELETE", Path: "/books/:id", Description: "Desactiva un libro sin eliminarlo", RequiresJWT: true, Roles: []string{"ADMINISTRADOR", "BIBLIOTECARIO"}},
		{Method: "POST", Path: "/downloads", Description: "Registra una descarga si el libro está disponible", RequestBody: "DownloadInput", RequiresJWT: true},
		{Method: "GET", Path: "/downloads/history", Description: "Consulta el historial de descargas del usuario", RequiresJWT: true},
		{Method: "GET", Path: "/users", Description: "Lista todos los usuarios", RequiresJWT: true, Roles: []string{"ADMINISTRADOR"}},
		{Method: "GET", Path: "/users/:id", Description: "Consulta un usuario por ID", RequiresJWT: true, Roles: []string{"ADMINISTRADOR"}},
		{Method: "DELETE", Path: "/users/:id", Description: "Elimina un usuario por ID", RequiresJWT: true, Roles: []string{"ADMINISTRADOR"}},
	}

	// Check if client wants HTML (e.g. web browser)
	accept := c.GetHeader("Accept")
	format := c.Query("format")

	if strings.Contains(accept, "text/html") && format != "json" {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, renderHTML(services))
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"module":   "Sistema de Gestión de Libros Electrónicos",
		"type":     "REST JSON API",
		"services": services,
	})
}

func renderHTML(services []WebServiceInfo) string {
	var items strings.Builder
	for _, s := range services {
		rolesHTML := ""
		if len(s.Roles) > 0 {
			rolesHTML = `<span class="badge badge-roles">👥 ` + strings.Join(s.Roles, ", ") + `</span>`
		} else if s.RequiresJWT {
			rolesHTML = `<span class="badge badge-roles-all">👥 Cualquier Rol</span>`
		}

		jwtHTML := ""
		if s.RequiresJWT {
			jwtHTML = `<span class="badge badge-jwt">🔒 JWT</span>`
		} else {
			jwtHTML = `<span class="badge badge-public">🔓 Público</span>`
		}

		bodyHTML := ""
		if s.RequestBody != "" {
			bodyHTML = `<div class="body-info">
				<span class="body-label">Payload:</span>
				<code class="code-body">` + s.RequestBody + `</code>
			</div>`
		}

		methodClass := "method-" + strings.ToLower(s.Method)

		items.WriteString(`
		<div class="card" data-method="` + s.Method + `" data-path="` + s.Path + `" data-desc="` + s.Description + `">
			<div class="card-top">
				<div class="endpoint-ident">
					<span class="method ` + methodClass + `">` + s.Method + `</span>
					<span class="path">` + s.Path + `</span>
				</div>
				<div class="badges">
					` + jwtHTML + `
					` + rolesHTML + `
				</div>
			</div>
			<div class="card-mid">
				<p class="desc">` + s.Description + `</p>
				` + bodyHTML + `
			</div>
		</div>
		`)
	}

	return `<!DOCTYPE html>
<html lang="es">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Catálogo de Servicios | Gestión de Libros</title>
	<meta name="description" content="Servicios Web expuestos para el Sistema de Gestión de Libros Electrónicos">
	<link rel="preconnect" href="https://fonts.googleapis.com">
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
	<link href="https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700&family=JetBrains+Mono:wght@500&display=swap" rel="stylesheet">
	<style>
		:root {
			--bg-primary: #0b0f19;
			--bg-secondary: #161b26;
			--bg-card: #1f2638;
			--border-color: #2e3748;
			--text-primary: #f3f4f6;
			--text-secondary: #9ca3af;
			--primary-color: #6366f1;
			--primary-hover: #4f46e5;
			--success: #10b981;
			--info: #3b82f6;
			--warning: #f59e0b;
			--danger: #ef4444;
		}

		* {
			box-sizing: border-box;
			margin: 0;
			padding: 0;
		}

		body {
			font-family: 'Plus Jakarta Sans', sans-serif;
			background-color: var(--bg-primary);
			color: var(--text-primary);
			line-height: 1.5;
			padding: 2rem 1rem;
		}

		.container {
			max-width: 1000px;
			margin: 0 auto;
		}

		header {
			margin-bottom: 2rem;
			display: flex;
			justify-content: space-between;
			align-items: center;
			flex-wrap: wrap;
			gap: 1rem;
		}

		h1 {
			font-size: 2rem;
			font-weight: 700;
			background: linear-gradient(to right, #818cf8, #c084fc);
			-webkit-background-clip: text;
			-webkit-text-fill-color: transparent;
		}

		.subtitle {
			color: var(--text-secondary);
			font-size: 0.95rem;
			margin-top: 0.25rem;
		}

		.status-badge {
			display: inline-flex;
			align-items: center;
			gap: 0.5rem;
			background: rgba(16, 185, 129, 0.1);
			color: var(--success);
			padding: 0.4rem 0.8rem;
			border-radius: 9999px;
			font-size: 0.85rem;
			font-weight: 600;
			border: 1px solid rgba(16, 185, 129, 0.2);
		}

		.status-dot {
			width: 8px;
			height: 8px;
			background-color: var(--success);
			border-radius: 50%;
			box-shadow: 0 0 8px var(--success);
			animation: pulse 1.5s infinite alternate;
		}

		@keyframes pulse {
			from { opacity: 0.4; }
			to { opacity: 1; }
		}

		.search-container {
			position: relative;
			margin-bottom: 2rem;
		}

		.search-input {
			width: 100%;
			padding: 0.9rem 1.2rem;
			background-color: var(--bg-secondary);
			border: 1px solid var(--border-color);
			border-radius: 12px;
			color: var(--text-primary);
			font-family: inherit;
			font-size: 1rem;
			transition: all 0.2s ease;
			outline: none;
		}

		.search-input:focus {
			border-color: var(--primary-color);
			box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.2);
		}

		.grid {
			display: flex;
			flex-direction: column;
			gap: 1rem;
		}

		.card {
			background-color: var(--bg-secondary);
			border: 1px solid var(--border-color);
			border-radius: 12px;
			padding: 1.25rem;
			transition: transform 0.2s ease, border-color 0.2s ease;
		}

		.card:hover {
			transform: translateY(-2px);
			border-color: var(--primary-color);
		}

		.card-top {
			display: flex;
			justify-content: space-between;
			align-items: center;
			flex-wrap: wrap;
			gap: 0.75rem;
			margin-bottom: 0.75rem;
		}

		.endpoint-ident {
			display: flex;
			align-items: center;
			gap: 0.75rem;
		}

		.method {
			font-family: 'JetBrains Mono', monospace;
			font-size: 0.75rem;
			font-weight: 700;
			padding: 0.25rem 0.6rem;
			border-radius: 6px;
			color: #fff;
			min-width: 60px;
			text-align: center;
		}

		.method-get { background-color: var(--info); }
		.method-post { background-color: var(--success); }
		.method-put { background-color: var(--warning); }
		.method-delete { background-color: var(--danger); }

		.path {
			font-family: 'JetBrains Mono', monospace;
			font-weight: 500;
			font-size: 1.05rem;
			color: var(--text-primary);
		}

		.badges {
			display: flex;
			gap: 0.5rem;
		}

		.badge {
			font-size: 0.75rem;
			font-weight: 600;
			padding: 0.25rem 0.5rem;
			border-radius: 6px;
			border: 1px solid transparent;
		}

		.badge-jwt {
			background: rgba(245, 158, 11, 0.1);
			color: var(--warning);
			border-color: rgba(245, 158, 11, 0.2);
		}

		.badge-public {
			background: rgba(16, 185, 129, 0.1);
			color: var(--success);
			border-color: rgba(16, 185, 129, 0.2);
		}

		.badge-roles {
			background: rgba(99, 102, 241, 0.1);
			color: #a5b4fc;
			border-color: rgba(99, 102, 241, 0.2);
		}

		.badge-roles-all {
			background: rgba(156, 163, 175, 0.1);
			color: var(--text-secondary);
			border-color: rgba(156, 163, 175, 0.2);
		}

		.card-mid {
			margin-top: 0.5rem;
		}

		.desc {
			color: var(--text-secondary);
			font-size: 0.95rem;
		}

		.body-info {
			margin-top: 0.75rem;
			display: inline-flex;
			align-items: center;
			gap: 0.5rem;
			background-color: var(--bg-card);
			padding: 0.35rem 0.6rem;
			border-radius: 6px;
			font-size: 0.8rem;
		}

		.body-label {
			color: var(--text-secondary);
		}

		.code-body {
			font-family: 'JetBrains Mono', monospace;
			color: #c084fc;
			font-weight: 500;
		}

		footer.page-footer {
			margin-top: 3rem;
			text-align: center;
			font-size: 0.8rem;
			color: var(--text-secondary);
			border-top: 1px solid var(--border-color);
			padding-top: 1.5rem;
		}

		footer a {
			color: var(--primary-color);
			text-decoration: none;
		}

		footer a:hover {
			text-decoration: underline;
		}
	</style>
</head>
<body>
	<div class="container">
		<header>
			<div>
				<h1>Catálogo de Servicios Web</h1>
				<p class="subtitle">Sistema de Gestión de Libros Electrónicos • API REST</p>
			</div>
			<div class="status-badge">
				<span class="status-dot"></span>
				Servidor en línea
			</div>
		</header>

		<div class="search-container">
			<input type="text" id="searchInput" class="search-input" placeholder="Filtrar por método, ruta o descripción..." autocomplete="off">
		</div>

		<div class="grid" id="endpointsGrid">
			` + items.String() + `
		</div>

		<footer class="page-footer">
			<p>Para obtener el catálogo en JSON formateado, solicita con la cabecera <code>Accept: application/json</code> o agrega <a href="?format=json">?format=json</a>.</p>
		</footer>
	</div>

	<script>
		const searchInput = document.getElementById('searchInput');
		const cards = document.querySelectorAll('.card');

		searchInput.addEventListener('input', function() {
			const query = searchInput.value.toLowerCase().trim();

			cards.forEach(card => {
				const method = card.getAttribute('data-method').toLowerCase();
				const path = card.getAttribute('data-path').toLowerCase();
				const desc = card.getAttribute('data-desc').toLowerCase();

				if (method.includes(query) || path.includes(query) || desc.includes(query)) {
					card.style.display = 'block';
				} else {
					card.style.display = 'none';
				}
			});
		});
	</script>
</body>
</html>`
}
