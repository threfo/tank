package rest

import (
	"context"
	"time"

	"github.com/eyebluecn/tank/code/core"
	"github.com/eyebluecn/tank/code/tool/result"
	uuid "github.com/nu7hatch/gouuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionDao struct {
	BaseDao
}

//find by uuid. if not found return nil.
func (this *SessionDao) FindByUuid(uuid string) *Session {
	var UserColl *mongo.Collection
	var entity = &Session{}

	UserColl = core.CONTEXT.GetMDB().Collection("session")
	filter := bson.M{"uuid": uuid}
	result := UserColl.FindOne(context.Background(), filter)
	err := result.Decode(&entity)
	if err != nil {
		panic(err)
	}
	return entity
}

//find by uuid. if not found panic NotFound error
func (this *SessionDao) CheckByUuid(uuid string) *Session {
	entity := this.FindByUuid(uuid)
	if entity == nil {
		panic(result.NotFound("not found record with uuid = %s", uuid))
	}
	return entity
}

func (this *SessionDao) Create(session *Session) *Session {
	var SessionColl *mongo.Collection

	timeUUID, _ := uuid.NewV4()
	session.Uuid = string(timeUUID.String())
	session.CreateTime = time.Now()
	session.UpdateTime = time.Now()
	session.Sort = time.Now().UnixNano() / 1e6
	session.Id_ = primitive.NewObjectID()

	SessionColl = core.CONTEXT.GetMDB().Collection("session")
	_, err := SessionColl.InsertOne(context.Background(), session)
	if err != nil {
		this.logger.Error(err.Error())
	}
	return session
}

func (this *SessionDao) Save(session *Session) *Session {
	var SessionColl *mongo.Collection

	session.UpdateTime = time.Now()
	filter := bson.M{"_id": session.Id_}
	update := bson.M{"$set": session}

	SessionColl = core.CONTEXT.GetMDB().Collection("session")
	_, err := SessionColl.UpdateOne(context.Background(), filter, update)
	if err != nil {
		this.logger.Error(err.Error())
	}
	return session
}

func (this *SessionDao) Delete(uuid string) {
	var SessionColl *mongo.Collection

	SessionColl = core.CONTEXT.GetMDB().Collection("session")

	session := this.CheckByUuid(uuid)
	filter := bson.M{"uuid": session.Uuid}

	_, err := SessionColl.DeleteOne(context.Background(), filter)
	if err != nil {
		this.logger.Error(err.Error())
	}
}

//System cleanup.
func (this *SessionDao) Cleanup() {
	this.logger.Info("[SessionDao] clean up. Delete all Session")
	var SessionColl *mongo.Collection

	SessionColl = core.CONTEXT.GetMDB().Collection("session")
	filter := bson.M{}
	_, err := SessionColl.DeleteMany(context.Background(), filter)
	if err != nil {
		this.logger.Error(err.Error())
	}
}
