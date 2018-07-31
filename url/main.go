package main

import (
	"ifth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ifth.InitSlotGenerator()
	_, err := ifth.InitMgo("mongo")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/:slot", GetHandle)

	router.Run(":80")
}

func GetHandle(c *gin.Context) {
	slot := c.Param("slot")
	url, err := ifth.FindUrlBySlot(slot)
	log.Printf("%#v\n", url)
	if err != nil || url == nil {
		c.Redirect(http.StatusSeeOther, "http://www.ifth.net")
		return
	}
	if url.Expired() {
		c.Redirect(http.StatusTemporaryRedirect, "http://www.ifth.net")
		return
	}
	log.Println(url.Origin)
	c.Redirect(http.StatusMovedPermanently, url.Origin)
}
