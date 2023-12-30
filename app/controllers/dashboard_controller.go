package controllers

import (
	"i1-go-starter-kit/app/config"
	"i1-go-starter-kit/app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type DashboardController struct {
	baseUrl string
	session *session.Store
}

func NewDashboardController(baseUrl string, session *session.Store) *DashboardController {
	return &DashboardController{
		baseUrl: baseUrl,
		session: session,
	}
}

func (d *DashboardController) Index(c *fiber.Ctx) error {
	session, err := d.session.Get(c)
	if err != nil {
		return err
	}
	username := session.Get("username")
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil
	}
	return c.Render("dashboard", fiber.Map{"user": user}, "layouts/admin")
}
