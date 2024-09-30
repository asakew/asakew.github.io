package server

import (
	"appFiber/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func Run() {
	engine := html.New("./web", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./web/assets")

	router.HTMLRendering(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
