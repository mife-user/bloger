package agenthandler

import (
	req "mifer/internal/api/dtos/request"
	resp "mifer/internal/api/dtos/response"
	"mifer/internal/domain"
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
	domainReq := domain.ChatRequest{
		Content: reqdto.Content,
	}
	response, err := h.service.Chat(c.Request.Context(), domainReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	respdto := resp.ChatResponse{
		Content: response.Content,
	}
	c.JSON(http.StatusOK, respdto)
}
