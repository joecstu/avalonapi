package data

import (
	"avalonapi/model"
	"github.com/globalsign/mgo"
)

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

func  CreateUser(user *model.UserRegis)  error {
	user.Nickname=""
	user.Status="offline"
	session, err := mgo.Dial("mongodb://test:test1234@ds026558.mlab.com:26558/avalon")
	defer session.Close()
	c := session.DB("avalon").C("user")

	err = c.Insert(user)
	return  err
}