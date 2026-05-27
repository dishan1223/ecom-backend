package AuthHandler

import (
	"ecom-backend/consts"
	"ecom-backend/internal/initializer"
	"ecom-backend/internal/models"
	"ecom-backend/internal/service"
	"errors"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Signup handler handles the sign up route
func Signup(c fiber.Ctx) error {
	body := new(models.SignUpRequest)

	// new() already returns a pointer
	if err := c.Bind().Body(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": consts.InvalidRequest,
		})
	}

	exists := new(models.User)
	err := initializer.DB.Where("email = ?", body.Email).First(exists).Error
	if err == nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"error": consts.UserExists,
		})
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Database error during email availability check", "error", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Bcrypt is not working properly")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": consts.AccountCreationFailed,
		})
	}

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hash),
	}

	if err := initializer.DB.Create(&user).Error; err != nil {
		log.Error("Failed to add user to database")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": consts.AccountCreationFailed,
		})
	}

	log.Infof("User created successfully. ID assigned by GORM: %v", user.ID)

	token, err := service.GenerateToken(user.ID)
	if err != nil {
		log.Error("Failed to generate token from service pkg.")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": consts.AccountCreationFailed,
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		// secure should be tuned on for production
		Secure:   false,
		SameSite: "lax",
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 365),
	})

	return c.SendStatus(http.StatusOK)
}
