// authMiddleware.go

package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func AuthMiddleware(session *session.Store) fiber.Handler {
    return func(c *fiber.Ctx) error {
        session, err := session.Get(c)
        if err != nil {
            return err
        }
        if session.Get("authenticated") != true {
            return c.Redirect("/login")
        }
        return c.Next()
    }
}
