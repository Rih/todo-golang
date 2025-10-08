# Todo List Web Interface

Una interfaz web moderna y responsiva para gestionar tareas que consume la API REST del Todo List desarrollada en Go.

## 🚀 Características

- **Interfaz Moderna**: Diseño limpio y profesional con gradientes y animaciones
- **Responsive**: Adaptable a dispositivos móviles y desktop
- **CRUD Completo**: Crear, leer, actualizar y eliminar tareas
- **Filtros**: Ver todas, pendientes o completadas
- **Estadísticas**: Contador de tareas en tiempo real
- **Validación**: Validación de formularios en tiempo real
- **Notificaciones**: Feedback visual para todas las acciones
- **Modal de Edición**: Interfaz intuitiva para editar tareas

## 📁 Estructura de Archivos

```
web/
├── index.html          # Página principal
├── styles.css          # Estilos CSS
├── script.js           # Lógica JavaScript
└── README.md           # Documentación
```

## 🛠️ Tecnologías Utilizadas

- **HTML5**: Estructura semántica
- **CSS3**: Estilos modernos con Flexbox y Grid
- **JavaScript ES6+**: Lógica de la aplicación
- **Font Awesome**: Iconos
- **Fetch API**: Comunicación con la API REST

## 🎨 Características del Diseño

### Colores y Temas
- Gradiente principal: `#667eea` → `#764ba2`
- Fondo degradado para mejor experiencia visual
- Tarjetas con sombras suaves
- Botones con efectos hover

### Responsive Design
- **Desktop**: Layout completo con sidebar
- **Tablet**: Adaptación de grid y formularios
- **Mobile**: Stack vertical optimizado

### Animaciones
- Transiciones suaves en hover
- Animaciones de entrada para modales
- Efectos de carga
- Notificaciones deslizantes

## 📱 Funcionalidades

### 1. Gestión de Tareas
- ✅ Crear nuevas tareas con título y descripción
- ✅ Marcar/desmarcar tareas como completadas
- ✅ Editar tareas existentes
- ✅ Eliminar tareas con confirmación

### 2. Filtros y Búsqueda
- 🔍 Ver todas las tareas
- ⏳ Ver solo tareas pendientes
- ✅ Ver solo tareas completadas

### 3. Estadísticas
- 📊 Contador total de tareas
- 📈 Tareas pendientes
- ✅ Tareas completadas

### 4. Interfaz de Usuario
- 🎨 Diseño moderno y limpio
- 📱 Completamente responsivo
- ⚡ Carga rápida y fluida
- 🔔 Notificaciones de estado

## 🚀 Instalación y Uso

### Prerrequisitos
- Servidor web (Apache, Nginx, o servidor de desarrollo)
- API del Todo List corriendo en `http://localhost:8080`

### Instalación

1. **Asegúrate de que la API esté corriendo:**
   ```bash
   cd /Users/rodrihgod/projects/golang/todo
   source /Users/rodrihgod/.gvm/init.sh
   go run main.go
   ```

2. **Abre la página web:**
   - Opción 1: Abrir `index.html` directamente en el navegador
   - Opción 2: Usar un servidor web local:
     ```bash
     # Con Python
     python -m http.server 8000
     # Luego ir a http://localhost:8000
     
     # Con Node.js (si tienes http-server instalado)
     npx http-server
     ```

3. **Acceder a la aplicación:**
   - Abrir `http://localhost:8000` (o el puerto que uses)
   - La aplicación se conectará automáticamente a la API

## 🔧 Configuración

### Cambiar URL de la API
Si tu API está en un puerto diferente, edita la variable en `script.js`:

```javascript
const API_BASE_URL = 'http://localhost:8080/api/v1';
```

### Personalizar Estilos
Los estilos están en `styles.css` y puedes modificar:
- Colores en las variables CSS
- Tamaños de fuente
- Espaciado y márgenes
- Animaciones y transiciones

## 📊 API Endpoints Utilizados

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/api/v1/todos` | Obtener todas las tareas |
| POST | `/api/v1/todos` | Crear nueva tarea |
| GET | `/api/v1/todos/{id}` | Obtener tarea por ID |
| PUT | `/api/v1/todos/{id}` | Actualizar tarea |
| DELETE | `/api/v1/todos/{id}` | Eliminar tarea |

## 🎯 Ejemplos de Uso

### Crear una Nueva Tarea
1. Escribir el título en el campo "Título de la tarea"
2. Opcionalmente agregar una descripción
3. Hacer clic en "Agregar Tarea"
4. La tarea aparecerá en la lista

### Editar una Tarea
1. Hacer clic en "Editar" en la tarea deseada
2. Modificar título, descripción o estado
3. Hacer clic en "Guardar Cambios"

### Filtrar Tareas
- **Todas**: Ver todas las tareas
- **Pendientes**: Solo tareas no completadas
- **Completadas**: Solo tareas completadas

### Marcar como Completada
1. Hacer clic en "Completar" en cualquier tarea
2. La tarea se marcará visualmente como completada
3. Las estadísticas se actualizarán automáticamente

## 🐛 Solución de Problemas

### Error de Conexión
- Verificar que la API esté corriendo en el puerto 8080
- Comprobar que no haya problemas de CORS
- Revisar la consola del navegador para errores

### Problemas de CORS
Si tienes problemas de CORS, asegúrate de que la API tenga configurado el middleware CORS correctamente.

### Tareas No Se Cargan
1. Verificar la conexión a la API
2. Revisar la consola del navegador
3. Comprobar que el endpoint `/api/v1/todos` responda correctamente

## 🔮 Mejoras Futuras

- [ ] Búsqueda de tareas por texto
- [ ] Ordenamiento por fecha/prioridad
- [ ] Categorías/etiquetas
- [ ] Exportar/importar tareas
- [ ] Modo oscuro
- [ ] PWA (Progressive Web App)
- [ ] Sincronización offline

## 📄 Licencia

Este proyecto está bajo la Licencia MIT.

## 🤝 Contribución

1. Fork el proyecto
2. Crea una rama para tu feature
3. Commit tus cambios
4. Push a la rama
5. Abre un Pull Request

---

**Desarrollado con ❤️ usando Go, HTML5, CSS3 y JavaScript ES6+**
