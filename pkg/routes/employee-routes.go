package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sk25469/go-mongodb-server/pkg/controllers"
)

func RegisterRoutes() {
	app := fiber.New()

	app.Get("/employee", controllers.GetAllEmployee)
	app.Post("/employee", controllers.AddNewEmployee)
	app.Put("/employee/:id", controllers.UpdateEmployeeById)
	app.Delete("/employee/:id", controllers.DeleteEmployeeById)

	log.Fatal(app.Listen(":3000"))
}
