# 人事档案管理系统

基于 Vue 3 + Go 的人事档案管理系统，支持 Docker 和 Podman 部署。

## 功能特性

- 员工信息管理（增删改查）
- 部门管理
- 数据统计和可视化
- 搜索和筛选功能

## 技术栈

- 前端：Vue 3 + Tailwind CSS
- 后端：Go + Gin + GORM
- 数据库：SQLite

## 快速开始

### 使用 Docker

```bash
# 构建和运行
docker-compose up --build

# 停止服务
docker-compose down

### 使用 Podman

# 配置代理（如果需要）
set HTTP_PROXY=http://127.0.0.1:1080

# 使用 podman-compose
podman-compose up --build

# 或者分别构建和运行
podman build -t traesys-frontend .
podman build -t traesys-backend ./backend
podman run -d -p 80:80 traesys-frontend
podman run -d -p 8080:8080 traesys-backend