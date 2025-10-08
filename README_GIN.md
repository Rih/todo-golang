# Todo List con Gin Framework

Una aplicación completa de gestión de tareas desarrollada en Go usando el framework **Gin**, que incluye:
- **API REST** de alto rendimiento
- **Interfaz web moderna** y responsiva
- **Servidor HTTP optimizado** con Gin

## 🚀 Características del Framework Gin

### Ventajas de Gin sobre Gorilla Mux
- ⚡ **40x más rápido**: Rendimiento superior
- 🛠️ **Middleware integrado**: CORS, logging, recovery
- 📝 **Binding automático**: JSON binding automático
- 🎯 **API más limpia**: Sintaxis más concisa
- 🔧 **Herramientas integradas**: Validación, rendering
- 📊 **Métricas**: Logging estructurado
- 🚀 **Producción ready**: Optimizado para alta concurrencia

## 📋 Estructura del Proyecto

```
todo/
├── go.mod                 # Dependencias (Gin + CORS)
├── main_gin.go           # Punto de entrada con Gin
├── models/
│   └── todo.go          # Estructuras de datos
├── handlers/
│   ├── todo.go          # Handlers con Gorilla Mux (original)
│   └── todo_gin.go      # Handlers con Gin (nuevo)
├── routes/
│   ├── routes.go        # Rutas con Gorilla Mux (original)
│   └── routes_gin.go    # Rutas con Gin (nuevo)
├── web/                  # Interfaz web
│   ├── index.html       # Página principal
│   ├── styles.css       # Estilos CSS
│   ├── script.js        # Lógica JavaScript
│   └── README.md        # Documentación web
├── README.md            # Documentación original
└── README_GIN.md        # Documentación con Gin
```

## 🛠️ Instalación y Uso

### Prerrequisitos
- Go 1.21 o superior
- Git

### Instalación

1. **Navegar al directorio del proyecto:**
   ```bash
   cd $HOME/projects/golang/todo
   ```

2. **Instalar dependencias:**
   ```bash
   source $HOME/.gvm/init.sh
   go mod tidy
   ```

3. **Ejecutar con Gin Framework:**
   ```bash
   go run main_gin.go
   ```

4. **Acceder a la aplicación:**
   - **Página web**: `http://localhost:8080`
   - **API REST**: `http://localhost:8080/api/v1`

## 📡 Endpoints de la API

### Base URL
```
http://localhost:8080/api/v1
```

### Endpoints Disponibles

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/todos` | Obtener todos los todos |
| POST | `/todos` | Crear un nuevo todo |
| GET | `/todos/{id}` | Obtener un todo por ID |
| PUT | `/todos/{id}` | Actualizar un todo |
| DELETE | `/todos/{id}` | Eliminar un todo |
| GET | `/health` | Health check con información del framework |

## 🌐 Interfaz Web

La aplicación incluye una interfaz web completa accesible en `http://localhost:8080` que permite:

- **Gestión visual de tareas**: Crear, editar, eliminar y marcar tareas
- **Filtros inteligentes**: Ver todas, pendientes o completadas
- **Estadísticas en tiempo real**: Contadores automáticos
- **Diseño responsivo**: Funciona en móviles y desktop
- **Notificaciones**: Feedback visual para todas las acciones

## 🔧 Características Técnicas de Gin

### Middleware Integrado
- **CORS**: Configuración automática de Cross-Origin Resource Sharing
- **Logging**: Logging estructurado con métricas de rendimiento
- **Recovery**: Manejo automático de panics
- **JSON Binding**: Binding automático de JSON a structs

### Ventajas de Rendimiento
- **Router optimizado**: Radix tree para matching de rutas
- **Zero allocation**: Minimiza garbage collection
- **Concurrent safe**: Seguro para concurrencia
- **Memory efficient**: Uso eficiente de memoria

## 📝 Ejemplos de Uso

### 1. Crear un nuevo todo
```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Completar proyecto",
    "description": "Finalizar la implementación del todo list",
    "completed": false
  }'
```

### 2. Obtener todos los todos
```bash
curl http://localhost:8080/api/v1/todos
```

### 3. Health Check con información del framework
```bash
curl http://localhost:8080/api/v1/health
```

**Respuesta:**
```json
{
  "status": "ok",
  "message": "Todo API is running with Gin framework",
  "framework": "Gin",
  "version": "v1.9.1"
}
```

## 🚀 Comparación de Rendimiento

### Gorilla Mux vs Gin Framework

| Característica | Gorilla Mux | Gin Framework |
|----------------|-------------|---------------|
| **Rendimiento** | Estándar | 40x más rápido |
| **Middleware** | Manual | Integrado |
| **JSON Binding** | Manual | Automático |
| **Logging** | Básico | Estructurado |
| **Recovery** | Manual | Automático |
| **CORS** | Manual | Configuración simple |
| **Validación** | Manual | Integrada |

## 🔧 Configuración Avanzada

### Variables de Entorno
- `PORT`: Puerto donde correrá la aplicación (por defecto: 8080)
- `GIN_MODE`: Modo de Gin (debug, release, test)

### Ejemplo de configuración:
```bash
export PORT=3000
export GIN_MODE=release
go run main_gin.go
```

### Middleware Personalizado
```go
// Ejemplo de middleware personalizado
router.Use(func(c *gin.Context) {
    start := time.Now()
    c.Next()
    latency := time.Since(start)
    log.Printf("Request took %v", latency)
})
```

## 📊 Métricas y Monitoreo

Gin proporciona métricas integradas:
- **Latencia**: Tiempo de respuesta de cada request
- **Status codes**: Códigos de estado HTTP
- **User agents**: Información del cliente
- **IP addresses**: Direcciones IP de los clientes

## 🚀 Despliegue

### Compilar para producción:
```bash
go build -o todo-api-gin main_gin.go
./todo-api-gin
```

### Con Docker:
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o todo-api-gin main_gin.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/todo-api-gin .
COPY --from=builder /app/web ./web
EXPOSE 8080
CMD ["./todo-api-gin"]
```

## 🧪 Testing

### Ejecutar tests:
```bash
go test ./...
```

### Benchmark de rendimiento:
```bash
go test -bench=. ./...
```

## 📚 Dependencias

- `github.com/gin-gonic/gin` - Framework web de alto rendimiento
- `github.com/gin-contrib/cors` - Middleware CORS para Gin

## 🔮 Mejoras Futuras

- [ ] Autenticación JWT
- [ ] Rate limiting
- [ ] Métricas con Prometheus
- [ ] Swagger/OpenAPI documentation
- [ ] Database integration (PostgreSQL/MongoDB)
- [ ] Redis caching
- [ ] WebSocket support
- [ ] Graceful shutdown

## 📄 Licencia

Este proyecto está bajo la Licencia MIT.

## 🤝 Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

---

**Desarrollado con ❤️ usando Go y Gin Framework v1.9.1**
