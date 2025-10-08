package routes

import (
	"fmt"
	"todo-list/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRoutesTempl configura todas las rutas usando Gin + Templ + HTMX
func SetupRoutesTempl() *gin.Engine {
	// Configurar Gin en modo debug para desarrollo
	// gin.SetMode(gin.ReleaseMode)
	
	// Crear router de Gin
	router := gin.Default()
	
	// Middleware de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "HX-Request"},
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
	todoHandler := handlers.NewTodoHandlerTempl()
	
	// Servir archivos estáticos
	router.Static("/static", "./web")
	
	// Ruta principal - página del todo list
	router.GET("/", todoHandler.GetHomePage)
	
	// Grupo de rutas para la API (HTMX)
	api := router.Group("/api")
	{
		// Rutas de todos para HTMX
		api.GET("/todos", todoHandler.GetAllTodos)
		api.POST("/todos", todoHandler.CreateTodo)
		api.GET("/todos/:id", todoHandler.GetTodoByID)
		api.PUT("/todos/:id", todoHandler.UpdateTodo)
		api.DELETE("/todos/:id", todoHandler.DeleteTodo)
		
		// Rutas para modales
		api.GET("/todos/:id/edit", todoHandler.GetEditModal)
		api.GET("/close-modal", todoHandler.CloseModal)
		
		// Ruta de health check
		api.GET("/health", todoHandler.HealthCheck)
	}
	
	return router
}
