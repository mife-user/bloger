# bloger

## 项目简介

bloger 是一个面向普通用户（尤其是非技术用户）的极简静态博客生成与部署工具，依托 AI + Hugo + GitHub Pages 实现零环境、零代码、零配置的一键建站体验。

核心目标
- 仅需 GitHub 账号授权即可完成建站全过程。
- AI 全自动生成博客内容、目录结构、主题配置、SEO 和持续更新方案。
- 完整集成 GitHub 仓库创建、代码提交、Pages 发布流程。

## 关键功能

1. GitHub 一键授权与仓库创建
2. Hugo 模板管理（内置多套主题），支持一键生成与切换
3. AI 文章与博客结构自动生成（建议、规划、排期）
4. 自动推送代码到 GitHub、自动触发 Pages 部署
5. 本地与远程双向同步：本地仓库(Git) + GitHub Pages
6. 对话记忆与用户偏好（短期/langchaingo 记忆 + 长期压缩存储）
7. 配置持久化（JSON/YAML）、多样化个性设置与SEO策略

## 项目架构

目录结构（主要模块）

- cmd/         : 启动模块
  - bootstrap/: 启动入口、初始化程序
  - main/     : 业务入口
- config/      : 配置文件目录
- internal/    : 核心业务与实现
  - api/      : HTTP REST API
    - handler/    : 处理器
      - agenthandler/: AI Agent 处理器
      - githandler/  : GitHub 相关处理器
    - middleware  : 中间件（鉴权、日志等）
    - route/      : 路由注册
  - service/  : 核心业务逻辑
    - agentservice/: AI Agent 业务逻辑
    - gitservice/  : GitHub 业务逻辑
  - domain/   : 领域接口定义
  - model/    : 数据模型层
    - gitmodel/   : GitHub 相关数据模型
  - repo/     : 数据持久化层
    - gitrepo/    : GitHub 数据持久化
  - ai/       : AI 能力层
    - agenter/    : Agent 逻辑、系统提示
    - exexcutor/  : 执行器，处理对话任务
    - llmer/      : 大模型封装与调用
    - memoryer/   : 记忆管理，向量存储
    - prompter/   : 提示词管理
    - tooler/     : 工具封装
      - gittool/    : Git 操作工具
      - hugotool/   : Hugo 操作封装
        - template/   : 内置默认 Hugo 风格模板目录
- pkg/         : 公共工具类库
  - conf/      : 配置管理
  - db/        : 数据库操作（JSON 文件存储）
  - err/       : 错误处理
  - exc/       : 异常处理和工具函数
  - logger/    : 日志系统
  - utils/     : 通用工具函数
- data/        : 用户数据存储
  - memory/talks/ : 聊天对话记录压缩存储目录
  - memory/paint/ : 用户喜好，行为习惯文本存储目录
  - lock/         : 用户隐私存储目录
- my_blog/     : 生成 Hugo 目录，在此创建 GitHub 仓库，部署在 GitHub Pages 上

## 运行与打包

1. 前端（Vue 3 + Vite + Element Plus）
   - 
pm install
   - 
pm run build -> 生成 web/dist
2. Wails 嵌入式打包
   - wails build -> 前端资源编译嵌入二进制
3. 启动程序
   - 运行 loger 可执行文件，或 go run ./cmd/main

## 实现方案亮点

- AI 原生：结合 LangChainGo，构建写作、目录、SEO、部署智能技能，自动化程度高。
- 零门槛：无须本地安装 Hugo/Node，后端封装 Hugo 生成逻辑。
- 零成本：仅使用 GitHub Pages 免费托管，自动化全链路发布。
- 轻量化：纯静态站点，低运维，高安全，便于长期迭代。

## 贡献指南

欢迎 Fork 并提交 PR：


## 联系方式

如需技术支持和定制化需求，请提交 issue 或联系项目维护人。
