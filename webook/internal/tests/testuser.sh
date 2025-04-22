# 这个文件里面的命令直接复制粘贴到终端执行
# 可以通过curl测试后端业务逻辑是否正确


# 注册逻辑测试
curl -X POST "http://localhost:8080/user/signup" \
     -H "Content-Type: application/json" \
     -d '{
           "email": "test@example.com",
           "password": "123456",
           "confirmPassword": "123456"
         }'
