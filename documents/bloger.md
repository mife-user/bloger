# 项目介绍
## 核心定位
只需要 GitHub 账号，不用装环境、不用写代码、不用配服务器，通过 AI/Hugo 模板，一键生成并自动部署到 GitHub Pages 的静态博客工具。
## 优势
1. 极致低门槛普惠优势，面向非技术人群与普通用户，无需安装本地开发环境、无需掌握命令行操作、无需编写代码与配置文件，仅依托 GitHub 账号即可完成全流程博客搭建与部署，大幅降低独立博客的使用门槛，覆盖更广泛的长尾用户群体。
2. 零成本
基于 GitHub Pages 实现免费静态站点托管，支持无限流量、HTTPS 加密与全球访问，用户无需支付服务器、域名、存储等费用，实现真正意义上的零成本独立建站。相较于付费建站平台与云服务，具备极强的价格竞争力与用户吸引力。
3. AI 原生核心壁垒
以大模型 API 为核心驱动，内置 Hugo 专业知识库与写作、结构规划、SEO 优化、自动化部署等技能体系，可根据用户需求全自动生成博客框架、栏目结构、文章内容与配置文件，实现从需求到上线的全链路智能化，区别于传统模板工具与单纯 AI 文本生成工具。
4. 极简闭环生态优势
全程仅依赖 GitHub 平台完成授权、仓库创建、代码提交、Pages 部署，无需接入第三方托管、构建或云服务平台，流程极简、信任度高、稳定性强，形成轻量化、高可靠的产品闭环。
5. 静态站点技术优势
基于 Hugo 生成纯静态站点，具备访问速度快、安全性高、抗流量冲击、SEO 友好、长期可迁移维护等特性，优于动态建站系统（如 WordPress）与轻量化托管方案，满足个人品牌、学习笔记、技术博客等长期使用场景。
6. 赛道稀缺与差异化优势
当前暂无同类一体化产品，本项目为“AI + Hugo + GitHub Pages 全自动部署”的创新品类，精准填补低代码/零代码 AI 建站工具的空白，具备显著差异化与先发优势。
## 分析：
### 1. WordPress.com
- **优势**：功能强大、插件丰富、用户基数大
- **劣势**：需要付费才能自定义域名、速度较慢、学习曲线陡峭
- **差异化**：本项目零成本、AI驱动、无需配置

### 2. Hexo
- **优势**：轻量、快速、主题丰富
- **劣势**：需要Node.js环境、命令行操作、手动部署
- **差异化**：本项目无需环境、一键部署、AI生成

### 3. Hugo 官方主题库
- **优势**：主题多样、性能极佳
- **劣势**：需要安装Hugo、手动配置、无AI辅助
- **差异化**：本项目内置Hugo、AI自动配置、自动部署

### 4. Netlify/Vercel
- **优势**：自动化部署、支持多框架
- **劣势**：需要配置文件、学习成本、第三方依赖
- **差异化**：本项目仅依赖GitHub、零配置、AI引导

### 5. GPT-4 直接生成网页
- **优势**：灵活、可定制性强
- **劣势**：需要手动部署、无持续维护、无博客结构
- **差异化**：本项目结构化输出、自动部署、持续更新
## 使用流程
1. 下载软件并启动
2. 指定仓库目录文件夹，获得github授权
3. 选择以模板创建还是AI生成
4. 自动推送
5. 后可利用AI更新模板
## 技术栈
### 前端:
Vue 3 + Vite + Element Plus
### 后端:
go,gin,json，viper
### AI 框架:
LangChainGo
### 向量存储：
Chroma
### 软件界面：
Wails
## 实现:
### github授权:
通过前端拉取，后端存储
### 持久化存储：
- 用户配置：JSON/YAML
- 博客内容：git仓库
- ai对话历史：
	1. 短时记忆langchaingo拥有
	2. 长时记忆用户对话结束后利用ai压缩对话后文本存储
### 架构：
- .github/workflows：Github Actions实现目录（待创建）
- web：前端文件夹（待创建）
- cmd/main：程序启动
- cmd/bootstrap：存放启动函数，简化main
- internal/api/handler：handler处理器
  - agenthandler：AI Agent 处理器
  - githandler：GitHub 相关处理器
- internal/api/middleware：中间件
- internal/api/route：路由注册
- internal/domain：领域层，定义接口
- internal/model：数据模型层
  - gitmodel：GitHub 相关数据模型
- internal/repo：数据持久化层
  - gitrepo：GitHub 数据持久化
- internal/service：业务逻辑处理层
  - agentservice：AI Agent 业务逻辑
  - gitservice：GitHub 业务逻辑
- internal/ai：AI 能力层
  - agenter：主 Agent 逻辑，系统提示
  - exexcutor：执行器，处理对话和任务
  - llmer：大模型封装
  - memoryer：记忆管理，向量存储
  - prompter：提示词管理
  - tooler：工具封装
    - gittool：Git 操作工具
    - hugotool：Hugo 操作封装
      - template：内置默认 Hugo 风格模板目录
- config：存放配置文件
- pkg：通用工具层
  - conf：配置管理
  - db：数据库操作（JSON 文件存储）
  - err：错误处理
  - exc：异常处理和工具函数
  - logger：日志系统
  - utils：通用工具函数
- data：用户数据存储
  - memory/talks：聊天对话记录压缩存储目录
  - memory/paint：用户喜好，行为习惯文本存储目录
  - lock：用户隐私存储目录
- my_blog：生成 Hugo 目录，在此创建 GitHub 仓库，部署在 GitHub Pages 上
### 前端打包：
采用 **Wails 嵌入式方案**，将前端资源编译进二进制文件

#### 打包流程：
1. **前端构建**：`npm run build` → 生成 `web/dist` 目录
2. **Wails 编译**：`wails build` → 自动嵌入前端资源
3. **资源加载**：Wails 通过 `wails:generate` 生成 Go 绑定

						————mife

