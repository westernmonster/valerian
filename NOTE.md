# 架构笔记


1. 要梳理好业务的边界，合理的拆分 microservice
2. 降级，如果一个功能出了问题，用户会疯狂刷，如果联系刷就返回对应的ttl值，客户端进行处理。



服务树：

* valerian 1
  * admin 1
    - config 1
  * infra 2
    - discovery 1
    - config 2
  * service 3
  * interface 4


