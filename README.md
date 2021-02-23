## go-csv导出

> 目的实现一个接收数据即可导出的导出服务器

#### 已实现:
- 多任务协程csv导出，单任务每秒处理8W-10W条数据
- 支持http和rpc协议(grpc)请求
- 文件上传阿里云
- redis记录导出进度

#### 待实现

- 错误日志
- 自有文件服务器上传
- 七牛云服务器上传
- 腾讯云服务器上传
- 监听服务健康
- kafka队列
- etcd远程配置化
- 微服务化
- 服务治理
