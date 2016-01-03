package models

import (
	"encoding/json"
	"io/ioutil"

	mgo "gopkg.in/mgo.v2"
)

type database struct {
	host     string
	database string
	username string
	password string
	port     string
}

var databaseConfig database

func Init() (*mgo.Session, error) {
	file, err := ioutil.ReadFile("../config/database.json")

	if err != nil {
		panic(err)
	}

	var databaseConfig database

	err = json.Unmarshal(file, databaseConfig)

	session, err := mgo.Dial(databaseConfig.host)

	if err != nil {
		panic(err)
	}

	session.DB(databaseConfig.database)

	return session, nil
}
