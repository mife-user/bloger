# Data Lock 目录说明

## 目录用途

此目录用于存储用户的隐私数据和敏感信息，这些文件**不会被上传到 GitHub**（已在 .gitignore 中配置忽略）。

## 存储内容

### secrets.json
存储用户的敏感凭证信息，包括：

- **GitHub OAuth Token**：用户的 GitHub 访问令牌
- **AI API Keys**：大模型 API 密钥
- **其他敏感配置**：用户自定义的隐私配置

## 文件结构示例

```json
{
  "github": {
    "access_token": "ghp_xxxxxxxxxxxxxxxxxxxx",
    "token_type": "bearer",
    "scope": "repo,user"
  },
  "ai": {
    "api_key": "sk-xxxxxxxxxxxxxxxxxxxx",
    "provider": "openai"
  },
  "user": {
    "preferences": {
      "theme": "dark",
      "language": "zh-CN"
    }
  }
}
```

## 安全说明

1. **切勿提交**：此目录下的所有 `.json` 文件已在 `.gitignore` 中配置忽略，不会被提交到 Git 仓库
2. **本地存储**：所有敏感数据仅存储在用户本地，确保隐私安全
3. **权限控制**：建议在生产环境中设置适当的文件权限，防止未授权访问
4. **数据备份**：用户应自行备份此目录中的重要数据

## 多人协作说明

由于此目录包含用户个人隐私数据，在多人协作开发时：

- **开发者**：需要自行创建 `secrets.json` 并填写自己的测试凭证
- **示例文件**：可参考 `config/prod.yml.example` 了解需要配置哪些参数
- **测试数据**：请勿在此目录中提交任何测试用的真实凭证信息

## 相关配置

- Git 忽略规则：`data/lock/*.json`（见项目根目录 `.gitignore`）
- 配置示例：`config/prod.yml.example`
