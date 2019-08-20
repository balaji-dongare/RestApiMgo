package mongocon

import (
	"gopkg.in/mgo.v2"
)

//Connect it connect to MongoDB
func Connect(uri string) (session *mgo.Session) {
	session, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
