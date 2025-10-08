package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo-list/routes"
)

func main() {
	// Configurar rutas
	router := routes.SetupRoutes()
	
	// Obtener puerto del entorno o usar 8080 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	// Mostrar información de inicio
	fmt.Printf("🚀 Todo List API iniciando en puerto %s\n", port)
	fmt.Println("📋 Endpoints disponibles:")
	fmt.Println("  GET    /api/v1/todos     - Obtener todos los todos")
	fmt.Println("  POST   /api/v1/todos     - Crear un nuevo todo")
	fmt.Println("  GET    /api/v1/todos/{id} - Obtener un todo por ID")
	fmt.Println("  PUT    /api/v1/todos/{id} - Actualizar un todo")
	fmt.Println("  DELETE /api/v1/todos/{id} - Eliminar un todo")
	fmt.Println("  GET    /api/v1/health    - Health check")
	fmt.Println("")
	fmt.Println("🌐 Página web disponible en:")
	fmt.Printf("  http://localhost:%s\n", port)
	fmt.Println("")
	fmt.Println("📁 Archivos estáticos servidos desde: ./web/")
	fmt.Printf("🌐 Servidor corriendo en: http://localhost:%s\n", port)
	
	// Iniciar servidor
	log.Fatal(http.ListenAndServe(":"+port, router))
}
