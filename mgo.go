package ifth

import (
	"github.com/globalsign/mgo"
	"github.com/pkg/errors"
)

type DB interface {
	Dial(dsn string) error
	Connect() (interface{}, error)
}

type Mgo struct {
	Session *mgo.Session
}

var Client DB

func InitMgo(dsn string) (DB, error) {
	Client = &Mgo{}
	err := Client.Dial(dsn)
	if err != nil {
		return nil, err
	}
	return Client, nil
}

func GetMgo() *mgo.Session {
	session, err := Client.Connect()
	if err != nil {
		return nil
	}
	return session.(*mgo.Session)
}

func (m *Mgo) Dial(dsn string) error {
	session, err := mgo.Dial(dsn)
	if err != nil {
		return err
	}
	m.Session = session
	return nil
}

func (m *Mgo) Connect() (interface{}, error) {
	if m.Session == nil {
		return nil, errors.New("doesn't connect to mongo db")
	}
	session := m.Session.Copy()
	return session, nil
}
