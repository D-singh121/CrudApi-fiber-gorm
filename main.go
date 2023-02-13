package main

import (
	"github.com/devesh/go-fiber-gorm-rest/user"

	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("this is fiber server")
}

// -----routing
func Routers(app *fiber.App) {
	app.Post("/user", user.SaveUser)
	app.Get("/user/:id", user.Getuser)
	app.Get("/users", user.Getusers)
	app.Delete("/user/:id", user.DeleteUser)
	app.Put("/user/:id", user.UpdateUser)
}

func main() {
	user.InitialMigration()
	app := fiber.New() //-----entry point
	Routers(app)       //------calling the routers
	app.Get("/", HelloWorld)
	app.Listen(":3000") //-----port setting

}
