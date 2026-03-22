package githandler

import (
	"bloger/pkg/exc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *GitHandler) Save(c *gin.Context) {
	// 保存token
	tokenany := c.ShouldBindJSON("token")
	tokenstr, ok := exc.IsString(tokenany)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token must be a string",
		})
		return
	}
	if tokenstr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token is required",
		})
		return
	}
	// 保存token
	err := h.Service.Save(tokenstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "token saved",
	})

}
