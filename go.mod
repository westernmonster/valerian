module valerian

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/ajg/form v1.5.1 // indirect
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190418113227-25233c783f4e
	github.com/asaskevich/govalidator v0.0.0-20180720115003-f9ffefc3facf
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/bouk/monkey v1.0.1 // indirect
	github.com/bwmarrin/snowflake v0.3.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dgryski/go-farm v0.0.0-20190423205320-6a90982ecee2
	github.com/fatih/structs v1.1.0 // indirect
	github.com/gavv/httpexpect v1.0.0 // indirect
	github.com/gavv/monotime v0.0.0-20190418164738-30dba4353424 // indirect
	github.com/go-ozzo/ozzo-validation v3.5.0+incompatible
	github.com/go-playground/locales v0.12.1 // indirect
	github.com/go-playground/universal-translator v0.16.0 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gobuffalo/packr v1.25.0
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.3.1
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.0.0 // indirect
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/jinzhu/copier v0.0.0-20180308034124-7e38e58719c3
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/joho/godotenv v1.3.0
	github.com/json-iterator/go v1.1.6
	github.com/leodido/go-urn v1.1.0 // indirect
	github.com/lib/pq v1.0.0
	github.com/mattn/go-sqlite3 v1.9.0
	github.com/mkideal/pkg v0.0.0-20170503154153-3e188c9e7ecc // indirect
	github.com/montanaflynn/stats v0.5.0
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/nicksnyder/go-i18n/v2 v2.0.0-beta.6
	github.com/opentracing-contrib/go-gin v0.0.0-20190301172248-2e18f8b9c7d4 // indirect
	github.com/opentracing-contrib/go-stdlib v0.0.0-20190324214902-3020fec0e66b
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v0.9.2
	github.com/samuel/go-zookeeper v0.0.0-20180130194729-c4fab1ac1bec
	github.com/satori/go.uuid v1.2.0
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/sirupsen/logrus v1.4.1
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a
	github.com/spf13/viper v1.3.2 // indirect
	github.com/stretchr/testify v1.3.0
	github.com/swaggo/swag v1.5.0
	github.com/tidwall/btree v0.0.0-20170113224114-9876f1454cf0 // indirect
	github.com/tidwall/buntdb v1.1.0 // indirect
	github.com/tidwall/gjson v1.2.1 // indirect
	github.com/tidwall/grect v0.0.0-20161006141115-ba9a043346eb // indirect
	github.com/tidwall/match v1.0.1 // indirect
	github.com/tidwall/pretty v0.0.0-20190325153808-1166b9ac2b65 // indirect
	github.com/tidwall/rtree v0.0.0-20180113144539-6cd427091e0e // indirect
	github.com/tidwall/tinyqueue v0.0.0-20180302190814-1e39f5511563 // indirect
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/uber/jaeger-lib v2.0.0+incompatible
	github.com/valyala/quicktemplate v1.0.2
	github.com/westernmonster/sqalx v0.3.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.1.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/ztrue/tracerr v0.3.0
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/net v0.0.0-20190322120337-addf6b3196f6
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/text v0.3.0
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
	google.golang.org/appengine v1.5.0 // indirect
	google.golang.org/grpc v1.22.0
	gopkg.in/eapache/go-resiliency.v1 v1.2.0 // indirect
	gopkg.in/go-playground/validator.v9 v9.28.0
	gopkg.in/h2non/gentleman-retry.v2 v2.0.1
	gopkg.in/h2non/gentleman.v2 v2.0.3
	gopkg.in/ini.v1 v1.42.0 // indirect
	gopkg.in/yaml.v2 v2.2.2
)
