package handlers

import (
	"github.com/gofiber/fiber/v3"
)

func UserInfo(c fiber.Ctx) error {
	return c.SendString("User info")
}

func ChangeUserInfo(c fiber.Ctx) error {
	return nil
}

func ChangeUserName(c fiber.Ctx) error {
	return nil
}
func ChangeUserEmail(c fiber.Ctx) error {
	return nil
}
func ChangeUserPassword(c fiber.Ctx) error {
	return nil
}

func LogoutUser(c fiber.Ctx) error {
	return nil
}
