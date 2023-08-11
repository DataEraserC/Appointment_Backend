# 接口文档

## 用户注册接口

接口地址：/register

请求方法：POST

请求参数：
- username：用户名称，类型为字符串
- password：用户密码，类型为字符串

请求示例：
```
POST /register
Content-Type: application/json

{
    "username": "testuser",
    "password": "123456"
}
```

返回数据：
- code：返回状态码，0 表示成功，非0 表示失败
- message：返回信息，注册成功或失败的提示信息

返回示例：
```
{
    "code": 0,
    "message": "注册成功"
}
```

## 用户登录接口

接口地址：/login

请求方法：POST

请求参数：
- username：用户名称，类型为字符串
- password：用户密码，类型为字符串

请求示例：
```
POST /login
Content-Type: application/json

{
    "username": "testuser",
    "password": "123456"
}
```

返回数据：
- code：返回状态码，0 表示成功，非0 表示失败
- message：返回信息，登录成功或失败的提示信息
- token：用户登录后生成的令牌，类型为字符串

返回示例：
```
{
    "code": 0,
    "message": "登录成功",
    "token": "abcd1234"
}
```

## 用户获取个人信息接口

接口地址：/userinfo

请求方法：POST

请求参数：
- token：用户登录后生成的令牌，类型为字符串

请求示例：
```
POST /userinfo
Content-Type: application/json

{
    "token": "abcd1234"
}
```

返回数据：
- code：返回状态码，0 表示成功，非0 表示失败
- message：返回信息，获取个人信息成功或失败的提示信息
- data：用户个人信息的数据，包含用户名和密码

返回示例：
```
{
    "code": 0,
    "message": "获取个人信息成功",
    "data": {
        "username": "testuser",
        "password": "123456"
    }
}
```

## 用户修改个人信息接口

接口地址：/updateuserinfo

请求方法：POST

请求参数：
- token：用户登录后生成的令牌，类型为字符串
- username：修改后的用户名，类型为字符串

请求示例：
```
POST /updateuserinfo
Content-Type: application/json

{
    "token": "abcd1234",
    "username": "newuser"
}
```

返回数据：
- code：返回状态码，0 表示成功，非0 表示失败
- message：返回信息，修改个人信息成功或失败的提示信息

返回示例：
```
{
    "code": 0,
    "message": "修改个人信息成功"
}
```