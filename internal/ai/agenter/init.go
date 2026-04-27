package agenter

import (
	"context"

	"mifer/pkg/logger"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
)

func InitAgent(ctx context.Context, llm model.ToolCallingChatModel,
	tools []tool.BaseTool, modifier react.MessageModifier) (*react.Agent, error) {
	logger.Info("初始化Agent...")
	return react.NewAgent(ctx, &react.AgentConfig{
		ToolCallingModel: llm,
		ToolsConfig:      compose.ToolsNodeConfig{Tools: tools},
		MessageModifier:  modifier,
		MaxStep:          12,
	})
}
