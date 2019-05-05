module github.com/magodo/shippy-service/consignment

go 1.12

require (
	github.com/EwanValentine/shippy-service-consignment v0.0.0-20190416075237-11e19626e7a1
	github.com/golang/protobuf v1.3.1
	github.com/magodo/shippy-service/user v0.0.0-20190505030531-49ad0199d0c6
	github.com/magodo/shippy-service/vessel v0.0.0-20190426032437-bcd5260b2215
	github.com/micro/cli v0.1.0
	github.com/micro/go-grpc v1.0.1
	github.com/micro/go-micro v1.1.0
	go.mongodb.org/mongo-driver v1.0.1
)

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f

replace github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1
