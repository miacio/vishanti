# 维山帝web服务系统
采用gin为基础web框架进行开发,数据库连接方向使用sqlx

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
name="" # 发送者邮件名称 xxx@mail.com的 xxx
mailAddr="" # 发送者邮件地址 xxx@mail.com
smtpAddr="" # 邮箱服务器(带端口协议) 例如: smtp.163.com:25
hostAddr="" # 邮箱服务器(不带端口协议) 例如: smtp.163.com
password="" # 邮箱授权码

[redis]
host="127.0.0.1:6379" # redis 服务器地址
password="" # redis 密码
db=0 # 默认连接库

[mysql]
host="127.0.0.1:3306" # 数据库地址
user="" # 数据库用户名
password="123456" # 数据库密码
database="" # 数据库名
charset="utf8mb4" # 数据库字符集
parseTime="True" # 是否分析时间 True或不给参数
loc="Local" # 时间地址
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