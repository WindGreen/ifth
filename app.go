package ifth

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/satori/go.uuid"
)

type App struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

func NewApp(name string) *App {
	app := App{
		Name: name,
	}
	id := uuid.NewV4()
	app.Token = id.String()
	return &app
}

func (a *App) Save() error {
	c, err := Client.Connect()
	if err != nil {
		return err
	}
	client := c.(*mgo.Session)
	_, err = client.DB("tuvstu").C("app").Upsert(bson.M{"name": a.Name}, a)
	return err
}
