package main

import (
	"ifth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	ifth.InitSlotGenerator()
	ifth.InitMgo("mongodb://localhost/")

	router:=gin.New()
	router.GET("/:slot",)
}

func GetHandle(c *gin.Context){
	slot:=c.Param("slot")
	url,err:=ifth.FindUrlBySlot(slot)
	if err!=nil{
		c.Redirect(status)
	}
	if url.Expired(){
		c.Redirect(http.StatusMovedPermanently,"www.ifth.net")
	}
	c.Redirect(http.StatusMovedPermanently,url.Origin)
}
