package lib

import (
//	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"net/http"
	)
var (
    clt *http.Client
	Fullsession *mgo.Session
	Readsession *mgo.Session
	FullDatabase *mgo.Database
	ReadDatabase *mgo.Database
)


func ConnectDB(conf *GConfig) {
	var err error
	
	domainAndPort := conf.DbSettings.Domain + ":" + conf.DbSettings.Port

	Fullsession, err = mgo.Dial(domainAndPort)
	check(err)
	Fullsession.SetSafe(&mgo.Safe{})
	Fullsession.SetMode(mgo.Monotonic, true)
	FullDatabase = Fullsession.DB(conf.DbSettings.Database)
	err = FullDatabase.Login(conf.DbSettings.Username, conf.DbSettings.Password)
	check(err)
}

func CloseDB() {
	FullDatabase.Logout()
	Fullsession.Close()
}
