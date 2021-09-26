package main

import "fmt"

func main() {
	// struct との違いは、key全て、value全てが、同じ型で統一されること

	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//var colors map[string]string
	//colors := map[string]string{}
	//colors := make(map[int]string)
	//
	//colors[10] = "#ffffff"

	// これは struct の構文。map では使えない。keyの型が分からないから。
	// structName.white

	//delete(colors, 10)

	fmt.Println(colors)

	printMap(colors)
}


func printMap(c map[string]string) { // mapはReference Type, struct は Value Type
	// struct で以下の書き方はできない
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// map は colors のように同じようなものを列挙するのに便利
// struct は 様々な要素を持った「もの」を表現するのに便利

// mapは key: value ペアを 追加したり、削除したりできる
// structは最初に、type として定義した枠は、後から変えられない

// 一般的に、struct のほうがよく使われる