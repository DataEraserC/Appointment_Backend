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
- data：用户个人信息的数据，包含用户名

返回示例：
```
{
    "code": 0,
    "message": "获取个人信息成功",
    "data": {
        "ID": 1,
        "Username": "username1",
        "Password": "",
        "Avatar": "/default_avatar.png",
        "NickName": "NickName1",
        "PhoneNumber": "123456789"
    }
}
```

## 用户修改个人信息接口

接口地址：/updateuserinfo

请求方法：POST

请求参数：
- token：用户登录后生成的令牌，类型为字符串
- username：修改后的用户名，类型为字符串
- password：修改后的密码，类型为字符串
- avatar：修改后的头像链接，类型为字符串
- nickname：修改后的昵称，类型为字符串
- phoneNumber：修改后的手机号，类型为字符串

请求示例：
```
POST /updateuserinfo
Content-Type: application/json

{
    "token": "abcd1234",
    "username": "newuser",
    "password": "newpassword",
    "avatar": "new_avatar.jpg",
    "nickname": "NewNickname",
    "phoneNumber": "1234567890"
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
## 添加地点接口

接口地址：/addlocation

请求方法：POST

请求参数：
- token：用户登录后生成的令牌，类型为字符串
- name：地点名称，类型为字符串
- description：地点描述，类型为字符串

请求示例：
```
POST /addlocation
Content-Type: application/json

{
    "token": "abcd1234",
    "name": "New Location",
    "description": "This is a new location"
}
```

返回数据：
- code：返回状态码，0 表示成功，非0 表示失败
- message：返回信息，添加地点成功或失败的提示信息
- data：返回的数据，添加成功后返回新创建的地点信息

返回示例：
```
{
    "code": 0,
    "message": "添加地点成功",
    "data": {
        "id": 1,
        "name": "New Location",
        "description": "This is a new location"
    }
}
```
## 预约地点搜索接口

接口地址：/searchlocation

请求方法：POST

请求参数：
- token：用户登录后生成的令牌，类型为字符串
- keyword：搜索关键词，类型为字符串

请求示例：
```
POST /searchlocation
Content-Type: application/json

{
    "token": "abcd1234",
    "keyword": "Location"
}
```

返回数据：
- code：返回状态码，0 表示成功，非0 表示失败
- message：返回信息，搜索地点成功或失败的提示信息
- data：返回的数据，搜索成功后返回符合搜索条件的地点列表

返回示例：
```
{
    "code": 0,
    "message": "搜索成功",
    "data": [
        {
            "id": 1,
            "name": "Location 1",
            "description": "This is location 1"
        },
        {
            "id": 2,
            "name": "Location 2",
            "description": "This is location 2"
        }
    ]
}
```

## 用户预约接口

接口地址：/reservation

请求方法：POST

请求参数：
- token：用户登录后生成的令牌，类型为字符串
- location_id：地点ID，类型为字符串
- date：预约日期，类型为字符串

请求示例：
```
POST /reservation
Content-Type: application/json

{
    "token": "abcd1234",
    "location_id": "1",
    "date": "2022-01-01"
}
```

返回数据：
- code：返回状态码，0 表示成功，非0 表示失败
- message：返回信息，预约成功或失败的提示信息
- data：返回的数据，预约成功后返回预约记录信息

返回示例：
```
{
    "code": 0,
    "message": "预约成功",
    "data": {
        "id": 1,
        "user_id": 123,
        "location_id": 1,
        "date": "2022-01-01"
    }
}
```