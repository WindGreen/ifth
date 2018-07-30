package ifth

import (
	"log"
	"testing"
)

func InitTest() {
	InitSlotGenerator()
	_, err := InitMgo("localhost")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("init ok")
}

func TestNewUrl(t *testing.T) {
	InitTest()
	url := NewUrl("http://test.tickpay.org", false)
	log.Println(url)
}
