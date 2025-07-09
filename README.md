# Web Example

Un proyecto de ejemplo de aplicación web moderna construida con Go, Gin y Templ.

## 📋 Descripción

Este proyecto demuestra cómo crear una aplicación web moderna utilizando:

- **Go 1.24.0** - Lenguaje de programación principal
- **Gin** - Framework web ligero y rápido para Go
- **Templ** - Sistema de plantillas moderno para Go que permite escribir componentes HTML en Go
- **Tailwind CSS** - Framework CSS utility-first para estilos modernos

## 🏗️ Arquitectura

```
web-example/
├── main.go                 # Punto de entrada de la aplicación
├── go.mod                  # Dependencias del módulo Go
├── go.sum                  # Checksums de las dependencias
└── internal/
    └── ui/
        ├── ui.go           # Configuración de generación de Templ
        ├── index.templ     # Plantilla HTML principal
        └── index_templ.go  # Código generado por Templ
```

## 🚀 Características

- **Servidor web con Gin**: Framework web rápido y eficiente
- **Templ para componentes**: Sistema de plantillas moderno que permite escribir HTML directamente en Go
- **Tailwind CSS**: Estilos modernos y responsivos
- **Fuentes personalizadas**: Integración con Google Fonts (Inter y Nunito)
- **Paleta de colores personalizada**: Tema oscuro con colores "groupie"
- **Configuración de entorno**: Puerto configurable mediante variable de entorno

## 🛠️ Tecnologías Utilizadas

### Backend
- **Go 1.24.0**: Lenguaje de programación
- **Gin v1.10.1**: Framework web
- **Templ v0.3.906**: Sistema de plantillas

### Frontend
- **Tailwind CSS**: Framework CSS
- **Google Fonts**: Inter y Nunito
- **HTML5**: Marcado semántico

## 📦 Instalación

### Prerrequisitos
- Go 1.24.0 o superior

### Pasos de instalación

1. **Clonar el repositorio**
   ```bash
   git clone <url-del-repositorio>
   cd web-example
   ```

2. **Instalar dependencias**
   ```bash
   go mod download
   ```

3. **Generar código de Templ** (si modificas archivos .templ)
   ```bash
   go generate ./...
   ```

4. **Ejecutar la aplicación**
   ```bash
   go run main.go
   ```

## 🌐 Uso

### Variables de entorno

- `PORT`: Puerto en el que se ejecutará el servidor (opcional, por defecto usa el puerto por defecto de Gin)

### Ejecutar en desarrollo

```bash
# Ejecutar directamente
go run main.go

# Con puerto específico
PORT=8080 go run main.go
```

### Construir para producción

```bash
# Construir binario
go build -o web-example main.go

# Ejecutar binario
./web-example
```

## 🎨 Personalización

### Colores del tema

El proyecto incluye una paleta de colores personalizada definida en `internal/ui/index.templ`:

- `groupie-primary`: #3a1c32
- `groupie-secondary`: #18181b
- `groupie-accent`: #a83250
- `groupie-success`: #01b796
- `groupie-warning`: #facc15
- `groupie-danger`: #ef4444

### Fuentes

- **Inter**: Fuente principal para texto general
- **Nunito**: Fuente secundaria con múltiples pesos

## 🔧 Desarrollo

### Estructura del proyecto

- `main.go`: Configuración del servidor Gin y rutas
- `internal/ui/`: Componentes de interfaz de usuario
  - `index.templ`: Plantilla principal con HTML y configuración de Tailwind
  - `ui.go`: Configuración para generación de código Templ

### Generación de código

Templ genera automáticamente código Go a partir de archivos `.templ`. Para regenerar el código:

```bash
go generate ./...
```

### Agregar nuevas rutas

Para agregar nuevas rutas, modifica `main.go`:

```go
s.GET("/nueva-ruta", func(c *gin.Context) {
    // Tu lógica aquí
})
```

## 📝 Licencia

Este proyecto es un ejemplo educativo y puede ser utilizado libremente.

## 🤝 Contribuciones

Las contribuciones son bienvenidas. Por favor, asegúrate de:

1. Seguir las convenciones de código de Go
2. Probar tus cambios antes de enviar un pull request
3. Documentar nuevas funcionalidades

## 📚 Recursos Adicionales

- [Documentación de Gin](https://gin-gonic.com/docs/)
- [Documentación de Templ](https://templ.guide/)
- [Documentación de Tailwind CSS](https://tailwindcss.com/docs)
- [Go Programming Language](https://golang.org/)

---

Desarrollado con ❤️ usando Go, Gin y Templ. 