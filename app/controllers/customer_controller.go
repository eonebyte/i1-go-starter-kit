package controllers

	import "github.com/gofiber/fiber/v2/middleware/session"


type customer_controller struct {
	baseUrl string
	session *session.Store
}

func Newcustomer_controller(baseUrl string, session *session.Store) *customer_controller {
	return &customer_controller{
		baseUrl: baseUrl,
		session: session,
	}
}