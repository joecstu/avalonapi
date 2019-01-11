package data

import (
	"crypto/tls"
	"net"

	"github.com/globalsign/mgo"
)

const DATABASE = "mongodb://test:test1234@cluster0-shard-00-00-ikbru.gcp.mongodb.net:27017,cluster0-shard-00-01-ikbru.gcp.mongodb.net:27017,cluster0-shard-00-02-ikbru.gcp.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin"

//CreateSessiion สร้างการเชื่อมต่อเป็นท่อให้
func CreateSession() (*mgo.Session, error) {
	var mongoURI = DATABASE
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	dialInfo, err := mgo.ParseURL(mongoURI)
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		return tls.Dial("tcp", addr.String(), tlsConfig)
	}
	session, err := mgo.DialWithInfo(dialInfo)
	return session, err
}
