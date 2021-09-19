package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"math/rand"
	"time"
)

type deck []string // オブジェクト指向でいうと、[]string の継承

func newDeck() deck { // コストラクタ的なものか
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards

}

func (d deck) print() {
	// (d deck) は reciever と呼ばれる。
	// d: python でいう self 的なもの。typeの最初の１〜２文字を使うのが慣例
	for i, card := range d { // d === main.go の cards
		fmt.Println(i, card)
	}
}

func deal (d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}


func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // seed
	r := rand.New(source) // number generator

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]

	}
}

