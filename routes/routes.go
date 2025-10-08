package routes

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"todo-list/handlers"

	"github.com/gorilla/mux"
)

// SetupRoutes configura todas las rutas de la aplicación
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	
	// Crear instancia del handler
	todoHandler := handlers.NewTodoHandler()
	
	// Middleware para logging
	router.Use(loggingMiddleware)
	
	// Middleware para CORS
	router.Use(corsMiddleware)
	
	// Rutas de la API
	api := router.PathPrefix("/api/v1").Subrouter()
	
	// Rutas de todos
	api.HandleFunc("/todos", todoHandler.GetAllTodos).Methods("GET")
	api.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	api.HandleFunc("/todos/{id}", todoHandler.GetTodoByID).Methods("GET")
	api.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PUT")
	api.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods("DELETE")
	
	// Ruta de health check
	api.HandleFunc("/health", healthCheck).Methods("GET")
	
	// Servir archivos estáticos de la página web
	router.PathPrefix("/").Handler(http.HandlerFunc(serveWebFiles))
	
	return router
}

// serveWebFiles sirve los archivos estáticos de la página web
func serveWebFiles(w http.ResponseWriter, r *http.Request) {
	// Si es una petición a la API, no servir archivos estáticos
	if strings.HasPrefix(r.URL.Path, "/api/") {
		http.NotFound(w, r)
		return
	}
	
	// Obtener la ruta del archivo
	path := r.URL.Path
	
	// Si la ruta es "/" o termina en "/", servir index.html
	if path == "/" || strings.HasSuffix(path, "/") {
		path = "/index.html"
	}
	
	// Construir la ruta completa del archivo
	filePath := filepath.Join("web", path)
	
	// Verificar si el archivo existe
	if !fileExists(filePath) {
		// Si no existe, servir index.html (para SPA routing)
		filePath = "web/index.html"
	}
	
	// Determinar el Content-Type basado en la extensión
	contentType := getContentType(filePath)
	w.Header().Set("Content-Type", contentType)
	
	// Servir el archivo
	http.ServeFile(w, r, filePath)
}

// fileExists verifica si un archivo existe
func fileExists(path string) bool {
	_, err := http.Dir(".").Open(path)
	return err == nil
}

// getContentType determina el Content-Type basado en la extensión del archivo
func getContentType(filePath string) string {
	ext := strings.ToLower(filepath.Ext(filePath))
	
	switch ext {
	case ".html":
		return "text/html; charset=utf-8"
	case ".css":
		return "text/css; charset=utf-8"
	case ".js":
		return "application/javascript; charset=utf-8"
	case ".json":
		return "application/json; charset=utf-8"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".svg":
		return "image/svg+xml"
	case ".ico":
		return "image/x-icon"
	case ".woff":
		return "font/woff"
	case ".woff2":
		return "font/woff2"
	case ".ttf":
		return "font/ttf"
	default:
		return "text/plain; charset=utf-8"
	}
}

// loggingMiddleware registra las peticiones HTTP
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// corsMiddleware maneja CORS
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

// healthCheck verifica el estado de la aplicación
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok","message":"Todo API is running"}`))
}