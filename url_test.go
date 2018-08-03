package ifth

import (
	"fmt"
	"log"
	"testing"
)

func InitTest() {
	InitSlotGenerator(4)
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

func BenchmarkNewUrl(b *testing.B) {
	InitTest()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = NewUrl("http://test.tickpay.org", false)
		}
	})
}

func TestFindUrls(t *testing.T) {
	InitTest()
	urls, err := FindHottestUrls(3)
	if err != nil {
		t.Fatal(err)
	}
	Dump(urls)
	urls, err = FindNewestUrls(3)
	if err != nil {
		t.Fatal(err)
	}
	Dump(urls)
}

func Dump(urls []Url) {
	for _, url := range urls {
		fmt.Printf("%s %d %s\n\n", url.Slot, url.Count, url.CreatedTime)
	}
}
