package main

import "fmt"

func main() {
	fmt.Println("Test yml")
	// app := fiber.New()
	// app.Use(cors.New(
	// /*cors.Config{
	// 	Next: nil,
	// 	// AllowOrigins:     "http://127.0.0.1",
	// 	AllowOrigins:     "",
	// 	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	// 	AllowHeaders:     "",
	// 	AllowCredentials: false,
	// 	ExposeHeaders:    "",
	// 	MaxAge:           0,
	// }*/))
	// app.Get("/test", func(c *fiber.Ctx) error {
	// 	// Get session from storage
	// 	sess, err := store.Get(c)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(sess)
	// 	// Get all Keys
	// 	keys := sess.Keys()
	// 	fmt.Println(keys)
	// 	// Get value
	// 	name := sess.Get("name")
	// 	fmt.Println(name)
	// 	// Set key/value
	// 	sess.Set("name", "john")
	// 	// Delete key
	// 	// sess.Delete("name")
	// 	// // Destroy session
	// 	// if err := sess.Destroy(); err != nil {
	// 	// 	panic(err)
	// 	// }
	// 	// Save session
	// 	if err := sess.Save(); err != nil {
	// 		panic(err)
	// 	}
	// 	return c.SendString(fmt.Sprintf("Welcome %v", name))
	// })
	// api.InitRouter(app)
	// app.Listen(fmt.Sprintf(":%s", common.Port))
}
