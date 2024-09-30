// myapp/internal/router/router.go
package router

import (
	"github.com/gofiber/fiber/v2"
)

// https://docs.gofiber.io/guide/templates
func HTMLRendering(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"HeadTitle":   "asaKew.GitHub.io",
			"Title":       "Go Fiber Template Example",
			"Description": "An example template",
			"Greeting":    "Hello, world!",
		})
	})
}
