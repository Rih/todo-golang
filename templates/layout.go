package templates

import (
	"html/template"
	"time"
	"todo-list/models"
)

// TodoStats representa las estadísticas del todo list
type TodoStats struct {
	Total     int
	Pending   int
	Completed int
}

// formatDate formatea una fecha para mostrar
func formatDate(t time.Time) string {
	return t.Format("02/01/2006 15:04")
}

// GetLayoutTemplate retorna el template principal
func GetLayoutTemplate() *template.Template {
	tmpl := `
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="/static/styles.css">
    <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css" rel="stylesheet">
    <script src="https://unpkg.com/htmx.org@1.9.10"></script>
</head>
<body>
    <div class="container">
        <header class="header">
            <h1><i class="fas fa-tasks"></i> Todo List</h1>
            <p>Gestiona tus tareas de manera eficiente</p>
        </header>

        <div class="todo-form">
            <form hx-post="/api/todos" hx-target="#todoList" hx-swap="outerHTML">
                <div class="form-group">
                    <input type="text" name="title" placeholder="Título de la tarea" required>
                </div>
                <div class="form-group">
                    <textarea name="description" placeholder="Descripción (opcional)"></textarea>
                </div>
                <div class="form-actions">
                    <button type="submit">
                        <i class="fas fa-plus"></i> Agregar Tarea
                    </button>
                </div>
            </form>
        </div>

        <div class="filters">
            <button class="filter-btn active" hx-get="/api/todos" hx-target="#todoList">
                <i class="fas fa-list"></i> Todas
            </button>
            <button class="filter-btn" hx-get="/api/todos?filter=pending" hx-target="#todoList">
                <i class="fas fa-clock"></i> Pendientes
            </button>
            <button class="filter-btn" hx-get="/api/todos?filter=completed" hx-target="#todoList">
                <i class="fas fa-check"></i> Completadas
            </button>
        </div>

        <div class="todo-stats">
            <div class="stat">
                <span class="stat-number">{{.Stats.Total}}</span>
                <span class="stat-label">Total</span>
            </div>
            <div class="stat">
                <span class="stat-number">{{.Stats.Pending}}</span>
                <span class="stat-label">Pendientes</span>
            </div>
            <div class="stat">
                <span class="stat-number">{{.Stats.Completed}}</span>
                <span class="stat-label">Completadas</span>
            </div>
        </div>

        <div id="todoList">
            {{if .Todos}}
                <div class="todo-list">
                    {{range .Todos}}
                        <div class="todo-item {{if .Completed}}todo-item-completed{{end}}">
                            <div class="todo-header">
                                <div>
                                    <div class="todo-title">{{.Title}}</div>
                                    {{if .Description}}
                                        <div class="todo-description">{{.Description}}</div>
                                    {{end}}
                                </div>
                            </div>
                            <div class="todo-meta">
                                <span><i class="fas fa-calendar"></i> {{formatDate .CreatedAt}}</span>
                                <span><i class="fas fa-clock"></i> {{formatDate .UpdatedAt}}</span>
                            </div>
                            <div class="todo-actions">
                                <button 
                                    class="btn {{if .Completed}}btn-secondary{{else}}btn-success{{end}}" 
                                    hx-put="/api/todos/{{.ID}}" 
                                    hx-vals='{"completed": {{if .Completed}}false{{else}}true{{end}}}'
                                    hx-target="#todoList"
                                    hx-swap="outerHTML"
                                >
                                    <i class="fas {{if .Completed}}fa-undo{{else}}fa-check{{end}}"></i>
                                    {{if .Completed}}Desmarcar{{else}}Completar{{end}}
                                </button>
                                <button 
                                    class="btn btn-primary" 
                                    hx-get="/api/todos/{{.ID}}/edit"
                                    hx-target="body"
                                    hx-swap="beforeend"
                                >
                                    <i class="fas fa-edit"></i> Editar
                                </button>
                                <button 
                                    class="btn btn-danger" 
                                    hx-delete="/api/todos/{{.ID}}"
                                    hx-target="#todoList"
                                    hx-swap="outerHTML"
                                    hx-confirm="¿Estás seguro de que quieres eliminar esta tarea?"
                                >
                                    <i class="fas fa-trash"></i> Eliminar
                                </button>
                            </div>
                        </div>
                    {{end}}
                </div>
            {{else}}
                <div class="empty-state">
                    <i class="fas fa-clipboard-list"></i>
                    <h3>No hay tareas</h3>
                    <p>Agrega tu primera tarea para comenzar</p>
                </div>
            {{end}}
        </div>
    </div>
</body>
</html>`

	return template.Must(template.New("layout").Funcs(template.FuncMap{
		"formatDate": formatDate,
	}).Parse(tmpl))
}

// GetTodoListTemplate retorna el template para la lista de todos (HTMX)
func GetTodoListTemplate() *template.Template {
	tmpl := `
{{if .Todos}}
    <div class="todo-list">
        {{range .Todos}}
            <div class="todo-item {{if .Completed}}todo-item-completed{{end}}">
                <div class="todo-header">
                    <div>
                        <div class="todo-title">{{.Title}}</div>
                        {{if .Description}}
                            <div class="todo-description">{{.Description}}</div>
                        {{end}}
                    </div>
                </div>
                <div class="todo-meta">
                    <span><i class="fas fa-calendar"></i> {{formatDate .CreatedAt}}</span>
                    <span><i class="fas fa-clock"></i> {{formatDate .UpdatedAt}}</span>
                </div>
                <div class="todo-actions">
                    <button 
                        class="btn {{if .Completed}}btn-secondary{{else}}btn-success{{end}}" 
                        hx-put="/api/todos/{{.ID}}" 
                        hx-vals='{"completed": {{if .Completed}}false{{else}}true{{end}}}'
                        hx-target="#todoList"
                        hx-swap="outerHTML"
                    >
                        <i class="fas {{if .Completed}}fa-undo{{else}}fa-check{{end}}"></i>
                        {{if .Completed}}Desmarcar{{else}}Completar{{end}}
                    </button>
                    <button 
                        class="btn btn-primary" 
                        hx-get="/api/todos/{{.ID}}/edit"
                        hx-target="body"
                        hx-swap="beforeend"
                    >
                        <i class="fas fa-edit"></i> Editar
                    </button>
                    <button 
                        class="btn btn-danger" 
                        hx-delete="/api/todos/{{.ID}}"
                        hx-target="#todoList"
                        hx-swap="outerHTML"
                        hx-confirm="¿Estás seguro de que quieres eliminar esta tarea?"
                    >
                        <i class="fas fa-trash"></i> Eliminar
                    </button>
                </div>
            </div>
        {{end}}
    </div>
{{else}}
    <div class="empty-state">
        <i class="fas fa-clipboard-list"></i>
        <h3>No hay tareas</h3>
        <p>Agrega tu primera tarea para comenzar</p>
    </div>
{{end}}`

	return template.Must(template.New("todoList").Funcs(template.FuncMap{
		"formatDate": formatDate,
	}).Parse(tmpl))
}

// GetEditModalTemplate retorna el template para el modal de edición
func GetEditModalTemplate() *template.Template {
	tmpl := `
<div class="modal" id="editModal">
    <div class="modal-content">
        <div class="modal-header">
            <h3><i class="fas fa-edit"></i> Editar Tarea</h3>
            <button class="close-btn" hx-get="/api/close-modal" hx-target="#editModal" hx-swap="outerHTML">
                <i class="fas fa-times"></i>
            </button>
        </div>
        <div class="modal-body">
            <form hx-put="/api/todos/{{.ID}}" hx-target="#todoList" hx-swap="outerHTML">
                <div class="form-group">
                    <label for="editTitle">Título:</label>
                    <input type="text" id="editTitle" name="title" value="{{.Title}}" required>
                </div>
                <div class="form-group">
                    <label for="editDescription">Descripción:</label>
                    <textarea id="editDescription" name="description">{{.Description}}</textarea>
                </div>
                <div class="form-group">
                    <label class="checkbox-label">
                        <input type="checkbox" name="completed" {{if .Completed}}checked{{end}}>
                        <span class="checkmark"></span>
                        Tarea completada
                    </label>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" hx-get="/api/close-modal" hx-target="#editModal" hx-swap="outerHTML">Cancelar</button>
                    <button type="submit" class="btn btn-primary">Guardar Cambios</button>
                </div>
            </form>
        </div>
    </div>
</div>`

	return template.Must(template.New("editModal").Parse(tmpl))
}

// PageData representa los datos para la página
type PageData struct {
	Title string
	Todos []models.Todo
	Stats TodoStats
}

// TodoListData representa los datos para la lista de todos
type TodoListData struct {
	Todos []models.Todo
}
