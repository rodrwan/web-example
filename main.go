package main

import (
	"os"

	"web-example/internal/ui"

	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()
	s.GET("/", func(c *gin.Context) {
		component := ui.Layout()
		component.Render(c.Request.Context(), c.Writer)
	})
	s.Run(":" + os.Getenv("PORT"))
}
