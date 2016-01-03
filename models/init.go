package models

import (
	"encoding/json"
	"io/ioutil"

	mgo "gopkg.in/mgo.v2"
)

type Models struct {
		Connect *mgo.Session
}

type database struct {
		host     string
		database string
		username string
		password string
		port     string
}

var databaseConfig database

func (m *Models) Init() {
		file, err := ioutil.ReadFile("../config/database.json")

		if err != nil {
			panic(err)
		}

		var databaseConfig database

		err = json.Unmarshal(file, databaseConfig)

		session, err := mgo.Dial(databaseConfig.host)

		defer session.Close()

		if err != nil {
			panic(err)
		}

		m.Connect = session.DB(databaseConfig.database)
}
