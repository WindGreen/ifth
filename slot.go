package ifth

import (
	"math/rand"
	"time"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
const hLetters = "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz23456789" //humanity letters

const (
	Random = iota
	AutoIncrement
)

var kv map[byte]int //map rune to index of letters

var SlotGenerator *Slot

type GenFunc func() string

func InitSlotGenerator(length, algorithm int, humanity bool) {
	SlotGenerator = NewSlot()
	SlotGenerator.Length = length
	SlotGenerator.Algorithm = algorithm
	SlotGenerator.Humanity = humanity
	SlotGenerator.Run()
}

type Slot struct {
	queue     chan string
	Humanity  bool
	Length    int
	Algorithm int
	Letters   string
	Start     string
}

func NewSlot() *Slot {
	s := &Slot{
		Length:    6,
		Humanity:  false,
		Algorithm: Random,
	}
	s.queue = make(chan string, 10)
	return s
}

func (s *Slot) Run() {
	if s.Humanity {
		s.Letters = hLetters
	} else {
		s.Letters = letters
	}
	initMap()
	start := make([]byte, s.Length)
	for i := 0; i < s.Length; i++ {
		start[i] = letters[0]
	}
	s.Start = string(start)
	go s.run()
}

func (s *Slot) run() {
	for {
		var result string
		if s.Algorithm == AutoIncrement {
			result = increate(s.Start)
			s.Start = result
		} else {
			result = random(s.Length)
		}
		s.queue <- result
		time.Sleep(1)
	}
}

func (s *Slot) Get() string {
	return <-s.queue
}

func initMap() {
	kv = make(map[byte]int)
	for i, r := range letters {
		kv[byte(r)] = i
	}
}

func random(length int) string {
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

func increate(current string) string {
	n := len(current)
	m := len(letters)
	next := make([]byte, n+1)
	carry := true
	for i := n - 1; i >= 0; i-- {
		r := current[i]
		index := kv[r]
		if carry && index == m-1 {
			next[i+1] = letters[0]
			carry = true
		} else if carry {
			next[i+1] = letters[index+1]
			carry = false
		} else {
			next[i+1] = letters[index]
			carry = false
		}
		if i == 0 {
			if carry {
				next[0] = letters[0]
			} else {
				next = next[1:]
			}
		}
	}
	return string(next)
}
