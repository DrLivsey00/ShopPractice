package server

import (
	"github.com/DrLivsey00/ShopPraactice/internal/handlers"
	"github.com/DrLivsey00/ShopPraactice/pkg/database"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"log"
)

var Sessions *session.Store

func Run(db *database.DB) {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,X-CSRF-Token,Authorization"},
		AllowMethods: []string{"GET,POST,PATCH,DELETE"},
	}))

	app.Get("/", handlers.HomeHandler) //main page of shop

	//user routes

	user := app.Group("/user")
	user.Get("", handlers.HomeHandler)                   // get user info
	user.Patch("/change_name", handlers.ChangeUserInfo)  // change name or bio
	user.Patch("/change_bio", handlers.ChangeUserInfo)   // change name or bio
	user.Patch("/change_email", handlers.ChangeUserInfo) // change name or bio
	user.Post("/register", handlers.RegisterUser)        // change name or bio
	user.Post("/login", handlers.LoginUser)              //login to system
	user.Post("/logout", handlers.LogoutUser)            //logout

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

	Sessions = session.New()

	log.Fatal(app.Listen(":8080"))

}
