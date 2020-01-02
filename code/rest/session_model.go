package rest

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	Id_        primitive.ObjectID    `json:"id" bson:"_id"`
	Uuid       string    `json:"Uuid" bson:"uuid"`
	UserUuid   string    `json:"userUuid" bson:"user_uuid"`
	UpdateTime time.Time `json:"updateTime" bson:"_updated"`
	CreateTime time.Time `json:"createTime" bson:"_created"`
	Sort       int64     `json:"sort" bson:"sort"`
	UserId 	   primitive.ObjectID    `json:"userId" bson:"user_id"`
	ManagerId  primitive.ObjectID	 `josn:"managerId" bson:"manager_id"`
	Token      string	`json:"token" bson:"token"`
	UserType   int    `json:"userType" bson:"user_type"`
}

// set Session's table name to be `session`
func (this *Session) TableName() string {
	return "session"
}
