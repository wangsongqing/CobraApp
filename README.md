## 一个干净的Golang初始化项目
 - 自己集成的一个Golang项目，常用功能已经集成，基本可以达到开箱即用
 - 可以作为基础项目使用

## 已经集成的功能
- MySQL
- Redis
- Logger 日志
- env 配置
- cobra 命令模式

## 运行方式
- 在目录下面把 env.example 修改成 .env 文件，把相应的配置写进去，然后就可以愉快的玩耍了
- 方式 1：go run main.go city beijing shanghai
- 方式 2：go run main.go user -u miaowing -a 20