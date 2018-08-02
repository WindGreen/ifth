package main

import (
	"fmt"
	"ifth"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	yaml "gopkg.in/yaml.v2"
)

var config Config

func main() {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	_, err = ifth.InitMgo(config.MongoDB.Host)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/:slot", GetHandle)

	router.Run(fmt.Sprintf(":%d", config.Url.Port))
}

func GetHandle(c *gin.Context) {
	slot := c.Params.ByName("slot")
	url, err := ifth.FindUrlBySlot(slot)
	log.Printf("%#v\n", url)
	if err != nil || url == nil {
		c.Redirect(http.StatusSeeOther, config.WWW.Home)
		return
	}
	if url.Expired() {
		c.Redirect(http.StatusTemporaryRedirect, config.WWW.Home)
		return
	}
	log.Println(url.Origin)
	c.Redirect(http.StatusMovedPermanently, url.Origin)
}
