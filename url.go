package ifth

import (
	"time"
	"github.com/globalsign/mgo/bson"
)

const (
	RedirectProxy = "proxy"
	Redirect301   = "301"
	Redirect302   = "302"
	Redirect303   = "303"
	Redirect307   = "307"
)

type Url struct {
	Slot         string    `json:"slot"`
	Origin       string    `json:"origin"`
	RedirectType string    `json:"redirect_type" bson:"redirect_type"`
	CreatedTime   time.Time `json:"created_time" bson:"created_time"`
	ExpiresIn    int       `json:"expires_in" bson:"expires_in"`
	Count        int       `json:"count"`
}

func NewUrl(origin string, duplicated bool) *Url {
	//check origin exist
	if !duplicated && OriginExist(origin) {
		url, err := FindUrlByOrigin(origin)
		if err != nil {
			return nil
		}
		return url
	}
	u := Url{
		Slot:         SlotGenerator.Get(),
		Origin:       origin,
		RedirectType: Redirect301,
		CreatedTime:   time.Now(),
		ExpiresIn:    0,
		Count:        0,
	}
	for {
		// check slot exist
		if !SlotExist(u.Slot) {
			break
		}
		u.Slot = SlotGenerator.Get()
	}
	return &u
}

func FindUrlByOrigin(origin string) (*Url, error) {
	session := GetMgo()
	var url Url
	err := session.DB("ifth").C("url").Find(bson.M{"origin": origin}).One(&url)
	if err != nil {
		return nil, err
	}
	return &url, nil
}

func FindUrlBySlot(slot string) (*Url, error) {
	session := GetMgo()
	var url Url
	err := session.DB("ifth").C("url").Find(bson.M{"slot": slot}).One(&url)
	if err != nil {
		return nil, err
	}
	return &url, nil
}

func OriginExist(origin string) bool {
	session := GetMgo()
	ct, err := session.DB("ifth").C("url").Find(bson.M{"origin": origin}).Count()
	if err != nil {
		return false //may cause problem?
	}
	if ct > 0 {
		return true
	}
	return false
}

func SlotExist(slot string) bool {
	session := GetMgo()
	ct, err := session.DB("ifth").C("url").Find(bson.M{"slot": slot}).Count()
	if err != nil {
		return false //may cause problem?
	}
	if ct > 0 {
		return true
	}
	return false
}

func (u *Url)Expired()bool{
	if u.ExpiresIn>0{
		if int(time.Now().Sub(u.CreatedTime).Seconds()) > u.ExpiresIn{
			return false
		}
	}
	return true
}