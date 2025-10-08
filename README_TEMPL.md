# Todo List con Templates + HTMX

Una aplicación completa de gestión de tareas desarrollada en Go usando **Templates nativos** y **HTMX**, que incluye:
- **API REST** de alto rendimiento con Gin
- **Templates HTML** nativos de Go
- **Interactividad** con HTMX (sin JavaScript vanilla)
- **Servidor HTTP optimizado** con Gin

## 🚀 Características de Templates + HTMX

### Ventajas de esta implementación
- ⚡ **Sin JavaScript**: Interactividad pura con HTMX
- 🛠️ **Templates nativos**: Usando html/template de Go
- 📝 **Server-side rendering**: Renderizado en el servidor
- 🎯 **API más limpia**: Sintaxis más concisa
- 🔧 **Herramientas integradas**: Validación, binding automático
- 📊 **Métricas**: Logging estructurado
- 🚀 **Producción ready**: Optimizado para alta concurrencia
- 🌐 **Progressive Enhancement**: Funciona sin JavaScript

## 📋 Estructura del Proyecto

```
todo/
├── go.mod                 # Dependencias (Gin + CORS)
├── main_templ.go          # Punto de entrada con Templates + HTMX
├── models/
│   └── todo.go           # Estructuras de datos
├── handlers/
│   ├── todo.go           # Handlers con Gorilla Mux (original)
│   ├── todo_gin.go       # Handlers con Gin (Gin puro)
│   └── todo_templ.go     # Handlers con Templates + HTMX (nuevo)
├── routes/
│   ├── routes.go         # Rutas con Gorilla Mux (original)
│   ├── routes_gin.go     # Rutas con Gin (Gin puro)
│   └── routes_templ.go   # Rutas con Templates + HTMX (nuevo)
├── templates/
│   └── layout.go         # Templates HTML nativos
├── web/                   # Archivos estáticos
│   ├── styles.css        # Estilos CSS
│   └── README.md         # Documentación web
├── README.md             # Documentación original
├── README_GIN.md         # Documentación con Gin
└── README_TEMPL.md       # Documentación con Templates + HTMX
```

## 🛠️ Instalación y Uso

### Prerrequisitos
- Go 1.21 o superior
- Git

### Instalación

1. **Navegar al directorio del proyecto:**
   ```bash
   cd /Users/rodrihgod/projects/golang/todo
   ```

2. **Instalar dependencias:**
   ```bash
   source /Users/rodrihgod/.gvm/init.sh
   go mod tidy
   ```

3. **Ejecutar con Templates + HTMX:**
   ```bash
   go run main_templ.go
   ```

4. **Acceder a la aplicación:**
   - **Página web**: `http://localhost:8080`
   - **API REST**: `http://localhost:8080/api`

## 📡 Endpoints de la API

### Base URL
```
http://localhost:8080/api
```

### Endpoints Disponibles

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/` | Página principal del Todo List |
| GET | `/api/todos` | Obtener todos los todos (HTMX) |
| POST | `/api/todos` | Crear un nuevo todo (HTMX) |
| GET | `/api/todos/{id}` | Obtener un todo por ID |
| PUT | `/api/todos/{id}` | Actualizar un todo (HTMX) |
| DELETE | `/api/todos/{id}` | Eliminar un todo (HTMX) |
| GET | `/api/todos/{id}/edit` | Modal de edición (HTMX) |
| GET | `/api/close-modal` | Cerrar modal (HTMX) |
| GET | `/api/health` | Health check |

## 🌐 Interfaz Web con HTMX

La aplicación incluye una interfaz web completamente funcional accesible en `http://localhost:8080` que permite:

- **Gestión visual de tareas**: Crear, editar, eliminar y marcar tareas
- **Filtros inteligentes**: Ver todas, pendientes o completadas
- **Estadísticas en tiempo real**: Contadores automáticos
- **Diseño responsivo**: Funciona en móviles y desktop
- **Notificaciones**: Feedback visual para todas las acciones
- **Sin JavaScript**: Todo funciona con HTMX

### Características de HTMX
- ✅ **Progressive Enhancement**: Funciona sin JavaScript
- ✅ **Server-side rendering**: Templates renderizados en Go
- ✅ **Interactividad**: HTMX maneja todas las interacciones
- ✅ **Performance**: Menos JavaScript, más velocidad
- ✅ **SEO friendly**: Contenido renderizado en servidor

## 🔧 Características Técnicas

### Templates HTML Nativos
- **html/template**: Motor de templates nativo de Go
- **Funciones personalizadas**: formatDate, etc.
- **Estructuras de datos**: PageData, TodoListData
- **Reutilización**: Templates modulares

### HTMX Integration
- **hx-post**: Envío de formularios
- **hx-get**: Carga de contenido
- **hx-put**: Actualización de datos
- **hx-delete**: Eliminación de elementos
- **hx-target**: Selector de destino
- **hx-swap**: Estrategia de intercambio

### Ventajas de Rendimiento
- **Server-side rendering**: Contenido inicial rápido
- **Partial updates**: Solo actualiza lo necesario
- **No JavaScript bundle**: Menos peso
- **Caching**: Templates compilados
- **Memory efficient**: Uso eficiente de memoria

## 📝 Ejemplos de Uso

### 1. Crear un nuevo todo (HTMX)
```html
<form hx-post="/api/todos" hx-target="#todoList" hx-swap="outerHTML">
    <input type="text" name="title" placeholder="Título de la tarea" required>
    <textarea name="description" placeholder="Descripción (opcional)"></textarea>
    <button type="submit">Agregar Tarea</button>
</form>
```

### 2. Actualizar un todo (HTMX)
```html
<button hx-put="/api/todos/1" 
      hx-vals='{"completed": true}' 
      hx-target="#todoList" 
      hx-swap="outerHTML">
    Completar
</button>
```

### 3. Eliminar un todo (HTMX)
```html
<button hx-delete="/api/todos/1" 
        hx-target="#todoList" 
        hx-swap="outerHTML"
        hx-confirm="¿Estás seguro?">
    Eliminar
</button>
```

### 4. Filtrar todos (HTMX)
```html
<button hx-get="/api/todos?filter=pending" 
        hx-target="#todoList">
    Pendientes
</button>
```

## 🚀 Comparación de Implementaciones

### Vanilla JavaScript vs Templates + HTMX

| Característica | Vanilla JS | Templates + HTMX |
|----------------|------------|------------------|
| **JavaScript** | Requerido | Opcional |
| **Bundle size** | Grande | Mínimo |
| **SEO** | Limitado | Excelente |
| **Performance** | Dependiente | Optimizado |
| **Mantenimiento** | Complejo | Simple |
| **Progressive Enhancement** | No | Sí |
| **Server-side rendering** | No | Sí |

## 🔧 Configuración Avanzada

### Variables de Entorno
- `PORT`: Puerto donde correrá la aplicación (por defecto: 8080)
- `GIN_MODE`: Modo de Gin (debug, release, test)

### Ejemplo de configuración:
```bash
export PORT=3000
export GIN_MODE=release
go run main_templ.go
```

### Templates Personalizados
```go
// Agregar función personalizada
funcMap := template.FuncMap{
    "formatDate": formatDate,
    "uppercase": strings.ToUpper,
}

tmpl := template.Must(template.New("layout").Funcs(funcMap).Parse(tmplStr))
```

## 📊 Métricas y Monitoreo

Gin proporciona métricas integradas:
- **Latencia**: Tiempo de respuesta de cada request
- **Status codes**: Códigos de estado HTTP
- **User agents**: Información del cliente
- **IP addresses**: Direcciones IP de los clientes
- **HTMX requests**: Requests específicos de HTMX

## 🧪 Testing

### Ejecutar tests:
```bash
go test ./...
```

### Test de templates:
```go
func TestTemplate(t *testing.T) {
    tmpl := templates.GetLayoutTemplate()
    data := templates.PageData{
        Title: "Test",
        Todos: []models.Todo{},
        Stats: templates.TodoStats{},
    }
    
    var buf bytes.Buffer
    err := tmpl.Execute(&buf, data)
    assert.NoError(t, err)
}
```

## 🚀 Despliegue

### Compilar para producción:
```bash
go build -o todo-api-templ main_templ.go
./todo-api-templ
```

### Con Docker:
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o todo-api-templ main_templ.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/todo-api-templ .
COPY --from=builder /app/web ./web
EXPOSE 8080
CMD ["./todo-api-templ"]
```

## 🔮 Mejoras Futuras

- [ ] **Templ framework**: Migrar a Templ para mejor performance
- [ ] **WebSockets**: Actualizaciones en tiempo real
- [ ] **Caching**: Redis para templates
- [ ] **CDN**: Servir assets estáticos
- [ ] **PWA**: Progressive Web App
- [ ] **Offline support**: Service workers
- [ ] **Database**: Persistencia de datos
- [ ] **Authentication**: Sistema de usuarios

## 📚 Dependencias

- `github.com/gin-gonic/gin` - Framework web de alto rendimiento
- `github.com/gin-contrib/cors` - Middleware CORS para Gin
- `html/template` - Motor de templates nativo de Go
- `HTMX` - Biblioteca para interactividad (CDN)

## 🎯 Casos de Uso

### Ideal para:
- ✅ **Aplicaciones simples**: CRUD básico
- ✅ **SEO importante**: Contenido renderizado en servidor
- ✅ **Performance**: Menos JavaScript, más velocidad
- ✅ **Mantenimiento**: Código más simple
- ✅ **Progressive Enhancement**: Funciona sin JavaScript

### No ideal para:
- ❌ **SPAs complejas**: Aplicaciones de una página complejas
- ❌ **Real-time**: Chat, notificaciones en tiempo real
- ❌ **Animaciones complejas**: Efectos visuales avanzados
- ❌ **Offline-first**: Aplicaciones que funcionan offline

## 📄 Licencia

Este proyecto está bajo la Licencia MIT.

## 🤝 Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

---

**Desarrollado con ❤️ usando Go, Gin, Templates nativos y HTMX**

## 🔗 Enlaces Útiles

- [HTMX Documentation](https://htmx.org/docs/)
- [Go Templates](https://pkg.go.dev/html/template)
- [Gin Framework](https://gin-gonic.com/)
- [Progressive Enhancement](https://developer.mozilla.org/en-US/docs/Glossary/Progressive_Enhancement)
