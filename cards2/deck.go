package main

import "fmt"

type deck []string // オブジェクト指向でいうと、[]string の継承

func (d deck) print() {
	// (d deck) は reciever と呼ばれる。
	// d: python でいう self 的なもの。typeの最初の１〜２文字を使うのが慣例
	for i, card := range d { // d === main.go の cards
		fmt.Println(i, card)
	}
}
