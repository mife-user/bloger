# LangChainGo 完全指南

> LangChainGo 是 LangChain 的 Go 语言实现，为构建大语言模型（LLM）应用提供强大的框架支持

---

## 📖 目录

1. [简介](#简介)
2. [核心架构](#核心架构)
3. [主要组件](#主要组件)
4. [项目中的使用](#项目中的使用)
5. [实战示例](#实战示例)
6. [最佳实践](#最佳实践)

---

## 简介

### 什么是 LangChainGo？

LangChainGo 是一个开源框架，专门用于开发基于大语言模型的应用程序。它是 Python 版 LangChain 的 Go 语言实现，提供了：

- **统一接口**：抽象不同 LLM 提供商的 API 差异
- **组件化设计**：模块化的组件可以灵活组合
- **链式调用**：支持复杂的 AI 工作流
- **工具集成**：轻松集成外部工具和 API

### 为什么选择 LangChainGo？

| 特性 | Python LangChain | LangChainGo |
|------|-----------------|-------------|
| 性能 | 较慢 | 高性能 |
| 并发 | GIL 限制 | 原生并发 |
| 部署 | 需要解释器 | 单一二进制 |
| 内存占用 | 较大 | 较小 |
| 类型安全 | 动态类型 | 静态类型 |

---

## 核心架构

### 架构图

```
┌─────────────────────────────────────────────────────────┐
│                      Application                         │
└─────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────┐
│                       Agents                             │
│  (智能代理：决策、工具调用、多轮对话)                      │
└─────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────┐
│                       Chains                             │
│  (链式调用：串联多个组件形成工作流)                        │
└─────────────────────────────────────────────────────────┘
                            │
                            ▼
┌──────────┬──────────┬──────────┬──────────┬────────────┐
│   LLMs   │  Memory  │  Tools   │ Prompts  │  Callbacks │
│ (大模型)  │  (记忆)  │  (工具)  │  (提示)  │  (回调)    │
└──────────┴──────────┴──────────┴──────────┴────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────┐
│              Model Providers (模型提供商)                │
│  OpenAI | DeepSeek | Anthropic | Ollama | ...          │
└─────────────────────────────────────────────────────────┘
```

### 核心概念

#### 1. **Models (模型)**
与大语言模型交互的统一接口

#### 2. **Prompts (提示)**
管理和优化提示词模板

#### 3. **Memory (记忆)**
存储和管理对话历史

#### 4. **Chains (链)**
将多个组件串联成工作流

#### 5. **Agents (代理)**
能够自主决策和使用工具的智能体

#### 6. **Tools (工具)**
执行特定功能的外部能力

#### 7. **Callbacks (回调)**
监控和记录执行过程

---

## 主要组件

### 1. LLMs (大语言模型)

#### 1.1 模型接口

LangChainGo 定义了统一的模型接口：

```go
// llms.Model 接口
type Model interface {
    // 生成文本
    Generate(ctx context.Context, prompts []string, options ...CallOption) ([]*Generation, error)
    
    // 流式生成
    Stream(ctx context.Context, prompt string, options ...CallOption) (*Stream, error)
}
```

#### 1.2 支持的模型提供商

**OpenAI**
```go
import "github.com/tmc/langchaingo/llms/openai"

llm, err := openai.New(
    openai.WithToken("your-api-key"),
    openai.WithModel("gpt-4"),
)
```

**DeepSeek (OpenAI 兼容)**
```go
llm, err := openai.New(
    openai.WithToken("your-deepseek-api-key"),
    openai.WithModel("deepseek-chat"),
    openai.WithBaseURL("https://api.deepseek.com/v1"),
)
```

**Ollama (本地模型)**
```go
import "github.com/tmc/langchaingo/llms/ollama"

llm, err := ollama.New(
    ollama.WithModel("llama2"),
)
```

#### 1.3 调用选项

```go
response, err := llm.Generate(ctx, []string{"你好"},
    llms.WithTemperature(0.7),      // 温度参数
    llms.WithMaxTokens(1000),       // 最大 Token 数
    llms.WithStopWords([]string{"\n"}), // 停止词
    llms.WithTopP(0.9),             // Top-P 采样
    llms.WithFrequencyPenalty(0.5), // 频率惩罚
    llms.WithPresencePenalty(0.5),  // 存在惩罚
)
```

### 2. Prompts (提示模板)

#### 2.1 提示模板

```go
import "github.com/tmc/langchaingo/prompts"

// 创建提示模板
template := prompts.NewPromptTemplate(
    "请写一篇关于 {{.topic}} 的博客文章，字数约 {{.words}} 字。",
    []string{"topic", "words"}, // 必需变量
)

// 格式化提示
prompt, err := template.Format(map[string]any{
    "topic": "Go 并发编程",
    "words": "1000",
})
```

#### 2.2 聊天提示模板

```go
import "github.com/tmc/langchaingo/prompts"

chatTemplate := prompts.NewChatPromptTemplate([]prompts.MessageFormatter{
    prompts.NewSystemMessagePromptTemplate(
        "你是一个专业的 {{.role}}。",
        []string{"role"},
    ),
    prompts.NewHumanMessagePromptTemplate(
        "{{.question}}",
        []string{"question"},
    ),
})

messages, err := chatTemplate.FormatMessages(map[string]any{
    "role":     "Hugo 博客助手",
    "question": "如何创建新博客？",
})
```

### 3. Memory (记忆)

#### 3.1 对话缓冲记忆

```go
import "github.com/tmc/langchaingo/memory"

// 创建对话记忆
memory := memory.NewConversationBuffer()

// 保存对话
memory.SaveContext(ctx, map[string]any{"input": "用户问题"}, map[string]any{"output": "AI回答"})

// 加载历史
history, _ := memory.LoadMemoryVariables(ctx, map[string]any{})
```

#### 3.2 对话摘要记忆

```go
// 自动压缩长对话
memory := memory.NewConversationSummary(llm)
```

#### 3.3 向量存储记忆

```go
// 基于向量检索的记忆
memory := memory.NewVectorStore(llm, vectorStore)
```

### 4. Chains (链)

#### 4.1 LLM 链

```go
import "github.com/tmc/langchaingo/chains"

// 创建 LLM 链
llmChain := chains.NewLLMChain(llm, promptTemplate)

// 执行链
result, err := chains.Call(ctx, llmChain, map[string]any{
    "topic": "Go 并发编程",
    "words": "1000",
})
```

#### 4.2 顺序链

```go
// 串联多个链
sequentialChain := chains.NewSequentialChain(
    []chains.Chain{chain1, chain2, chain3},
    []string{"input"},   // 输入变量
    []string{"output"},  // 输出变量
)

result, err := chains.Call(ctx, sequentialChain, map[string]any{
    "input": "初始输入",
})
```

#### 4.3 路由链

```go
// 根据条件选择不同的链
routerChain := chains.NewRouterChain(
    llm,
    []chains.Chain{mathChain, historyChain, codeChain},
    []string{"math", "history", "code"},
)
```

### 5. Agents (代理)

#### 5.1 会话代理

```go
import "github.com/tmc/langchaingo/agents"

// 创建会话代理
agent := agents.NewConversationalAgent(llm, tools, memory)

// 执行代理
result, err := agent.Call(ctx, map[string]any{
    "input": "帮我创建一个 Hugo 博客",
})
```

#### 5.2 工具调用代理

```go
// 创建工具调用代理
agent := agents.NewOpenAIFunctionsAgent(llm, tools)

// 执行代理
result, err := agents.Execute(ctx, agent, map[string]any{
    "input": "查询今天的天气",
})
```

### 6. Tools (工具)

#### 6.1 定义工具

```go
import "github.com/tmc/langchaingo/tools"

// 创建自定义工具
weatherTool := tools.Tool{
    Name:        "weather",
    Description: "查询指定城市的天气",
    ArgsSchema: map[string]any{
        "type": "object",
        "properties": map[string]any{
            "city": map[string]any{
                "type":        "string",
                "description": "城市名称",
            },
        },
        "required": []string{"city"},
    },
    Func: func(ctx context.Context, args string) (string, error) {
        // 实现工具逻辑
        return "北京今天晴，温度 25°C", nil
    },
}
```

#### 6.2 内置工具

```go
// 搜索工具
searchTool := tools.NewSearchTool(searchAPI)

// 计算器工具
calculatorTool := tools.NewCalculatorTool()

// Python REPL 工具
pythonTool := tools.NewPythonREPLTool()
```

### 7. Callbacks (回调)

#### 7.1 回调处理器

```go
import "github.com/tmc/langchaingo/callbacks"

// 创建回调处理器
handler := &callbacks.Handler{
    HandleLLMStart: func(ctx context.Context, prompts []string) {
        fmt.Println("LLM 开始执行")
    },
    HandleLLMEnd: func(ctx context.Context, output string) {
        fmt.Println("LLM 执行完成:", output)
    },
    HandleLLMError: func(ctx context.Context, err error) {
        fmt.Println("LLM 执行错误:", err)
    },
}

// 使用回调
result, err := llm.Generate(ctx, prompts, llms.WithCallback(handler))
```

---

## 项目中的使用

### 当前项目结构

```
internal/ai/
├── llm/
│   └── init.go          # LLM 初始化
├── agent/
│   ├── init.go          # Agent 初始化
│   └── talk.go          # 对话功能
├── skill/
│   └── init.go          # 技能模块（待实现）
└── knowledge/
    ├── init.go          # 知识库（待实现）
    └── data/            # 知识库数据
```

### 1. LLM 初始化

**文件**：[internal/ai/llm/init.go](file:///d:/vscode/VsCodeWork/bloger/internal/ai/llm/init.go)

```go
package llm

import (
    "bloger/pkg/conf"
    "bloger/pkg/err"
    "bloger/pkg/logger"
    
    "github.com/tmc/langchaingo/llms"
    "github.com/tmc/langchaingo/llms/openai"
)

// LLM 大语言模型客户端
type LLM struct {
    Client llms.Model
}

// InitLLM 初始化LLM客户端
func InitLLM(config *conf.Config) (*LLM, error) {
    logger.Info("初始化LLM...")
    
    // 验证配置
    if config.Ai.ApiKey == "" {
        return nil, err.New("api_key is empty")
    }
    if config.Ai.BaseURL == "" {
        return nil, err.New("base_url is empty")
    }
    
    // 创建 OpenAI 兼容客户端（支持 DeepSeek）
    client, err := openai.New(
        openai.WithToken(config.Ai.ApiKey),
        openai.WithModel(config.Ai.Model),
        openai.WithBaseURL(config.Ai.BaseURL),
    )
    if err != nil {
        logger.Error("创建LLM客户端失败", logger.C(err))
        return nil, err
    }
    
    return &LLM{Client: client}, nil
}
```

### 2. Agent 初始化

**文件**：[internal/ai/agent/init.go](file:///d:/vscode/VsCodeWork/bloger/internal/ai/agent/init.go)

```go
package agent

import (
    "bloger/internal/ai/llm"
    "bloger/pkg/logger"
    
    "github.com/tmc/langchaingo/agents"
)

// Agent 智能代理
type Agent struct {
    agent *agents.ConversationalAgent
}

// InitAgent 初始化Agent
func InitAgent(llm *llm.LLM) *Agent {
    logger.Info("初始化Agent...")
    
    // 创建会话代理
    agent := agents.NewConversationalAgent(llm.Client, nil, nil)
    
    return &Agent{agent: agent}
}
```

### 3. 对话功能

**文件**：[internal/ai/agent/talk.go](file:///d:/vscode/VsCodeWork/bloger/internal/ai/agent/talk.go)

```go
package agent

import "context"

// Chat 聊天
func (a *Agent) Chat(ctx context.Context, input map[string]any) (map[string]any, error) {
    return a.agent.Chain.Call(ctx, input)
}
```

---

## 实战示例

### 示例 1：博客内容生成

```go
package service

import (
    "bloger/internal/ai/llm"
    "context"
    "fmt"
    
    "github.com/tmc/langchaingo/llms"
    "github.com/tmc/langchaingo/prompts"
    "github.com/tmc/langchaingo/chains"
)

type BlogService struct {
    llm *llm.LLM
}

// GenerateBlogPost 生成博客文章
func (s *BlogService) GenerateBlogPost(ctx context.Context, topic, style string) (string, error) {
    // 1. 创建提示模板
    template := prompts.NewPromptTemplate(
        `请写一篇关于 {{.topic}} 的技术博客文章。

要求：
1. 标题吸引人
2. 内容结构清晰，包含引言、正文、总结
3. 包含代码示例
4. 适合 Hugo 博客格式
5. 写作风格：{{.style}}

请直接输出 Markdown 格式的文章内容。`,
        []string{"topic", "style"},
    )
    
    // 2. 创建 LLM 链
    chain := chains.NewLLMChain(s.llm.Client, template)
    
    // 3. 执行链
    result, err := chains.Call(ctx, chain, map[string]any{
        "topic": topic,
        "style": style,
    }, llms.WithTemperature(0.8), llms.WithMaxTokens(3000))
    if err != nil {
        return "", err
    }
    
    // 4. 提取结果
    text := result["text"].(string)
    return text, nil
}

// 使用示例
func Example() {
    // 初始化
    config := conf.GetConfig()
    llmClient, _ := llm.InitLLM(config)
    service := &BlogService{llm: llmClient}
    
    // 生成博客
    content, err := service.GenerateBlogPost(
        context.Background(),
        "Go 并发编程",
        "通俗易懂",
    )
    if err != nil {
        panic(err)
    }
    
    fmt.Println(content)
}
```

### 示例 2：SEO 优化建议

```go
package service

import (
    "context"
    
    "github.com/tmc/langchaingo/llms"
)

// GenerateSEOSuggestions 生成 SEO 优化建议
func (s *BlogService) GenerateSEOSuggestions(ctx context.Context, content string) (string, error) {
    prompt := fmt.Sprintf(`请分析以下博客内容，并提供 SEO 优化建议：

%s

建议包括：
1. 标题优化建议
2. 关键词建议（5-10个）
3. Meta 描述建议
4. 内容结构优化建议
5. 内链和外链建议

请以结构化的方式输出建议。`, content)
    
    response, err := llms.GenerateFromSinglePrompt(
        ctx,
        s.llm.Client,
        prompt,
        llms.WithTemperature(0.7),
        llms.WithMaxTokens(1500),
    )
    
    return response, err
}
```

### 示例 3：多轮对话助手

```go
package service

import (
    "context"
    
    "github.com/tmc/langchaingo/memory"
    "github.com/tmc/langchaingo/agents"
)

// ChatAssistant 对话助手
type ChatAssistant struct {
    agent  *agents.ConversationalAgent
    memory *memory.ConversationBuffer
}

// NewChatAssistant 创建对话助手
func NewChatAssistant(llm *llm.LLM) *ChatAssistant {
    // 创建对话记忆
    mem := memory.NewConversationBuffer()
    
    // 创建会话代理
    agent := agents.NewConversationalAgent(llm.Client, nil, mem)
    
    return &ChatAssistant{
        agent:  agent,
        memory: mem,
    }
}

// Chat 对话
func (a *ChatAssistant) Chat(ctx context.Context, input string) (string, error) {
    result, err := a.agent.Call(ctx, map[string]any{
        "input": input,
    })
    if err != nil {
        return "", err
    }
    
    return result["output"].(string), nil
}

// 使用示例
func Example() {
    config := conf.GetConfig()
    llmClient, _ := llm.InitLLM(config)
    assistant := NewChatAssistant(llmClient)
    
    // 多轮对话
    ctx := context.Background()
    
    // 第一轮
    response1, _ := assistant.Chat(ctx, "你好，我想创建一个 Hugo 博客")
    fmt.Println(response1)
    
    // 第二轮（有上下文记忆）
    response2, _ := assistant.Chat(ctx, "应该选择什么主题？")
    fmt.Println(response2)
    
    // 第三轮
    response3, _ := assistant.Chat(ctx, "如何优化 SEO？")
    fmt.Println(response3)
}
```

### 示例 4：工具调用代理

```go
package service

import (
    "context"
    "fmt"
    
    "github.com/tmc/langchaingo/agents"
    "github.com/tmc/langchaingo/tools"
)

// HugoAssistant Hugo 助手
type HugoAssistant struct {
    agent *agents.Executor
}

// NewHugoAssistant 创建 Hugo 助手
func NewHugoAssistant(llm *llm.LLM) *HugoAssistant {
    // 定义工具
    hugoTools := []tools.Tool{
        {
            Name:        "create_post",
            Description: "创建新的博客文章",
            ArgsSchema: map[string]any{
                "type": "object",
                "properties": map[string]any{
                    "title": map[string]any{
                        "type":        "string",
                        "description": "文章标题",
                    },
                    "content": map[string]any{
                        "type":        "string",
                        "description": "文章内容",
                    },
                },
                "required": []string{"title", "content"},
            },
            Func: func(ctx context.Context, args string) (string, error) {
                // 实现创建文章逻辑
                return "文章创建成功", nil
            },
        },
        {
            Name:        "optimize_seo",
            Description: "优化文章 SEO",
            ArgsSchema: map[string]any{
                "type": "object",
                "properties": map[string]any{
                    "content": map[string]any{
                        "type":        "string",
                        "description": "文章内容",
                    },
                },
                "required": []string{"content"},
            },
            Func: func(ctx context.Context, args string) (string, error) {
                // 实现 SEO 优化逻辑
                return "SEO 优化建议：...", nil
            },
        },
    }
    
    // 创建代理
    agent := agents.NewOpenAIFunctionsAgent(llm.Client, hugoTools)
    executor := agents.NewExecutor(agent)
    
    return &HugoAssistant{agent: executor}
}

// Execute 执行任务
func (a *HugoAssistant) Execute(ctx context.Context, task string) (string, error) {
    result, err := agents.Execute(ctx, a.agent, map[string]any{
        "input": task,
    })
    if err != nil {
        return "", err
    }
    
    return result["output"].(string), nil
}

// 使用示例
func Example() {
    config := conf.GetConfig()
    llmClient, _ := llm.InitLLM(config)
    assistant := NewHugoAssistant(llmClient)
    
    // 执行任务（代理会自动选择合适的工具）
    result, _ := assistant.Execute(
        context.Background(),
        "帮我创建一篇关于 Go 并发的博客文章，并优化其 SEO",
    )
    
    fmt.Println(result)
}
```

---

## 最佳实践

### 1. 错误处理

```go
func SafeGenerate(ctx context.Context, llm llms.Model, prompt string) (string, error) {
    // 添加超时控制
    ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
    defer cancel()
    
    // 添加重试机制
    var lastErr error
    for i := 0; i < 3; i++ {
        response, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
        if err == nil {
            return response, nil
        }
        lastErr = err
        
        // 指数退避
        time.Sleep(time.Duration(math.Pow(2, float64(i))) * time.Second)
    }
    
    return "", fmt.Errorf("重试3次后失败: %w", lastErr)
}
```

### 2. 上下文管理

```go
// 使用 context 传递请求级别的信息
type RequestContext struct {
    UserID    string
    SessionID string
    TraceID   string
}

func ProcessWithTracing(ctx context.Context, reqCtx *RequestContext, llm llms.Model, prompt string) (string, error) {
    // 添加追踪信息到日志
    logger.Info("处理请求", 
        logger.S("user_id", reqCtx.UserID),
        logger.S("trace_id", reqCtx.TraceID),
    )
    
    // 执行 LLM 调用
    response, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
    if err != nil {
        logger.Error("LLM调用失败",
            logger.S("trace_id", reqCtx.TraceID),
            logger.C(err),
        )
        return "", err
    }
    
    return response, nil
}
```

### 3. 性能优化

```go
// 使用并发处理多个请求
func BatchGenerate(ctx context.Context, llm llms.Model, prompts []string) ([]string, error) {
    results := make([]string, len(prompts))
    errors := make([]error, len(prompts))
    
    var wg sync.WaitGroup
    for i, prompt := range prompts {
        wg.Add(1)
        go func(idx int, p string) {
            defer wg.Done()
            results[idx], errors[idx] = llms.GenerateFromSinglePrompt(ctx, llm, p)
        }(i, prompt)
    }
    wg.Wait()
    
    // 检查错误
    for _, err := range errors {
        if err != nil {
            return nil, err
        }
    }
    
    return results, nil
}
```

### 4. 成本控制

```go
// Token 计数和成本估算
type CostTracker struct {
    totalTokens int
    costPer1k   float64
}

func (t *CostTracker) TrackUsage(usage llms.TokenUsage) {
    t.totalTokens += usage.TotalTokens
    cost := float64(usage.TotalTokens) / 1000 * t.costPer1k
    logger.Info("Token 使用统计",
        logger.I("tokens", usage.TotalTokens),
        logger.F("cost", cost),
    )
}

// 使用示例
tracker := &CostTracker{costPer1k: 0.002} // DeepSeek 价格
response, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt,
    llms.WithCallback(&callbacks.Handler{
        HandleLLMEnd: func(ctx context.Context, output string, usage llms.TokenUsage) {
            tracker.TrackUsage(usage)
        },
    }),
)
```

### 5. 安全性

```go
// 输入验证和清理
func SanitizeInput(input string) string {
    // 移除潜在的危险字符
    input = strings.TrimSpace(input)
    
    // 限制输入长度
    if len(input) > 10000 {
        input = input[:10000]
    }
    
    return input
}

// 敏感信息过滤
func FilterSensitiveInfo(response string) string {
    // 过滤 API Key
    response = regexp.MustCompile(`sk-[a-zA-Z0-9]{48}`).ReplaceAllString(response, "[API_KEY]")
    
    // 过滤邮箱
    response = regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`).ReplaceAllString(response, "[EMAIL]")
    
    return response
}
```

---

## 📚 参考资料

- [LangChainGo GitHub](https://github.com/tmc/langchaingo)
- [LangChainGo 文档](https://pkg.go.dev/github.com/tmc/langchaingo)
- [LangChain 官方文档](https://python.langchain.com/)
- [DeepSeek API 文档](https://platform.deepseek.com/docs)
- [OpenAI API 文档](https://platform.openai.com/docs)

---

## 🚀 下一步

1. **实现 Skill 模块**：封装具体的业务技能
2. **集成知识库**：使用向量存储实现 RAG
3. **添加工具支持**：为 Agent 添加工具调用能力
4. **实现记忆管理**：持久化对话历史
5. **性能优化**：添加缓存和并发处理

---

**最后更新**：2026-03-22
