package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"todo-list/models"

	"github.com/gorilla/mux"
)

// TodoHandler maneja las operaciones CRUD de todos
type TodoHandler struct {
	todos []models.Todo
	nextID int
}

// NewTodoHandler crea una nueva instancia del handler
func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		todos:  make([]models.Todo, 0),
		nextID: 1,
	}
}

// GetAllTodos obtiene todos los todos
func (h *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := models.Response{
		Success: true,
		Message: "Todos obtenidos exitosamente",
		Data:    h.todos,
	}
	
	json.NewEncoder(w).Encode(response)
}

// GetTodoByID obtiene un todo por ID
func (h *TodoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	
	for _, todo := range h.todos {
		if todo.ID == id {
			response := models.Response{
				Success: true,
				Message: "Todo encontrado",
				Data:    todo,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	
	response := models.Response{
		Success: false,
		Message: "Todo no encontrado",
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

// CreateTodo crea un nuevo todo
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var todoReq models.TodoRequest
	if err := json.NewDecoder(r.Body).Decode(&todoReq); err != nil {
		response := models.Response{
			Success: false,
			Message: "Datos inválidos",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	if todoReq.Title == "" {
		response := models.Response{
			Success: false,
			Message: "El título es requerido",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
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
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// UpdateTodo actualiza un todo existente
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	
	var todoReq models.TodoRequest
	if err := json.NewDecoder(r.Body).Decode(&todoReq); err != nil {
		response := models.Response{
			Success: false,
			Message: "Datos inválidos",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
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
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	
	response := models.Response{
		Success: false,
		Message: "Todo no encontrado",
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

// DeleteTodo elimina un todo
func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	
	for i, todo := range h.todos {
		if todo.ID == id {
			h.todos = append(h.todos[:i], h.todos[i+1:]...)
			
			response := models.Response{
				Success: true,
				Message: "Todo eliminado exitosamente",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	
	response := models.Response{
		Success: false,
		Message: "Todo no encontrado",
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}
