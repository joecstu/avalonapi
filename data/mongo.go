package data

import (
	"avalonapi/model"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"fmt"
)
type Test struct {
	Id string `bson:"_id,omitempty"`
}

/*
session := mongo.getSession()
defer session.Close()
c := session.DB(mongo.Config.MongoDb).C(USERS)

user.Id = bson.NewObjectId()
err := c.Insert(user)
return user, err


	session.SetMode(mgo.Monotonic, true)

	c := session.DB("pokemondb").C("pokemons")
	pokemons := []Pokemon{}
	err = c.Find(nil).All(&pokemons)
	if err != nil {
		log.Fatal(err)
	}

 */

func  CreateUser(user *model.User)  error {
	user.Nickname=""
	user.Status="offline"
	session, err := mgo.Dial("mongodb://test:test1234@ds026558.mlab.com:26558/avalon")
	defer session.Close()
	c := session.DB("avalon").C("user")

	err = c.Insert(user)
	return  err
}
func  Login(user *model.User)  (error,string,string) {
	session, err := mgo.Dial("mongodb://test:test1234@ds026558.mlab.com:26558/avalon")
	defer session.Close()
	c := session.DB("avalon").C("user")

	result := model.User{}

	//err = c.Insert(user)
	err = c.Find(bson.M{"email": user.Email, "password":user.Password}).One(&result)
	if err != nil {
		return err,"",""
	}

	//err = coll.Update(bson.M{"_id": person.ID}, bson.M{"$set": bson.M{"ph": "+99 99 9999 1111"}})
	err = c.Update(bson.M{"email":result.Email},bson.M{"$set":bson.M{"status":"online"}})
	c = session.DB("avalon").C("session")
	key := bson.NewObjectId()
	err = c.Insert(bson.M{"_id":key,"email":result.Email})
	fmt.Println(key.Hex())
	//test := bson.ObjectId(key.Hex())
	return  err,result.Nickname,key.Hex()
}
func  Logout(key string) error {
	session, err := mgo.Dial("mongodb://test:test1234@ds026558.mlab.com:26558/avalon")
	defer session.Close()
	c := session.DB("avalon").C("session")
	search := model.Session{}
	err = c.FindId(bson.ObjectIdHex(key)).One(&search)
	err = c.RemoveId(bson.ObjectIdHex(key))
	c = session.DB("avalon").C("user")
	err = c.Update(bson.M{"email":search.Email},bson.M{"$set":bson.M{"status":"offline"}})

	//fmt.Println(key.Hex())
	//test := bson.ObjectId(key.Hex())
	return err
}