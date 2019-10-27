module valerian

go 1.12

replace (
	git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
)

require (
	github.com/Azure/azure-sdk-for-go v21.1.0+incompatible
	github.com/BurntSushi/toml v0.3.1
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190912105415-4fb175c277cc
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a
	github.com/bwmarrin/snowflake v0.3.0
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/dgryski/go-farm v0.0.0-20190423205320-6a90982ecee2
	github.com/fortytw2/leaktest v1.3.0 // indirect
	github.com/go-openapi/validate v0.19.3 // indirect
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/go-playground/locales v0.12.1 // indirect
	github.com/go-playground/universal-translator v0.16.0 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gobuffalo/packr v1.30.1
	github.com/gogo/protobuf v1.3.0
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/copier v0.0.0-20190625015134-976e0346caa8
	github.com/jmoiron/sqlx v1.2.0
	github.com/kamilsk/retry/v4 v4.3.1
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/leodido/go-urn v1.1.0 // indirect
	github.com/levigross/grequests v0.0.0-20190908174114-253788527a1a
	github.com/mattn/go-sqlite3 v1.11.0 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/montanaflynn/stats v0.5.0
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/naoina/toml v0.1.1
	github.com/nats-io/nats-server/v2 v2.1.0 // indirect
	github.com/nats-io/nats-streaming-server v0.16.2 // indirect
	github.com/nats-io/stan.go v0.5.0
	github.com/nicksnyder/go-i18n/v2 v2.0.2
	github.com/olivere/elastic v6.2.23+incompatible
	github.com/opentracing-contrib/go-stdlib v0.0.0-20190519235532-cf7a6c988dc9
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.1.0
	github.com/samuel/go-zookeeper v0.0.0-20190810000440-0ceca61e4d75
	github.com/satori/go.uuid v1.2.0
	github.com/sergi/go-diff v1.0.0
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337
	github.com/sony/gobreaker v0.4.1
	github.com/spf13/pflag v1.0.3
	github.com/stretchr/testify v1.4.0
	github.com/swaggo/swag v1.6.2
	github.com/uber-go/atomic v1.4.0 // indirect
	github.com/uber/jaeger-client-go v2.17.0+incompatible
	github.com/uber/jaeger-lib v2.1.1+incompatible
	github.com/valyala/quicktemplate v1.2.0
	github.com/ztrue/tracerr v0.3.0
	go.uber.org/zap v1.10.0
	golang.org/x/net v0.0.0-20190912160710-24e19bdeb0f2
	golang.org/x/text v0.3.2
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
	google.golang.org/grpc v1.23.1
	gopkg.in/eapache/go-resiliency.v1 v1.2.0 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1
	gopkg.in/h2non/gentleman-retry.v2 v2.0.1
	gopkg.in/h2non/gentleman.v2 v2.0.3
	gopkg.in/h2non/gock.v1 v1.0.15
	gopkg.in/olivere/elastic.v6 v6.2.25
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/test-infra v0.0.0-20190914015041-e1cbc3ccd91c // indirect
)
