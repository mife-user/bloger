package prompter

import (
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/prompts"
)

// InitPrompter 初始化提示词模板
func InitPrompter(llm llms.Model, prompt string) prompts.PromptTemplate {
	return prompts.NewPromptTemplate(prompt, []string{"input"})
}
