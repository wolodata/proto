# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目概述

这是 Wolodata 微服务系统的 **Protocol Buffer API 定义仓库**，使用 Buf 工具进行管理，支持 Go 和 TypeScript 代码自动生成。

- **存储库**: https://github.com/wolodata/proto.git
- **Buf Registry**: buf.build/wolodata/api
- **代码生成**: Go (ConnectRPC) + TypeScript
- **版本管理**: Buf v2 + buf.lock

## 常用开发命令

```bash
# 格式化 + Lint 检查（开发必用）
make lint

# 首次设置：安装 Buf 工具
make init

# 更新依赖（如 googleapis）
make update
```

## 代码架构

### 模块组织结构

所有服务模块采用标准结构：`module_name/v1/module_name.proto`

**核心服务分组**:

1. **认证与授权**
   - `auth/v1` - 登录、SMS 验证、多因素认证
   - `permission/v1` - RBAC 系统（角色、方法、用户分配、权限限额）

2. **用户与关注**
   - `user/v1` - 用户信息、设备管理、权限查询
   - `follow/v1` - 用户关注主题管理

3. **内容管理**
   - `article/v1` - 文章 CRUD、翻译、总结、流式总结、智能问题生成
   - `source/v1` - 信源 + 目录树管理（支持层级、排序）
   - `topic/v1` - 领域/主题/命中统计

4. **关键词与分类**
   - `keyword/v1` - 关键词管理（强/弱）、重索引
   - `category/v1` - 分类管理、LLM 参数配置

5. **AI 对话系统**
   - `chat/v1` - 文章相关的流式 AI 对话
   - `brain/v1` - 通用智能对话、对话历史管理
   - `assistant/v1` - AI 助手 CRUD、对话管理

6. **配置与任务**
   - `config/v1` - 提示词配置 + 系统配置
   - `task/v1` - 后台任务优先级、状态追踪
   - `notify/v1` - 通知消息（新文章推送）

7. **特殊集成**
   - `crawler/v1` - 爬虫配置、代理管理
   - `report/v1` - 报告生成（日/周/月/自定义）
   - `zsxq/v1` - ZSXQ 平台集成（书籍、PDF 管理）
   - `cspan/v1` - C-SPAN 数据源
   - `federalregister/v1` - 美国联邦公报服务
   - `convert/v1` - 转换服务

### 后台管理 API 识别

多个服务包含"后台管理"专用 RPC（通过注释标识），主要包括：
- `user.v1`: ListUsers, CreateUser, AdminUpdate*
- `article.v1`: UpdateArticle, CreateArticle, DeleteArticles, BatchOperate*
- `permission.v1`: 所有角色和方法管理 RPC
- `source.v1`: CreateSource, DeleteSource, UpdateSource, Folder*

### 权限系统架构重点

`permission/v1` 是最近重构的核心模块：
- 使用 `enabled_methods` 替代旧的 `function_sets_enabled`
- 支持动态方法发现（通过 protoregistry）
- 角色包含 LLM 模型禁用、关键词/范畴/免费提问次数限额

## 代码生成配置

**buf.gen.yaml 生成目标**:
```yaml
1. Go 官方编译器 (paths=source_relative)
2. ConnectRPC Go 插件 (gRPC 兼容)
3. TypeScript/ES (bufbuild/es)

Go 包前缀: github.com/wolodata/proto/gen
```

**生成代码分发**:
- Go: 通过 Go modules 导入
- TypeScript: 通过 npm 包分发
- BSR 自动生成: 推送到 buf.build/wolodata/api 后自动触发

## 开发工作流

```bash
1. 编辑 proto 定义
   vim module_name/v1/module_name.proto

2. 格式化 + 检查
   make lint

3. 提交更改
   git add .
   git commit -m "feat: describe changes"
   git push origin master

4. 代码生成（自动）
   Buf BSR 自动生成到 gen/ 目录
   其他项目通过 Go modules/npm 导入使用
```

## Buf 配置说明

**buf.yaml 核心配置**:
- 模块名: `buf.build/wolodata/api`
- 依赖: `googleapis/googleapis` (Google 标准库)
- Lint 规则: STANDARD（例外：FIELD_NOT_REQUIRED, PACKAGE_NO_IMPORT_CYCLE）
- 兼容性检查: FILE 级（例外：EXTENSION_NO_DELETE, FIELD_SAME_DEFAULT）

## 注意事项

1. **不要在本地生成代码**：所有需要生成的代码由 Buf 在服务器上自动管理
2. **修改后必须运行 make lint**：任何 proto 文件修改后必须执行 `make lint` 进行格式化和检查，这是强制性的开发规范
3. **Proto 注释使用中文**：现有代码库采用中文注释风格
4. **版本管理**：所有模块使用 v1 版本
5. **依赖更新**：修改 buf.yaml 依赖后需运行 `make update` 更新 buf.lock
6. **命名约定**：RPC 方法采用 PascalCase，message 字段采用 snake_case
7. **流式 RPC**：AI 对话相关服务广泛使用 `stream` 响应（如 StreamSummary, StreamChat）
