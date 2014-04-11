package data

import (
	"fmt"
	"log"
	"os"

	"github.com/aminjam/onemember"
	"github.com/couchbaselabs/go-couchbase"
)

var CouchbaseDB onemember.DataConnector
var bucket *couchbase.Bucket

type couchbaseDBWrite struct{}

func init() {
	CouchbaseDB = &couchbaseDBWrite{}
}

func (couchbaseDB *couchbaseDBWrite) OnememberCreate(in *onemember.Account) error {
	key := "account:" + in.Username
	err := bucket.Set(key, 0, in)
	if err != nil {
		log.Println("Write returned error %v", err)
		return err
	}
	return nil
}

func (couchbaseDB *couchbaseDBWrite) OnememberRead(uid string) (map[string]interface{}, error) {
	out := map[string]interface{}{}
	key := "account:" + uid
	err := bucket.Get(key, &out)
	if err != nil {
		log.Println("Read returned error %v", err)
		return nil, err
	}
	return out, nil
}

func (couchbaseDB *couchbaseDBWrite) OnememberUpdate(in *onemember.Account) error {
	key := "account:" + in.Username
	err := bucket.Set(key, 0, in)
	if err != nil {
		log.Println("Update returned err %v", err)
		return err
	}
	return nil
}

func SetCouchbaseDB() {
	server := os.Getenv("ONEMEMBER_DBURL")
	if server == "" {
		server = "http://127.0.0.1:8091"
	}
	c, err := couchbase.Connect(server)
	if err != nil {
		log.Fatalf("Error connecting:  %v", err)
	}
	fmt.Printf("Connected to ver=%s\n", c.Info.ImplementationVersion)

	poolName := os.Getenv("ONEMEMBER_DBPOOL")
	if poolName == "" {
		poolName = "default"
	}
	pool, err := c.GetPool(poolName)
	if err != nil {
		log.Fatalf("Can't get pool %q:  %v", poolName, err)
	}
	bucketName := os.Getenv("ONEMEMBER_DBNAME")
	if bucketName == "" {
		bucketName = "default"
	}
	bucket, err = pool.GetBucket(bucketName)
	if err != nil {
		log.Fatalf("Can't get bucket %q:  %v", bucketName, err)
	}
}
