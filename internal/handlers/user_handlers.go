package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/DrLivsey00/ShopPraactice/internal/models"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"time"
)

func RegisterUser(c fiber.Ctx) error {
	user := new(models.User)
	createdAt := time.Now()

	err := c.Bind().Body(user)
	fmt.Println(user.Username)

	password := sha256.Sum256([]byte(user.Password))
	err = Db.RegisterUser(user.Username, hex.EncodeToString(password[:]), user.Bio, user.Email, createdAt)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	sessionToken := uuid.New()

	cookie := new(fiber.Cookie)
	cookie.Name = "session_token"
	cookie.Value = sessionToken.String()
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Secure = true
	cookie.HTTPOnly = true
	c.Cookie(cookie)

	return c.SendString("Successfully registered user")
}

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

func LoginUser(c fiber.Ctx) error {

	return nil
}
func LogoutUser(c fiber.Ctx) error {
	return nil
}
