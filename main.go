package main

import (
	"fmt"

	"leo_go_api/api"
	"leo_go_api/api/common"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(
	/*cors.Config{
		Next: nil,
		// AllowOrigins:     "http://127.0.0.1",
		AllowOrigins:     "",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}*/))
	api.InitRouter(app)
	app.Listen(fmt.Sprintf(":%s", common.Port))
}
