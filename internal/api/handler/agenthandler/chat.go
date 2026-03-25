package agenthandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Chat 聊天
func (h *AgentHandler) Chat(c *gin.Context) {
	// 处理聊天逻辑
	var req map[string]any
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 调用服务层处理聊天逻辑
	response, err := h.service.Chat(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": response})
}
