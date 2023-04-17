# 维山帝web服务系统
## 产品结构
采用gin为基础web框架进行开发,数据库连接方向使用sqlx,日志采用zap,配置文件使用viper

运行时需要准备以下服务,并基于下列服务参数配置对应的配置文件内容 ↓↓↓
```
数据库: mysql
缓存: redis
对象存储: minio
```

## 产品内容 - 持续开发中
集成邮箱模式的用户注册登录功能

基于minio的对象存储服务

## 运行
下载源码到本地某一路径
``` powershell
git clone https://github.com/miacio/vishanti.git
```

### 编译
```
// unix
make build

// windows
build.bat
```

### 配置
``` toml
// bin目录下新建一个config.toml文件,然后写入下述内容

[email]
name="miajio" # 发送者邮件名称
mailAddr="miajio@163.com" # 发送者邮件地址
smtpAddr="smtp.163.com:25"
hostAddr="smtp.163.com"
password="XXXXXXXXXXXXXXX" # 邮箱授权码

[redis]
host="127.0.0.1:6379" # redis 服务器地址
password="" # redis 密码
db=0 # 默认连接库

[mysql]
host="127.0.0.1:3306" # 数据库地址
user="miajio" # 数据库用户名
password="123456" # 数据库密码
database="miajiodb" # 数据库名
charset="utf8mb4" # 字符集
parseTime="True" # 是否格式化时间
loc="Local" # 时区

[minio]
endPoint="127.0.0.1:9000" #  minio服务器地址 服务ID-地址
accessKey="Vvw9WyJlGWSBDbKl" #  minio 访问密钥
secretKey="Z0hXJiVGKKU2CmIqKGIrkrul0hEKYwCg" #  minio 保密密钥
useSSL=false # minio 服务协议是否使用https
```

### 运行
``` powershell
// unix
make start

// windows
start.bat

或者

// unix
bin/vishanti
// windows
bin\vishanti.exe
```

### 开发者指南
基于sqlx模式连接数据库,结构体字段采用基础数据类型(非指针模式),日期字段均使用指针模式,固需要注意:
数据库表设计时字段需设定默认值
