package main

import (
	"i1-go-starter-kit/app/config"
	"i1-go-starter-kit/app/helpers"
	"i1-go-starter-kit/app/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
)

func main() {

	// Arguments CLI for create awesome feature
	helpers.HandleCommandArgs()

	// Load environment variables from .env file
	config.InitEnvVariables()

	// Config DB
	dbConfig := config.GetDatabaseConfig()
	dsnMysql := dbConfig.GetDSN("mysql")

	// Connection DB for scalable or change adapter DB
	// config.InitDatabaseOracle(dsn)
	// config.InitDatabasePostgreSQL(dsn)
	config.InitDatabaseMySQL(dsnMysql)

	// Get the value of the BASE_URL environment variable
	baseUrl := os.Getenv("BASE_URL")

	// Initialize standard Go html template engine
	// Logger
	// Session
	// Serve Static
	// Routes
	// Listen Server
	engine := html.New("./app/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())
	store := session.New()
	app.Static("/", "./static")
	routes.SetupRoutes(app, baseUrl, store)
	err := app.Listen("localhost:" + "8000")
	if err != nil {
		return
	}

}
