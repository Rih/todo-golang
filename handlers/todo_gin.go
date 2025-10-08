package handlers

import (
	"net/http"
	"strconv"
	"time"
	"todo-list/models"

	"github.com/gin-gonic/gin"
)

// TodoHandlerGin maneja las operaciones CRUD de todos usando Gin
type TodoHandlerGin struct {
	todos  []models.Todo
	nextID int
}

// NewTodoHandlerGin crea una nueva instancia del handler con Gin
func NewTodoHandlerGin() *TodoHandlerGin {
	return &TodoHandlerGin{
		todos:  make([]models.Todo, 0),
		nextID: 1,
	}
}

// GetAllTodos obtiene todos los todos
func (h *TodoHandlerGin) GetAllTodos(c *gin.Context) {
	response := models.Response{
		Success: true,
		Message: "Todos obtenidos exitosamente",
		Data:    h.todos,
	}
	
	c.JSON(http.StatusOK, response)
}

// GetTodoByID obtiene un todo por ID
func (h *TodoHandlerGin) GetTodoByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "ID inválido",
		})
		return
	}
	
	for _, todo := range h.todos {
		if todo.ID == id {
			response := models.Response{
				Success: true,
				Message: "Todo encontrado",
				Data:    todo,
			}
			c.JSON(http.StatusOK, response)
			return
		}
	}
	
	c.JSON(http.StatusNotFound, models.Response{
		Success: false,
		Message: "Todo no encontrado",
	})
}

// CreateTodo crea un nuevo todo
func (h *TodoHandlerGin) CreateTodo(c *gin.Context) {
	var todoReq models.TodoRequest
	if err := c.ShouldBindJSON(&todoReq); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Datos inválidos: " + err.Error(),
		})
		return
	}
	
	if todoReq.Title == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "El título es requerido",
		})
		return
	}
	
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
	
	response := models.Response{
		Success: true,
		Message: "Todo creado exitosamente",
		Data:    todo,
	}
	
	c.JSON(http.StatusCreated, response)
}

// UpdateTodo actualiza un todo existente
func (h *TodoHandlerGin) UpdateTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "ID inválido",
		})
		return
	}
	
	var todoReq models.TodoRequest
	if err := c.ShouldBindJSON(&todoReq); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Datos inválidos: " + err.Error(),
		})
		return
	}
	
	for i, todo := range h.todos {
		if todo.ID == id {
			h.todos[i].Title = todoReq.Title
			h.todos[i].Description = todoReq.Description
			h.todos[i].Completed = todoReq.Completed
			h.todos[i].UpdatedAt = time.Now()
			
			response := models.Response{
				Success: true,
				Message: "Todo actualizado exitosamente",
				Data:    h.todos[i],
			}
			c.JSON(http.StatusOK, response)
			return
		}
	}
	
	c.JSON(http.StatusNotFound, models.Response{
		Success: false,
		Message: "Todo no encontrado",
	})
}

// DeleteTodo elimina un todo
func (h *TodoHandlerGin) DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "ID inválido",
		})
		return
	}
	
	for i, todo := range h.todos {
		if todo.ID == id {
			h.todos = append(h.todos[:i], h.todos[i+1:]...)
			
			response := models.Response{
				Success: true,
				Message: "Todo eliminado exitosamente",
			}
			c.JSON(http.StatusOK, response)
			return
		}
	}
	
	c.JSON(http.StatusNotFound, models.Response{
		Success: false,
		Message: "Todo no encontrado",
	})
}

// HealthCheck verifica el estado de la aplicación
func (h *TodoHandlerGin) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Todo API is running with Gin framework",
		"framework": "Gin",
		"version": "v1.9.1",
	})
}
