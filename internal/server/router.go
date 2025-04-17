package server

import (
	"github.com/gofiber/fiber/v2"
)

func setupMux(a *fiber.App) {
	// common
	//a.Get("/", func(c *fiber.Ctx) error { return nil })
	//a.All("/oauth/:provider/:flag", storeOAuth)
	//a.Get("/p/:id", getPage)
	//// form
	//a.Post("/form", postForm)
	//// page
	//a.Get("/page/:id/:flag", renderPage)
	//// agent
	//a.Post("/agent", agentData)
	//// webhook
	//a.All("/webhook/:flag", doWebhook)
	//// platform
	//a.All("/chatbot/:platform", platformCallback)
}

// handler
