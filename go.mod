module github.com/yuexclusive/utils

go 1.16

require (
	github.com/ahmetb/go-linq v3.0.0+incompatible
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gin-contrib/zap v0.0.1
	github.com/gin-gonic/gin v1.7.1
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/validator/v10 v10.6.1 // indirect
	github.com/go-redis/redis v6.15.8+incompatible
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.1.2
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/jinzhu/gorm v1.9.12
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/juju/ratelimit v1.0.1
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/magefile/mage v1.11.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/micro/go-micro/v2 v2.8.0
	github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2 v2.8.0
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.8.0
	github.com/micro/micro/v3 v3.0.1 // indirect
	github.com/miekg/dns v1.1.29 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/nats-io/nats.go v1.11.0
	github.com/olivere/elastic/v7 v7.0.16
	github.com/opentracing/opentracing-go v1.1.0
	github.com/prometheus/client_golang v1.6.0
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.1
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/uber/jaeger-client-go v2.23.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	github.com/ugorji/go v1.2.5 // indirect
	go.elastic.co/ecszap v1.0.0
	go.etcd.io/etcd v3.3.22+incompatible
	go.etcd.io/etcd/api/v3 v3.5.0-alpha.0
	go.etcd.io/etcd/client/v3 v3.5.0-alpha.0
	go.mongodb.org/mongo-driver v1.5.4
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sys v0.0.0-20210511113859-b0526f3d8744 // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/tools v0.1.1 // indirect
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/sohlich/elogrus.v7 v7.0.0
	gopkg.in/yaml.v2 v2.4.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.37.0
