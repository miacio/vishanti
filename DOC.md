# 维山帝接口文档
## 内置
### ping
接口地址: / 或 /ping

请求方式: GET

接口参数:  无

响应: 
```
{
    "code": int 200,
    "msg": string "pong"
}
```

### error 错误页
接口地址: /error

请求方式: GET|POST

接口参数: 使用json、uri、form表单方式均可
```
{
    "code": string 响应码 必填
    "msg": string 消息信息 必填
    "err": string 错误信息
}
```

响应:

错误界面

## 字典服务
### 依据字段组获取字典列表
接口地址: /dict/findByGroup

请求方式: GET

接口参数: string group 组名 必填

响应:
```
{
    "code": int 状态码,
    "data": [{
        "id": "180AB176DB6711ED88C504421A25E929", 字典ID
        "name": "普通用户", 当前字段类型
        "group": "USER_VIP", 组名
        "parent_group": "USER", 上级组名
        "describe": "用户VIP类别", 描述
        "val": "1", 值
        "create_time": "2023-04-15 16:25:34", 创建时间
        "create_by": "", 创建人
        "update_time": null, 修改时间
        "update_by": "" 修改人
    }...], 字典列表
    "msg": "获取成功"
}
```

### 批量写入字典数据
接口地址: /dict/inserts

请求方式: POST

请求参数: 
header: token = 登录后响应的token码
body: 
```
[{
    "name": string 当前字段类型名称,
    "group": string 当前字段组,
    "parentGroup": string 上级组,
    "describe": string 字段描述,
    "val": string 值,
}...]
```

响应:
```
{
  "code": 200,
  "msg": "写入成功"
}
```

## 文件服务
### 文件上传接口
接口地址: /file/upload

请求方式: POST

请求参数: 
header: token = 登录授权id

form表单参数: 
```
{
    "region": "beijing", // 基于业务的地区 - 用于存储minio使用
    "md5": "xxxxxxxxx", // 当前文件加密后的md5 用于校验文件的一致性
    "file": "file 文件"
}

提交方式前端可参考: page/upload_file_test.html
```

### 文件下载接口
接口地址: /file/load?id=文件id

请求方式: GET

## 邮箱服务
### 发送邮件接口
接口地址: /email/sendCheckCode

请求方式: POST

请求参数:
```
{
    "email": "xxxx@qq.com", 接收邮件的账号
    "emailType": "register" 邮件类型 register login update
}
```

响应:
```
{
    "code": 200,
    "msg": "发送成功",
    "data": "当前邮件id,由于后续校验邮件验证码对应的绑定id"
}
```

## 用户服务
### 获取token信息接口
接口地址: /user/token

请求方式: GET

请求参数: string token

响应:
```
{
  "code": 200, 响应码
  "data": {
    # 账号信息
    "accountInfo": {
      "id": "B178CCA3DB6311ED88C504421A25E929", 账号id
      "mobile": "18616220047", 手机号
      "email": "miajio@163.com", 邮箱
      "account": "miajio", 账号
      "password": "",
      "create_time": "2023-04-15 16:01:14", 创建时间
      "update_time": "2023-04-15 16:17:37", 修改时间
      "status": "1", 用户状态
      "lock_time": null 封号时间
    },
    # 用户信息
    "detailedInfo": {
      "id": "950908E4DB6711ED88C504421A25E929", 用户id
      "user_account_id": "B178CCA3DB6311ED88C504421A25E929", 用户对应的账号id
      "vip": "4", 用户vip等级
      "head_pic_id": "", 头像文件id
      "nick_name": "miajio", 昵称
      "sex": "1", 性别
      "birthday_year": 1995, 出生年
      "birthday_month": 11, 出生月
      "birthday_day": 2, 出生日
      "profile": "个人简介" 个人简介
    },
    "headPic": "" 头像地址
  },
  "msg": "获取成功"
}
```

### 邮箱注册
接口地址: /user/email/register

请求方式: POST

请求参数:
```
{
    "email": "xxxx@qq.com", 邮箱地址
    "code": "DB7ZN6", 验证码
    "uid": "291A4E0C605D4DEF8FE76ED96E47F005", 邮件id
    "nickName": "昵称", 昵称
    "account": "vishanti", 账号
    "password": "123456" 密码
}
```

响应:
```
{
    "code": 200,
    "msg": "登录成功",
    "data": "token"
}
```

### 邮箱验证码登录
接口地址: /user/email/login

请求方式: POST

请求参数:
```
{
    "email":  "xxxx@qq.com", 邮箱地址
    "code": "DB7ZN6", 验证码,
    "uid": "291A4E0C605D4DEF8FE76ED96E47F005" 邮件id
}
```

响应:
```
{
    "code": 200,
    "msg": "登录成功",
    "data": "token"
}
```

### 邮箱密码登录
接口地址: /user/email/loginPwd

请求方式: POST

请求参数:
```
{
    "email": "xxxx@qq.com", 邮箱地址
    "password": "password" 密码
}
```

响应:
```
{
    "code": 200,
    "msg": "登录成功",
    "data": "token"
}
```

## 用户信息模块
### 修改用户信息接口
接口地址: /user/detailed/update

请求方式: POST

请求参数:
```
token: 

{
    "nickName": "昵称",
    "sex": 1 男 2 女 0 未知,
    "birthdayYear": int 出生年
    "birthdayMonth": int 出生月
    "birthdayDay": int 出生日
}
```

响应:
```
{
    "code": 200,
    "msg": "修改成功"
}
```

### 修改用户头像接口(与文件上传模式一致)
接口地址: /user/detailed/updateHeadPic

请求方式: POST

请求参数: 
header: token = 登录授权id

form表单参数: 
```
{
    "region": "beijing", // 基于业务的地区 - 用于存储minio使用
    "md5": "xxxxxxxxx", // 当前文件加密后的md5 用于校验文件的一致性
    "file": "file 文件"
}

提交方式前端可参考: page/upload_file_test.html
```