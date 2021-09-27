package main

import (
	"fmt"
	"net/http"
	"os"
)

// struct の型に、interface を指定すると、「そのinterfaceを満たす型すべて」の意になる
// interface の中に別のinterfaceを列挙し、interfaceをまとめられる。
// 「列挙したすべてのinterfaceを満たすもの」が、そのinterfaceの型グループとなる

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(resp)
 }
