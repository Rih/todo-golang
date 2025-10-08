package handlers

import (
	"net/http"
	"strconv"
	"time"
	"todo-list/models"
	"todo-list/templates"
	"github.com/gin-gonic/gin"
)

// TodoHandlerTempl maneja las operaciones CRUD usando Templ y HTMX
type TodoHandlerTempl struct {
	todos  []models.Todo
	nextID int
}

// NewTodoHandlerTempl crea una nueva instancia del handler con Templ
func NewTodoHandlerTempl() *TodoHandlerTempl {
	return &TodoHandlerTempl{
		todos:  make([]models.Todo, 0),
		nextID: 1,
	}
}

// GetHomePage muestra la página principal
func (h *TodoHandlerTempl) GetHomePage(c *gin.Context) {
	stats := h.calculateStats()
	data := templates.PageData{
		Title: "Todo List - Gestor de Tareas",
		Todos: h.todos,
		Stats: stats,
	}
	
	tmpl := templates.GetLayoutTemplate()
	tmpl.Execute(c.Writer, data)
}

// GetAllTodos obtiene todos los todos (para HTMX)
func (h *TodoHandlerTempl) GetAllTodos(c *gin.Context) {
	filter := c.Query("filter")
	filteredTodos := h.getFilteredTodos(filter)
	
	data := templates.TodoListData{
		Todos: filteredTodos,
	}
	
	tmpl := templates.GetTodoListTemplate()
	tmpl.Execute(c.Writer, data)
}

// CreateTodo crea un nuevo todo (para HTMX)
func (h *TodoHandlerTempl) CreateTodo(c *gin.Context) {
	var todoReq models.TodoRequest
	var err error
	
	// Intentar bindear como JSON primero
	err = c.ShouldBindJSON(&todoReq)
	if err != nil {
		// Si falla JSON, intentar como form data
		err = c.ShouldBind(&todoReq)
		if err != nil {
			c.String(http.StatusBadRequest, "No se pudo procesar los datos: "+err.Error())
			return
		}
	}
	
	// Validar datos
	if todoReq.Title == "" {
		c.String(http.StatusBadRequest, "El título es requerido")
		return
	}
	
	// Crear el todo
	todo := models.Todo{
		ID:          h.nextID,
		Title:       todoReq.Title,
		Description: todoReq.Description,
		Completed:   todoReq.Completed,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	
	h.todos = append(h.todos, todo)
	h.nextID++
	
	// Redirigir para recargar la página
	c.Redirect(http.StatusSeeOther, "/")
}

// GetTodoByID obtiene un todo por ID
func (h *TodoHandlerTempl) GetTodoByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "ID inválido")
		return
	}
	
	for _, todo := range h.todos {
		if todo.ID == id {
			c.JSON(http.StatusOK, models.Response{
				Success: true,
				Message: "Todo encontrado",
				Data:    todo,
			})
			return
		}
	}
	
	c.String(http.StatusNotFound, "Todo no encontrado")
}

// UpdateTodo actualiza un todo existente (para HTMX)
func (h *TodoHandlerTempl) UpdateTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "ID inválido")
		return
	}
	
	var todoReq models.TodoRequest
	
	// Intentar bindear como JSON primero
	err = c.ShouldBindJSON(&todoReq)
	if err != nil {
		// Si falla JSON, intentar como form data
		err = c.ShouldBind(&todoReq)
		if err != nil {
			c.String(http.StatusBadRequest, "No se pudo procesar los datos: "+err.Error())
			return
		}
	}
	
	// Validar datos
	if todoReq.Title == "" {
		c.String(http.StatusBadRequest, "El título es requerido")
		return
	}
	
	// Buscar y actualizar el todo
	for i, todo := range h.todos {
		if todo.ID == id {
			h.todos[i].Title = todoReq.Title
			h.todos[i].Description = todoReq.Description
			h.todos[i].Completed = todoReq.Completed
			h.todos[i].UpdatedAt = time.Now()
			
			// Redirigir para recargar la página
			c.Redirect(http.StatusSeeOther, "/")
			return
		}
	}
	
	c.String(http.StatusNotFound, "Todo no encontrado")
}

// DeleteTodo elimina un todo (para HTMX)
func (h *TodoHandlerTempl) DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "ID inválido")
		return
	}
	
	for i, todo := range h.todos {
		if todo.ID == id {
			h.todos = append(h.todos[:i], h.todos[i+1:]...)
			
			// Redirigir para recargar la página
			c.Redirect(http.StatusSeeOther, "/")
			return
		}
	}
	
	c.String(http.StatusNotFound, "Todo no encontrado")
}

// GetEditModal muestra el modal de edición
func (h *TodoHandlerTempl) GetEditModal(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "ID inválido")
		return
	}
	
	for _, todo := range h.todos {
		if todo.ID == id {
			tmpl := templates.GetEditModalTemplate()
			tmpl.Execute(c.Writer, todo)
			return
		}
	}
	
	c.String(http.StatusNotFound, "Todo no encontrado")
}

// CloseModal cierra el modal
func (h *TodoHandlerTempl) CloseModal(c *gin.Context) {
	c.String(http.StatusOK, "")
}

// HealthCheck verifica el estado de la aplicación
func (h *TodoHandlerTempl) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"message":   "Todo API is running with Templ + HTMX",
		"framework": "Templ + HTMX",
		"version":   "v0.2.543",
		"formats":   []string{"JSON", "Form Data"},
	})
}

// CreateTodoFlexible crea un todo aceptando cualquier formato
func (h *TodoHandlerTempl) CreateTodoFlexible(c *gin.Context) {
	var todoReq models.TodoRequest
	var err error
	
	// Intentar bindear como JSON primero
	err = c.ShouldBindJSON(&todoReq)
	if err != nil {
		// Si falla JSON, intentar como form data
		err = c.ShouldBind(&todoReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "No se pudo procesar los datos. Acepta JSON o Form Data",
				"json_error": c.ShouldBindJSON(&todoReq).Error(),
				"form_error": err.Error(),
			})
			return
		}
	}
	
	// Validar datos
	if todoReq.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El título es requerido",
		})
		return
	}
	
	// Crear el todo
	todo := models.Todo{
		ID:          h.nextID,
		Title:       todoReq.Title,
		Description: todoReq.Description,
		Completed:   todoReq.Completed,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	
	h.todos = append(h.todos, todo)
	h.nextID++
	
	// Devolver el todo creado
	c.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Todo creado exitosamente",
		Data:    todo,
	})
}

// Funciones auxiliares

// calculateStats calcula las estadísticas del todo list
func (h *TodoHandlerTempl) calculateStats() templates.TodoStats {
	total := len(h.todos)
	completed := 0
	
	for _, todo := range h.todos {
		if todo.Completed {
			completed++
		}
	}
	
	return templates.TodoStats{
		Total:     total,
		Pending:   total - completed,
		Completed: completed,
	}
}

// getFilteredTodos obtiene todos filtrados
func (h *TodoHandlerTempl) getFilteredTodos(filter string) []models.Todo {
	switch filter {
	case "completed":
		var completed []models.Todo
		for _, todo := range h.todos {
			if todo.Completed {
				completed = append(completed, todo)
			}
		}
		return completed
	case "pending":
		var pending []models.Todo
		for _, todo := range h.todos {
			if !todo.Completed {
				pending = append(pending, todo)
			}
		}
		return pending
	default:
		return h.todos
	}
}
