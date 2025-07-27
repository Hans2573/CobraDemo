# CobraDemo - 多级命令行工具企业级架构

> **项目地址：** https://github.com/Hans2573/CobraDemo

## 💡 前言

在现代软件开发中，命令行工具（CLI）扮演着越来越重要的角色。无论是 Docker、Git、kubectl，还是各种开发框架的脚手架工具，优秀的命令行工具都能极大地提升开发效率和用户体验。

作为开发者，我们经常会遇到这样的需求：
- 🔧 **开发内部工具**：为团队开发自动化脚本和管理工具
- 📦 **构建产品 CLI**：为自己的产品提供命令行接口
- ⚡ **提升工作效率**：将重复性工作自动化
- 🎯 **学习最佳实践**：掌握企业级工具的设计模式

但是，如何设计一个结构清晰、易于维护、功能丰富的命令行工具呢？特别是当需求变得复杂，需要支持多级子命令、丰富的参数选项时，代码组织就成了一个挑战。

## 🎯 为什么选择 Cobra？

[Cobra](https://github.com/spf13/cobra) 是 Go 语言生态中最流行的 CLI 框架，被广泛应用于：
- **Kubernetes** - 容器编排平台的核心工具
- **Docker** - 容器技术的标准工具
- **Hugo** - 流行的静态网站生成器
- **Terraform** - 基础设施即代码工具

选择 Cobra 的理由：
- ✅ **简单易用**：API 设计直观，学习曲线平缓
- ✅ **功能完善**：支持子命令、参数、标志、别名等全部特性
- ✅ **自动生成**：自动生成帮助文档和 bash 补全
- ✅ **企业级**：已在大量生产环境中验证
- ✅ **社区活跃**：文档完善，问题响应及时

## 📖 项目介绍

**CobraDemo** 是一个完整的企业级命令行工具演示程序，采用最佳实践设计，展示了如何使用 Cobra 框架构建具有清晰层次结构和丰富参数功能的命令行应用程序。

### 🌟 项目特色

这个项目不仅仅是一个简单的 Demo，而是一个可以直接用于生产环境的完整解决方案：

- **🏗️ 完全模块化设计**：每个功能模块都有独立的目录和文件，便于团队协作
- **📋 多级命令结构**：支持复杂的命令层次，满足企业级应用需求
- **⚙️ 丰富的参数系统**：短参数、长参数、默认值、类型验证一应俱全
- **🎨 优雅的用户体验**：美观的输出格式、详细的帮助信息、友好的错误提示
- **🔧 实用的功能模块**：用户管理、配置管理、工具集等常见业务场景
- **📚 完整的文档**：详细的使用说明和开发指南

### 🎓 学习价值

通过这个项目，您将学会：

1. **架构设计**：如何设计可扩展的命令行工具架构
2. **代码组织**：如何优雅地组织多模块代码结构
3. **参数设计**：如何设计用户友好的参数系统
4. **用户体验**：如何提供出色的命令行用户体验
5. **最佳实践**：掌握企业级 CLI 工具的开发规范

### 🚀 适用场景

这个项目可以作为以下场景的起点：

- **企业内部工具**：员工管理、配置管理、运维工具等
- **产品 CLI 工具**：为您的 SaaS 产品提供命令行接口
- **开发者工具**：代码生成器、项目脚手架、部署工具等
- **学习项目**：Go 语言和 Cobra 框架的学习材料

## 🏗️ 项目结构

### 目录组织

```
cmd/
├── root.go              # 根命令
├── version.go           # 版本命令
├── user.go              # 用户命令注册
├── config.go            # 配置命令注册
├── tools.go             # 工具命令注册
├── user/                # 用户相关命令目录
│   ├── basic.go         # 基本用户命令 (list, create)
│   └── permission.go    # 权限管理命令
├── config/              # 配置相关命令目录
│   ├── basic.go         # 基本配置命令 (show, set)
│   └── backup.go        # 备份管理命令
└── tools/               # 工具相关命令目录
    └── basic.go         # 基本工具命令 (hash, time)
```

## 📋 命令层次结构

```
cobrademo
├── version                                    # 显示版本信息
├── user                                       # 用户管理
│   ├── list [--format|-f] [--active|-a] [--limit|-l]      # 列出用户
│   ├── create [name] [--email|-e] [--role|-r] [--force]   # 创建用户
│   └── permission                             # 权限管理
│       ├── list [user-id]                     #   列出权限
│       └── grant [user-id] [permission]       #   授予权限
├── config                                     # 配置管理
│   ├── show [--format|-f] [--section|-s]     #   显示配置
│   ├── set [key] [value] [--type|-t] [--validate|-v] [--backup|-b]  # 设置配置
│   └── backup                                 #   备份管理
│       ├── create [name]                      #     创建备份
│       └── list                               #     列出备份
└── tools                                      # 工具集
    ├── hash [text] [--algorithm|-a] [--uppercase|-u] [--file|-f]  # 计算哈希
    └── time [--format|-f] [--timezone|-z] [--unix|-u]            # 显示时间
```

## 🚀 使用示例

### 基本命令
```bash
# 显示版本信息
cobrademo version

# 查看帮助
cobrademo --help
cobrademo user --help
```

### 用户管理
```bash
# 查看用户管理概览
cobrademo user

# 列出所有用户 (基本用法)
cobrademo user list

# 高级用法：JSON格式显示活跃用户，限制5个
cobrademo user list --format json --active --limit 5

# YAML格式显示所有用户
cobrademo user list --format yaml

# 创建新用户 (基本用法)
cobrademo user create john

# 高级用法：创建管理员用户，设置邮箱，强制覆盖
cobrademo user create admin --email admin@example.com --role admin --force

# 创建普通用户，设置邮箱和角色
cobrademo user create alice --email alice@example.com --role user

# 权限管理
cobrademo user permission
cobrademo user permission list 1
cobrademo user permission grant 1 admin
```

### 配置管理
```bash
# 查看配置管理概览
cobrademo config

# 显示所有配置 (默认表格格式)
cobrademo config show

# JSON格式显示所有配置
cobrademo config show --format json

# 仅显示数据库配置节
cobrademo config show --section database

# YAML格式显示服务器配置
cobrademo config show --section server --format yaml

# 设置配置 (基本用法)
cobrademo config set database.host localhost

# 高级用法：设置配置并验证类型，创建备份
cobrademo config set database.port 3307 --type int --validate --backup

# 设置字符串类型配置
cobrademo config set server.host "0.0.0.0" --type string --validate

# 备份管理
cobrademo config backup
cobrademo config backup create prod_backup
cobrademo config backup list
```

### 工具集
```bash
# 查看工具概览
cobrademo tools

# 计算哈希值 (显示所有算法)
cobrademo tools hash "hello world"

# 指定算法计算哈希值
cobrademo tools hash "hello world" --algorithm sha256

# 大写输出SHA1哈希值
cobrademo tools hash "hello world" --algorithm sha1 --uppercase

# 从文件计算MD5哈希值
cobrademo tools hash --file ./example.txt --algorithm md5

# 显示时间信息 (默认格式)
cobrademo tools time

# 显示指定时区的时间
cobrademo tools time --timezone Asia/Shanghai

# 自定义时间格式
cobrademo tools time --format "2006年01月02日 15:04:05"

# 仅显示Unix时间戳
cobrademo tools time --unix

# 北京时间的自定义格式
cobrademo tools time --timezone Asia/Shanghai --format "2006-01-02 15:04:05 MST"
```

## ⚙️ 参数和标志详解

### 用户管理参数

#### `user list` 参数
- `--format, -f`: 输出格式 (`table`|`json`|`yaml`) [默认: table]
- `--active, -a`: 仅显示活跃用户 [布尔值]
- `--limit, -l`: 限制显示数量，0表示不限制 [默认: 0]

#### `user create` 参数
- `--email, -e`: 用户邮箱地址
- `--role, -r`: 用户角色 (`admin`|`user`|`guest`)
- `--force`: 强制创建，覆盖已存在用户 [布尔值]

### 配置管理参数

#### `config show` 参数
- `--format, -f`: 输出格式 (`table`|`json`|`yaml`) [默认: table]
- `--section, -s`: 仅显示指定配置节 (`database`|`server`|`logging`)

#### `config set` 参数
- `--type, -t`: 值类型 (`string`|`int`|`bool`|`float`)
- `--validate, -v`: 验证配置值格式 [布尔值]
- `--backup, -b`: 设置前创建备份 [布尔值]

### 工具集参数

#### `tools hash` 参数
- `--algorithm, -a`: 哈希算法 (`md5`|`sha1`|`sha256`|`sha512`)，留空显示所有
- `--uppercase, -u`: 输出大写哈希值 [布尔值]
- `--file, -f`: 从文件读取内容计算哈希

#### `tools time` 参数
- `--format, -f`: 自定义时间格式 (Go时间格式)
- `--timezone, -z`: 指定时区 (如: `UTC`, `Asia/Shanghai`)
- `--unix, -u`: 仅显示Unix时间戳 [布尔值]

## 🎯 设计原则

### 1. **完全模块化**
- 所有命令都在对应的子目录中实现
- 主文件只负责注册，不包含具体实现
- 每个功能模块都有独立的目录

### 2. **职责分离**
- **主文件（注册）**: `cmd/user.go`, `cmd/config.go`, `cmd/tools.go`
- **实现文件（逻辑）**: `cmd/user/basic.go`, `cmd/config/basic.go`, `cmd/tools/basic.go`

### 3. **易于扩展**
- 添加新命令只需在对应目录创建新文件
- 目录结构完全反映命令层次
- 结构清晰，便于理解和维护

### 4. **参数设计原则**
- 提供短参数和长参数形式
- 合理的默认值，减少必需参数
- 丰富的输出格式选择
- 安全操作选项（如备份、验证）

## 🔧 代码组织

### 导入方式
```go
// 在 user.go 中导入用户命令
import "github.com/Hans2573/CobraDemo/cmd/user"

// 在 config.go 中导入配置命令
import "github.com/Hans2573/CobraDemo/cmd/config"

// 在 tools.go 中导入工具命令
import "github.com/Hans2573/CobraDemo/cmd/tools"
```

### 注册方式
```go
// 在 init() 函数中注册命令
userCmd.AddCommand(user.ListCmd)
userCmd.AddCommand(user.CreateCmd)
userCmd.AddCommand(user.PermissionCmd)
```

### 参数定义方式
```go
// 定义参数变量
var (
    listFormat string
    listActive bool
    listLimit  int
)

// 在 init() 函数中绑定参数
func init() {
    ListCmd.Flags().StringVarP(&listFormat, "format", "f", "table", "输出格式")
    ListCmd.Flags().BoolVarP(&listActive, "active", "a", false, "仅显示活跃用户")
    ListCmd.Flags().IntVarP(&listLimit, "limit", "l", 0, "限制显示数量")
}
```

## 📦 添加新命令

### 步骤 1: 在对应目录创建新文件
```go
// cmd/user/role.go
package user

import (
    "fmt"
    "github.com/spf13/cobra"
)

var (
    roleActive bool
    roleFormat string
)

var RoleCmd = &cobra.Command{
    Use:   "role",
    Short: "角色管理",
    Long:  "用户角色管理相关的命令",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("🎭 用户角色管理")
        fmt.Println("查看子命令获取更多帮助")
    },
}

var RoleListCmd = &cobra.Command{
    Use:   "list",
    Short: "列出角色",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("角色列表 (格式: %s):\n", roleFormat)
        fmt.Println("1. admin")
        fmt.Println("2. user")
    },
}

func init() {
    RoleCmd.AddCommand(RoleListCmd)
    
    // 添加参数
    RoleListCmd.Flags().StringVarP(&roleFormat, "format", "f", "table", "输出格式")
    RoleListCmd.Flags().BoolVarP(&roleActive, "active", "a", false, "仅显示活跃角色")
}
```

### 步骤 2: 在主文件中注册
```go
// 在 cmd/user.go 中添加
userCmd.AddCommand(user.RoleCmd)
```

## 🌟 特性

### 1. **交互式帮助**
- 每个命令都有详细的帮助信息
- 支持 `--help` 参数获取使用说明
- 主命令提供功能概览和使用示例
- 参数说明清晰，包含默认值和可选项

### 2. **美观的输出**
- 使用表情符号和格式化输出
- 支持多种输出格式 (table、json、yaml)
- 清晰的命令层次展示
- 友好的用户体验

### 3. **完整的功能覆盖**
- **用户管理**: 创建、列表、权限管理，支持过滤和格式化
- **配置管理**: 显示、设置、备份，支持类型验证和安全操作
- **实用工具**: 哈希计算、时间显示，支持多种算法和格式

### 4. **丰富的参数支持**
- **输出格式**: 支持 table、json、yaml 多种格式
- **过滤选项**: 支持按条件筛选数据
- **安全功能**: 提供备份、验证、强制模式等安全选项
- **国际化**: 支持时区转换和自定义时间格式
- **文件操作**: 支持从文件读取内容处理

### 5. **用户体验优化**
- 合理的默认值，减少必需参数
- 短参数和长参数并存
- 详细的错误提示和帮助信息
- 渐进式功能，从简单到复杂

## 🛠️ 开发指南

### 安装和运行

#### 1. 安装依赖
```bash
go mod tidy
```

#### 2. 编译程序
```bash
go build -o cobrademo
```

#### 3. 运行程序
```bash
./cobrademo
```

### 文件职责

#### 主文件（仅负责注册）
- **cmd/user.go**: 导入并注册用户相关命令
- **cmd/config.go**: 导入并注册配置相关命令  
- **cmd/tools.go**: 导入并注册工具相关命令

#### 实现文件（包含具体逻辑）
- **cmd/user/basic.go**: 用户基本操作实现，包含丰富的参数处理
- **cmd/user/permission.go**: 用户权限管理实现
- **cmd/config/basic.go**: 配置基本操作实现，支持多格式输出和验证
- **cmd/config/backup.go**: 配置备份管理实现
- **cmd/tools/basic.go**: 工具基本操作实现，支持多种哈希算法和时间格式

### 优势总结

1. **结构清晰**: 所有子命令都在对应的目录中，层次分明
2. **易于维护**: 每个模块独立，修改不影响其他模块
3. **便于扩展**: 添加新功能只需在对应目录下创建新文件
4. **参数丰富**: 每个命令都有实用的参数选项
5. **代码复用**: 子命令可以在不同地方复用
6. **团队协作**: 不同开发者可以专注于不同的模块
7. **职责分离**: 每个文件都有明确的职责分工
8. **用户友好**: 提供多种输出格式和便捷选项

## 📚 依赖

- [github.com/spf13/cobra](https://github.com/spf13/cobra) - 命令行框架
- [github.com/spf13/viper](https://github.com/spf13/viper) - 配置管理

## 📄 许可证

MIT License

---

这个demo不仅展示了如何完全模块化地组织多级命令的目录结构，还演示了如何为命令行工具添加丰富的参数功能，提供了完整的企业级命令行工具开发模式供学习和参考。 