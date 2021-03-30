package user

import (
	"CloudGameServer/db"
	bilicoin "CloudGameServer/utils"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Uid      string             `json:"uid"`
	Username string             `json:"username" valid:"length(3|16)"`
	Mobile   string             `json:"mobile" valid:"required,numeric,length(11|11)"`
	Avatar   string             `json:"avatar"`
	Password string             `json:"password,omitempty" valid:"required,ascii,length(5|16)"`
	Status   int                `json:"status"`
	Role     int                `json:"role,omitempty"`
}

//type UInfo struct {
//	Mobile   string `json:"mobile" valid:"alphanum,length(5|18)"`
//	Password string `json:"password" valid:"ascii,length(3|16)"`
//	// CSRFImpl
//}

func (u *User) Save() error {
	u.Id = primitive.NewObjectID()
	return db.MDB().InsertOne(u)
}

func (u *User) Update() error {
	return db.MDB().UpdateOne(u.Id, u)
}

func (u *User) Delete() error {
	return db.MDB().DeleteOne(u)
}

// exist
func (u *User) IsExist() bool {
	var users []User
	err := db.MDB().FindAll(bson.M{"mobile": u.Mobile}, &users)
	if users == nil || err != nil && err.Error() == "not found" {
		return false
	}
	return true
}

// validate
func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	return err
}

// check token
func CheckToken(token string) bool {
	result, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return bilicoin.GetConfig().GetJwtSecret(), nil
	})
	if err != nil {
		if err.Error() == "Token is expired" && result != nil {
			bilicoin.Info("token is expired", logrus.Fields{"uid": result.Claims.(jwt.MapClaims)["jti"]})
		} else {
			bilicoin.Warn("parse with claims failed")
		}
		return false
	}
	return true
}

// create token
func (u *User) CreateToken() string {
	claims := &jwt.StandardClaims{
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * bilicoin.GetConfig().GetJwtTimeout()).Unix(),
		Issuer:    "r3inb",
		Id:        u.Uid,
	}
	bilicoin.Info("create token", logrus.Fields{"uid": u.Uid})
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(bilicoin.GetConfig().GetJwtSecret())
	if err != nil {
		return ""
	}
	return ss
}

func (u *User) CreateUser() *User {
	u.Uid = bilicoin.CreateMD5(u.Mobile)
	u.Username = u.Mobile
	u.Password = bilicoin.CreateMD5(u.Password)
	u.Status = 0
	u.Role = 0
	return u
}
