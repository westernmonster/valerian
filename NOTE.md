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
  * interface 4
    - account  1401
    - article 1402
    - certification 1403
    - discuss 1404
    - dm 1405
    - editor 1406
    - fav 1407
    - feed 1408
    - feedback 1409
    - file 1410
    - locale 1411
    - location 1412
    - passport-auth 1413
    - passport-login 1414
    - passport-register 1415
    - recent 1416
    - topic 1417
    - valcode 1418
