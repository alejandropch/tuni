package interfaces

import (
	"github.com/gofiber/fiber/v2"
)

type AuthServiceInterface interface {
	CreateUser(c *fiber.Ctx, body byte)
	LogoutUser(c *fiber.Ctx, userID int)
	DeleteUser(c *fiber.Ctx, userID int)
	GetUserByID(c *fiber.Ctx, userID int)
	LoginUser(c *fiber.Ctx, body byte)
	GenerateTokens(c *fiber.Ctx, userID int, oldRefreshToken string)
}
