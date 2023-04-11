package routes

import (
	"html/template"

	"github.com/gofiber/fiber/v2"

	swagger "github.com/gofiber/swagger"
)

func SwaggerRoute(a *fiber.App) {
	route := a.Group("/docs")

	route.Get("*", swagger.New(swagger.Config{
		Title:  "Api boilerplate",
		Layout: "BaseLayout",
		Presets: []template.JS{
			template.JS("SwaggerUIBundle.presets.apis"),
			template.JS("SwaggerUIStandalonePreset"),
		},
	}))
}
