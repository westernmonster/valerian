# 说明文档


## 技术选型

* 数据库采用 MYSQL


* SQL 版本管理采用 sql-migrate
> github.com/rubenv/sql-migrate

* ID 生成采用预先生成，不使用数据库自增长
> github.com/bwmarrin/snowflake


* 本地开发采用 realize 自动编译、重启
> github.com/oxequa/realize

* 静态文件嵌入使用 packr
> github.com/gobuffalo/packr

* 追求效率，不使用 ORM，采用 sqlx 和 sqalx 处理数据库访问
> github.com/jmoiron/sqlx
>
> github.com/heetch/sqalx

* 采用 jennifer 作为代码生成器，数据访问层可以统一通过代码生成器生成。
> github.com/dave/jennifer

* 应用程序监控采用 Prometheus
> github.com/prometheus/prometheus

* 配置管理采用 viper
> github.com/spf13/viper
