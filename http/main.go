package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

// struct の型に、interface を指定すると、「そのinterfaceを満たす型すべて」の意になる
// interface の中に別のinterfaceを列挙し、interfaceをまとめられる。
// 「列挙したすべてのinterfaceを満たすもの」が、そのinterfaceの型グループとなる


type logWritter struct {}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// bs := make([]byte, 99999) // make(スライス, 要素の数を指定)
	// resp.Body.Read(bs) // からのbyteスライスに、bodyのデータを格納してくれる

	lw := logWritter{}

	io.Copy(lw, resp.Body)


    // fmt.Println(string(bs))
 }

 func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
 }