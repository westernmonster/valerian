# MVP API变更记录


1. 获取用户粉丝有变更

GET /api/v1/account/list/follow  更改为 GET /api/v1/account/list/following

2. 获取用户资料中 follow\_count 更改为 following\_count

2. 增加接口

POST  /api/v1/account/follow
POST  /api/v1/account/unfollow

