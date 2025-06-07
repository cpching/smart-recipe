package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MiddlewareValidateProduct validate the product in the request and calls next if ok
func (h *Handler) MiddlewareValidateUser(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.validator.Validate(input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Set("registerInput", input)
	c.Next()
}
