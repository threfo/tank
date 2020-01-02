package rest

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	//guest
	USER_ROLE_GUEST = "GUEST"
	//normal user
	USER_ROLE_USER = "USER"
	//administrator
	USER_ROLE_ADMINISTRATOR = "ADMINISTRATOR"
)

const (
	//ok
	USER_STATUS_OK = "OK"
	//disabled
	USER_STATUS_DISABLED = "DISABLED"
)

const (
	//username pattern
	USERNAME_PATTERN = `^[0-9a-z_]+$`
	USERNAME_DEMO    = "demo"
)

type User struct {
	Username   string             `json:"username" bson:"name"`
	Password   string             `json:"-" bson:"password"`
	Mobile     string             `json:"mobile" bson:"mobile"`
	Id_        primitive.ObjectID `json:"id" bson:"_id"`
	UpdateTime time.Time          `json:"updateTime" bson:"_updated"`
	CreateTime time.Time          `json:"createTime" bson:"_created"`

	// cloud
	Uuid           string    `json:"uuid" bson:"uuid"`
	Sort           int64     `json:"sort" bson:"sort"`
	Role           string    `json:"role" bson:"role"`
	AvatarUrl      string    `json:"avatarUrl" bson:"avatar_url"`
	LastIp         string    `json:"lastIp" bson:"last_ip"`
	LastTime       time.Time `json:"lastTime" bson:"last_time"`
	SizeLimit      int64     `json:"sizeLimit" bson:"size_limit"`
	TotalSizeLimit int64     `json:"totalSizeLimit" bson:"total_size_limit"`
	TotalSize      int64     `json:"totalSize" bson:"total_size"`
	Status         string    `json:"status" bson:"status"`
}

// set User's table name to be `profiles`
func (this *User) TableName() string {
	return "user"
}
