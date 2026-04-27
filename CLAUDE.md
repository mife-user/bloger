# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Test

```bash
go build ./cmd/main/...          # build
go run ./cmd/main                # run (starts HTTP server on configured port)
go test ./...                    # all tests
go test -run TestName ./pkg/...  # single test
go vet ./...                     # vet
```

## Architecture

Layered architecture with dependency injection wired in `cmd/bootstrap/`. The app starts a Gin HTTP server exposing Agent chat and Git endpoints.

```
cmd/main → bootstrap.Init()
  ├─ LoadConfig()  (Viper, loads config/dev.yml)
  ├─ LoadLogger()  (Zap, file + console)
  └─ LoadRoute()   (Gin engine, CORS middleware)
       └─ route.NewRoute(config)
            ├─ gitrepo → gitservice → githandler  (POST /git/save)
            └─ executor → agentservice → agenthandler (POST /agent/chat)
```

Go module: `mifer`. All internal imports use the `mifer/...` prefix.

## Domain Interfaces

Defined in `internal/domain/birge.go`: `GitService`, `GitRepo`, `AgentService`, `Agent`. All layers depend on interfaces, not concrete types.

## AI Layer (`internal/ai/`)

Uses LangChainGo to assemble a ConversationalAgent. The wiring chain in `executor/init.go`:

```
Executor
  ├─ llmer     — DeepSeek API via OpenAI-compatible client (openai.New with custom BaseURL)
  ├─ prompter  — text/template with {chat_history} + {input} variables
  ├─ memoryer  — ConversationTokenBuffer (2048 token window)
  └─ tooler    — Tool registry
       ├─ gittool — go-git based Git ops (init, add, commit, push, pull, clone, branch, checkout, remote, status)
       └─ hugotool — stub (returns empty)
```

The Executor's `Chat` method calls `chains.Call(ctx, executor, input)` then saves context via `executor.Memory.SaveContext`.

Tools implement the LangChainGo `tools.Tool` interface (Name, Description, Call methods). The GitTool's Call method dispatches by `operation` field parsed from JSON input. No actual `git` CLI is shelled out — everything uses `go-git`.

## Config

`config/dev.yml` — loaded by Viper at startup. Key sections:
- `gin`: mode (debug/release), port, CORS origins/methods
- `ai`: base_url, api_key, model, system_prompt
- `git`: lock_path (JSON file for storing secrets)

Config search paths: `./config`, `<workdir>/config`, falls back to `../../config`.

## Key Dependencies

- `github.com/tmc/langchaingo` — Agent framework (ConversationalAgent, chains, memory, tools)
- `github.com/gin-gonic/gin` — HTTP framework
- `github.com/go-git/go-git/v5` — Pure Go Git library (no CLI dependency)
- `github.com/spf13/viper` — Configuration
- `go.uber.org/zap` — Structured logging
