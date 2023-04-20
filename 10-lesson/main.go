package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	 taskOne()
	//taskTwo()
	//taskThree()
	// ch:=make(chan string)
	// wg:=sync.WaitGroup{}
	// downloadFile("https://raw.githubusercontent.com/maxmurjon/Golang-bootcamp/main/lesson10/fisrt_Part/go.mod",ch,&wg)
	// fmt.Println(ch)
}
func taskOne() {
	var wg sync.WaitGroup
	toqSonlar := func() {
		for i := 1; i < 11; i++ {
			if i%2 != 0 {
				fmt.Println(i)
			}
		}
		wg.Done()
	}

	juftSonlar := func() {
		for i := 1; i < 11; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
		wg.Done()
	}
	wg.Add(2)
	go toqSonlar()
	// time.Sleep(time.Second *5)
	go juftSonlar()
	wg.Wait()
}
	func taskTwo() {
		ch := make(chan int)
		var input = func() {
			for {
				var number int
				fmt.Printf("Son kiriting ")
				_, err := fmt.Scan(&number)
				if err != nil {
					close(ch) // kanalni yopish
					return
				}
				ch <- number
			}
		}
		channel := func() {
			sum := 0
			for num := range ch {
				sum += num
			}
			fmt.Println(sum)
		}
		go input()
		channel()
	}

func taskThree() {
	ch := make(chan int)
	rand.Seed(time.Now().UnixNano())
	go func() {
		randNumber := rand.Intn(101)
		ch <- randNumber

	}()

	func() {
		fmt.Println(<-ch)
	}()
}

func downloadFile(url string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error downloading %s: %v", url, err)
		return
	}
	defer resp.Body.Close()
	ch <- fmt.Sprintf("%s downloaded %d bytes.", url, resp)
}
