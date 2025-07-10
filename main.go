package main

import (
	"embed"
	"io/fs"
	"log/slog"
	"net/http"
	"os"

	"web-example/internal/ui"

	"github.com/gin-gonic/gin"
)

//go:generate ./...

//go:embed static
var staticFS embed.FS

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	s := gin.Default()

	// Serve static files
	staticSubFS, err := fs.Sub(staticFS, "static")
	if err != nil {
		slog.Error("Error configurando archivos estáticos", "error", err)
		panic(err)
	}
	s.StaticFS("/static", http.FS(staticSubFS))
	s.GET("/", func(c *gin.Context) {
		component := ui.Layout()
		component.Render(c.Request.Context(), c.Writer)
	})
	s.Run(":" + os.Getenv("PORT"))
}
