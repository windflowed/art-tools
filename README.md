# 艺术平台管理员工具

这是艺术平台的管理员工具项目，提供创建管理员账号等实用工具。

## 功能

- 创建管理员账号：自动向数据库添加管理员用户，包括用户名、邮箱和密码（安全哈希处理）

## 技术栈

- 语言：Go
- 数据库驱动：github.com/go-sql-driver/mysql
- 密码加密：golang.org/x/crypto/bcrypt

## 安装

### 前置条件

- Go 1.16+
- MySQL 数据库（艺术平台数据库已创建）

### 安装依赖

```bash
go mod download
```

## 使用方法

### 创建管理员账号

1. 打开 `create_admin.go` 文件，根据需要修改以下配置：

   ```go
   // 数据库连接配置
   dsn := "root:art123456@tcp(192.168.1.47:20000)/artplatform?charset=utf8mb4&parseTime=True&loc=Local"
   
   // 管理员账号信息，自己填充
   email := "******"      // 管理员的邮箱
   password := "******"   // 请使用一个强密码，并妥善保管
   name := "******"       //管理员的姓名
   ```

   请务必修改为安全的密码！

2. 运行创建管理员账号的命令：

   ```bash
   go run create_admin.go
   ```

3. 如果执行成功，将显示以下信息：

   ```
   Successfully connected to the database!
   Admin user 'Super Admin' with email 'admin@example.com' has been successfully added.
   ```

   如果管理员账号已存在，将显示：

   ```
   Successfully connected to the database!
   Admin user with email admin@example.com already exists. Skipping creation.
   ```

## 自定义配置

### 数据库连接

修改 `create_admin.go` 文件中的 DSN (Data Source Name) 字符串：

```go
dsn := "username:password@tcp(host:port)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
```

- `username`: 数据库用户名
- `password`: 数据库密码
- `host`: 数据库主机地址
- `port`: 数据库端口
- `database_name`: 数据库名称

### 管理员信息

修改 `create_admin.go` 文件中的管理员信息：

```go
email := "your_admin_email@example.com"
password := "your_secure_password"
name := "Your Admin Name"
```

## 安全注意事项

- 不要在代码中硬编码生产环境的数据库凭据
- 使用强密码并妥善保管
- 在创建管理员账号后，立即更改默认密码
- 限制对此工具的访问，仅允许授权人员使用

## 故障排除

### 数据库连接错误

如果遇到数据库连接错误，请检查：

1. 数据库服务器是否运行
2. 数据库连接字符串是否正确
3. 数据库用户是否有足够的权限
4. 网络连接是否正常

### 用户创建失败

如果用户创建失败，可能的原因包括：

1. 数据库表结构不匹配
2. 数据库用户没有写入权限
3. 邮箱已被使用

## 开发

### 添加新工具

1. 在项目根目录创建新的 Go 文件
2. 实现所需功能
3. 更新 README.md 文件，添加新工具的使用说明

## 许可证

[MIT License](LICENSE)
