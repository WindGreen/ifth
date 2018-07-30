package ifth

import (
	"log"
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
	CreatedTime  time.Time `json:"created_time" bson:"created_time"`
	ExpiresIn    int       `json:"expires_in" bson:"expires_in"`
	Count        int64     `json:"count"`
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
	url := &Url{
		Slot:         SlotGenerator.Get(),
		Origin:       origin,
		RedirectType: Redirect301,
		CreatedTime:  time.Now(),
		ExpiresIn:    0,
		Count:        0,
	}
	for {
		log.Println(url.Slot)
		// check slot exist
		if !SlotExist(url.Slot) {
			break
		}
		url.Slot = SlotGenerator.Get()
	}
	url.Save()
	return url
}

func NewCustomUrl(slot, origin string) *Url {
	if SlotExist(slot) {
		return nil
	}
	url := &Url{
		Slot:         slot,
		Origin:       origin,
		RedirectType: Redirect301,
		CreatedTime:  time.Now(),
		ExpiresIn:    0,
		Count:        0,
	}
	url.Save()
	return url
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
	url.Count++
	url.Save()
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

func (u *Url) Expired() bool {
	if u.ExpiresIn > 0 {
		if int(time.Now().Sub(u.CreatedTime).Seconds()) > u.ExpiresIn {
			return true
		}
	}
	return false
}

func (u *Url) Save() error {
	session := GetMgo()
	_, err := session.DB("ifth").C("url").Upsert(bson.M{"slot": u.Slot}, u)
	return err
}
