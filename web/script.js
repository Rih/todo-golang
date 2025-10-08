// Configuración de la API
const API_BASE_URL = 'http://localhost:8080/api/v1';

// Estado global de la aplicación
let todos = [];
let currentFilter = 'all';
let editingTodoId = null;

// Elementos del DOM
const todoForm = document.getElementById('todoForm');
const todoTitle = document.getElementById('todoTitle');
const todoDescription = document.getElementById('todoDescription');
const submitBtn = document.getElementById('submitBtn');
const cancelBtn = document.getElementById('cancelBtn');
const todoList = document.getElementById('todoList');
const emptyState = document.getElementById('emptyState');
const loading = document.getElementById('loading');
const errorMessage = document.getElementById('errorMessage');
const errorText = document.getElementById('errorText');
const filterBtns = document.querySelectorAll('.filter-btn');
const totalTodos = document.getElementById('totalTodos');
const pendingTodos = document.getElementById('pendingTodos');
const completedTodos = document.getElementById('completedTodos');
const editModal = document.getElementById('editModal');
const editForm = document.getElementById('editForm');
const editTitle = document.getElementById('editTitle');
const editDescription = document.getElementById('editDescription');
const editCompleted = document.getElementById('editCompleted');
const closeModal = document.getElementById('closeModal');
const cancelEdit = document.getElementById('cancelEdit');
const saveEdit = document.getElementById('saveEdit');

// Inicialización
document.addEventListener('DOMContentLoaded', function() {
    loadTodos();
    setupEventListeners();
});

// Configurar event listeners
function setupEventListeners() {
    // Formulario principal
    todoForm.addEventListener('submit', handleSubmit);
    cancelBtn.addEventListener('click', cancelEditMode);
    
    // Filtros
    filterBtns.forEach(btn => {
        btn.addEventListener('click', () => handleFilter(btn.dataset.filter));
    });
    
    // Modal
    closeModal.addEventListener('click', closeEditModal);
    cancelEdit.addEventListener('click', closeEditModal);
    saveEdit.addEventListener('click', handleSaveEdit);
    
    // Cerrar modal al hacer clic fuera
    editModal.addEventListener('click', (e) => {
        if (e.target === editModal) {
            closeEditModal();
        }
    });
}

// Cargar todos desde la API
async function loadTodos() {
    showLoading(true);
    hideError();
    
    try {
        const response = await fetch(`${API_BASE_URL}/todos`);
        const data = await response.json();
        
        if (data.success) {
            todos = data.data || [];
            renderTodos();
            updateStats();
        } else {
            showError('Error al cargar las tareas: ' + data.message);
        }
    } catch (error) {
        showError('Error de conexión: ' + error.message);
    } finally {
        showLoading(false);
    }
}

// Manejar envío del formulario
async function handleSubmit(e) {
    e.preventDefault();
    
    const title = todoTitle.value.trim();
    const description = todoDescription.value.trim();
    
    if (!title) {
        showError('El título es requerido');
        return;
    }
    
    if (editingTodoId) {
        await updateTodo(editingTodoId, title, description);
    } else {
        await createTodo(title, description);
    }
}

// Crear nuevo todo
async function createTodo(title, description) {
    showLoading(true);
    hideError();
    
    try {
        const response = await fetch(`${API_BASE_URL}/todos`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                title: title,
                description: description,
                completed: false
            })
        });
        
        const data = await response.json();
        
        if (data.success) {
            todos.push(data.data);
            renderTodos();
            updateStats();
            resetForm();
            showSuccess('Tarea creada exitosamente');
        } else {
            showError('Error al crear la tarea: ' + data.message);
        }
    } catch (error) {
        showError('Error de conexión: ' + error.message);
    } finally {
        showLoading(false);
    }
}

// Actualizar todo
async function updateTodo(id, title, description) {
    showLoading(true);
    hideError();
    
    try {
        const response = await fetch(`${API_BASE_URL}/todos/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                title: title,
                description: description,
                completed: false
            })
        });
        
        const data = await response.json();
        
        if (data.success) {
            const index = todos.findIndex(todo => todo.id === id);
            if (index !== -1) {
                todos[index] = data.data;
            }
            renderTodos();
            updateStats();
            resetForm();
            showSuccess('Tarea actualizada exitosamente');
        } else {
            showError('Error al actualizar la tarea: ' + data.message);
        }
    } catch (error) {
        showError('Error de conexión: ' + error.message);
    } finally {
        showLoading(false);
    }
}

// Eliminar todo
async function deleteTodo(id) {
    if (!confirm('¿Estás seguro de que quieres eliminar esta tarea?')) {
        return;
    }
    
    showLoading(true);
    hideError();
    
    try {
        const response = await fetch(`${API_BASE_URL}/todos/${id}`, {
            method: 'DELETE'
        });
        
        const data = await response.json();
        
        if (data.success) {
            todos = todos.filter(todo => todo.id !== id);
            renderTodos();
            updateStats();
            showSuccess('Tarea eliminada exitosamente');
        } else {
            showError('Error al eliminar la tarea: ' + data.message);
        }
    } catch (error) {
        showError('Error de conexión: ' + error.message);
    } finally {
        showLoading(false);
    }
}

// Toggle completado
async function toggleTodo(id) {
    const todo = todos.find(t => t.id === id);
    if (!todo) return;
    
    showLoading(true);
    hideError();
    
    try {
        const response = await fetch(`${API_BASE_URL}/todos/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                title: todo.title,
                description: todo.description,
                completed: !todo.completed
            })
        });
        
        const data = await response.json();
        
        if (data.success) {
            const index = todos.findIndex(t => t.id === id);
            if (index !== -1) {
                todos[index] = data.data;
            }
            renderTodos();
            updateStats();
        } else {
            showError('Error al actualizar la tarea: ' + data.message);
        }
    } catch (error) {
        showError('Error de conexión: ' + error.message);
    } finally {
        showLoading(false);
    }
}

// Editar todo
function editTodo(id) {
    const todo = todos.find(t => t.id === id);
    if (!todo) return;
    
    editingTodoId = id;
    editTitle.value = todo.title;
    editDescription.value = todo.description;
    editCompleted.checked = todo.completed;
    
    editModal.style.display = 'block';
    document.body.style.overflow = 'hidden';
}

// Guardar edición
async function handleSaveEdit() {
    const title = editTitle.value.trim();
    const description = editDescription.value.trim();
    const completed = editCompleted.checked;
    
    if (!title) {
        showError('El título es requerido');
        return;
    }
    
    showLoading(true);
    hideError();
    
    try {
        const response = await fetch(`${API_BASE_URL}/todos/${editingTodoId}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                title: title,
                description: description,
                completed: completed
            })
        });
        
        const data = await response.json();
        
        if (data.success) {
            const index = todos.findIndex(todo => todo.id === editingTodoId);
            if (index !== -1) {
                todos[index] = data.data;
            }
            renderTodos();
            updateStats();
            closeEditModal();
            showSuccess('Tarea actualizada exitosamente');
        } else {
            showError('Error al actualizar la tarea: ' + data.message);
        }
    } catch (error) {
        showError('Error de conexión: ' + error.message);
    } finally {
        showLoading(false);
    }
}

// Manejar filtros
function handleFilter(filter) {
    currentFilter = filter;
    
    // Actualizar botones de filtro
    filterBtns.forEach(btn => {
        btn.classList.remove('active');
        if (btn.dataset.filter === filter) {
            btn.classList.add('active');
        }
    });
    
    renderTodos();
}

// Renderizar todos
function renderTodos() {
    const filteredTodos = getFilteredTodos();
    
    if (filteredTodos.length === 0) {
        todoList.style.display = 'none';
        emptyState.style.display = 'block';
        return;
    }
    
    todoList.style.display = 'block';
    emptyState.style.display = 'none';
    
    todoList.innerHTML = filteredTodos.map(todo => `
        <div class="todo-item ${todo.completed ? 'completed' : ''}" data-id="${todo.id}">
            <div class="todo-header">
                <div>
                    <div class="todo-title">${escapeHtml(todo.title)}</div>
                    ${todo.description ? `<div class="todo-description">${escapeHtml(todo.description)}</div>` : ''}
                </div>
            </div>
            <div class="todo-meta">
                <span><i class="fas fa-calendar"></i> ${formatDate(todo.created_at)}</span>
                <span><i class="fas fa-clock"></i> ${formatDate(todo.updated_at)}</span>
            </div>
            <div class="todo-actions">
                <button class="btn ${todo.completed ? 'btn-secondary' : 'btn-success'}" 
                        onclick="toggleTodo(${todo.id})">
                    <i class="fas ${todo.completed ? 'fa-undo' : 'fa-check'}"></i>
                    ${todo.completed ? 'Desmarcar' : 'Completar'}
                </button>
                <button class="btn btn-primary" onclick="editTodo(${todo.id})">
                    <i class="fas fa-edit"></i> Editar
                </button>
                <button class="btn btn-danger" onclick="deleteTodo(${todo.id})">
                    <i class="fas fa-trash"></i> Eliminar
                </button>
            </div>
        </div>
    `).join('');
}

// Obtener todos filtrados
function getFilteredTodos() {
    switch (currentFilter) {
        case 'completed':
            return todos.filter(todo => todo.completed);
        case 'pending':
            return todos.filter(todo => !todo.completed);
        default:
            return todos;
    }
}

// Actualizar estadísticas
function updateStats() {
    const total = todos.length;
    const completed = todos.filter(todo => todo.completed).length;
    const pending = total - completed;
    
    totalTodos.textContent = total;
    completedTodos.textContent = completed;
    pendingTodos.textContent = pending;
}

// Resetear formulario
function resetForm() {
    todoForm.reset();
    editingTodoId = null;
    submitBtn.innerHTML = '<i class="fas fa-plus"></i> Agregar Tarea';
    cancelBtn.style.display = 'none';
}

// Cancelar modo edición
function cancelEditMode() {
    resetForm();
}

// Cerrar modal de edición
function closeEditModal() {
    editModal.style.display = 'none';
    document.body.style.overflow = 'auto';
    editingTodoId = null;
}

// Mostrar/ocultar loading
function showLoading(show) {
    loading.style.display = show ? 'block' : 'none';
}

// Mostrar error
function showError(message) {
    errorText.textContent = message;
    errorMessage.style.display = 'block';
    setTimeout(() => {
        hideError();
    }, 5000);
}

// Ocultar error
function hideError() {
    errorMessage.style.display = 'none';
}

// Mostrar éxito
function showSuccess(message) {
    // Crear notificación temporal
    const notification = document.createElement('div');
    notification.style.cssText = `
        position: fixed;
        top: 20px;
        right: 20px;
        background: #28a745;
        color: white;
        padding: 15px 20px;
        border-radius: 8px;
        box-shadow: 0 5px 15px rgba(0,0,0,0.2);
        z-index: 1001;
        animation: slideIn 0.3s ease;
    `;
    notification.innerHTML = `<i class="fas fa-check"></i> ${message}`;
    
    document.body.appendChild(notification);
    
    setTimeout(() => {
        notification.remove();
    }, 3000);
}

// Utilidades
function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}

function formatDate(dateString) {
    const date = new Date(dateString);
    return date.toLocaleDateString('es-ES', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
    });
}

// Agregar estilos para animaciones
const style = document.createElement('style');
style.textContent = `
    @keyframes slideIn {
        from {
            transform: translateX(100%);
            opacity: 0;
        }
        to {
            transform: translateX(0);
            opacity: 1;
        }
    }
`;
document.head.appendChild(style);
