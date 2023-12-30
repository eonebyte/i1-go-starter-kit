package routes

import (
	"i1-go-starter-kit/app/controllers"
	"i1-go-starter-kit/app/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SetupRoutes(app *fiber.App, baseUrl string, store *session.Store) {
	// Wire Controllers
	dashboardController := controllers.NewDashboardController(baseUrl, store)
	authController := controllers.NewAuthController(baseUrl, store)

	// Middleware auth
	authMiddleware := middlewares.AuthMiddleware(store)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("welcome", fiber.Map{}, "layouts/main")
	})
	app.Get("/login", func(c *fiber.Ctx) error {
		message := c.Query("message")
		return c.Render("login", fiber.Map{"message": message})
	})
	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("register", fiber.Map{})
	})

	// Routes with Controller
	app.Get("/dashboard", authMiddleware, dashboardController.Index)
	app.Post("/handle-register", authController.HandleRegister)
	app.Post("/handle-login", authController.HandleLogin)
	app.Get("/logout", authController.Logout)
}
