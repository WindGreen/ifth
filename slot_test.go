package ifth

import (
	"log"
	"testing"
	"time"
)

func TestNewSlot(t *testing.T) {
	s := NewSlot(6)
	s.Run()
	for i := 0; i <= 5; i++ {
		log.Print(s.Get(), ",")
		time.Sleep(1 * time.Second)
	}
}

func TestSlotCarry(t *testing.T) {
	initMap()
	s := increate("A")
	if s != "B" {
		t.Fatal("A->B:", s)
	}
	s = increate("9")
	if s != "AA" {
		t.Fatal("9->AA", s)
	}
	s = increate("AA")
	if s != "AB" {
		t.Fatal("AA->AB", s)
	}
	s = increate("A9")
	if s != "BA" {
		t.Fatal("A9->BA", s)
	}
	s = increate("99")
	if s != "AAA" {
		t.Fatal("99->AAA", s)
	}
	s = increate("99ABcd99")
	if s != "99ABceAA" {
		t.Fatal("99ABcd99->99ABceAA", s)
	}
	s = increate("9999999999999999999999999999999999999999999999")
	t.Log(s)
}
