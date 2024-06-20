package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/DrLivsey00/ShopPraactice/internal/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

func RegisterUser(c *fiber.Ctx) error {
	user := new(models.User)
	createdAt := time.Now()

	err := c.BodyParser(&user)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	password := sha256.Sum256([]byte(user.Password))

	err = Db.RegisterUser(user.Username, hex.EncodeToString(password[:]), user.Bio, user.Email, createdAt)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	c.SendString("User registered successfully")
	return nil
}

func UserInfo(c *fiber.Ctx) error {
	return nil
}

func ChangeUserInfo(c *fiber.Ctx) error {
	return nil
}

func LoginUser(c *fiber.Ctx) error {
	return nil
}
func LogoutUser(c *fiber.Ctx) error {
	return nil
}
