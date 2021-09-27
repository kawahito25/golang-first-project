package main

import "fmt"

// 抽象的な型
type bot interface {
	// getGreeting() string を持つ type はすべて、bot型のメンバーとなる。
	getGreeting() string
}

// 具体的な型
type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}


func (englishBot) getGreeting() string {
	// （仮定）VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string { // sb spanishBot としてもOKだが、sbを使わないため省略できる
	return "Hola!"
}



