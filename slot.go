package ifth

import (
	"math/rand"
	"time"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
const hLetters = "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz23456789" //humanity letters

var SlotGenerator *Slot

func InitSlotGenerator() {
	SlotGenerator = NewSlot()
	SlotGenerator.Run()
}

type Slot struct {
	queue    chan string
	Humanity bool
	Length   int
}

func NewSlot() *Slot {
	s := &Slot{
		Length:   6,
		Humanity: false,
	}
	s.queue = make(chan string, 10)
	return s
}

func (s *Slot) Run() {
	go s.run()
}

func (s *Slot) run() {
	for {
		s.queue <- generate(s.Length)
		time.Sleep(1)
	}
}

func (s *Slot) Get() string {
	return <-s.queue
}

func generate(length int) string {
	rand.Seed(time.Now().UnixNano())
	mask := int64(1<<6 - 1)
	result := make([]byte, length)
	count := len(hLetters)
	for i, film, remain := 0, rand.Int63(), 10; i < length; {
		if remain == 0 {
			film = rand.Int63()
			remain = 10
		}
		if index := int(film & mask); index < count {
			result[i] = hLetters[index]
			i++
		}
		film >>= 6
		remain--
	}
	// log.Println("genereate：", string(result))
	return string(result)
}
