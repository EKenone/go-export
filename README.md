## go-csv导出

> 目的实现一个接收数据即可导出的导出服务器

#### 项目测试

1. 把http.yaml.loc和rpc.yaml.loc去掉后缀，并且配置自身的配置参数
2. http运行：到api/http目录下运行 go run http.go
3. rpc运行：到api/rpc目录下运行 go run rpc.go
4. 分别运行test下面的http_test.go和rpc_test.go可调试


#### 项目运行
> 依赖makefile管理
1. 打包项目: ``make build``
2. 运行项目: ``make run``
3. 终止项目: ``make stop``

#### 已实现:

- 多任务协程csv导出，单任务每秒处理8W-10W条数据
- 支持http和rpc协议(grpc)请求
- 文件上传至阿里云
- 文件上传至腾讯云
- redis记录导出进度

#### 待实现

- 错误日志、异常通知
- kafka队列
- 文件上传至七牛云
- socket通讯实时导出进度
- 监听服务健康、调用追踪
- etcd远程配置化
- 服务安全
