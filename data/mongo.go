package data

import (
	"avalonapi/model"

	"github.com/globalsign/mgo/bson"
)

type Test struct {
	Id string `bson:"_id,omitempty"`
}

func CreateUser(user *model.User) error {

	session, err := CreateSession()

	defer session.Close()
	c := session.DB("avalon").C("user")

	err = c.Insert(user)
	return err
}

func Login(user *model.User) (error, string, string) {

	session, err := CreateSession()

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
	//fmt.Println(key.Hex())
	return err, result.Nickname, key.Hex()
}

func Logout(key string) error {
	session, err := CreateSession()
	defer session.Close()
	c := session.DB("avalon").C("session")
	search := model.Session{}
	err = c.FindId(bson.ObjectIdHex(key)).One(&search)
	err = c.RemoveId(bson.ObjectIdHex(key))

	c = session.DB("avalon").C("user")
	err = c.Update(bson.M{"email": search.Email}, bson.M{"$set": bson.M{"status": "offline"}})
	return err
}

func LogoutAll() error {
	session, err := CreateSession()
	defer session.Close()
	c := session.DB("avalon").C("session")
	search := []*model.Session{}
	err = c.Find(nil).All(&search)
	c.RemoveAll(nil)
	c = session.DB("avalon").C("user")
	for i := 0; i < len(search); i++ {
		err = c.Update(bson.M{"email": search[i].Email}, bson.M{"$set": bson.M{"status": "offline"}})
	}
	return err
}

func ChangeNickName(nickname string, key string) (error, int) {
	session, err := CreateSession()
	defer session.Close()
	c := session.DB("avalon").C("session")
	search := model.Session{}
	err = c.FindId(bson.ObjectIdHex(key)).One(&search)

	usersearch := []model.User{}
	c = session.DB("avalon").C("user")
	err = c.Find(bson.M{"nickname": nickname}).All(&usersearch)
	if len(usersearch) == 1 {
		return err, 0
	} else {
		err = c.Update(bson.M{"email": search.Email}, bson.M{"$set": bson.M{"nickname": nickname}})
		return err, 1
	}

}
func Useronline() (error, []model.User) {
	session, err := CreateSession()
	defer session.Close()
	c := session.DB("avalon").C("session")
	search := []model.Session{}
	err = c.Find(nil).All(&search)

	usersearch := model.User{}
	result := []model.User{}

	c = session.DB("avalon").C("user")

	for i := 0; i < len(search); i++ {
		c.Find(bson.M{"email": search[i].Email}).One(&usersearch)
		usersearch.Password = ""
		result = append(result, usersearch)
		if err != nil {
		}
	}

	if err != nil {
		return err, []model.User{}
	} else {
		return err, result
	}

}
func CreateRoom(nickname string, key string) (model.Room, error, int) {
	session, err := CreateSession()
	defer session.Close()
	c := session.DB("avalon").C("session")
	checklogin, err := c.FindId(bson.ObjectIdHex(key)).Count()
	if checklogin != 1 {
		return model.Room{}, err, 0
	}
	c = session.DB("avalon").C("room")
	room := []model.Room{}
	c.Find(nil).All(&room)
	roomnew := model.Room{}
	roomnew.Roomid = (len(room) + 1)
	roomnew.Numberplayer = 1
	roomnew.Maxplayer = 5
	roomnew.Nameplayer = append(roomnew.Nameplayer, nickname)
	roomnew.Host = nickname
	err = c.Insert(roomnew)
	if err != nil {
		return roomnew, nil, 0
	} else {
		return roomnew, nil, 1
	}

}

func Joinroom(nickname string, key string, roomid int) (model.Room, error, int) {
	session, err := CreateSession()
	defer session.Close()
	c := session.DB("avalon").C("session")
	checklogin, err := c.FindId(bson.ObjectIdHex(key)).Count()
	if checklogin != 1 {
		return model.Room{}, err, 0
	}
	c = session.DB("avalon").C("room")
	room := model.Room{}
	c.Find(bson.M{"roomid": roomid}).One(&room)
	room.Nameplayer = append(room.Nameplayer, nickname)

	c.Update(bson.M{"roomid": roomid}, bson.M{"$set": bson.M{"nameplayer": room.Nameplayer}})

	c.Find(bson.M{"roomid": roomid}).One(&room)

	if err != nil {
		return room, nil, 0
	} else {
		return room, nil, 1
	}

}
func Getroom(roomid int) (model.Room, error, int) {
	session, err := CreateSession()
	defer session.Close()
	c := session.DB("avalon").C("room")
	room := model.Room{}
	c.Find(bson.M{"roomid": roomid}).One(&room)

	if err != nil {
		return room, nil, 0
	} else {
		return room, nil, 1
	}

}
func GetAllroom() (error, []model.Room, int) {
	session, err := CreateSession()
	defer session.Close()
	c := session.DB("avalon").C("room")
	room := []model.Room{}
	c.Find(nil).All(&room)

	if err != nil {
		return nil, room, 0
	} else {
		return nil, room, 1
	}

}


func DeleteAllroom() (error, []model.Room, int) {
	session, err := CreateSession()
	defer session.Close()
	c := session.DB("avalon").C("room")
	room := []model.Room{}
	c.RemoveAll(nil)

	if err != nil {
		return nil, room, 0
	} else {
		return nil, room, 1
	}

}
