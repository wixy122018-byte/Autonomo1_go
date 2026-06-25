package handlers

import (
	_ "embed"
	"strings"
)

//go:embed templates/catalog.html
var catalogHTML string

// renderHTML genera el HTML interactivo inyectando los endpoints dinámicamente en el template embebido.
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

		btnHTML := `<button class="btn-test" onclick="openPlayground('` + s.Method + `', '` + s.Path + `', '` + s.RequestBody + `')">
			<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polygon points="5 3 19 12 5 21 5 3"></polygon></svg>
			Probar
		</button>`

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
				<div style="margin-top: 1rem; display: flex; justify-content: flex-end;">
					` + btnHTML + `
				</div>
			</div>
		</div>
		`)
	}

	return strings.Replace(catalogHTML, "<!-- {{ENDPOINTS_GRID}} -->", items.String(), 1)
}
