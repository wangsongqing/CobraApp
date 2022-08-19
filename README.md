## 一个干净的Golang初始化项目
 - 自己集成的一个Golang项目，常用功能已经集成，基本可以达到开箱即用
 - 可以作为基础项目使用

## 已经集成的功能
- MySQL
- Redis
- Logger 日志
- env 配置
- cobra 命令模式
- 文件操作

## 运行方式
- 在目录下面把 env.example 修改成 .env 文件，把相应的配置写进去，然后就可以愉快的玩耍了
- 方式 1：go run main.go city beijing shanghai
- 方式 2： go run main.go index -i 1 -e 11@qq.com

## 日志级别配置
.env 文件的 LOG_LEVEL=
```
"debug" —— 信息量大，一般调试时打开。系统模块详细运行的日志，例如 HTTP 请求、数据库请求、发送邮件、发送短信
"info" —— 业务级别的运行日志，如用户登录、用户退出、订单撤销。
"warn" —— 感兴趣、需要引起关注的信息。 例如，调试时候打印调试信息（命令行输出会有高亮）。
"error" —— 记录错误信息。Panic 或者 Error。如数据库连接错误、HTTP 端口被占用等。一般生产环境使用的等级。
```

## 打印执行的sql
修改 bootstrap/databases.go 
```
logger.Error 只打印错误信息
logger.Info  打印详细信息（包括执行的sql语句）
loggrr.Warn  打印提示信息
database.Connect(dbConfig, logger.Default.LogMode(logger.Error))
```