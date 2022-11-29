## 一个干净的Golang CLI初始化项目
 - 自己集成的一个Golang CLI项目，常用功能已经集成，基本可以达到开箱即用
 - 可以作为基础项目使用
 - 目录结构类Laravel

## 已经集成的功能
- MySQL
- Redis
- Logger 日志
- env 配置
- cobra 命令模式
- 文件操作
- RabbitMQ
- Excel读写操作
- Elasticsearch

## 运行方式
- 执行命令：go mod tidy
- 在目录下面把 env.example 修改成 .env 文件，把相应的配置写进去，然后就可以愉快的玩耍了
- 方式 1: go run main.go index -i 1 -e 11@qq.com
- 方式 2: go run main.go excel (read || write)
- 方式 3: go run main.go json (decode || encode)
- 方式 4: go run main.go rabbitmq simple send
- 方式 5: go run main.go time
- 方式 6: go run main.go es (create || select || update || delete || add || search)

## 日志级别配置
.env 文件的 LOG_LEVEL=
```
开发时推荐使用 "debug" 或者 "info" ，生产环境下使用 "error"

"debug" —— 信息量大，一般调试时打开。系统模块详细运行的日志，例如 HTTP 请求、数据库请求、发送邮件、发送短信
"info" —— 业务级别的运行日志，如用户登录、用户退出、订单撤销。
"warn" —— 感兴趣、需要引起关注的信息。 例如，调试时候打印调试信息（命令行输出会有高亮）。
"error" —— 记录错误信息。Panic 或者 Error。如数据库连接错误、HTTP 端口被占用等。一般生产环境使用的等级。
```