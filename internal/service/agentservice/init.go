package agentservice

import "bloger/internal/domain"

type AgentService struct {
	agent domain.Agent
}

func NewAgentService(agent domain.Agent) *AgentService {
	return &AgentService{agent: agent}
}
