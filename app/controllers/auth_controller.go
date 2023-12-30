package controllers

import (
	"fmt"
	"html"
	"i1-go-starter-kit/app/config"
	"i1-go-starter-kit/app/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	baseUrl string
	session *session.Store
}

func NewAuthController(baseUrl string, session *session.Store) *AuthController {
	return &AuthController{
		baseUrl: baseUrl,
		session: session,
	}
}

func (a *AuthController) HandleRegister(c *fiber.Ctx) error {
	var err error
	newUser := new(models.User)
	err = c.BodyParser(newUser)
	if err != nil {
		return err
	}

	username := newUser.Username
	password := newUser.Password

	var users []models.User
	err = config.DB.Find(&users).Error
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.Username == username {
			return c.Redirect("/login")
		}
	}
	//daftarkan pengguna baru dengan password yang di-hash
	hashedPassword := hashPassword(password)
	newUser.Password = string(hashedPassword)
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	err = config.DB.Create(newUser).Error
	if err != nil {
		return err
	}

	session, err := a.session.Get(c)
	if err != nil {
		return err
	}
	session.Set("authenticated", true)
	session.Set("username", username)
	if err := session.Save(); err != nil {
        return err
    }

	return c.Redirect("/dashboard")

}

func (a *AuthController) HandleLogin(c *fiber.Ctx) error {
	rawUsername := c.FormValue("username")
    cleanedUsername := html.EscapeString(rawUsername)

    rawPassword := c.FormValue("password")
    cleanedPassword := html.EscapeString(rawPassword)

    // Validasi jika username atau password kosong setelah membersihkan
    if cleanedUsername == "" || cleanedPassword == "" {
        message := "Username atau password tidak boleh kosong"
        redirectURL := fmt.Sprintf("/login?message=%s", message)
        return c.Redirect(redirectURL, fiber.StatusSeeOther)
    }

	var user models.User
    if err := config.DB.Where("username = ?", cleanedUsername).First(&user).Error; err != nil {
        // Jika username tidak ditemukan dalam database, kembalikan pesan umum
        message := "Username atau Password salah"
        redirectURL := fmt.Sprintf("/login?message=%s", message)
        return c.Redirect(redirectURL, fiber.StatusSeeOther)
    }

	// Bandingkan password yang dimasukkan dengan password di database
    if !comparePassword([]byte(user.Password), cleanedPassword) {
        message := "Username atau Password salah"
        redirectURL := fmt.Sprintf("/login?message=%s", message)
        return c.Redirect(redirectURL, fiber.StatusSeeOther)
    }

	 // Set session setelah berhasil login
	 session, err := a.session.Get(c)
	 if err != nil {
		 return err
	 }
	 session.Set("authenticated", true)
	 session.Set("username", cleanedUsername)
	 if err := session.Save(); err != nil {
		 return err
	 }
 
	 return c.Redirect("/dashboard")

}

func (a *AuthController) Logout(c *fiber.Ctx) error {
	session, err := a.session.Get(c)
	if err != nil {
		return err
	}
	session.Destroy()
	return c.Redirect("/login")
}

func hashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return hashedPassword
}

func comparePassword(hashedPassword []byte, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(plainPassword))
	return err == nil
}
