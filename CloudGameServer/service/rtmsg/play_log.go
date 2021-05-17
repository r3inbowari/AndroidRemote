package rtmsg

import "go.mongodb.org/mongo-driver/bson/primitive"

type PlayLog struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
	StartTime int64              `json:"st"`
	Duration  int64              `json:"duration"`
	AID       string             `json:"aid"`
	Title     string             `json:"title"`
	UID       string             `json:"uid"`
	Mobile    string             `json:"mobile"`
	DeviceID  string             `json:"device_id"`
}
