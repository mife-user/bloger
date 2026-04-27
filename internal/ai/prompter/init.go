package prompter

import (
	"context"

	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
)

type ModifierBuilder struct{}

func (ModifierBuilder) Build(systemPrompt string) react.MessageModifier {
	return func(ctx context.Context, input []*schema.Message) []*schema.Message {
		return append([]*schema.Message{schema.SystemMessage(systemPrompt)}, input...)
	}
}
