package domain

import "context"

type GitService interface {
	Save(ctx context.Context, token string) error
}

type GitRepo interface {
	Save(ctx context.Context, token string) error
}

type AgentService interface {
	Chat(ctx context.Context, input ChatRequest) (ChatResponse, error)
}

type Agent interface {
	Chat(ctx context.Context, input ChatRequest) (ChatResponse, error)
}
