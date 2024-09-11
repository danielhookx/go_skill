package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `form:"name"`
}

func main() {
	g := gin.Default()
	g.GET("/test", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindQuery(&person); err != nil {
			c.String(http.StatusBadRequest, "err: %w", err)
		}
		c.String(http.StatusOK, "success: %s", person.Name)
	})
	http.ListenAndServe(":1831", g)
}
