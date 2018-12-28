package data

import (
	"avalonapi/model"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Test struct {
	Id string `bson:"_id,omitempty"`
}

const DATABASE = "mongodb://test:test1234@ds026558.mlab.com:26558/avalon"

func CreateUser(user *model.User) error {
	user.Nickname = ""
	user.Status = "offline"
	session, err := mgo.Dial(DATABASE)
	defer session.Close()
	c := session.DB("avalon").C("user")

	err = c.Insert(user)
	return err
}

func Login(user *model.User) (error, string, string) {
	session, err := mgo.Dial(DATABASE)
	defer session.Close()
	c := session.DB("avalon").C("user")
	result := model.User{}
	err = c.Find(bson.M{"email": user.Email, "password": user.Password}).One(&result)
	if err != nil {
		return err, "", ""
	}
	err = c.Update(bson.M{"email": result.Email}, bson.M{"$set": bson.M{"status": "online"}})
	c = session.DB("avalon").C("session")
	key := bson.NewObjectId()
	err = c.Insert(bson.M{"_id": key, "email": result.Email})
	fmt.Println(key.Hex())
	return err, result.Nickname, key.Hex()
}

func Logout(key string) error {
	session, err := mgo.Dial(DATABASE)
	defer session.Close()
	c := session.DB("avalon").C("session")
	search := model.Session{}
	err = c.FindId(bson.ObjectIdHex(key)).One(&search)
	err = c.RemoveId(bson.ObjectIdHex(key))

	c = session.DB("avalon").C("user")
	err = c.Update(bson.M{"email": search.Email}, bson.M{"$set": bson.M{"status": "offline"}})
	return err
}
func CreateNickname(nickname string,key string) (error, int) {
	session, err := mgo.Dial(DATABASE)
	defer session.Close()
	c := session.DB("avalon").C("session")
	search := model.Session{}
	err = c.FindId(bson.ObjectIdHex(key)).One(&search)

	usersearch := []model.User{}
	c = session.DB("avalon").C("user")
	err = c.Find(bson.M{"nickname":nickname}).All(&usersearch)
	if(len(usersearch)==1){
		return err,0
	}else{
		err = c.Update(bson.M{"email": search.Email}, bson.M{"$set": bson.M{"nickname": nickname}})
		return err,1
	}

}
