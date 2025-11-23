# 测试

- 完整的数据库连接和连接池配置

- 基本的 CRUD 操作

- HTTP API 接口

- 错误处理

## 创建用户

curl -X POST http://localhost:8080/users/create \
 -d "name=张三" \
 -d "email=zhangsan@example.com"

## 获取所有用户

curl http://localhost:8080/users

## 获取用户

curl "http://localhost:8080/user?id=1"
