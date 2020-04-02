module github.com/zhangchong5566/manba

go 1.13

require (
	code.yunzhanghu.com/be/yos v0.0.0-20200116063411-3275fcd3b92f
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/buger/jsonparser v0.0.0-20180318095312-2cac668e8456
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fagongzi/goetty v0.0.0-20180427060148-8f06d410550f
	github.com/fagongzi/log v0.0.0-20170831135209-9a647df25e0e
	github.com/fagongzi/util v0.0.0-20180330021808-4acf02da76a9
	github.com/fullstorydev/grpcurl v1.1.0
	github.com/garyburd/redigo v0.0.0-20180228092057-a69d19351219
	github.com/gogo/protobuf v1.2.0
	github.com/golang/protobuf v1.3.2
	github.com/gorilla/websocket v1.4.0
	github.com/jhump/protoreflect v1.1.0
	github.com/juju/ratelimit v1.0.1
	github.com/koding/websocketproxy v0.0.0-20180716164433-0fa3f994f6e7
	github.com/labstack/echo v0.0.0-20180412143600-6d227dfea4d2
	github.com/labstack/gommon v0.0.0-20180613044413-d6898124de91 // indirect
	github.com/prometheus/client_golang v0.9.3-0.20190127221311-3c4408c8b829
	github.com/prometheus/common v0.2.0
	github.com/robertkrimen/otto v0.0.0-20180617131154-15f95af6e78d
	github.com/satori/go.uuid v1.2.0
	github.com/soheilhy/cmux v0.1.4
	github.com/stretchr/testify v1.3.0
	github.com/ugorji/go/codec v0.0.0-20181209151446-772ced7fd4c2 // indirect
	github.com/valyala/fasthttp v1.2.0
	github.com/valyala/fastrand v1.0.0
	github.com/valyala/fasttemplate v0.0.0-20170224212429-dcecefd839c4 // indirect
	github.com/yuin/goldmark v1.1.26 // indirect
	go.etcd.io/etcd v3.3.12+incompatible
	go.uber.org/atomic v1.3.2 // indirect
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59 // indirect
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
	golang.org/x/sys v0.0.0-20200327173247-9dae0f8f5775 // indirect
	golang.org/x/tools v0.0.0-20200330040139-fa3cc9eebcfe // indirect
	google.golang.org/grpc v1.24.0
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
)

replace (
	github.com/YunzhanghuOpen/glog => github.com/zhanghongnian/glog v0.0.0-20191006035400-c890ddc19451
	github.com/ugorji/go/codec => github.com/ugorji/go v1.1.2
	go.etcd.io/etcd => github.com/YunzhanghuOpen/etcd v0.0.0-20190530103243-54cdb2605d64
	gopkg.in/DATA-DOG/go-sqlmock.v1 => github.com/DATA-DOG/go-sqlmock v0.0.0-20190621103928-e98392b8111b
)
