package server

import (
	"github.com/DrLivsey00/ShopPraactice/internal/handlers"
	"github.com/DrLivsey00/ShopPraactice/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func Run(db *database.DB) {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,X-CSRF-Token,Authorization",
		AllowMethods: "GET,POST,PATCH,DELETE",
	}))

	app.Get("/", handlers.HomeHandler) //main page of shop

	//user routes

	user := app.Group("/user")
	user.Get("", handlers.UserInfo)                // get user info
	user.Patch("/change", handlers.ChangeUserInfo) // change name or bio
	user.Post("/register", handlers.RegisterUser)  // change name or bio
	user.Post("/login", handlers.LoginUser)        //login to system
	user.Post("/logout", handlers.LogoutUser)      //logout

	//admin routes
	admin := app.Group("/admin")
	admin.Get("/:uuid", handlers.HomeHandler)                     //get admin info
	admin.Post("/login/:username:password", handlers.HomeHandler) //login admin
	admin.Post("/add/:file", handlers.HomeHandler)                //update devices table
	admin.Post("/logout", handlers.HomeHandler)                   //logout

	//dish routes
	device := app.Group("/device")
	device.Get("/:id", handlers.HomeHandler) //get info about device

	handlers.Handlersinit(db)

	log.Fatal(app.Listen(":8080"))

}
