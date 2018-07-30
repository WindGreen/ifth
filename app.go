package ifth

import (
	"github.com/satori/go.uuid"
	"log"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type App struct {
	Name string `json:"name"`
	Token string `json:"token"`
}

func NewApp(name string)*App{
	app:=App{
		Name:name,
	}
	id,err:=uuid.NewV4()
	if err!=nil{
		log.Fatal(err)
	}
	app.Token=id.String()
	return &app
}

func (a *App)Save()error{
	c,err:=Client.Connect()
	if err!=nil{
	return err
	}
	client:=c.(*mgo.Session)
	_,err=client.DB("tuvstu").C("app").Upsert(bson.M{"name":a.Name},a)
	return err
}