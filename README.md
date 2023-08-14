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

## 修改地点接口

接口地址：/updatelocation

请求方法：POST

请求参数：
- token：用户登录后生成的令牌，类型为字符串
- location_id：地点ID，类型为整数
- name：地点名称，类型为字符串
- description：地点描述，类型为字符串

请求示例：
```
POST /updatelocation
Content-Type: application/json

{
    "token": "abcd1234",
    "location_id": 1,
    "name": "Updated Location",
    "description": "This is an updated location"
}
```

返回数据：
- code：返回状态码，0 表示成功，非0 表示失败
- message：返回信息，更新地点成功或失败的提示信息

返回示例：
```
{
    "code": 0,
    "message": "更新地点成功"
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
- data：返回的数据，搜索成功后返回符合搜索条件的地点列表(注意:如果要列出所有可以将keyword设为空)

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
- location_id：地点ID，类型为整数
- date：预约日期，类型为字符串
- time：预约时间，类型为字符串

请求示例：
```
POST /reservation
Content-Type: application/json

{
    "token": "abcd1234",
    "location_id": 1,
    "date": "2022-01-01",
    "time": "17:00-19:00"
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

## 用户预约记录列表接口

接口地址：/listrecord

请求方法：POST

请求参数：
- token：用户登录后生成的令牌，类型为字符串

请求示例：
```
POST /listrecord
Content-Type: application/json

{
    "token": "abcd1234"
}
```

返回数据：
- code：返回状态码，0 表示成功，非0 表示失败
- message：返回信息，搜索成功或失败的提示信息
- data：返回的数据，搜索成功后返回预约记录列表

返回示例：
```
{
    "code": 0,
    "message": "搜索成功",
    "data": [
        {
            "ID": 1,
            "UserID": 1,
            "LocationID": 1,
            "Date": "2022-01-01",
            "Time": "17:00 - 19:00"
        },
        {
            "ID": 2,
            "UserID": 1,
            "LocationID": 1,
            "Date": "2022-01-01",
            "Time": "17:00 - 19:00"
        }
    ]
}
```

# 用户预约记录详细列表接口

接口地址：/listrecorddetail

请求方法：POST

请求参数：
- token：用户登录后生成的令牌，类型为字符串

请求示例：
```
POST /listrecorddetail
Content-Type: application/json

{
    "token": "abcd1234"
}
```

返回数据：
- code：返回状态码，0 表示成功，非0 表示失败
- message：返回信息，搜索成功或失败的提示信息
- data：返回的数据，搜索成功后返回预约记录详细列表

返回示例：
```
{
    "code": 0,
    "message": "搜索成功",
    "data": [
        {
            "ID": 1,
            "UserID": 1,
            "LocationID": 1,
            "Date": "2022-01-01",
            "Time": "17:00 - 19:00",
            "LocationName": "会议室A",
            "LocationDescription": "适用于小型会议"
        },
        {
            "ID": 2,
            "UserID": 1,
            "LocationID": 2,
            "Date": "2022-01-01",
            "Time": "17:00 - 19:00",
            "LocationName": "会议室B",
            "LocationDescription": "适用于大型会议"
        }
    ]
}
```

注意事项：
- 需要用户登录后才能调用该接口，所以需要先调用登录接口获取到token，并在请求参数中带上token进行请求。
- 在这个接口中，我们联合查询了预约记录和地点表，通过地点表的信息来获取预约记录的详细信息。
- 返回的数据中，增加了地点名称和地点描述的字段，方便用户查看预约记录时能够直接看到预约的地点信息。

## 查询地点信息接口

接口地址：/locationinfo

请求方法：POST

请求参数：
- token：用户登录后生成的令牌，类型为字符串
- location_id：地点ID，类型为整数

请求示例：
```
POST /locationinfo
Content-Type: application/json

{
    "token": "abcd1234",
    "location_id": 1
}
```

返回数据：
- code：返回状态码，0 表示成功，非0 表示失败
- message：返回信息，查询成功或失败的提示信息
- data：返回的数据，查询成功后返回地点的详细信息

返回示例：
```
{
    "code": 0,
    "message": "查询成功",
    "data": {
        "id": 1,
        "name": "New Location",
        "description": "This is a new location"
    }
}
```