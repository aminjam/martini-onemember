package data

import (
	"log"
	"os"

	"github.com/aminjam/onemember"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var MongoDB onemember.DataConnector

type mongoDBWrite struct{}

func init() {
	MongoDB = &mongoDBWrite{}
}

func (mongoDB *mongoDBWrite) OnememberCreate(in *onemember.Account) error {
	session,collection,err := mgoInit()
	if err != nil {
		return err
	}
	defer session.Close()
	err = collection.Insert(in)
	if err != nil {
		log.Println("Write returned error %v", err)
		return err
	}
	return nil
}

func (mongoDB *mongoDBWrite) OnememberRead(username string) (map[string]interface{}, error) {
	session,collection,err := mgoInit()
	if err != nil {
		return nil,err
	}
	defer session.Close()
	out := map[string]interface{}{}
	err = collection.Find(bson.M{"username":username}).One(&out)
	if err != nil {
		log.Println("Read returned error %v", err)
		return nil, err
	}
	return out, nil
}

func (mongoDB *mongoDBWrite) OnememberUpdate(in *onemember.Account) error {
	session,collection,err := mgoInit()
	if err != nil {
		return err
	}
	defer session.Close()
	err =collection.Update(bson.M{"username":in.Username},*in) 
	if err != nil {
		log.Println("Update returned err %v", err)
		return err
	}
	return nil
}

func mgoInit()(session *mgo.Session,collection *mgo.Collection, err error) {
	session, collection = nil, nil
	server := os.Getenv("ONEMEMBER_DBURL")
	if server == "" {
		server = "localhost"
	}
	session, err = mgo.Dial(server)
	if err != nil {
		log.Fatalf("Error connecting:  %v", err)
		return
	}
	dbName := os.Getenv("ONEMEMBER_DBNAME")
	if dbName == "" {
		dbName = "default"
	}
	c := os.Getenv("ONEMEMBER_DBCOLLECTION")
	if c == "" {
		c = "onemember"
	}
	collection = session.DB(dbName).C(c)
	return
}
