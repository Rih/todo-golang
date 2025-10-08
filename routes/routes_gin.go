package routes

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"todo-list/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRoutesGin configura todas las rutas usando Gin framework
func SetupRoutesGin() *gin.Engine {
	// Configurar Gin en modo release para producción
	// gin.SetMode(gin.ReleaseMode)
	
	// Crear router de Gin
	router := gin.Default()
	
	// Middleware de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	
	// Middleware de logging personalizado
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s %s %s %d %s %s\n",
			param.ClientIP,
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
		)
	}))
	
	// Middleware de recuperación
	router.Use(gin.Recovery())
	
	// Crear instancia del handler
	todoHandler := handlers.NewTodoHandlerGin()
	
	// Grupo de rutas para la API
	api := router.Group("/api/v1")
	{
		// Rutas de todos
		api.GET("/todos", todoHandler.GetAllTodos)
		api.POST("/todos", todoHandler.CreateTodo)
		api.GET("/todos/:id", todoHandler.GetTodoByID)
		api.PUT("/todos/:id", todoHandler.UpdateTodo)
		api.DELETE("/todos/:id", todoHandler.DeleteTodo)
		
		// Ruta de health check
		api.GET("/health", todoHandler.HealthCheck)
	}
	
	// Servir archivos estáticos de la página web
	router.NoRoute(gin.WrapF(serveWebFilesGin))
	
	return router
}

// serveWebFilesGin sirve los archivos estáticos de la página web
func serveWebFilesGin(w http.ResponseWriter, r *http.Request) {
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
	if !fileExistsGin(filePath) {
		// Si no existe, servir index.html (para SPA routing)
		filePath = "web/index.html"
	}
	
	// Determinar el Content-Type basado en la extensión
	contentType := getContentTypeGin(filePath)
	w.Header().Set("Content-Type", contentType)
	
	// Servir el archivo
	http.ServeFile(w, r, filePath)
}

// fileExistsGin verifica si un archivo existe
func fileExistsGin(path string) bool {
	_, err := http.Dir(".").Open(path)
	return err == nil
}

// getContentTypeGin determina el Content-Type basado en la extensión del archivo
func getContentTypeGin(filePath string) string {
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
