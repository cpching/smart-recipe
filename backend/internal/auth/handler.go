package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)

type RegisterInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type Handler struct {
	service   AuthService
	validator *Validation
}

func NewHandler(service AuthService) *Handler {
	return &Handler{
		service:   service,
		validator: NewValidation(),
	}
}

func (h *Handler) Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.validator.Validate(input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// if err := h.validator.Struct(input); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	// 	return
	// }

	user, err := h.service.Register(c.Request.Context(), input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    user.ID,
		"email": user.Email,
	})
}
