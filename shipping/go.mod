module github.com/steph4nn/microservices/shipping

go 1.25.2

require (
	github.com/sirupsen/logrus v1.9.3
	github.com/steph4nn/microservices-proto/golang/shipping v0.0.0
	google.golang.org/grpc v1.78.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251029180050-ab9386a59fda // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace github.com/steph4nn/microservices-proto/golang/shipping => ../../microservices-proto/golang/shipping
