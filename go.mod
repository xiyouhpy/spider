module github.com/xiyouhpy/spider

go 1.16

replace github.com/xiyouhpy/spider => ./

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/go-redis/redis/v8 v8.10.0
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.4.0
)
