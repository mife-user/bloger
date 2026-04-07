package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/memory"
)

func main() {
	// 创建 LLM
	llm, err := openai.New(
		openai.WithToken("test-key"),
		openai.WithModel("deepseek-chat"),
		openai.WithBaseURL("https://api.deepseek.com/v1"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 测试 1: Token 限制太小 (10 tokens)
	fmt.Println("=== 测试 1: Token 限制 = 10 ===")
	smallMemory := memory.NewConversationTokenBuffer(llm, 10)
	ctx := context.Background()

	// 添加第一轮对话
	smallMemory.ChatHistory.AddUserMessage(ctx, "你好，我叫张三")
	smallMemory.ChatHistory.AddAIMessage(ctx, "你好张三！很高兴认识你。")

	// 查看保存的消息
	messages, _ := smallMemory.ChatHistory.Messages(ctx)
	fmt.Printf("保存的消息数量: %d\n", len(messages))
	for i, msg := range messages {
		fmt.Printf("消息 %d: %s\n", i+1, msg.GetContent())
	}

	// 测试 2: 合理的 Token 限制 (2048 tokens)
	fmt.Println("\n=== 测试 2: Token 限制 = 2048 ===")
	largeMemory := memory.NewConversationTokenBuffer(llm, 2048)

	// 添加多轮对话
	largeMemory.ChatHistory.AddUserMessage(ctx, "你好，我叫张三")
	largeMemory.ChatHistory.AddAIMessage(ctx, "你好张三！很高兴认识你。")
	largeMemory.ChatHistory.AddUserMessage(ctx, "我喜欢编程")
	largeMemory.ChatHistory.AddAIMessage(ctx, "编程是一项很有趣的技能！你喜欢什么编程语言？")
	largeMemory.ChatHistory.AddUserMessage(ctx, "我喜欢 Go 语言")
	largeMemory.ChatHistory.AddAIMessage(ctx, "Go 是一门很棒的语言！简洁高效。")

	// 查看保存的消息
	messages2, _ := largeMemory.ChatHistory.Messages(ctx)
	fmt.Printf("保存的消息数量: %d\n", len(messages2))
	for i, msg := range messages2 {
		fmt.Printf("消息 %d: %s\n", i+1, msg.GetContent())
	}

	// 测试 3: 计算 token 数量
	fmt.Println("\n=== Token 计算 ===")
	// 注意：实际的 token 计算需要调用 LLM 的 API
	// 这里只是演示概念
	fmt.Println("10 tokens ≈ 1-2 个英文单词")
	fmt.Println("2048 tokens ≈ 1500-2000 个中文字符")
}
