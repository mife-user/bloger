package agenthandler

import (
	req "bloger/internal/api/dtos/request"
	resp "bloger/internal/api/dtos/response"
	"bloger/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Chat 聊天
func (h *AgentHandler) Chat(c *gin.Context) {
	// 处理聊天逻辑
	var reqdto req.ChatRequest
	if err := c.ShouldBindJSON(&reqdto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req := domain.ChatRequest{
		Message: reqdto.Message,
	}
	// 调用服务层处理聊天逻辑
	response, err := h.service.Chat(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	respdto := resp.ChatResponse{
		Message: response.Message,
	}
	c.JSON(http.StatusOK, respdto)
}
