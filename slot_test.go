package ifth

import (
	"testing"
	"log"
	"time"
)

func TestNewSlot(t *testing.T) {
	s:=NewSlot()
	s.Run()
	for i:=0;i<=5;i++{
		log.Print(s.Get(),",")
		time.Sleep(1*time.Second)
	}
}
