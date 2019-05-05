module github.com/magodo/shippy-service/email

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.3.1
	github.com/jinzhu/gorm v1.9.5
	github.com/magodo/shippy-service/user v0.0.0-20190505033607-d682d780bcc0
	github.com/micro/cli v0.1.0
	github.com/micro/go-micro v1.1.0
	github.com/micro/go-plugins v1.1.0
	github.com/satori/go.uuid v1.2.0
	go.mongodb.org/mongo-driver v1.0.1
	golang.org/x/crypto v0.0.0-20190404164418-38d8ce5564a5
)

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f

replace github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1
