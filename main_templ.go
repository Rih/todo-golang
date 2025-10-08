package main

import (
	"fmt"
	"log"
	"os"
	"todo-list/routes"
)

func main() {
	// Configurar rutas con Gin + Templ + HTMX
	router := routes.SetupRoutesTempl()
	
	// Obtener puerto del entorno o usar 8080 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	// Mostrar información de inicio
	fmt.Printf("🚀 Todo List con Templ + HTMX iniciando en puerto %s\n", port)
	fmt.Println("📋 Endpoints disponibles:")
	fmt.Println("  GET    /                 - Página principal del Todo List")
	fmt.Println("  GET    /api/todos        - Obtener todos los todos (HTMX)")
	fmt.Println("  POST   /api/todos        - Crear un nuevo todo (HTMX)")
	fmt.Println("  POST   /api/todos/flexible - Crear todo (acepta JSON y Form Data)")
	fmt.Println("  GET    /api/todos/{id}    - Obtener un todo por ID")
	fmt.Println("  PUT    /api/todos/{id}    - Actualizar un todo (HTMX)")
	fmt.Println("  DELETE /api/todos/{id}   - Eliminar un todo (HTMX)")
	fmt.Println("  GET    /api/todos/{id}/edit - Modal de edición (HTMX)")
	fmt.Println("  GET    /api/close-modal  - Cerrar modal (HTMX)")
	fmt.Println("  GET    /api/health       - Health check")
	fmt.Println("")
	fmt.Println("🌐 Página web disponible en:")
	fmt.Printf("  http://localhost:%s\n", port)
	fmt.Println("")
	fmt.Println("📁 Archivos estáticos servidos desde: ./web/")
	fmt.Println("⚡ Framework: Gin + Templ + HTMX")
	fmt.Println("🎨 Templates: Templ v0.2.543")
	fmt.Println("🔄 Interactividad: HTMX v1.9.10")
	fmt.Printf("🌐 Servidor corriendo en: http://localhost:%s\n", port)
	
	// Iniciar servidor con Gin
	log.Fatal(router.Run(":" + port))
}
