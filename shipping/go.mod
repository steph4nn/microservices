module github.com/steph4nn/microservices/shipping

go 1.22

require (
	github.com/sirupsen/logrus v1.9.3
	github.com/steph4nn/microservices-proto/golang/shipping v0.0.0
	google.golang.org/grpc v1.70.0
	google.golang.org/protobuf v1.36.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
)

replace github.com/steph4nn/microservices-proto/golang/shipping => ../../microservices-proto/golang/shipping
