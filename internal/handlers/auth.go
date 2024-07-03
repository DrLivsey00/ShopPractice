package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/DrLivsey00/ShopPraactice/internal/models"
	"github.com/DrLivsey00/ShopPraactice/pkg"
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

	sessionToken := uuid.New().String()
	cookie := pkg.CreateCookie(sessionToken)
	jwtToken, err := pkg.CreateJWT(user.Username)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	err = Db.RegisterUser(user.Username, hex.EncodeToString(password[:]), user.Bio, user.Email, sessionToken, jwtToken, createdAt)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	c.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
	c.Cookie(cookie)

	return c.SendString("Successfully registered user")
}

func LoginUser(c fiber.Ctx) error {

	return nil
}
