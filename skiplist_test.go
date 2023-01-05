package mycode

import (
	"fmt"
	"os"
	"os/signal"
	"testing"
	"time"
)

func Test_SkipList(t *testing.T) {
	sl := NewSkipList(10)
	t.Log(sl.Search(1))

	sl.Insert(1, 1)
	time.Sleep(time.Millisecond)
	t.Log(sl.Search(1))

	sl.Insert(1, 2)
	time.Sleep(time.Millisecond)
	t.Log(sl.Search(1))

	sl.Insert(2, 3)
	time.Sleep(time.Millisecond)
	t.Log(sl.Search(1))
	t.Log(sl.Search(2))

	sl.Delete(1)
	t.Log(sl.Search(1))
	t.Log(sl.Search(2))

	sl.Insert(1, 1)
	time.Sleep(time.Millisecond)
	sl.Insert(2, 2)
	time.Sleep(time.Millisecond)
	sl.Insert(3, 3)
	time.Sleep(time.Millisecond)
	sl.Insert(4, 4)
	time.Sleep(time.Millisecond)
	sl.Insert(5, 5)
	time.Sleep(time.Millisecond)
	sl.Insert(6, 6)
	time.Sleep(time.Millisecond)
	sl.Insert(7, 7)
	time.Sleep(time.Millisecond)
	sl.Insert(8, 8)
	time.Sleep(time.Millisecond)
	sl.Insert(9, 9)
	time.Sleep(time.Millisecond)
	sl.Insert(10, 10)
	time.Sleep(time.Millisecond)
	sl.Insert(11, 11)
	time.Sleep(time.Millisecond)
	sl.Insert(12, 12)
	time.Sleep(time.Millisecond)
	sl.Print()

	sl.Delete(1)
	sl.Delete(3)
	sl.Delete(5)
	sl.Delete(7)
	sl.Print()
}

func AlternatelyPrint() {
	c1 := make(chan struct{})
	c2 := make(chan struct{})

	printfunc := func(recive, send chan struct{}, s string) {
		defer close(recive)
		for {
			select {
			case <-recive:
				fmt.Println(s)
				time.Sleep(500 * time.Millisecond)
				send <- struct{}{}
			}
		}
	}

	go printfunc(c1, c2, "1")
	go printfunc(c2, c1, "2")

	c1 <- struct{}{}

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Kill, os.Interrupt)
	for {
		select {
		case <-sig:
		}
	}
}

func Test_AlternatelyPrint(t *testing.T) {
	AlternatelyPrint()
}
