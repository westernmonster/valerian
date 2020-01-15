# 架构笔记


1. 要梳理好业务的边界，合理的拆分 microservice
2. 降级，如果一个功能出了问题，用户会疯狂刷，如果联系刷就返回对应的ttl值，客户端进行处理。



服务树：

* valerian 1
  * admin 1
    - config 1101
  * infra 2
    - discovery 1201
    - config 2 1202
  * service 3
    - account 1301
    - identify 1302
    - msm 1303
    - relation 1304
    - topic 1305
    - discuss 1306
    - feed 1307
    - fav 1308
    - like 1309
    - article 1310
    - recent 1311
    - topic-feed 1312
    - account-feed 1313
    - message 1314
    - comment 1315
    - search 1316
  * interface 4
    - account  1401
    - article 1402
    - discuss 1404
    - dm 1405
    - feed 1408
    - feedback 1409
    - passport-auth 1413
    - passport-login 1414
    - passport-register 1415
    - recent 1416
    - topic 1417
    - valcode 1418
    - search 1420
    - comment 1421
    - common 1423
