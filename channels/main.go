package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) // string type の channel を作成
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // この l を引数として渡せば、値がコピーされる。渡さずにグローバルなlを参照すれば、メインルーチンと全く同じメモリ番地の値がとれてしまう
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		time.Sleep(5 * time.Second)
		fmt.Println(link + "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}