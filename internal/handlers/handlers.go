package handlers

import (
	"github.com/DrLivsey00/ShopPraactice/pkg/database"
	"github.com/gofiber/fiber/v2"
)

var Db *database.DB

func Handlersinit(db *database.DB) {
	Db = db
}

func HomeHandler(c *fiber.Ctx) error {

	return c.SendString("this is main page")
}
