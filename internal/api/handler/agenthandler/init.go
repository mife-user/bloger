package agenthandler

import "mifer/internal/domain"

type AgentHandler struct {
	service domain.AgentService
}

func NewAgentHandler(service domain.AgentService) *AgentHandler {
	return &AgentHandler{service: service}
}
