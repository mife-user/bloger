# mifer

## 项目简介

mifer 是一个类 Claude Code 的 CLI Agent 工具，基于 Go + LangChainGo + 大模型 API 构建，通过命令行终端与 AI 进行多轮对话，自动执行代码编写、Git 操作、文件管理等软件工程任务。

核心理念
- **终端原生**：CLI 优先，类 Claude Code 的交互体验，在终端中完成所有操作
- **AI Agent 驱动**：内置对话代理（Conversational Agent），支持工具调用、记忆管理、任务分解
- **工具生态**：内置 Git 操作、文件系统操作等工具，Agent 可自动调用完成复杂任务
- **可扩展**：插件化的工具系统，方便扩展新的工具能力

## 关键功能

1. **多轮对话**：基于 LangChainGo 的对话代理，支持上下文记忆
2. **工具调用**：Agent 自动选择合适的工具执行任务（Git 操作、Hugo 博客生成等）
3. **记忆管理**：短期记忆（Token 缓冲）+ 长期记忆（对话压缩存储）
4. **配置持久化**：YAML/JSON 配置，支持多模型切换
5. **Git 集成**：内置完整 Git 操作工具集（init、add、commit、push、pull、clone、branch、checkout、remote、status）
6. **Docker 部署**：支持 Docker 多阶段构建，容器化运行

## 项目架构

```
- cmd/           : 启动模块
  - bootstrap/   : 启动入口、初始化程序
  - main/        : 业务入口
- config/        : 配置文件（dev.yml, prod.yml.example）
- internal/      : 核心业务实现
  - api/         : HTTP REST API 层
    - handler/   : 请求处理器
    - middleware/ : 中间件（CORS 等）
    - route/     : 路由注册
    - dtos/      : 数据传输对象
  - ai/          : AI 能力层
    - agenter/   : 对话代理（ConversationalAgent）
    - executor/  : 执行器（Chat 执行）
    - llmer/     : 大模型封装（DeepSeek via OpenAI 兼容接口）
    - memoryer/  : 记忆管理（Token 缓冲窗口）
    - prompter/  : 提示词模板管理
    - tooler/    : 工具注册与封装
      - gittool/ : Git 操作工具集
      - hugotool/: Hugo 博客工具
  - domain/      : 领域接口定义
  - model/       : 数据模型
  - repo/        : 数据持久化层
  - service/     : 业务逻辑层
- pkg/           : 公共工具库
  - conf/        : 配置管理（Viper）
  - db/          : JSON 文件数据库
  - errs/        : 自定义错误类型
  - exc/         : 异常处理与工具函数
  - logger/      : 日志系统（Zap）
  - task/        : 任务运行器
  - utils/       : 通用工具函数
- data/          : 用户数据存储
  - memory/     : 对话记忆存储
  - lock/       : 隐私数据
- logs/          : 日志文件
```

## 快速开始

```bash
# 安装依赖
go mod tidy

# 配置 API Key
# 编辑 config/dev.yml，填入 DeepSeek API Key

# 启动服务
go run ./cmd/main

# API 端点
# POST /agent/chat  - AI 对话接口
# POST /git/save    - Git Token 保存接口
```

## Docker 部署

```bash
# 启动 mifer 服务
docker-compose up -d
```

## 技术栈

- **语言**：Go 1.25
- **Web 框架**：Gin
- **AI 框架**：LangChainGo
- **大模型**：DeepSeek（OpenAI 兼容接口，可切换）
- **配置管理**：Viper
- **日志**：Zap
- **Git 操作**：go-git
- **部署**：Docker / Docker Compose

## 工具系统

mifer 内置以下工具供 Agent 调用：

### Git 工具
- `git_init` - 初始化仓库
- `git_add` - 暂存文件
- `git_commit` - 提交变更
- `git_push` - 推送到远程
- `git_pull` - 拉取远程更新
- `git_clone` - 克隆仓库
- `git_branch` - 分支管理
- `git_checkout` - 切换分支
- `git_remote` - 远程仓库管理
- `git_status` - 查看状态

### Hugo 工具
- `hugo_tool` - Hugo 博客生成操作（开发中）

## 配置说明

参见 `config/dev.yml`（开发环境）和 `config/prod.yml.example`（生产环境示例）：
- `ai.base_url` - 大模型 API 地址
- `ai.api_key` - API 密钥
- `ai.model` - 模型名称
- `ai.system_prompt` - 系统提示词模板
- `git.lock_path` - Git Token 存储路径

## 项目状态

项目处于架构搭建阶段，核心 AI Agent 对话链路已打通，工具系统持续完善中。

## 联系方式

如有技术问题和定制化需求，请提交 issue 或联系项目维护人。
