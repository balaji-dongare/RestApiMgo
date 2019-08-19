package mongocon

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

//Connect it connect to MongoDB
func Connect(uri string) (session *mgo.Session) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Printf("Error:%T", err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
