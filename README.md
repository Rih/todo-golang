# Todo List - Aplicación Completa

Una aplicación completa de gestión de tareas desarrollada en Go que incluye:
- **API REST** para gestión de datos
- **Interfaz web moderna** y responsiva
- **Servidor HTTP nativo** de Go

## 🚀 Características

### Backend (API REST)
- **CRUD completo**: Crear, leer, actualizar y eliminar todos
- **API REST**: Endpoints HTTP estándar
- **JSON**: Comunicación mediante JSON
- **CORS**: Soporte para Cross-Origin Resource Sharing
- **Logging**: Registro de peticiones HTTP
- **Health Check**: Endpoint para verificar el estado de la API

### Frontend (Página Web)
- **Interfaz moderna**: Diseño limpio y profesional
- **Responsive**: Adaptable a móviles y desktop
- **CRUD completo**: Gestión visual de tareas
- **Filtros**: Ver todas, pendientes o completadas
- **Estadísticas**: Contadores en tiempo real
- **Notificaciones**: Feedback visual para todas las acciones

## 📋 Estructura del Proyecto

```
todo/
├── go.mod                 # Dependencias del proyecto
├── main.go               # Punto de entrada de la aplicación
├── models/
│   └── todo.go          # Estructuras de datos
├── handlers/
│   └── todo.go          # Lógica de negocio y handlers HTTP
├── routes/
│   └── routes.go        # Configuración de rutas y middleware
├── web/                  # Interfaz web
│   ├── index.html       # Página principal
│   ├── styles.css       # Estilos CSS
│   ├── script.js        # Lógica JavaScript
│   └── README.md        # Documentación web
└── README.md            # Documentación principal
```

## 🛠️ Instalación y Uso

### Prerrequisitos

- Go 1.21 o superior
- Git

### Instalación

1. **Clonar o navegar al directorio del proyecto:**
   ```bash
   cd /Users/rodrihgod/projects/golang/todo
   ```

2. **Instalar dependencias:**
   ```bash
   go mod tidy
   ```

3. **Ejecutar la aplicación:**
   ```bash
   go run main.go
   ```

4. **Acceder a la aplicación:**
   - **Página web**: `http://localhost:8080`
   - **API REST**: `http://localhost:8080/api/v1`

## 📡 Endpoints de la API

### Base URL
```
http://localhost:8080/api/v1
```

## 🌐 Interfaz Web

La aplicación incluye una interfaz web completa accesible en `http://localhost:8080` que permite:

- **Gestión visual de tareas**: Crear, editar, eliminar y marcar tareas
- **Filtros inteligentes**: Ver todas, pendientes o completadas
- **Estadísticas en tiempo real**: Contadores automáticos
- **Diseño responsivo**: Funciona en móviles y desktop
- **Notificaciones**: Feedback visual para todas las acciones

### Características de la Interfaz
- ✅ **Sin dependencias externas**: Solo HTML, CSS y JavaScript vanilla
- ✅ **Servidor nativo de Go**: No requiere servidor web adicional
- ✅ **Diseño moderno**: Gradientes, animaciones y efectos hover
- ✅ **Completamente funcional**: Todas las operaciones CRUD

### Endpoints Disponibles

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/todos` | Obtener todos los todos |
| POST | `/todos` | Crear un nuevo todo |
| GET | `/todos/{id}` | Obtener un todo por ID |
| PUT | `/todos/{id}` | Actualizar un todo |
| DELETE | `/todos/{id}` | Eliminar un todo |
| GET | `/health` | Health check |

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

### 3. Obtener un todo específico
```bash
curl http://localhost:8080/api/v1/todos/1
```

### 4. Actualizar un todo
```bash
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Completar proyecto",
    "description": "Finalizar la implementación del todo list",
    "completed": true
  }'
```

### 5. Eliminar un todo
```bash
curl -X DELETE http://localhost:8080/api/v1/todos/1
```

### 6. Health check
```bash
curl http://localhost:8080/api/v1/health
```

## 📊 Estructura de Datos

### Todo
```json
{
  "id": 1,
  "title": "Título de la tarea",
  "description": "Descripción de la tarea",
  "completed": false,
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

### Respuesta de la API
```json
{
  "success": true,
  "message": "Mensaje descriptivo",
  "data": { ... }
}
```

## 🔧 Configuración

### Variables de Entorno

- `PORT`: Puerto donde correrá la aplicación (por defecto: 8080)

### Ejemplo de configuración:
```bash
export PORT=3000
go run main.go
```

## 🚀 Despliegue

### Compilar para producción:
```bash
go build -o todo-api main.go
./todo-api
```

### Con Docker (opcional):
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o todo-api main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/todo-api .
EXPOSE 8080
CMD ["./todo-api"]
```

## 🧪 Testing

Para probar la API puedes usar herramientas como:
- **curl** (línea de comandos)
- **Postman** (interfaz gráfica)
- **Insomnia** (interfaz gráfica)
- **httpie** (línea de comandos)

## 📚 Dependencias

- `github.com/gorilla/mux` - Router HTTP
- `github.com/gorilla/handlers` - Middleware para HTTP

## 🤝 Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT.
