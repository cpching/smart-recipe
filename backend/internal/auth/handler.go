package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service   AuthService
	validator *Validation
}

func NewHandler(service AuthService, validator *Validation) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

// @Summary      Create user
// @Description  Create a new user in the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body  RegisterInput  true  "User data"
// @Success      201   {object}  domain.User
// @Router       /register [post]
func (h *Handler) Register(c *gin.Context) {
	value, exists := c.Get("registerInput")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Input not found in context"})
		return
	}
	input := value.(RegisterInput)
	log.Println(input.Email)
	log.Println(input.Password)

	user, err := h.service.Register(c.Request.Context(), input.Email, input.Password)
	if err != nil {
		if err.Error() == "Email already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			// c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    user.ID,
		"email": user.Email,
	})
}
