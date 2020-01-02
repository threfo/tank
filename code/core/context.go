package core

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
)

type Context interface {
	http.Handler

	//get the gorm.DB. all the db connection will use this
	GetDB() *gorm.DB

	//get the mongo.Db. all the mongodb connection will use this
	GetMDB() *mongo.Database

	GetBean(bean Bean) Bean

	//get the global session cache
	// GetSessionCache() *cache.Table

	GetControllerMap() map[string]Controller

	//when application installed. this method will invoke every bean's Bootstrap method
	InstallOk()

	//this method will invoke every bean's Cleanup method
	Cleanup()
}
