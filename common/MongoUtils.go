package common

import (
	"log"
	"time"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{mgoConfig.DBHost},
			Username: mgoConfig.DBUser,
			Password: mgoConfig.DBPwd,
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}
	session.DB(mgoConfig.Database)
	return session
}
func createDbSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{mgoConfig.DBHost},
		Username: mgoConfig.DBUser,
		Password: mgoConfig.DBPwd,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}

}

// Add indexes into MongoDB
func addIndexes() {
	var err error
	passengerIndex := mgo.Index{
		Key:        []string{"passenger_id"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
	//
	pathOrderIndex := mgo.Index{
		Key:        []string{"order_id"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	// Add indexes into MongoDB
	session := GetSession().Copy()
	defer session.Close()
	passenger_lbs := session.DB(mgoConfig.Database).C("passenger_lbs")
	order_path := session.DB(mgoConfig.Database).C("order_path")

	err = passenger_lbs.EnsureIndex(passengerIndex)

	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
	err = order_path.EnsureIndex(pathOrderIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
}

type MongoAppConfig struct {
	DBHost   string
	DBUser   string
	DBPwd    string
	Database string
}

var mgoConfig = MongoAppConfig{}

func (m *MongoAppConfig) initConfig() {
	m.DBHost = GetCfg().Datasource["mgo"]["host"]
	m.DBUser = GetCfg().Datasource["mgo"]["user"]
	m.DBPwd = GetCfg().Datasource["mgo"]["pass"]
	m.Database = GetCfg().Datasource["mgo"]["database"]
}
func initMgo() {
	mgoConfig.initConfig()
	createDbSession()
	addIndexes()
}
