# 这个文件里面的命令直接复制粘贴到终端执行
# 可以通过curl测试后端业务逻辑是否正确


# 注册逻辑测试
### 1. 注册成功
curl -X POST http://localhost:8080/user/signup \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "Password123!",
    "confirmPassword": "Password123!"
}'

### 2. 注册失败
# 2.1 密码格式错误
curl -X POST http://localhost:8080/user/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"123","confirmPassword":"123"}'
# 2.2 两次密码不一致
curl -X POST http://localhost:8080/user/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"Password123!","confirmPassword":"Password321!"}'
# 2.3 邮箱格式错误
curl -X POST http://localhost:8080/user/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"bademail.com","password":"GoodPass123!","confirmPassword":"GoodPass123!"}'

