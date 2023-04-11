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