package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo // contactInfo contactInfo の省略
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "sample@gmail.com",
			zipCode: 94000, // 最後の要素でも、コンマが必要
		},
	}

	// Goは参照渡しではなく、値渡し
	// ValueType: int, float, string, bool, structs
	// Reference Type : slice, map, function, pointer, channel
	// Goは基本的に値渡しなので、Reference Type でも、値のコピーはされる。内部で別の値を参照しているだけ。

	// 参照を渡したい場合は、以下のように明示的に
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")

	jim.updateName("jimmy") // 上記の shortcut
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	// このメソッドは、person からも &person からも 呼び出せる
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}