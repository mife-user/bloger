package domain

import "context"

type GitService interface {
	Save(token string) error
}

type GitRepo interface {
	Save(token string) error
}

type AgentService interface {
	Chat(ctx context.Context, input map[string]any) (map[string]any, error)
}

type Agent interface {
	Chat(ctx context.Context, input map[string]any) (map[string]any, error)
}
