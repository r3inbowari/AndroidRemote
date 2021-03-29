package user

import (
	"CloudGameServer/db"
	bilicoin "CloudGameServer/utils"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Username string             `json:"username" valid:"required,length(3|16)"`
	Mobile   string             `json:"mobile"`
	Avatar   string             `json:"avatar"`
	Password string             `json:"password,omitempty" valid:"ascii,length(3|16)"`
	Status   int                `json:"status"`
}

func (u *User) Save() error {
	u.Id = primitive.NewObjectID()
	return db.MDB().InsertOne(u)
}

func (u *User) Update() error {
	return db.MDB().UpdateOne(u.Id, u)
}

// 查找用户
//func (u *User) FoundUser() bool {
//	t := User{}
//	err := db.MDB().C("User").Find(bson.M{"mobile": u.Mobile}).One(&t)
//	if err != nil && err.Error() == "not found" {
//		return false
//	}
//	return true
//}

// 格式校验
func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	return err
}

// token 检查
func CheckToken(token string) bool {
	result, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return bilicoin.GetConfig().JwtSecret, nil
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

// 创建token
func (u *User) CreateToken(uid string) string {
	claims := &jwt.StandardClaims{
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    "r3inb",
		Id:        uid,
	}
	bilicoin.Info("create token "+uid, logrus.Fields{"uid": uid})
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(bilicoin.GetConfig().JwtSecret)
	if err != nil {
		println(err)
		return ""
	}
	return ss
}
